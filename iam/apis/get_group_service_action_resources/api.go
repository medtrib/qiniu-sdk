// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 列举用户分组指定服务操作下的可访问资源
package get_group_service_action_resources

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	GroupAlias  string                          // 用户分组别名
	Service     string                          // 资源操作关联的服务
	ActionAlias string                          // 资源操作别名
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Data GetGroupServiceActionResources // 用户分组指定服务操作下的可访问资源列表信息
}

// 可用资源列表
type AllowedResources = []string

// 禁用资源列表
type DeniedResources = []string

// 用户分组指定服务操作下的可访问资源列表信息
type Data struct {
	AllowedResources AllowedResources // 可用资源
	DeniedResources  DeniedResources  // 禁用资源
}

// 返回的用户分组指定服务操作下的可访问资源列表信息
type GetGroupServiceActionResources = Data
type jsonData struct {
	AllowedResources AllowedResources `json:"allow"` // 可用资源
	DeniedResources  DeniedResources  `json:"deny"`  // 禁用资源
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{AllowedResources: j.AllowedResources, DeniedResources: j.DeniedResources})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.AllowedResources = nj.AllowedResources
	j.DeniedResources = nj.DeniedResources
	return nil
}
func (j *Data) validate() error {
	if len(j.AllowedResources) == 0 {
		return errors.MissingRequiredFieldError{Name: "AllowedResources"}
	}
	if len(j.DeniedResources) == 0 {
		return errors.MissingRequiredFieldError{Name: "DeniedResources"}
	}
	return nil
}

// 返回的用户分组指定服务操作下的可访问资源列表响应
type GetGroupServiceActionResourcesResp = Response
type jsonResponse struct {
	Data GetGroupServiceActionResources `json:"data"` // 用户分组指定服务操作下的可访问资源列表信息
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
