// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 对于设置了镜像存储的空间，从镜像源站抓取指定名称的对象并存储到该空间中，如果该空间中已存在该名称的对象，则会将镜像源站的对象覆盖空间中相同名称的对象
package prefetch_object

import credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Entry       string                          // 指定目标对象空间与目标对象名称，格式为 <目标对象空间>:<目标对象名称>
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}