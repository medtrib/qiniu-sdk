// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	"encoding/json"
	auth "github.com/sulwan/qiniu-sdk/v7/auth"
	modifypolicy "github.com/sulwan/qiniu-sdk/v7/iam/apis/modify_policy"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/sulwan/qiniu-sdk/v7/storagev2/uptoken"
	"strings"
	"time"
)

type innerModifyPolicyRequest modifypolicy.Request

func (path *innerModifyPolicyRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 1)
	if path.Alias != "" {
		allSegments = append(allSegments, path.Alias)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Alias"}
	}
	return allSegments, nil
}
func (j *innerModifyPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal((*modifypolicy.Request)(j))
}
func (j *innerModifyPolicyRequest) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*modifypolicy.Request)(j))
}

type ModifyPolicyRequest = modifypolicy.Request
type ModifyPolicyResponse = modifypolicy.Response

// 修改授权策略
func (iam *Iam) ModifyPolicy(ctx context.Context, request *ModifyPolicyRequest, options *Options) (*ModifyPolicyResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerModifyPolicyRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && iam.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 4)
	pathSegments = append(pathSegments, "iam", "v1", "policies")
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
	uplogInterceptor, err := uplog.NewRequestUplog("modifyPolicy", "", "", func() (string, error) {
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
	req := httpclient.Request{Method: "PATCH", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, BufferResponse: true, RequestBody: body, OnRequestProgress: options.OnRequestProgress}
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
	var respBody ModifyPolicyResponse
	if err := iam.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}