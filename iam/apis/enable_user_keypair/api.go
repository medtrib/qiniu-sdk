// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 启用 IAM 子账号密钥
package enable_user_keypair

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Alias       string                          // 子账号别名
	AccessKey   string                          // IAM 子账号 Access Key
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
