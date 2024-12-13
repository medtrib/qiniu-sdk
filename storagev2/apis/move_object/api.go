// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 将源空间的指定对象移动到目标空间，或在同一空间内对对象重命名
package move_object

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	SrcEntry    string                          // 指定源对象空间与源对象名称，格式为 <源对象空间>:<源对象名称>
	DestEntry   string                          // 指定目标对象空间与目标对象名称，格式为 <目标对象空间>:<目标对象名称>
	IsForce     bool                            // 如果目标对象名已被占用，则返回错误码 614，且不做任何覆盖操作；如果指定为 true，会强制覆盖目标对象
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
