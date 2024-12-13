package http_client

import compatible_io "github.com/medtrib/qiniu-sdk/v7/internal/io"

type MultipartFormBinaryData struct {
	Data        compatible_io.ReadSeekCloser
	Name        string
	ContentType string
}
