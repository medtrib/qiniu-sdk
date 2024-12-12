// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 查询指定授权策略详情
package get_policy

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Alias       string                          // 授权策略别名
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Data GetPolicyData // 授权策略信息
}

// 授权策略规则的操作集合
type Actions = []string

// 授权策略规则的资源集合
type Resources = []string

// 授权策略规则
type Statement struct {
	Actions   Actions   // 授权策略规则的操作集合
	Resources Resources // 授权策略规则的资源集合
	Effect    string    // 授权策略规则的生效类型，允许访问或拒绝访问
}
type jsonStatement struct {
	Actions   Actions   `json:"action"`   // 授权策略规则的操作集合
	Resources Resources `json:"resource"` // 授权策略规则的资源集合
	Effect    string    `json:"effect"`   // 授权策略规则的生效类型，允许访问或拒绝访问
}

func (j *Statement) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonStatement{Actions: j.Actions, Resources: j.Resources, Effect: j.Effect})
}
func (j *Statement) UnmarshalJSON(data []byte) error {
	var nj jsonStatement
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Actions = nj.Actions
	j.Resources = nj.Resources
	j.Effect = nj.Effect
	return nil
}
func (j *Statement) validate() error {
	if len(j.Actions) == 0 {
		return errors.MissingRequiredFieldError{Name: "Actions"}
	}
	if len(j.Resources) == 0 {
		return errors.MissingRequiredFieldError{Name: "Resources"}
	}
	if j.Effect == "" {
		return errors.MissingRequiredFieldError{Name: "Effect"}
	}
	return nil
}

// 授权策略规则集合
type Statements = []Statement

// 授权策略信息
type Data struct {
	Id          string     // 记录 ID
	RootUid     int64      // 根用户 uid
	Alias       string     // 授权策略别名
	Description string     // 授权策略描述
	Enabled     bool       // 授权策略是否启用
	CreatedAt   string     // 授权策略创建时间
	UpdatedAt   string     // 授权策略上次更新时间
	Statement   Statements // 授权策略规则集合
}

// 返回的授权策略信息
type GetPolicyData = Data
type jsonData struct {
	Id          string     `json:"id"`          // 记录 ID
	RootUid     int64      `json:"root_uid"`    // 根用户 uid
	Alias       string     `json:"alias"`       // 授权策略别名
	Description string     `json:"description"` // 授权策略描述
	Enabled     bool       `json:"enabled"`     // 授权策略是否启用
	CreatedAt   string     `json:"created_at"`  // 授权策略创建时间
	UpdatedAt   string     `json:"updated_at"`  // 授权策略上次更新时间
	Statement   Statements `json:"statement"`   // 授权策略规则集合
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{Id: j.Id, RootUid: j.RootUid, Alias: j.Alias, Description: j.Description, Enabled: j.Enabled, CreatedAt: j.CreatedAt, UpdatedAt: j.UpdatedAt, Statement: j.Statement})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.RootUid = nj.RootUid
	j.Alias = nj.Alias
	j.Description = nj.Description
	j.Enabled = nj.Enabled
	j.CreatedAt = nj.CreatedAt
	j.UpdatedAt = nj.UpdatedAt
	j.Statement = nj.Statement
	return nil
}
func (j *Data) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.RootUid == 0 {
		return errors.MissingRequiredFieldError{Name: "RootUid"}
	}
	if j.Alias == "" {
		return errors.MissingRequiredFieldError{Name: "Alias"}
	}
	if j.Description == "" {
		return errors.MissingRequiredFieldError{Name: "Description"}
	}
	if j.CreatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "CreatedAt"}
	}
	if j.UpdatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "UpdatedAt"}
	}
	if len(j.Statement) == 0 {
		return errors.MissingRequiredFieldError{Name: "Statement"}
	}
	for _, value := range j.Statement {
		if err := value.validate(); err != nil {
			return err
		}
	}
	return nil
}

// 返回的授权策略响应
type GetPolicyResp = Response
type jsonResponse struct {
	Data GetPolicyData `json:"data"` // 授权策略信息
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{Data: j.Data})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Data = nj.Data
	return nil
}
func (j *Response) validate() error {
	if err := j.Data.validate(); err != nil {
		return err
	}
	return nil
}
