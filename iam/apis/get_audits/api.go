// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

// 查询审计日志列表
package get_audits

import (
	"encoding/json"
	credentials "github.com/sulwan/qiniu-sdk/v7/storagev2/credentials"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
)

// 调用 API 所用的请求
type Request struct {
	Iuid        int64                           // IAM 子账号 UID
	Service     string                          // 操作对应的服务别名
	Action      string                          // 操作别名
	Resource    string                          // 操作别名
	StartTime   string                          // 操作开始时间
	EndTime     string                          // 操作截止时间
	Marker      string                          // 下页标记
	Limit       int64                           // 分页大小，默认 20，最大 2000
	Credentials credentials.CredentialsProvider // 鉴权参数，用于生成鉴权凭证，如果为空，则使用 HTTPClientOptions 中的 CredentialsProvider
}

// 获取 API 所用的响应
type Response struct {
	Data GetAuditLogsData // 审计日志信息
}

// 返回的审计日志
type GetAuditLog struct {
	Id           string // 记录 ID
	RootUid      int64  // 根用户 uid
	Iuid         int64  // 子账号 uid
	Service      string // 接口操作对应的服务
	Action       string // 接口操作别名
	CreatedAt    string // 日志创建时间
	EventTime    string // 请求发生时间
	DurationMs   int64  // 请求持续时间，毫秒
	SourceIp     string // 源 IP
	UserEvent    string // 用户代理
	ErrorCode    int64  // 错误码
	ErrorMessage string // 错误消息
}
type jsonGetAuditLog struct {
	Id           string `json:"id"`            // 记录 ID
	RootUid      int64  `json:"root_uid"`      // 根用户 uid
	Iuid         int64  `json:"iuid"`          // 子账号 uid
	Service      string `json:"service"`       // 接口操作对应的服务
	Action       string `json:"action"`        // 接口操作别名
	CreatedAt    string `json:"created_at"`    // 日志创建时间
	EventTime    string `json:"event_time"`    // 请求发生时间
	DurationMs   int64  `json:"duration_ms"`   // 请求持续时间，毫秒
	SourceIp     string `json:"source_ip"`     // 源 IP
	UserEvent    string `json:"user_event"`    // 用户代理
	ErrorCode    int64  `json:"error_code"`    // 错误码
	ErrorMessage string `json:"error_message"` // 错误消息
}

func (j *GetAuditLog) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonGetAuditLog{Id: j.Id, RootUid: j.RootUid, Iuid: j.Iuid, Service: j.Service, Action: j.Action, CreatedAt: j.CreatedAt, EventTime: j.EventTime, DurationMs: j.DurationMs, SourceIp: j.SourceIp, UserEvent: j.UserEvent, ErrorCode: j.ErrorCode, ErrorMessage: j.ErrorMessage})
}
func (j *GetAuditLog) UnmarshalJSON(data []byte) error {
	var nj jsonGetAuditLog
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Id = nj.Id
	j.RootUid = nj.RootUid
	j.Iuid = nj.Iuid
	j.Service = nj.Service
	j.Action = nj.Action
	j.CreatedAt = nj.CreatedAt
	j.EventTime = nj.EventTime
	j.DurationMs = nj.DurationMs
	j.SourceIp = nj.SourceIp
	j.UserEvent = nj.UserEvent
	j.ErrorCode = nj.ErrorCode
	j.ErrorMessage = nj.ErrorMessage
	return nil
}
func (j *GetAuditLog) validate() error {
	if j.Id == "" {
		return errors.MissingRequiredFieldError{Name: "Id"}
	}
	if j.RootUid == 0 {
		return errors.MissingRequiredFieldError{Name: "RootUid"}
	}
	if j.Iuid == 0 {
		return errors.MissingRequiredFieldError{Name: "Iuid"}
	}
	if j.Service == "" {
		return errors.MissingRequiredFieldError{Name: "Service"}
	}
	if j.Action == "" {
		return errors.MissingRequiredFieldError{Name: "Action"}
	}
	if j.CreatedAt == "" {
		return errors.MissingRequiredFieldError{Name: "CreatedAt"}
	}
	if j.EventTime == "" {
		return errors.MissingRequiredFieldError{Name: "EventTime"}
	}
	if j.DurationMs == 0 {
		return errors.MissingRequiredFieldError{Name: "DurationMs"}
	}
	if j.SourceIp == "" {
		return errors.MissingRequiredFieldError{Name: "SourceIp"}
	}
	if j.UserEvent == "" {
		return errors.MissingRequiredFieldError{Name: "UserEvent"}
	}
	if j.ErrorCode == 0 {
		return errors.MissingRequiredFieldError{Name: "ErrorCode"}
	}
	if j.ErrorMessage == "" {
		return errors.MissingRequiredFieldError{Name: "ErrorMessage"}
	}
	return nil
}

// 返回的审计日志列表
type GetAuditLogs = []GetAuditLog

// 审计日志信息
type Data struct {
	Marker string       // 下页标记
	List   GetAuditLogs // 审计日志列表
}

// 返回的审计日志列表信息
type GetAuditLogsData = Data
type jsonData struct {
	Marker string       `json:"marker"` // 下页标记
	List   GetAuditLogs `json:"list"`   // 审计日志列表
}

func (j *Data) MarshalJSON() ([]byte, error) {
	if err := j.validate(); err != nil {
		return nil, err
	}
	return json.Marshal(&jsonData{Marker: j.Marker, List: j.List})
}
func (j *Data) UnmarshalJSON(data []byte) error {
	var nj jsonData
	if err := json.Unmarshal(data, &nj); err != nil {
		return err
	}
	j.Marker = nj.Marker
	j.List = nj.List
	return nil
}
func (j *Data) validate() error {
	if j.Marker == "" {
		return errors.MissingRequiredFieldError{Name: "Marker"}
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

// 返回的审计日志列表响应
type GetAuditLogsResp = Response
type jsonResponse struct {
	Data GetAuditLogsData `json:"data"` // 审计日志信息
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