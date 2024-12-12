// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置存储空间的标签列表，包括新增和修改
package set_bucket_taggings

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Bucket      string                          // 空间名称
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Tags        Tags                            // 标签列表
}

// 标签键值对
type TagInfo struct {
	Key   string // 标签名称，最大 64 Byte，不能为空且大小写敏感，不能以 kodo 为前缀(预留), 不支持中文字符，可使用的字符有：字母，数字，空格，+ - = . _ : / @
	Value string // 标签值，最大 128 Byte，不能为空且大小写敏感，不支持中文字符，可使用的字符有：字母，数字，空格，+ - = . _ : / @
}
type jsonTagInfo struct {
	Key   string `json:"Key"`   // 标签名称，最大 64 Byte，不能为空且大小写敏感，不能以 kodo 为前缀(预留), 不支持中文字符，可使用的字符有：字母，数字，空格，+ - = . _ : / @
	Value string `json:"Value"` // 标签值，最大 128 Byte，不能为空且大小写敏感，不支持中文字符，可使用的字符有：字母，数字，空格，+ - = . _ : / @
}

func (j *TagInfo) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonTagInfo{Key: j.Key, Value: j.Value})
}
func (j *TagInfo) UnmarshalJSON(data []byte) error {
	var nj jsonTagInfo
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Key = nj.Key
	j.Value = nj.Value
	return nil
}
func (j *TagInfo) validate() error {
	if j.Key == "" {
		return errors.MissingRequiredFieldError{Name: "Key"}
	}
	if j.Value == "" {
		return errors.MissingRequiredFieldError{Name: "Value"}
	}
	return nil
}

// 标签列表
type Tags = []TagInfo

// 存储空间标签信息
type TagsInfo = Request
type jsonRequest struct {
	Tags Tags `json:"Tags"` // 标签列表
}

func (j *Request) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRequest{Tags: j.Tags})
}
func (j *Request) UnmarshalJSON(data []byte) error {
	var nj jsonRequest
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Tags = nj.Tags
	return nil
}
func (j *Request) validate() error {
	if len(j.Tags) == 0 {
		return errors.MissingRequiredFieldError{Name: "Tags"}
	}
	for _, value := range j.Tags {
		if err := value.validate(); err != nil {
			return err
		}
	}
	return nil
}

// 获取 API 所用的响应
type Response struct{}
