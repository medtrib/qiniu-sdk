// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 设置存储空间的防盗链模式
package set_bucket_refer_anti_leech

import credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"

// 调用 API 所用的请求
type Request struct {
	Bucket            string                          // 存储空间名称
	Mode              int64                           // 设置防盗链模式，0：表示关闭 Referer(使用此选项将会忽略以下参数并将恢复默认值); 1：表示设置 Referer 白名单; 2：表示设置 Referer 黑名单
	Pattern           string                          // 规则字符串，当前允许格式分为三种：一种为空主机头域名，比如 `foo.com`; 一种是泛域名，比如 `*.bar.com`; 一种是完全通配符，即一个 `*`; 多个规则之间用`;`隔开，比如：`foo.com;*.bar.com;sub.foo.com;*.sub.bar.com`
	AllowEmptyReferer int64                           // 0：表示不允许空 Refer 访问; 1：表示允许空 Refer 访问
	SourceEnabled     int64                           // 源站是否支持，默认为 0 只给 CDN 配置, 设置为 1 表示开启源站防盗链
	Credentials       credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct{}