//go:build integration
// +build integration

package storage

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	clientV1 "github.com/medtrib/qiniu-sdk/v7/client"
	"github.com/medtrib/qiniu-sdk/v7/storagev2/uplog"
)

func TestUploadManagerUplogForm(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test-uplog-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	uplog.SetUplogFileBufferDirPath(tmpDir)
	defer uplog.SetUplogFileBufferDirPath("")

	if err = uplog.FlushBuffer(); err != nil {
		t.Fatal(err)
	}

	writeFileBufferInterval := uplog.GetWriteFileBufferInterval()
	uplog.SetFlushFileBufferInterval(1 * time.Hour)
	defer uplog.SetFlushFileBufferInterval(writeFileBufferInterval)

	clientV1.DebugMode = true
	defer func() {
		clientV1.DebugMode = false
	}()

	data := []byte("hello, 七牛！！！")
	dataLen := int64(len(data))

	tempFile, err := ioutil.TempFile("", "TestUploadManagerFormPut-*")
	if err != nil {
		t.Fatalf("create temp file error:%v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	tempFile.Write(data)

	uploadManager := getUploadManagerV2([]string{"mock03.qiniu.com", "mock04.qiniu.com"})
	var ret Ret

	// 上传 file
	source, err := NewUploadSourceFile(tempFile.Name())
	if err != nil {
		t.Fatalf("upload source file error:%v", err)
	}
	err = uploadManager.Put(context.Background(), &ret, getUploadToken(), nil, source, &UploadExtra{
		TryTimes: 1,
	})
	if err != nil {
		t.Fatalf("form upload file error:%v", err)
	}
	if len(ret.Key) == 0 || len(ret.Hash) == 0 {
		t.Fatal("form upload file error, key or hash is empty")
	}

	time.Sleep(1 * time.Second)

	uplogFile, err := os.Open(filepath.Join(tmpDir, "uplog_v4_01.buffer"))
	if err != nil {
		t.Fatalf("uplog file error:%v", err)
	}
	defer uplogFile.Close()

	uplogFileJsonDecoder := json.NewDecoder(uplogFile)
	uplogs := make([]map[string]interface{}, 0, 4)
	for {
		var uplog map[string]interface{}
		if err := uplogFileJsonDecoder.Decode(&uplog); err != nil {
			break
		}
		uplogs = append(uplogs, uplog)
	}
	if len(uplogs) != 4 {
		t.Fatalf("unexpected uplog count:%v", len(uplogs))
	}
	if uplogs[0]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[0]["log_type"])
	}
	if uplogs[0]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[0]["api_type"])
	}
	if uplogs[0]["api_name"] != "postObject" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[0]["api_name"])
	}
	if uplogs[0]["error_type"] != "unknown_host" {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[0]["error_type"])
	}
	if uplogs[0]["host"] != "mock03.qiniu.com" {
		t.Fatalf("unexpected uplog host:%v", uplogs[0]["host"])
	}
	if uplogs[0]["path"] != "/" {
		t.Fatalf("unexpected uplog path:%v", uplogs[0]["path"])
	}
	if uplogs[0]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[0]["method"])
	}
	if uplogs[0]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[0]["target_bucket"])
	}
	if uplogs[1]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[1]["log_type"])
	}
	if uplogs[1]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[1]["api_type"])
	}
	if uplogs[1]["api_name"] != "postObject" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[1]["api_name"])
	}
	if uplogs[1]["error_type"] != "unknown_host" {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[1]["error_type"])
	}
	if uplogs[1]["host"] != "mock04.qiniu.com" {
		t.Fatalf("unexpected uplog host:%v", uplogs[1]["host"])
	}
	if uplogs[1]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[1]["target_bucket"])
	}
	if uplogs[1]["path"] != "/" {
		t.Fatalf("unexpected uplog path:%v", uplogs[1]["path"])
	}
	if uplogs[1]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[1]["method"])
	}
	if uplogs[2]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[2]["log_type"])
	}
	if uplogs[2]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[2]["api_type"])
	}
	if uplogs[2]["api_name"] != "postObject" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[2]["api_name"])
	}
	if uplogs[2]["error_type"] != nil {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[2]["error_type"])
	}
	if uplogs[2]["port"] != float64(443) {
		t.Fatalf("unexpected uplog port:%v", uplogs[2]["port"])
	}
	if uplogs[2]["remote_ip"] == nil {
		t.Fatalf("unexpected uplog remote_ip:%v", uplogs[2]["remote_ip"])
	}
	if uplogs[2]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[2]["target_bucket"])
	}
	if uplogs[2]["path"] != "/" {
		t.Fatalf("unexpected uplog path:%v", uplogs[2]["path"])
	}
	if uplogs[2]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[2]["method"])
	}
	if uplogs[2]["status_code"] != float64(200) {
		t.Fatalf("unexpected uplog status_code:%v", uplogs[2]["status_code"])
	}
	if uplogs[3]["log_type"] != "quality" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[3]["log_type"])
	}
	if uplogs[3]["result"] != "ok" {
		t.Fatalf("unexpected uplog result:%v", uplogs[3]["result"])
	}
	if uplogs[3]["up_type"] != "form" {
		t.Fatalf("unexpected uplog up_type:%v", uplogs[3]["up_type"])
	}
	if uplogs[3]["regions_count"] != float64(2) {
		t.Fatalf("unexpected uplog regions_count:%v", uplogs[3]["regions_count"])
	}
	if uplogs[3]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[3]["api_type"])
	}
	if uplogs[3]["file_size"] != float64(dataLen) {
		t.Fatalf("unexpected uplog file_size:%v", uplogs[3]["file_size"])
	}
	if uplogs[3]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[3]["target_bucket"])
	}
}

func TestUploadManagerUplogResumableV1(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "test-uplog-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	uplog.SetUplogFileBufferDirPath(tmpDir)
	defer uplog.SetUplogFileBufferDirPath("")

	if err = uplog.FlushBuffer(); err != nil {
		t.Fatal(err)
	}

	writeFileBufferInterval := uplog.GetWriteFileBufferInterval()
	uplog.SetFlushFileBufferInterval(1 * time.Hour)
	defer uplog.SetFlushFileBufferInterval(writeFileBufferInterval)

	clientV1.DebugMode = true
	defer func() {
		clientV1.DebugMode = false
	}()

	length := 1024 * 1024 * 4
	data := make([]byte, length, length)
	data[0] = 8
	data[length-1] = 8
	tempFile, err := ioutil.TempFile("", "TestUploadManagerResumeV1Upload-*")
	if err != nil {
		t.Fatalf("create temp file error:%v", err)
	}
	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()
	tempFile.Write(data)

	uploadManager := getUploadManagerV2([]string{"mock05.qiniu.com", "mock06.qiniu.com"})
	var ret Ret

	// 上传 file
	source, err := NewUploadSourceFile(tempFile.Name())
	if err != nil {
		t.Fatalf("upload source file error:%v", err)
	}
	err = uploadManager.Put(context.Background(), &ret, getUploadToken(), nil, source, &UploadExtra{
		TryTimes:            1,
		UploadResumeVersion: UploadResumeV1,
		UploadThreshold:     1024 * 1024 * 2,
		PartSize:            1024 * 1024,
	})
	if err != nil {
		t.Fatalf("form upload file error:%v", err)
	}
	if len(ret.Key) == 0 || len(ret.Hash) == 0 {
		t.Fatal("form upload file error, key or hash is empty")
	}

	time.Sleep(1 * time.Second)

	uplogFile, err := os.Open(filepath.Join(tmpDir, "uplog_v4_01.buffer"))
	if err != nil {
		t.Fatalf("uplog file error:%v", err)
	}
	defer uplogFile.Close()

	uplogFileJsonDecoder := json.NewDecoder(uplogFile)
	uplogs := make([]map[string]interface{}, 0, 12)
	for {
		var uplog map[string]interface{}
		if err := uplogFileJsonDecoder.Decode(&uplog); err != nil {
			break
		}
		uplogs = append(uplogs, uplog)
	}
	if len(uplogs) != 10 {
		t.Fatalf("unexpected uplog count:%v", len(uplogs))
	}
	if uplogs[0]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[0]["log_type"])
	}
	if uplogs[0]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[0]["api_type"])
	}
	if uplogs[0]["api_name"] != "resumableUploadV1MakeBlock" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[0]["api_name"])
	}
	if uplogs[0]["error_type"] != "unknown_host" {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[0]["error_type"])
	}
	if uplogs[0]["host"] != "mock05.qiniu.com" {
		t.Fatalf("unexpected uplog host:%v", uplogs[0]["host"])
	}
	if uplogs[0]["path"] != "/mkblk/4194304" {
		t.Fatalf("unexpected uplog path:%v", uplogs[0]["path"])
	}
	if uplogs[0]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[0]["method"])
	}
	if uplogs[0]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[0]["target_bucket"])
	}
	if uplogs[1]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[1]["log_type"])
	}
	if uplogs[1]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[1]["api_type"])
	}
	if uplogs[1]["api_name"] != "resumableUploadV1MakeBlock" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[1]["api_name"])
	}
	if uplogs[1]["error_type"] != "unknown_host" {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[1]["error_type"])
	}
	if uplogs[1]["host"] != "mock06.qiniu.com" {
		t.Fatalf("unexpected uplog host:%v", uplogs[1]["host"])
	}
	if uplogs[1]["path"] != "/mkblk/4194304" {
		t.Fatalf("unexpected uplog path:%v", uplogs[1]["path"])
	}
	if uplogs[1]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[1]["method"])
	}
	if uplogs[1]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[1]["target_bucket"])
	}
	if uplogs[2]["log_type"] != "block" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[2]["log_type"])
	}
	if uplogs[2]["up_api_version"] != float64(1) {
		t.Fatalf("unexpected uplog up_api_version:%v", uplogs[2]["up_api_version"])
	}
	if uplogs[2]["requests_count"] != float64(2) {
		t.Fatalf("unexpected uplog requests_count:%v", uplogs[2]["requests_count"])
	}
	if uplogs[2]["error_type"] != "unknown_host" {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[2]["error_type"])
	}
	if uplogs[2]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[2]["api_type"])
	}
	if uplogs[2]["file_size"] != float64(length) {
		t.Fatalf("unexpected uplog file_size:%v", uplogs[2]["file_size"])
	}
	if uplogs[2]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[2]["target_bucket"])
	}
	if uplogs[3]["log_type"] != "request" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[3]["log_type"])
	}
	if uplogs[3]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[3]["api_type"])
	}
	if uplogs[3]["api_name"] != "resumableUploadV1MakeBlock" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[3]["api_name"])
	}
	if uplogs[3]["error_type"] != nil {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[3]["error_type"])
	}
	if uplogs[3]["port"] != float64(443) {
		t.Fatalf("unexpected uplog port:%v", uplogs[3]["port"])
	}
	if uplogs[3]["remote_ip"] == nil {
		t.Fatalf("unexpected uplog remote_ip:%v", uplogs[3]["remote_ip"])
	}
	if uplogs[3]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[3]["target_bucket"])
	}
	if uplogs[3]["path"] != "/mkblk/4194304" {
		t.Fatalf("unexpected uplog path:%v", uplogs[3]["path"])
	}
	if uplogs[3]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[3]["method"])
	}
	if uplogs[3]["status_code"] != float64(200) {
		t.Fatalf("unexpected uplog status_code:%v", uplogs[3]["status_code"])
	}
	if uplogs[3]["bytes_sent"] != float64(1024*1024) {
		t.Fatalf("unexpected uplog bytes_sent:%v", uplogs[3]["bytes_sent"])
	}
	for i := 4; i < 7; i++ {
		if uplogs[i]["log_type"] != "request" {
			t.Fatalf("unexpected uplog log_type:%v", uplogs[i]["log_type"])
		}
		if uplogs[i]["api_type"] != "kodo" {
			t.Fatalf("unexpected uplog api_type:%v", uplogs[i]["api_type"])
		}
		if uplogs[i]["api_name"] != "resumableUploadV1Bput" {
			t.Fatalf("unexpected uplog api_name:%v", uplogs[i]["api_name"])
		}
		if uplogs[i]["error_type"] != nil {
			t.Fatalf("unexpected uplog error_type:%v", uplogs[i]["error_type"])
		}
		if uplogs[i]["port"] != float64(443) {
			t.Fatalf("unexpected uplog port:%v", uplogs[i]["port"])
		}
		if uplogs[i]["remote_ip"] == nil {
			t.Fatalf("unexpected uplog remote_ip:%v", uplogs[i]["remote_ip"])
		}
		if uplogs[i]["target_bucket"] != testBucket {
			t.Fatalf("unexpected uplog target_bucket:%v", uplogs[i]["target_bucket"])
		}
		if !strings.HasPrefix(uplogs[i]["path"].(string), "/bput/") {
			t.Fatalf("unexpected uplog path:%v", uplogs[i]["path"])
		}
		if uplogs[i]["method"] != "POST" {
			t.Fatalf("unexpected uplog method:%v", uplogs[i]["method"])
		}
		if uplogs[i]["status_code"] != float64(200) {
			t.Fatalf("unexpected uplog status_code:%v", uplogs[i]["status_code"])
		}
		if uplogs[i]["bytes_sent"] != float64(1024*1024) {
			t.Fatalf("unexpected uplog bytes_sent:%v", uplogs[i]["bytes_sent"])
		}
	}
	if uplogs[7]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[7]["api_type"])
	}
	if uplogs[7]["api_name"] != "resumableUploadV1MakeFile" {
		t.Fatalf("unexpected uplog api_name:%v", uplogs[7]["api_name"])
	}
	if uplogs[7]["error_type"] != nil {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[7]["error_type"])
	}
	if uplogs[7]["port"] != float64(443) {
		t.Fatalf("unexpected uplog port:%v", uplogs[7]["port"])
	}
	if uplogs[7]["remote_ip"] == nil {
		t.Fatalf("unexpected uplog remote_ip:%v", uplogs[7]["remote_ip"])
	}
	if uplogs[7]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[7]["target_bucket"])
	}
	if uplogs[7]["path"] != "/mkfile/4194304" {
		t.Fatalf("unexpected uplog path:%v", uplogs[7]["path"])
	}
	if uplogs[7]["method"] != "POST" {
		t.Fatalf("unexpected uplog method:%v", uplogs[7]["method"])
	}
	if uplogs[7]["status_code"] != float64(200) {
		t.Fatalf("unexpected uplog status_code:%v", uplogs[7]["status_code"])
	}
	if uplogs[8]["log_type"] != "block" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[8]["log_type"])
	}
	if uplogs[8]["up_api_version"] != float64(1) {
		t.Fatalf("unexpected uplog up_api_version:%v", uplogs[8]["up_api_version"])
	}
	if uplogs[8]["requests_count"] != float64(5) {
		t.Fatalf("unexpected uplog requests_count:%v", uplogs[8]["requests_count"])
	}
	if uplogs[8]["error_type"] != nil {
		t.Fatalf("unexpected uplog error_type:%v", uplogs[8]["error_type"])
	}
	if uplogs[8]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[8]["api_type"])
	}
	if uplogs[8]["file_size"] != float64(length) {
		t.Fatalf("unexpected uplog file_size:%v", uplogs[8]["file_size"])
	}
	if uplogs[8]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[8]["target_bucket"])
	}
	if uplogs[9]["log_type"] != "quality" {
		t.Fatalf("unexpected uplog log_type:%v", uplogs[9]["log_type"])
	}
	if uplogs[9]["requests_count"] != float64(7) {
		t.Fatalf("unexpected uplog requests_count:%v", uplogs[9]["requests_count"])
	}
	if uplogs[9]["result"] != "ok" {
		t.Fatalf("unexpected uplog result:%v", uplogs[9]["result"])
	}
	if uplogs[9]["up_type"] != "resumable_v1" {
		t.Fatalf("unexpected uplog up_type:%v", uplogs[9]["up_type"])
	}
	if uplogs[9]["regions_count"] != float64(2) {
		t.Fatalf("unexpected uplog regions_count:%v", uplogs[9]["regions_count"])
	}
	if uplogs[9]["api_type"] != "kodo" {
		t.Fatalf("unexpected uplog api_type:%v", uplogs[9]["api_type"])
	}
	if uplogs[9]["file_size"] != float64(length) {
		t.Fatalf("unexpected uplog file_size:%v", uplogs[9]["file_size"])
	}
	if uplogs[9]["target_bucket"] != testBucket {
		t.Fatalf("unexpected uplog target_bucket:%v", uplogs[9]["target_bucket"])
	}
}
