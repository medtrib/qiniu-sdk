// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	"encoding/base64"
	"encoding/json"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	resumableuploadv2completemultipartupload "github.com/sulwan/qiniu-sdk/v7/storagev2/apis/resumable_upload_v2_complete_multipart_upload"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	"strings"
)

type innerResumableUploadV2CompleteMultipartUploadRequest resumableuploadv2completemultipartupload.Request

func (request *innerResumableUploadV2CompleteMultipartUploadRequest) getBucketName(ctx context.Context) (string, error) {
	if request.UpToken != nil {
		if putPolicy, err := request.UpToken.GetPutPolicy(ctx); err != nil {
			return "", err
		} else {
			return putPolicy.GetBucketName()
		}
	}
	return "", nil
}
func (path *innerResumableUploadV2CompleteMultipartUploadRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 5)
	if path.BucketName != "" {
		allSegments = append(allSegments, path.BucketName)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "BucketName"}
	}
	if path.ObjectName != nil {
		allSegments = append(allSegments, "objects", base64.URLEncoding.EncodeToString([]byte(*path.ObjectName)))
	} else {
		allSegments = append(allSegments, "objects", "~")
	}
	if path.UploadId != "" {
		allSegments = append(allSegments, "uploads", path.UploadId)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "UploadId"}
	}
	return allSegments, nil
}
func (j *innerResumableUploadV2CompleteMultipartUploadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal((*resumableuploadv2completemultipartupload.Request)(j))
}
func (j *innerResumableUploadV2CompleteMultipartUploadRequest) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*resumableuploadv2completemultipartupload.Request)(j))
}
func (request *innerResumableUploadV2CompleteMultipartUploadRequest) getAccessKey(ctx context.Context) (string, error) {
	if request.UpToken != nil {
		return request.UpToken.GetAccessKey(ctx)
	}
	return "", nil
}

type ResumableUploadV2CompleteMultipartUploadRequest = resumableuploadv2completemultipartupload.Request
type ResumableUploadV2CompleteMultipartUploadResponse = resumableuploadv2completemultipartupload.Response

// 在将所有数据分片都上传完成后，必须调用 completeMultipartUpload API 来完成整个文件的 Multipart Upload。用户需要提供有效数据的分片列表（包括 PartNumber 和调用 uploadPart API 服务端返回的 Etag）。服务端收到用户提交的分片列表后，会逐一验证每个数据分片的有效性。当所有的数据分片验证通过后，会把这些数据分片组合成一个完整的对象
func (storage *Storage) ResumableUploadV2CompleteMultipartUpload(ctx context.Context, request *ResumableUploadV2CompleteMultipartUploadRequest, options *Options) (*ResumableUploadV2CompleteMultipartUploadResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerResumableUploadV2CompleteMultipartUploadRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceUp}
	if innerRequest.UpToken == nil {
		return nil, errors.MissingRequiredFieldError{Name: "UpToken"}
	}
	pathSegments := make([]string, 0, 6)
	pathSegments = append(pathSegments, "buckets")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	body, err := httpclient.GetJsonRequestBody(&innerRequest)
	if err != nil {
		return nil, err
	}
	bucketName := options.OverwrittenBucketName
	if bucketName == "" {
		var err error
		if bucketName, err = innerRequest.getBucketName(ctx); err != nil {
			return nil, err
		}
	}
	uplogInterceptor, err := uplog.NewRequestUplog("resumableUploadV2CompleteMultipartUpload", bucketName, "", func() (string, error) {
		return innerRequest.UpToken.GetUpToken(ctx)
	})
	if err != nil {
		return nil, err
	}
	req := httpclient.Request{Method: "POST", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, UpToken: innerRequest.UpToken, BufferResponse: true, RequestBody: body, OnRequestProgress: options.OnRequestProgress}
	if options.OverwrittenEndpoints == nil && options.OverwrittenRegion == nil && storage.client.GetRegions() == nil {
		bucketHosts := httpclient.DefaultBucketHosts()
		if bucketName != "" {
			query := storage.client.GetBucketQuery()
			if query == nil {
				if options.OverwrittenBucketHosts != nil {
					if bucketHosts, err = options.OverwrittenBucketHosts.GetEndpoints(ctx); err != nil {
						return nil, err
					}
				}
				queryOptions := region.BucketRegionsQueryOptions{UseInsecureProtocol: storage.client.UseInsecureProtocol(), AccelerateUploading: storage.client.AccelerateUploadingEnabled(), HostFreezeDuration: storage.client.GetHostFreezeDuration(), Resolver: storage.client.GetResolver(), Chooser: storage.client.GetChooser(), BeforeResolve: storage.client.GetBeforeResolveCallback(), AfterResolve: storage.client.GetAfterResolveCallback(), ResolveError: storage.client.GetResolveErrorCallback(), BeforeBackoff: storage.client.GetBeforeBackoffCallback(), AfterBackoff: storage.client.GetAfterBackoffCallback(), BeforeRequest: storage.client.GetBeforeRequestCallback(), AfterResponse: storage.client.GetAfterResponseCallback()}
				if hostRetryConfig := storage.client.GetHostRetryConfig(); hostRetryConfig != nil {
					queryOptions.RetryMax = hostRetryConfig.RetryMax
					queryOptions.Backoff = hostRetryConfig.Backoff
				}
				if query, err = region.NewBucketRegionsQuery(bucketHosts, &queryOptions); err != nil {
					return nil, err
				}
			}
			if query != nil {
				var accessKey string
				var err error
				if accessKey, err = innerRequest.getAccessKey(ctx); err != nil {
					return nil, err
				}
				if accessKey == "" {
					if credentialsProvider := storage.client.GetCredentials(); credentialsProvider != nil {
						if creds, err := credentialsProvider.Get(ctx); err != nil {
							return nil, err
						} else if creds != nil {
							accessKey = creds.AccessKey
						}
					}
				}
				if accessKey != "" {
					req.Region = query.Query(accessKey, bucketName)
				}
			}
		}
	}
	respBody := ResumableUploadV2CompleteMultipartUploadResponse{Body: innerRequest.ResponseBody}
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
