// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 从指定 URL 抓取指定名称的对象并存储到该空间中
package fetch_object

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	FromUrl     string                          // 指定抓取的 URL
	ToEntry     string                          // 指定目标对象空间与目标对象名称，格式为 <目标对象空间>:<目标对象名称>
	Host        string                          // 指定抓取 URL 请求用的 HOST 参数
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Hash       string // 抓取的对象内容的 Etag 值
	ObjectName string // 抓取后保存的对象名称
	Size       int64  // 对象大小，单位为字节
	MimeType   string // 对象 MIME 类型
}

// 抓取到的文件元信息
type FetchedObjectMetadata = Response
type jsonResponse struct {
	Hash       string `json:"hash"`            // 抓取的对象内容的 Etag 值
	ObjectName string `json:"key"`             // 抓取后保存的对象名称
	Size       int64  `json:"fsize,omitempty"` // 对象大小，单位为字节
	MimeType   string `json:"mimeType"`        // 对象 MIME 类型
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{Hash: j.Hash, ObjectName: j.ObjectName, Size: j.Size, MimeType: j.MimeType})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Hash = nj.Hash
	j.ObjectName = nj.ObjectName
	j.Size = nj.Size
	j.MimeType = nj.MimeType
	return nil
}
func (j *Response) validate() error {
	if j.Hash == "" {
		return errors.MissingRequiredFieldError{Name: "Hash"}
	}
	if j.ObjectName == "" {
		return errors.MissingRequiredFieldError{Name: "ObjectName"}
	}
	if j.MimeType == "" {
		return errors.MissingRequiredFieldError{Name: "MimeType"}
	}
	return nil
}
