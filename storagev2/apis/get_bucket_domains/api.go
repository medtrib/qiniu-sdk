// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 获取存储空间的域名列表
package get_bucket_domains

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
)

// 调用 API 所用的请求
type Request struct {
	BucketName  string                          // 要获取域名列表的目标空间名称
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Domains Domains // 存储空间的域名列表
}

// 存储空间的域名列表
type Domains []string

func (j *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Domains)
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var array Domains
	if err := json.Unmarshal(data, &array); err != nil {
		return err
	}
	j.Domains = array
	return nil
}
