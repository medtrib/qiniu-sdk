// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 删除用户分组
package delete_group

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Alias       string                          // 用户分组别名
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
