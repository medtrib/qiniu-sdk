// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	getbucketeventrules "github.com/medtrib/qiniu-sdk/v7/storagev2/apis/get_bucket_event_rules"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strings"
	"time"
)

type innerGetBucketEventRulesRequest getbucketeventrules.Request

func (query *innerGetBucketEventRulesRequest) getBucketName(ctx context.Context) (string, error) {
	return query.Bucket, nil
}
func (query *innerGetBucketEventRulesRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Bucket != "" {
		allQuery.Set("bucket", query.Bucket)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Bucket"}
	}
	return allQuery, nil
}

type GetBucketEventRulesRequest = getbucketeventrules.Request
type GetBucketEventRulesResponse = getbucketeventrules.Response

// 获取存储空间事件通知规则
func (storage *Storage) GetBucketEventRules(ctx context.Context, request *GetBucketEventRulesRequest, options *Options) (*GetBucketEventRulesResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetBucketEventRulesRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 2)
	pathSegments = append(pathSegments, "events", "get")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	bucketName := options.OverwrittenBucketName
	if bucketName == "" {
		var err error
		if bucketName, err = innerRequest.getBucketName(ctx); err != nil {
			return nil, err
		}
	}
	uplogInterceptor, err := uplog.NewRequestUplog("getBucketEventRules", bucketName, "", func() (string, error) {
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
	req := httpclient.Request{Method: "GET", ServiceNames: serviceNames, Path: path, RawQuery: rawQuery, Endpoints: options.OverwrittenEndpoints, Region: options.OverwrittenRegion, Interceptors: []httpclient.Interceptor{uplogInterceptor}, AuthType: auth.TokenQiniu, Credentials: innerRequest.Credentials, BufferResponse: true, OnRequestProgress: options.OnRequestProgress}
	if options.OverwrittenEndpoints == nil && options.OverwrittenRegion == nil && storage.client.GetRegions() == nil {
		bucketHosts := httpclient.DefaultBucketHosts()
		if options.OverwrittenBucketHosts != nil {
			req.Endpoints = options.OverwrittenBucketHosts
		} else {
			req.Endpoints = bucketHosts
		}
	}
	var respBody GetBucketEventRulesResponse
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
