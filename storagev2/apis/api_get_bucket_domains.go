// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/sulwan/qiniu-sdk/v7/auth"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	getbucketdomains "github.com/sulwan/qiniu-sdk/v7/storagev2/apis/get_bucket_domains"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/sulwan/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strings"
	"time"
)

type innerGetBucketDomainsRequest getbucketdomains.Request

func (query *innerGetBucketDomainsRequest) getBucketName(ctx context.Context) (string, error) {
	return query.BucketName, nil
}
func (query *innerGetBucketDomainsRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.BucketName != "" {
		allQuery.Set("tbl", query.BucketName)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "BucketName"}
	}
	return allQuery, nil
}

type GetBucketDomainsRequest = getbucketdomains.Request
type GetBucketDomainsResponse = getbucketdomains.Response

// 获取存储空间的域名列表
func (storage *Storage) GetBucketDomains(ctx context.Context, request *GetBucketDomainsRequest, options *Options) (*GetBucketDomainsResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerGetBucketDomainsRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 2)
	pathSegments = append(pathSegments, "v2", "domains")
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
	uplogInterceptor, err := uplog.NewRequestUplog("getBucketDomains", bucketName, "", func() (string, error) {
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
	var respBody GetBucketDomainsResponse
	if err := storage.client.DoAndAcceptJSON(ctx, &req, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
