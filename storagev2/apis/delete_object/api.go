// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 删除指定对象
package delete_object

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Entry       string                          // 指定目标对象空间与目标对象名称，格式为 <目标对象空间>:<目标对象名称>
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
