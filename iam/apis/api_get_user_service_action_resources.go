// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	getuserserviceactionresources "github.com/medtrib/qiniu-sdk/v7/iam/apis/get_user_service_action_resources"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"strings"
	"time"
)

type innerGetUserServiceActionResourcesRequest getuserserviceactionresources.Request

func (path *innerGetUserServiceActionResourcesRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 5)
	if path.UserAlias != "" {
		allSegments = append(allSegments, path.UserAlias)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "UserAlias"}
	}
	if path.Service != "" {
		allSegments = append(allSegments, "services", path.Service)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Service"}
	}
	if path.ActionAlias != "" {
		allSegments = append(allSegments, "actions", path.ActionAlias)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "ActionAlias"}
	}
	return allSegments, nil
}

type GetUserServiceActionResourcesRequest = getuserserviceactionresources.Request
type GetUserServiceActionResourcesResponse = getuserserviceactionresources.Response

// 列举子账号指定服务操作下的可访问资源
func (iam *Iam) GetUserServiceActionResources(ctx context.Context, request *GetUserServiceActionResourcesRequest, options *Options) (*GetUserServiceActionResourcesResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetUserServiceActionResourcesRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && iam.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 9)
	pathSegments = append(pathSegments, "iam", "v1", "users")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	pathSegments = append(pathSegments, "resources")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	uplogInterceptor, err := uplog.NewRequestUplog("getUserServiceActionResources", "", "", func() (string, error) {
		credentials := innerRequest.Credentials
		if credentials == nil {
			credentials = iam.client.GetCredentials()
		}
		putPolicy, err := uptoken.NewPutPolicy("", time.Now().Add(time.Hour))
		if err != nil {
			return "", err
		}
		return uptoken.NewSigner(putPolicy, credentials).GetUpToken(ctx)
	})
	if err != nil {
		return nil, err
	}
	req := httpclient.Request{Method: "GET", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, BufferResponse: true, OnRequestProgress: options.OnRequestProgress}
	if options.OverwrittenEndpoints == nil && options.OverwrittenRegion == nil && iam.client.GetRegions() == nil {
		bucketHosts := httpclient.DefaultBucketHosts()

		req.Region = iam.client.GetAllRegions()
		if req.Region == nil {
			if options.OverwrittenBucketHosts != nil {
				if bucketHosts, err = options.OverwrittenBucketHosts.GetEndpoints(ctx); err != nil {
					return nil, err
				}
			}
			allRegionsOptions := region.AllRegionsProviderOptions{UseInsecureProtocol: iam.client.UseInsecureProtocol(), HostFreezeDuration: iam.client.GetHostFreezeDuration(), Resolver: iam.client.GetResolver(), Chooser: iam.client.GetChooser(), BeforeSign: iam.client.GetBeforeSignCallback(), AfterSign: iam.client.GetAfterSignCallback(), SignError: iam.client.GetSignErrorCallback(), BeforeResolve: iam.client.GetBeforeResolveCallback(), AfterResolve: iam.client.GetAfterResolveCallback(), ResolveError: iam.client.GetResolveErrorCallback(), BeforeBackoff: iam.client.GetBeforeBackoffCallback(), AfterBackoff: iam.client.GetAfterBackoffCallback(), BeforeRequest: iam.client.GetBeforeRequestCallback(), AfterResponse: iam.client.GetAfterResponseCallback()}
			if hostRetryConfig := iam.client.GetHostRetryConfig(); hostRetryConfig != nil {
				allRegionsOptions.RetryMax = hostRetryConfig.RetryMax
				allRegionsOptions.Backoff = hostRetryConfig.Backoff
			}
			credentials := innerRequest.Credentials
			if credentials == nil {
				credentials = iam.client.GetCredentials()
			}
			if req.Region, err = region.NewAllRegionsProvider(credentials, bucketHosts, &allRegionsOptions); err != nil {
				return nil, err
			}
		}
	}
	var respBody GetUserServiceActionResourcesResponse
	if err := iam.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
