// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	getbuckets "github.com/medtrib/qiniu-sdk/v7/storagev2/apis/get_buckets"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strings"
	"time"
)

type innerGetBucketsRequest getbuckets.Request

func (query *innerGetBucketsRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Shared != "" {
		allQuery.Set("shared", query.Shared)
	}
	return allQuery, nil
}

type GetBucketsRequest = getbuckets.Request
type GetBucketsResponse = getbuckets.Response

// 获取拥有的所有存储空间列表
func (storage *Storage) GetBuckets(ctx context.Context, request *GetBucketsRequest, options *Options) (*GetBucketsResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetBucketsRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 1)
	pathSegments = append(pathSegments, "buckets")
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	if query, err := innerRequest.buildQuery(); err != nil {
		return nil, err
	} else {
		rawQuery += query.Encode()
	}
	bucketName := options.OverwrittenBucketName
	uplogInterceptor, err := uplog.NewRequestUplog("getBuckets", bucketName, "", func() (string, error) {
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
	var respBody GetBucketsResponse
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
