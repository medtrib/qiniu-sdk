//go:build integration
// +build integration

package storage

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	imgDataStr = "/9j/4AAQSkZJRgABAQAASABIAAD/4QBARXhpZgAATU0AKgAAAAgAAYdpAAQAAAABAAAAGgAAAAAAAqACAAQAAAABAAAAgKADAAQAAAABAAAAgAAAAAD/7QA4UGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAAAA4QklNBCUAAAAAABDUHYzZjwCyBOmACZjs+EJ+/+ICOElDQ19QUk9GSUxFAAEBAAACKEFEQkUCEAAAbW50clJHQiBYWVogB9AACAALABMANAAYYWNzcEFQUEwAAAAAbm9uZQAAAAAAAAAAAAAAAAAAAAAAAPbWAAEAAAAA0y1BREJFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKY3BydAAAAPwAAAAyZGVzYwAAATAAAABkd3RwdAAAAZQAAAAUYmtwdAAAAagAAAAUclRSQwAAAbwAAAAOZ1RSQwAAAcwAAAAOYlRSQwAAAdwAAAAOclhZWgAAAewAAAAUZ1hZWgAAAgAAAAAUYlhZWgAAAhQAAAAUdGV4dAAAAABDb3B5cmlnaHQgMjAwMCBBZG9iZSBTeXN0ZW1zIEluY29ycG9yYXRlZAAAAGRlc2MAAAAAAAAACkFwcGxlIFJHQgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABYWVogAAAAAAAA81EAAQAAAAEWzFhZWiAAAAAAAAAAAAAAAAAAAAAAY3VydgAAAAAAAAABAc0AAGN1cnYAAAAAAAAAAQHNAABjdXJ2AAAAAAAAAAEBzQAAWFlaIAAAAAAAAHm9AABBUgAABLlYWVogAAAAAAAAVvgAAKwvAAAdA1hZWiAAAAAAAAAmIgAAEn8AALFw/8AAEQgAgACAAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/bAEMAAgICAgICAwICAwUDAwMFBgUFBQUGCAYGBgYGCAoICAgICAgKCgoKCgoKCgwMDAwMDA4ODg4ODw8PDw8PDw8PD//bAEMBAgICBAQEBwQEBxALCQsQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEP/dAAQACP/aAAwDAQACEQMRAD8A/fyiiigAoopGYKMngUALQSB1rAudcXzmtNPja6nXqqdF/wB5jwP5+1cx4l17TPDenNrHjnxDbaDYL1Z5UgXPXb5khyx9lAJ7VUYuTslqVGLk7JanfyXMEQ3SOFHqTWe+vaPGcPeRL9XUf1r4i8T/ALbv7NfhWVodNN54nnXI321sXXd7yXbRnHuoI9K8mvP+Clvhy2Zl0jwBPMgI2mS+SEkd8hYJMfma9uhwzj6ivGk/nZfnY9yhwxj6ivGi/np+dj9Nk17R5DhLyJvo6n+taEdzBKN0bhh6g1+XNn/wUt8OXLKur+AJ4UJO4x3yTEDtgNBHn8xXrPhj9t39mvxVKsOpG88MTtgb7m2KLuPpJaNIce7AD1or8M4+mrypP5Wf5XFX4Zx9NXlSfy1/K596Ag9KK858N67pviPT11jwP4httesDjBSVJ1yRnb5kZyp9mBI7109triGZbS/ia1nbor9G/wB1hwfw59q8SUXF2a1PElFp2a1N+ikVgwyORS1JIUUUUAf/0P38oopruEUsegoAZNMkKF3OAK4661BtRhlvJbgWOlQKzyTswTcijLMGPCoAOWP4cc1BqF7DqTXEt3MsGlWIZp5HYIjbBlgWPARRy56dvWvxi/a0/a21D4p6hceAvAdw1p4MtHCu6go+oOh++/cQg8oh64DMM4VPZyXJauNq8kNEt32/4J6+TZPUxlXkholu+x9EfHX9vbS/DRuPB3wLgiupocxvq0q7oFboTbxn/WEH/lo/ykjhWBDV+Xfi7x34v8e6tJrfjHV7nV72TP7y4lZyAedqg8Ko7KuAOwrjd2e+aN3vX7LleSYfCRtSjr36s/Y8qyfD4ONqUde/Un3Ubqg3e9G73r1j2fak+6jdUG73o3e9Ae1Oy8I+O/F/gLVk1zwdq9zpF7Hj95bytGSM52sBwynurZB7iv1E+BX7e2l+JTb+DvjpBFazTYjTVolCwM3QG4jH+rJP/LRPlBxlUALV+RO73pdxHfFeTmmSYfFxtVjr36o8bNcnw+MjarHXv1P6kbXUG06CK8juBfaVOqvHOrB9qMMqxYcMhB4YduvrXYwypMgdDkGvws/ZK/a2vvhZfweAfHs7Xfg28fbHI2XfT5HP309YSeZE7ffXncr/ALN6fexaa9vJaTLcaVfBWgkRg6LvGVAYZBRhyh6dvSvxrOslq4KryT1T2ff/AIJ+OZzk1TB1eSeqez7ne0U1HDqGHenV4x5B/9H9/K5TxHfTBY9Os2xcXR2Kf7o6s3/ARz9cCuokYIhY9q8g1/xNp/h3TvEHjrWG22OiW0rk5wdkCmSTbnuxAUepGKqMXJpLccYtuyPgD9vP47t4a0qD4GeDrkxTXUSy6s8Z+ZYG5jtyeoMn35OhK7R0ZhX5JbjXR+OfGGq+PPF+r+MNbfzL3VrmS4k54Uu2dq56Ko+VR2AArlN1fvmSZXHCYeNJb9fNn7Vk2CjhaEaa36+p+lH7BfwS8D/Ea28VeJfHukRaxb2bW9paxz5Mau4Z5WwCPmACAexNfRvxR/4J+/C/xPbS3fw6uJfCmpAErGzPc2bnrhlcmVMn+JXIA/gNcT/wTY12CfwZ4v8ADo4ltL63ujx1W4jKDn2MX61+lm6vzXiDOcXQzGp7ObVradNl02Pgc6zfFUsdN05tWtp027bH4BXX7GX7RMPiO58OW3haS6a3IxdJLEtpIp6Mk8jIhyO2Qw6FQeK9A0//AIJ8/Hu7jV7o6VYscZSW7LEZHfykkHHfB+lft7uo3U58d4xqyUV8n/mVPjXGNWSS+X/BPxF1D/gnx8ebONntX0q+KkgLFdlS2M8jzY4xg+5B5+uPOk/Yz/aK/wCEhtfD1x4Vkha5bH2kyxPaRr1LSTRu6KAOcZ3HoAW4r9/91G6inx3jFuov5P8AzCHGuMS1SfyPgP4V/wDBPz4beGreK++JdzJ4p1LALQRu9tZIfQbCsr4PcsoPdK8K/bx+BngL4eeHPC/irwBosOixPczWd0sG7bIzoJIidxPICSc98+1frjur86f+CkOvQW3wx8M+HG/11/qpul4/htYHRufrOtRkWc4uvmFP2k27vbps+mxnk+cYqrjqbnNu726bdtj8ctxr9bf2DPju3iXSrj4GeMbnzZrWJpdJeQ/M0CjMluCeSY/vx9wu7oFUD8it1dZ4F8Yar4C8YaR4x0R9l5pFzFcR5JwxRgSrYxlWGVYdwSK/TM7yuOLw8qT36eTP0DOcFHFUJU3v09T+n/w5fSlZNOvGzcWp2Mf7w6q34jn68V1deO6B4m0/xFp3h/x1o7brHW7aJwc5OyZBJHuxxlSSp9zivX42DoGHevwOUXFtPc/FZRadmf/S/ejWJxb2Mkh4wCa/Pv8Aba8VzeF/2bzpsTbZ/Et3bWz4OGCsWupD9MxhT9cV91+MpDHol0w7Rt/Kvyz/AOCk9+1v4c8BaUrEJNPfSFccEwpAoOfbefzr3eGaKqY+lF97/cr/AKHqZJTUsXTT7/lqfk5u96kiR5pUhjBZnIAA6kmqu416d8FLKHVPjD4I025AMV1renROCAw2vcIpyDweD3r9yrVVCDn2P1ariOWLl2P3y/Z6+Emm/Bj4Y6X4Wt4lGpSxrc6lL/FJdyKC4J7rH9xP9kZ6kmvb91V91G6v52xGInVqSqTd29T8arVpVJuc3qyxuo3VX3UbqxMixuo3VX3UbqALG6vIvjl8KdJ+M3w41TwZqKKLmRDLYzt1gu0BMT5wSFJ+V8dULDrgj1XdRurWhXlTmqkHZrU0pVZQkpxeqP5Zru3nsbqayuVMc0DsjqwwVZTggj1Bqvu969b/AGg7SKw+OXjy1gAWMa1fsqqAoUNO7YAHQDOBXj241/ROHrc9OM+6TP2WjiOaCl3P3K/Yl8VzeKP2b106Zt0/hq7uLZOfmKqVuoz9MyFR9K/QTR5xcWMcgOcgGvyb/wCCbF+9x4b8eaSzEpBPYyBOwMyTKT+Plj8q/UvwbIZNEtmPeNf5V+HcS0VTx9WK73+9X/U/K87pqOLqJd/z1P/T/cvxqpfRLlR3jb+Vflb/AMFKLRrjw/4C1VeUgmv4ycjGZlgYcdf+WZ/ziv1k8RQefp0qYzkGvzu/ba8MT+Kf2cTqcA3TeG7u2un7sVXdayY/GQMfYZr3eGqyp46lJ97ferfqellFXkxMJPv+eh+Iu6ur8Ca+3hfxroPiRMg6VfW10ME5zDIr9sHt2NcbuoDkHI6iv3GceZOL6n6PKd00z+qeG4huYY7i3cSRSqHRlOQysMgg9wRUu6vir9i745af8S/hpZ+EdRuVHiPwvCltJEzfPNaRgLDMo7gLhHPJDAE43Cvs3dX8+Y/Bzw9aVGe6Py3E0JUpuEuhY3Ubqr7qN1chgWN1G6q+6jdQBY3UjSKil3IVVGSTwABUG6vkf9sL446d8Kfhne6HZXK/8JJ4khe1tIlPzxQyApLOQOVCrlUPdyMZ2tjqwWEnXqxo01qzbD0JVJqEd2fiv8U/Ekfi74k+KPE8Tbo9V1O7uUPONksrOoGc8AHA9q4LdUJcsSx6nmk3V/QdOCjFRWyP1KE+VKK6H63/APBNi0a38PePNVOQs89hHnIx+5Wdjx1/j/zzX6o+ClK6JbKeyL/KvgH9ibwzL4W/ZxXUpl2TeJLu5uk/vbWK2seR9Yyw9jmv0Q8OweRp0SYxgCvw/iWuqmPqyXe33K36H5zm9XnxM5ef5aH/1P3zvIvOt3T1FfPmveHNP8Rab4h8Ba0ubHWbeaNhgEhJ0Mblc5AKkhgexOa+jCMjFeUeNNOmtpo9WtFzJAd2P7y9GX8R+tVCTi1JboabTuj+Yjxx4T1fwH4u1bwfrsflX2k3EkEgHQlDjcp7qw5U9wQa+nfg5+yLffGzwG3ivwp4y0xNSjZ0l02USiSBgSFErKpK7wAykKykHrkED6c/bx+BreI9Mg+OHhGDzZ7SJItVjjXLPAvEdwQOcxj5H9F2nACsa/N34VfFDxN8JPGdh4w8M3DRTWrgSxZIjnhJG+KQd1YD8DgjBAI/a8Nj6uNwSq4WVp9euq3XzPtoYydaipUnaRqX1n8Tv2e/iIIpxP4e8SaPJuR1I5BGNysMpJG68ZGVYcciv0c+Fn/BRPw/eWsOnfFnSZbK8VQpvrBRJDIR/E8LEMnvtL89FA6eufG34S+Hf2tvhPonjPwfLFb619nW506eU4Vkk5ltJmUHGGyAcHY4PQMxr8kNb/Z7+N/h/Un0rUPBOrGZWKgxWks0b4OMpJGrI491JFebRqYPNKVsUlGpHR62a9PL77HKqlHFR/fK0kfvJ4J/aB+DvxFvIdN8H+KbW+vbjPl27b4ZnIBYhY5lRjgAngV7Dur84/2E/gR4m+Hkeu+NvH+hzaTqt6sdrYpdLsmSDl5iYz8ybmCD5gD8pxwTX6J76/Oc5wtCjiJUsPK8V10/Q+cxlOEKjjTd0Wd1eR+Nvj38IPh1eTab4y8U2mn3tuFMlvl5p03AMN0cSu4yCCOOhzXqe+vzJ/4KO+CbCTw/4Z+IsMIS8huG0yeQDmSORGmiDf7hSTH+8fbDyTBUsRiY0araT7BgqMalRQnszV+KP/BRPwxp9rNYfCnSZdSvWG1by+XyrdD/AHliU739txTnscYP5YeNPHHif4heIrrxT4v1CTUdSvG3PLIew4CqBgKqjhVUAAcAVx26jdX7JlmSYbCL9zHXu9z7LC4alR+BE+73rq/A3hPVvHni7SfB+hx+be6tcR28Y7AucFmPZVGWY9gCa43ca/Wr9g74Gt4c0yf44eLoPKnu4ni0qORcMkDcSXAB5BkGUj7ldxwQymnnWZxwmHlVe/TzZWMxypU3Lr0P0J0Dw5p/h3TfD3gLRV22OjW8MajgEpAojQtjgsxBYnuRmvoO0iENuiDsK8y8F6dNczSatdriSc7sf3V6Kv4D9a9WAwMV+DTk5Nye7Pgm23dn/9X9/KoajZpeW7ROM5FX6KAPnrUbJNFnuNP1GJZtLvAyukih4wHGGDKcgowOGB49eCa/Gn9rD9lDUPhfqFz498B2z3Xgy5cM6KS76c7nGx8/MYi3CSHOMhHO7aX/AH61zRLfVLdkdQSRXid/p9xoay6df24vNMnVo3jdQ4CONrKVPDIQcFT2/KvZyXOquCq88NU913/4J1YTFypSutj8Xf2Zf2wNS+Cdk3g/xNZPq/hmSUyosbAXFq7ffMW75WVupQkDPIIyc/oTZft1fs83Vsk8ur3dozjmOSzlLr9fLDL+TGvA/jn+wlpXiYz+LvgbPFZ3EpLyaTM4WB2PUW8rcRn/AKZyHaOcMoAWvzA8XeB/GHgLVpND8Y6Tc6RfRcmK4jaMlT0Zcj5lPZhkHsa+9hl2WZpL20G1J7pOz+a/VHqunQrvmW5/Rd8Kvi34R+Mnh+68T+C2mksLW7ezLTx+WWkSOOQlRknbiQcnBznjvXpm6vgT/gnc+PgjqqnOf7duDyCBzbW3f8K+891fnWb4WNDEzpQ2TPGxEFGbiixur4M/4KJN/wAWP0r/ALD1t/6S3Vfde6vg3/gohk/A/S8DONetj/5LXVdPDn+/UvU0wTtViz8Wt1G6uk8JeCfF/jvVY9E8H6Rc6vfSciK2iaRgo6s20Hao7scAdzX6ffAv9hPTPDbW/i/44zRXc0WHj0mJ90KPnI+0yDiTH/POP5T3ZhlT+w5nnWHwkb1Za9urPpK2NhBXkzwj9k79lDUPijqNt478eWz2vg21YsiMSj6i6HGxCMERBh+8cYzgoh3ZZP2W06yTWp7fT9OhWHS7MKqJGoSMhBhQqjACKBhQOPTgCq9hp9xrgi0+wgFnpkKrGkaKEBRAAqhRgKgAwFHb8q9s0PRLfS7dURQCBX47nWdVcbV556JbLt/wT5vF4uVWV3saWnWaWdusSDGBV+iivGOU/9b9/KKKKACsvUNKtr+MpIoOa1KKAPEtX8EXVlM15pLtC567ejfUdDXn/iPSdL8Q6edE8eaBb6vZc/JNCk8YJ4LBJASh91OfQ19VsqsMMM1lXWjWV0D5kYOaqMnF3i7Madtj5x8Lad4K8KaVFoXhK0tdHsYiStvCggG49WIIBZj3Y5J7mutEitypB/Gu4vPAmnXBJCDn2rDk+GWnsf8AVr+VEpOTvJ3YN33MIyKoyxAHua5TxRYeDPFOkzaB4rtbbV7CfG+3mQTAkdCFAJDDswwR2NekR/DLT1P+rX8q3LPwJp1uQSg49qIycXdPUE7bHiPhrR9J8OaeuieAfD9vo9kMfLDCsCEjozKgy592OT3NegaR4Iur2ZbvVnaZx03dF+gHAr1e10aytQNkY4rVVVUYUYonNyd5O7Bu+5m6fpVtYRhIlAxWpRRUiCiiigD/2Q=="
	imgData    = []byte(imgDataStr)
	imgHash    = "FqKXVdTvIx_mPjOYdjDyUSy_H1jr"
)

func TestBase64Uploader(t *testing.T) {
	var putRet PutRet
	ctx := context.TODO()
	putPolicy := PutPolicy{
		Scope:           testBucket,
		DeleteAfterDays: 7,
	}
	upToken := putPolicy.UploadToken(mac)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	testKey := fmt.Sprintf("testPutFileKey_%d", r.Int())

	extra := Base64PutExtra{}
	extra.MimeType = "image/png"
	extra.Params = map[string]string{
		"x:qiniu":   "Qiniu Cloud",
		"x:storage": "KODO Storage",
	}

	err := base64Uploader.Put(ctx, &putRet, upToken, testKey, imgData, &extra)
	if err != nil {
		t.Fatalf("Base64Uploader#Put() error, %s", err)
	}
	if putRet.Hash != imgHash {
		t.Fatalf("Base64Uploader#Put() error, unmatch hash")
	}
}

func TestBase64UploaderWithoutKey(t *testing.T) {
	var putRet PutRet
	ctx := context.TODO()
	putPolicy := PutPolicy{
		Scope:           testBucket,
		DeleteAfterDays: 7,
	}
	upToken := putPolicy.UploadToken(mac)

	extra := Base64PutExtra{}
	extra.MimeType = "image/png"
	extra.Params = map[string]string{
		"x:qiniu":   "Qiniu Cloud",
		"x:storage": "KODO Storage",
	}

	err := base64Uploader.PutWithoutKey(ctx, &putRet, upToken, imgData, &extra)
	if err != nil {
		t.Fatalf("Base64Uploader#PutWithoutKey() error, %s", err)
	}
	if putRet.Hash != imgHash {
		t.Fatalf("Base64Uploader#PutWithoutKey() error, unmatch hash")
	}
}

func TestBase64UploaderWithBackup(t *testing.T) {
	var putRet PutRet
	ctx := context.TODO()
	putPolicy := PutPolicy{
		Scope:           testBucket,
		DeleteAfterDays: 7,
	}
	upToken := putPolicy.UploadToken(mac)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	testKey := fmt.Sprintf("testPutFileKey_%d", r.Int())

	extra := Base64PutExtra{}
	extra.MimeType = "image/png"
	extra.Params = map[string]string{
		"x:qiniu":   "Qiniu Cloud",
		"x:storage": "KODO Storage",
	}

	region, err := GetRegion(mac.AccessKey, testBucket)
	if err != nil {
		t.Fatal("get region error:", err)
	}

	customizedHost := []string{"mock.qiniu.com"}
	customizedHost = append(customizedHost, region.SrcUpHosts...)
	region.SrcUpHosts = customizedHost
	cfg := Config{}
	cfg.UseCdnDomains = false
	cfg.Region = region
	uploader := NewBase64UploaderEx(&cfg, &clt)
	err = uploader.Put(ctx, &putRet, upToken, testKey, imgData, &extra)
	if err != nil {
		t.Fatalf("Base64Uploader#Put() error, %s", err)
	}
	if putRet.Hash != imgHash {
		t.Fatalf("Base64Uploader#Put() error, unmatch hash")
	}
}