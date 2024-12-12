// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 查询 IAM 的操作
package get_actions

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Service     string                          // 操作对应的服务别名
	Page        int64                           // 分页页号，从 1 开始，默认 1
	PageSize    int64                           // 分页大小，默认 20，最大 2000
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Data GetActionsData // 接口操作信息
}

// 返回的接口操作
type GetAction struct {
	Id        string // 记录 ID
	Name      string // 接口操作名称
	Alias     string // 接口操作别名
	Service   string // 接口操作对应的服务
	Scope     int64  // 接口操作权限粒度，0: 操作级，不限制资源，1: 资源级，只能访问特定资源
	Enabled   bool   // 接口操作是否启用
	CreatedAt string // 接口操作创建时间
	UpdatedAt string // 接口操作上次更新时间
}
type jsonGetAction struct {
	Id        string `json:"id"`         // 记录 ID
	Name      string `json:"name"`       // 接口操作名称
	Alias     string `json:"alias"`      // 接口操作别名
	Service   string `json:"service"`    // 接口操作对应的服务
	Scope     int64  `json:"scope"`      // 接口操作权限粒度，0: 操作级，不限制资源，1: 资源级，只能访问特定资源
	Enabled   bool   `json:"enabled"`    // 接口操作是否启用
	CreatedAt string `json:"created_at"` // 接口操作创建时间
	UpdatedAt string `json:"updated_at"` // 接口操作上次更新时间
}

func (j *GetAction) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonGetAction{Id: j.Id, Name: j.Name, Alias: j.Alias, Service: j.Service, Scope: j.Scope, Enabled: j.Enabled, CreatedAt: j.CreatedAt, UpdatedAt: j.UpdatedAt})
}
func (j *GetAction) UnmarshalJSON(data []byte) error {
	var nj jsonGetAction
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.Name = nj.Name
	j.Alias = nj.Alias
	j.Service = nj.Service
	j.Scope = nj.Scope
	j.Enabled = nj.Enabled
	j.CreatedAt = nj.CreatedAt
	j.UpdatedAt = nj.UpdatedAt
	return nil
}
func (j *GetAction) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.Name == "" {
		return errors.MissingRequiredFieldError{Name: "Name"}
	}
	if j.Alias == "" {
		return errors.MissingRequiredFieldError{Name: "Alias"}
	}
	if j.Service == "" {
		return errors.MissingRequiredFieldError{Name: "Service"}
	}
	if j.Scope == 0 {
		return errors.MissingRequiredFieldError{Name: "Scope"}
	}
	if j.CreatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "CreatedAt"}
	}
	if j.UpdatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "UpdatedAt"}
	}
	return nil
}

// 返回的接口操作列表
type GetActions = []GetAction

// 接口操作信息
type Data struct {
	Count int64      // 接口操作数量
	List  GetActions // 接口操作列表
}

// 返回的接口操作列表信息
type GetActionsData = Data
type jsonData struct {
	Count int64      `json:"count"` // 接口操作数量
	List  GetActions `json:"list"`  // 接口操作列表
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{Count: j.Count, List: j.List})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Count = nj.Count
	j.List = nj.List
	return nil
}
func (j *Data) validate() error {
	if j.Count == 0 {
		return errors.MissingRequiredFieldError{Name: "Count"}
	}
	if len(j.List) == 0 {
		return errors.MissingRequiredFieldError{Name: "List"}
	}
	for _, value := range j.List {
		if err := value.validate(); err != nil {
			return err
		}
	}
	return nil
}

// 返回的接口操作列表响应
type GetActionsResp = Response
type jsonResponse struct {
	Data GetActionsData `json:"data"` // 接口操作信息
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