// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置存储空间的访问权限
package set_bucket_private

import credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Bucket      string                          // 空间名称
	IsPrivate   int64                           // `0`: 公开，`1`: 私有
}

// 获取 API 所用的响应
type Response struct{}