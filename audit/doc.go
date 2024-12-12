// audit 包提供了账号审计等功能。
package audit

//go:generate go run ../internal/api-generator -- --api-specs=../api-specs/audit --output=apis/ --struct-name=Audit --api-package=github.com/sulwan/qiniu-sdk/v7/audit/apis
//go:generate go build ./apis/...
