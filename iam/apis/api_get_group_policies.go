// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	getgrouppolicies "github.com/medtrib/qiniu-sdk/v7/iam/apis/get_group_policies"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type innerGetGroupPoliciesRequest getgrouppolicies.Request

func (path *innerGetGroupPoliciesRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 1)
	if path.Alias != "" {
		allSegments = append(allSegments, path.Alias)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Alias"}
	}
	return allSegments, nil
}
func (query *innerGetGroupPoliciesRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Page != 0 {
		allQuery.Set("page", strconv.FormatInt(query.Page, 10))
	}
	if query.PageSize != 0 {
		allQuery.Set("page_size", strconv.FormatInt(query.PageSize, 10))
	}
	return allQuery, nil
}

type GetGroupPoliciesRequest = getgrouppolicies.Request
type GetGroupPoliciesResponse = getgrouppolicies.Response

// 查询用户分组下分配的授权策略
func (iam *Iam) GetGroupPolicies(ctx context.Context, request *GetGroupPoliciesRequest, options *Options) (*GetGroupPoliciesResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetGroupPoliciesRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && iam.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 5)
	pathSegments = append(pathSegments, "iam", "v1", "groups")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	pathSegments = append(pathSegments, "policies")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	uplogInterceptor, err := uplog.NewRequestUplog("getGroupPolicies", "", "", func() (string, error) {
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
	var respBody GetGroupPoliciesResponse
	if err := iam.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
