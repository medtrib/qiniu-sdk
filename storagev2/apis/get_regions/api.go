// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 获取所有区域信息
package get_regions

import (
	"encoding/json"
	credentials "github.com/medtrib/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Regions Regions // 区域列表
}

// 区域信息
type Region struct {
	Id          string // 区域 ID
	Description string // 区域描述信息
}
type jsonRegion struct {
	Id          string `json:"id"`          // 区域 ID
	Description string `json:"description"` // 区域描述信息
}

func (j *Region) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonRegion{Id: j.Id, Description: j.Description})
}
func (j *Region) UnmarshalJSON(data []byte) error {
	var nj jsonRegion
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.Description = nj.Description
	return nil
}
func (j *Region) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.Description == "" {
		return errors.MissingRequiredFieldError{Name: "Description"}
	}
	return nil
}

// 区域列表
type Regions = []Region

// 所有区域信息
type RegionsInfo = Response
type jsonResponse struct {
	Regions Regions `json:"regions"` // 区域列表
}

func (j *Response) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonResponse{Regions: j.Regions})
}
func (j *Response) UnmarshalJSON(data []byte) error {
	var nj jsonResponse
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Regions = nj.Regions
	return nil
}
func (j *Response) validate() error {
	if len(j.Regions) == 0 {
		return errors.MissingRequiredFieldError{Name: "Regions"}
	}
	for _, value := range j.Regions {
		if err := value.validate(); err != nil {
			return err
		}
	}
	return nil
}
