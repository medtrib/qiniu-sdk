// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 查询子账号所属用户分组列表
package get_user_groups

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Alias       string                          // 子账号别名
	Page        int64                           // 分页页号，从 1 开始，默认 1
	PageSize    int64                           // 分页大小，默认 20，最大 2000
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Data GetIamUserGroupsData // IAM 子账号分组信息
}

// 返回的 IAM 子账号分组
type IamUserGroup struct {
	Id          string // 记录 ID
	RootUid     int64  // 根用户 uid
	Alias       string // 用户分组别名
	Description string // 用户分组描述
	Enabled     bool   // 用户分组是否启用
	CreatedAt   string // 用户分组创建时间
	UpdatedAt   string // 用户分组上次更新时间
}
type jsonIamUserGroup struct {
	Id          string `json:"id"`          // 记录 ID
	RootUid     int64  `json:"root_uid"`    // 根用户 uid
	Alias       string `json:"alias"`       // 用户分组别名
	Description string `json:"description"` // 用户分组描述
	Enabled     bool   `json:"enabled"`     // 用户分组是否启用
	CreatedAt   string `json:"created_at"`  // 用户分组创建时间
	UpdatedAt   string `json:"updated_at"`  // 用户分组上次更新时间
}

func (j *IamUserGroup) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonIamUserGroup{Id: j.Id, RootUid: j.RootUid, Alias: j.Alias, Description: j.Description, Enabled: j.Enabled, CreatedAt: j.CreatedAt, UpdatedAt: j.UpdatedAt})
}
func (j *IamUserGroup) UnmarshalJSON(data []byte) error {
	var nj jsonIamUserGroup
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
	return nil
}
func (j *IamUserGroup) validate() error {
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
	return nil
}

// 返回的 IAM 子账号分组列表
type GetIamUserGroups = []IamUserGroup

// IAM 子账号分组信息
type Data struct {
	Count int64            // IAM 子账号分组数量
	List  GetIamUserGroups // IAM 子账号分组列表
}

// 返回的 IAM 子账号分组列表信息
type GetIamUserGroupsData = Data
type jsonData struct {
	Count int64            `json:"count"` // IAM 子账号分组数量
	List  GetIamUserGroups `json:"list"`  // IAM 子账号分组列表
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

// 返回的 IAM 子账号分组列表响应
type GetIamUserGroupsResp = Response
type jsonResponse struct {
	Data GetIamUserGroupsData `json:"data"` // IAM 子账号分组信息
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
