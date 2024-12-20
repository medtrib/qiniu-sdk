// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 将上传好的所有数据块按指定顺序合并成一个资源文件
package resumable_upload_v1_make_file

import (
	"encoding/json"
	io "github.com/medtrib/qiniu-sdk/v7/internal/io"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
)

// 调用 API 所用的请求
type Request struct {
	Size         int64             // 对象大小
	ObjectName   *string           // 对象名称
	FileName     string            // 文件名称，若未指定，则魔法变量中无法使用fname，ext，fprefix
	MimeType     string            // 文件 MIME 类型，若未指定，则根据文件内容自动检测 MIME 类型
	CustomData   map[string]string // 自定义元数据（需要以 `x-qn-meta-` 作为前缀）或自定义变量（需要以 `x:` 作为前缀）
	UpToken      uptoken.Provider  // 上传凭证，如果为空，则使用 HTTPClientOptions 中的 UpToken
	Body         io.ReadSeekCloser // 请求体
	ResponseBody interface{}       // 响应体，如果为空，则 Response.Body 的类型由 encoding/json 库决定
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
