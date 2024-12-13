// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	"encoding/base64"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	modifyobjectstatus "github.com/medtrib/qiniu-sdk/v7/storagev2/apis/modify_object_status"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"strconv"
	"strings"
	"time"
)

type innerModifyObjectStatusRequest modifyobjectstatus.Request

func (pp *innerModifyObjectStatusRequest) getBucketName(ctx context.Context) (string, error) {
	return strings.SplitN(pp.Entry, ":", 2)[0], nil
}
func (pp *innerModifyObjectStatusRequest) getObjectName() string {
	parts := strings.SplitN(pp.Entry, ":", 2)
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
func (path *innerModifyObjectStatusRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 3)
	if path.Entry != "" {
		allSegments = append(allSegments, base64.URLEncoding.EncodeToString([]byte(path.Entry)))
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Entry"}
	}
	allSegments = append(allSegments, "status", strconv.FormatInt(path.Status, 10))
	return allSegments, nil
}
func (request *innerModifyObjectStatusRequest) getAccessKey(ctx context.Context) (string, error) {
	if request.Credentials != nil {
		if credentials, err := request.Credentials.Get(ctx); err != nil {
			return "", err
		} else {
			return credentials.AccessKey, nil
		}
	}
	return "", nil
}

type ModifyObjectStatusRequest = modifyobjectstatus.Request
type ModifyObjectStatusResponse = modifyobjectstatus.Response

// 修改文件的存储状态，即禁用状态和启用状态间的的互相转换
func (storage *Storage) ModifyObjectStatus(ctx context.Context, request *ModifyObjectStatusRequest, options *Options) (*ModifyObjectStatusResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerModifyObjectStatusRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceRs}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 4)
	pathSegments = append(pathSegments, "chstatus")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	bucketName := options.OverwrittenBucketName
	if bucketName == "" {
		var err error
		if bucketName, err = innerRequest.getBucketName(ctx); err != nil {
			return nil, err
		}
	}
	objectName := innerRequest.getObjectName()
	uplogInterceptor, err := uplog.NewRequestUplog("modifyObjectStatus", bucketName, objectName, func() (string, error) {
		credentials := innerRequest.Credentials
		if credentials == nil {
			credentials = storage.client.GetCredentials()
		}
		putPolicy, err := uptoken.NewPutPolicy(bucketName, time.Now().Add(time.Hour))
		if err != nil {
			return "", err
		}
		return uptoken.NewSigner(putPolicy, credentials).GetUpToken(ctx)
	})
	if err != nil {
		return nil, err
	}
	req := httpclient.Request{Method: "POST", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, OnRequestProgress: options.OnRequestProgress}
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
	resp, err := storage.client.Do(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &ModifyObjectStatusResponse{}, resp.Body.Close()
}
