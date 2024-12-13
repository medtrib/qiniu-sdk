// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 检查目录分享
package check_share

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	ShareId     string                          // 分享 ID
	Token       string                          // 分享 Token
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
