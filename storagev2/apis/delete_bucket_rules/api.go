// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 删除空间规则
package delete_bucket_rules

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Bucket      string                          // 空间名称
	Name        string                          // 要删除的规则名称
}

// 获取 API 所用的响应
type Response struct{}
