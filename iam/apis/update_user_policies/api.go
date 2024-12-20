// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 为子账号重新分配授权策略
package update_user_policies

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
)

// 调用 API 所用的请求
type Request struct {
	Alias         string                          // 子账号别名
	Credentials   credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	PolicyAliases PolicyAliases                   // 授权策略别名集合
}

// 重新分配给用户的授权策略别名集合
type PolicyAliases = []string

// 为子账号重新分配授权策略参数
type UpdatedIamUserPoliciesParam = Request
type jsonRequest struct {
	PolicyAliases PolicyAliases `json:"policy_aliases,omitempty"` // 授权策略别名集合
}

func (j *Request) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRequest{PolicyAliases: j.PolicyAliases})
}
func (j *Request) UnmarshalJSON(data []byte) error {
	var nj jsonRequest
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.PolicyAliases = nj.PolicyAliases
	return nil
}
func (j *Request) validate() error {
	return nil
}

// 获取 API 所用的响应
type Response struct{}
