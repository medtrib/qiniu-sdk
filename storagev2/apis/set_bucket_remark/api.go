// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置空间备注
package set_bucket_remark

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Bucket      string                          // 空间名称
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Remark      string                          // 空间备注信息, 字符长度不能超过 100, 允许为空
}

// 空间备注信息
type BucketRemark = Request
type jsonRequest struct {
	Remark string `json:"remark"` // 空间备注信息, 字符长度不能超过 100, 允许为空
}

func (j *Request) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRequest{Remark: j.Remark})
}
func (j *Request) UnmarshalJSON(data []byte) error {
	var nj jsonRequest
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Remark = nj.Remark
	return nil
}
func (j *Request) validate() error {
	if j.Remark == "" {
		return errors.MissingRequiredFieldError{Name: "Remark"}
	}
	return nil
}

// 获取 API 所用的响应
type Response struct{}
