// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 在一次 HTTP 会话中上传单一的一个文件
package post_object

import (
	"encoding/json"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
)

// 调用 API 所用的请求
type Request struct {
	ObjectName   *string
	UploadToken  uptoken.Provider
	Crc32        int64
	File         httpclient.MultipartFormBinaryData
	CustomData   map[string]string
	ResponseBody interface{} // 响应体，如果为空，则 Response.Body 的类型由 encoding/json 库决定
}

// 获取 API 所用的响应
type Response struct {
	Body interface{}
}

func (j *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Body)
}
func (j *Response) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.Body)
}
