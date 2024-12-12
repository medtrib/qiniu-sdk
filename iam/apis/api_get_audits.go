// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/sulwan/qiniu-sdk/v7/auth"
	getaudits "github.com/sulwan/qiniu-sdk/v7/iam/apis/get_audits"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/sulwan/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type innerGetAuditsRequest getaudits.Request

func (query *innerGetAuditsRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Iuid != 0 {
		allQuery.Set("iuid", strconv.FormatInt(query.Iuid, 10))
	}
	if query.Service != "" {
		allQuery.Set("service", query.Service)
	}
	if query.Action != "" {
		allQuery.Set("action", query.Action)
	}
	if query.Resource != "" {
		allQuery.Set("resource", query.Resource)
	}
	if query.StartTime != "" {
		allQuery.Set("start_time", query.StartTime)
	}
	if query.EndTime != "" {
		allQuery.Set("end_time", query.EndTime)
	}
	if query.Marker != "" {
		allQuery.Set("marker", query.Marker)
	}
	if query.Limit != 0 {
		allQuery.Set("limit", strconv.FormatInt(query.Limit, 10))
	}
	return allQuery, nil
}

type GetAuditsRequest = getaudits.Request
type GetAuditsResponse = getaudits.Response

// 查询审计日志列表
func (iam *Iam) GetAudits(ctx context.Context, request *GetAuditsRequest, options *Options) (*GetAuditsResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetAuditsRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceApi}
	if innerRequest.Credentials == nil && iam.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 3)
	pathSegments = append(pathSegments, "iam", "v1", "audits")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	uplogInterceptor, err := uplog.NewRequestUplog("getAudits", "", "", func() (string, error) {
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
	var respBody GetAuditsResponse
	if err := iam.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
