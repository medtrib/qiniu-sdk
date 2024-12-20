// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 创建 IAM 子账号
package create_user

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
	Alias       string                          // 子账号别名
	Password    string                          // 子账号密码
}

// 创建 IAM 子账号参数
type CreateIamUserParam = Request
type jsonRequest struct {
	Alias    string `json:"alias"`    // 子账号别名
	Password string `json:"password"` // 子账号密码
}

func (j *Request) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRequest{Alias: j.Alias, Password: j.Password})
}
func (j *Request) UnmarshalJSON(data []byte) error {
	var nj jsonRequest
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Alias = nj.Alias
	j.Password = nj.Password
	return nil
}
func (j *Request) validate() error {
	if j.Alias == "" {
		return errors.MissingRequiredFieldError{Name: "Alias"}
	}
	if j.Password == "" {
		return errors.MissingRequiredFieldError{Name: "Password"}
	}
	return nil
}

// 获取 API 所用的响应
type Response struct {
	Data CreatedIamUserData // IAM 子账号信息
}

// IAM 子账号信息
type Data struct {
	Id            string // 记录 ID
	RootUid       int64  // 根用户 uid
	Iuid          int64  // 子账号 uid
	Alias         string // 子账号别名
	CreatedAt     string // 子账号创建时间
	UpdatedAt     string // 子账号上次更新时间
	LastLoginTime string // 子账号上次更新时间
	Enabled       bool   // 子账号是否启用
}

// 返回的 IAM 子账号信息
type CreatedIamUserData = Data
type jsonData struct {
	Id            string `json:"id"`              // 记录 ID
	RootUid       int64  `json:"root_uid"`        // 根用户 uid
	Iuid          int64  `json:"iuid"`            // 子账号 uid
	Alias         string `json:"alias"`           // 子账号别名
	CreatedAt     string `json:"created_at"`      // 子账号创建时间
	UpdatedAt     string `json:"updated_at"`      // 子账号上次更新时间
	LastLoginTime string `json:"last_login_time"` // 子账号上次更新时间
	Enabled       bool   `json:"enabled"`         // 子账号是否启用
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{Id: j.Id, RootUid: j.RootUid, Iuid: j.Iuid, Alias: j.Alias, CreatedAt: j.CreatedAt, UpdatedAt: j.UpdatedAt, LastLoginTime: j.LastLoginTime, Enabled: j.Enabled})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.RootUid = nj.RootUid
	j.Iuid = nj.Iuid
	j.Alias = nj.Alias
	j.CreatedAt = nj.CreatedAt
	j.UpdatedAt = nj.UpdatedAt
	j.LastLoginTime = nj.LastLoginTime
	j.Enabled = nj.Enabled
	return nil
}
func (j *Data) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.RootUid == 0 {
		return errors.MissingRequiredFieldError{Name: "RootUid"}
	}
	if j.Iuid == 0 {
		return errors.MissingRequiredFieldError{Name: "Iuid"}
	}
	if j.Alias == "" {
		return errors.MissingRequiredFieldError{Name: "Alias"}
	}
	if j.CreatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "CreatedAt"}
	}
	if j.UpdatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "UpdatedAt"}
	}
	if j.LastLoginTime == "" {
		return errors.MissingRequiredFieldError{Name: "LastLoginTime"}
	}
	return nil
}

// 返回的 IAM 子账号响应
type CreatedIamUserResp = Response
type jsonResponse struct {
	Data CreatedIamUserData `json:"data"` // IAM 子账号信息
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
