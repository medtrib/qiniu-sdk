// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 创建目录分享
package create_share

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Credentials     credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Bucket          string                          // 分享的空间名称
	Prefix          string                          // 分享的对象名称前缀
	DurationSeconds int64                           // 分享过期时间，单位为秒，最小 900 秒，最长 7200 秒
	ExtractCode     string                          // 提取码，必须给出 6 个字符，仅大小写字母和数字
	Permission      string                          // 权限，目前仅支持 READONLY，未来会支持 READWRITE
}

// 新创建的分享参数
type NewShareParams = Request
type jsonRequest struct {
	Bucket          string `json:"bucket"`           // 分享的空间名称
	Prefix          string `json:"prefix"`           // 分享的对象名称前缀
	DurationSeconds int64  `json:"duration_seconds"` // 分享过期时间，单位为秒，最小 900 秒，最长 7200 秒
	ExtractCode     string `json:"extract_code"`     // 提取码，必须给出 6 个字符，仅大小写字母和数字
	Permission      string `json:"permission"`       // 权限，目前仅支持 READONLY，未来会支持 READWRITE
}

func (j *Request) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRequest{Bucket: j.Bucket, Prefix: j.Prefix, DurationSeconds: j.DurationSeconds, ExtractCode: j.ExtractCode, Permission: j.Permission})
}
func (j *Request) UnmarshalJSON(data []byte) error {
	var nj jsonRequest
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Bucket = nj.Bucket
	j.Prefix = nj.Prefix
	j.DurationSeconds = nj.DurationSeconds
	j.ExtractCode = nj.ExtractCode
	j.Permission = nj.Permission
	return nil
}
func (j *Request) validate() error {
	if j.Bucket == "" {
		return errors.MissingRequiredFieldError{Name: "Bucket"}
	}
	if j.Prefix == "" {
		return errors.MissingRequiredFieldError{Name: "Prefix"}
	}
	if j.DurationSeconds == 0 {
		return errors.MissingRequiredFieldError{Name: "DurationSeconds"}
	}
	if j.ExtractCode == "" {
		return errors.MissingRequiredFieldError{Name: "ExtractCode"}
	}
	if j.Permission == "" {
		return errors.MissingRequiredFieldError{Name: "Permission"}
	}
	return nil
}

// 获取 API 所用的响应
type Response struct {
	Id      string // 分享 ID
	Token   string // 分享 Token
	Expires string // 分享过期时间，遵循 ISO8601 风格，使用 UTC 0 时区时间
}

// 返回的分享信息
type NewShareInfo = Response
type jsonResponse struct {
	Id      string `json:"id"`      // 分享 ID
	Token   string `json:"token"`   // 分享 Token
	Expires string `json:"expires"` // 分享过期时间，遵循 ISO8601 风格，使用 UTC 0 时区时间
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{Id: j.Id, Token: j.Token, Expires: j.Expires})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.Token = nj.Token
	j.Expires = nj.Expires
	return nil
}
func (j *Response) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.Token == "" {
		return errors.MissingRequiredFieldError{Name: "Token"}
	}
	if j.Expires == "" {
		return errors.MissingRequiredFieldError{Name: "Expires"}
	}
	return nil
}
