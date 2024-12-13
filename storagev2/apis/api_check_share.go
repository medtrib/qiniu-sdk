// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	checkshare "github.com/medtrib/qiniu-sdk/v7/storagev2/apis/check_share"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strings"
	"time"
)

type innerCheckShareRequest checkshare.Request

func (path *innerCheckShareRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 1)
	if path.ShareId != "" {
		allSegments = append(allSegments, path.ShareId)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "ShareId"}
	}
	return allSegments, nil
}
func (query *innerCheckShareRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Token != "" {
		allQuery.Set("token", query.Token)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Token"}
	}
	return allQuery, nil
}
func (request *innerCheckShareRequest) getAccessKey(ctx context.Context) (string, error) {
	if request.Credentials != nil {
		if credentials, err := request.Credentials.Get(ctx); err != nil {
			return "", err
		} else {
			return credentials.AccessKey, nil
		}
	}
	return "", nil
}

type CheckShareRequest = checkshare.Request
type CheckShareResponse = checkshare.Response

// 检查目录分享
func (storage *Storage) CheckShare(ctx context.Context, request *CheckShareRequest, options *Options) (*CheckShareResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerCheckShareRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 3)
	pathSegments = append(pathSegments, "shares")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	pathSegments = append(pathSegments, "check")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	bucketName := options.OverwrittenBucketName
	uplogInterceptor, err := uplog.NewRequestUplog("checkShare", bucketName, "", func() (string, error) {
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
		} else {

			req.Region = storage.client.GetAllRegions()
			if req.Region == nil {
				if options.OverwrittenBucketHosts != nil {
					if bucketHosts, err = options.OverwrittenBucketHosts.GetEndpoints(ctx); err != nil {
						return nil, err
					}
				}
				allRegionsOptions := region.AllRegionsProviderOptions{UseInsecureProtocol: storage.client.UseInsecureProtocol(), HostFreezeDuration: storage.client.GetHostFreezeDuration(), Resolver: storage.client.GetResolver(), Chooser: storage.client.GetChooser(), BeforeSign: storage.client.GetBeforeSignCallback(), AfterSign: storage.client.GetAfterSignCallback(), SignError: storage.client.GetSignErrorCallback(), BeforeResolve: storage.client.GetBeforeResolveCallback(), AfterResolve: storage.client.GetAfterResolveCallback(), ResolveError: storage.client.GetResolveErrorCallback(), BeforeBackoff: storage.client.GetBeforeBackoffCallback(), AfterBackoff: storage.client.GetAfterBackoffCallback(), BeforeRequest: storage.client.GetBeforeRequestCallback(), AfterResponse: storage.client.GetAfterResponseCallback()}
				if hostRetryConfig := storage.client.GetHostRetryConfig(); hostRetryConfig != nil {
					allRegionsOptions.RetryMax = hostRetryConfig.RetryMax
					allRegionsOptions.Backoff = hostRetryConfig.Backoff
				}
				credentials := innerRequest.Credentials
				if credentials == nil {
					credentials = storage.client.GetCredentials()
				}
				if req.Region, err = region.NewAllRegionsProvider(credentials, bucketHosts, &allRegionsOptions); err != nil {
					return nil, err
				}
			}
		}
	}
	resp, err := storage.client.Do(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &CheckShareResponse{}, resp.Body.Close()
}
