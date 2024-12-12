// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/sulwan/qiniu-sdk/v7/auth"
	enableuserkeypair "github.com/sulwan/qiniu-sdk/v7/iam/apis/enable_user_keypair"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/sulwan/qiniu-sdk/v7/storagev2/uptoken"
	"strings"
	"time"
)

type innerEnableUserKeypairRequest enableuserkeypair.Request

func (path *innerEnableUserKeypairRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 3)
	if path.Alias != "" {
		allSegments = append(allSegments, path.Alias)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Alias"}
	}
	if path.AccessKey != "" {
		allSegments = append(allSegments, "keypairs", path.AccessKey)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "AccessKey"}
	}
	return allSegments, nil
}

type EnableUserKeypairRequest = enableuserkeypair.Request
type EnableUserKeypairResponse = enableuserkeypair.Response

// 启用 IAM 子账号密钥
func (iam *Iam) EnableUserKeypair(ctx context.Context, request *EnableUserKeypairRequest, options *Options) (*EnableUserKeypairResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerEnableUserKeypairRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && iam.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 7)
	pathSegments = append(pathSegments, "iam", "v1", "users")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	pathSegments = append(pathSegments, "enable")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	uplogInterceptor, err := uplog.NewRequestUplog("enableUserKeypair", "", "", func() (string, error) {
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
	req := httpclient.Request{Method: "POST", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, OnRequestProgress: options.OnRequestProgress}
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
	resp, err := iam.client.Do(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &EnableUserKeypairResponse{}, resp.Body.Close()
}
