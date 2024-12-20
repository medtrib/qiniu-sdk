// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置用户存储空间配额限制
package set_bucket_quota

import credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Bucket      string                          // 指定存储空间
	Size        int64                           // 空间存储量配额，参数传入 0 或不传表示不更改当前配置，传入 -1 表示取消限额
	Count       int64                           // 空间文件数配额，参数传入 0 或不传表示不更改当前配置，传入 -1 表示取消限额
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}
