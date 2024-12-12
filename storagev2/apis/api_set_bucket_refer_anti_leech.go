// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	auth "github.com/sulwan/qiniu-sdk/v7/auth"
	uplog "github.com/sulwan/qiniu-sdk/v7/internal/uplog"
	setbucketreferantileech "github.com/sulwan/qiniu-sdk/v7/storagev2/apis/set_bucket_refer_anti_leech"
	errors "github.com/sulwan/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/sulwan/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/sulwan/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/sulwan/qiniu-sdk/v7/storagev2/uptoken"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type innerSetBucketReferAntiLeechRequest setbucketreferantileech.Request

func (query *innerSetBucketReferAntiLeechRequest) getBucketName(ctx context.Context) (string, error) {
	return query.Bucket, nil
}
func (query *innerSetBucketReferAntiLeechRequest) buildQuery() (url.Values, error) {
	allQuery := make(url.Values)
	if query.Bucket != "" {
		allQuery.Set("bucket", query.Bucket)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Bucket"}
	}
	allQuery.Set("mode", strconv.FormatInt(query.Mode, 10))
	allQuery.Set("pattern", query.Pattern)
	allQuery.Set("norefer", strconv.FormatInt(query.AllowEmptyReferer, 10))
	allQuery.Set("source_enabled", strconv.FormatInt(query.SourceEnabled, 10))
	return allQuery, nil
}

type SetBucketReferAntiLeechRequest = setbucketreferantileech.Request
type SetBucketReferAntiLeechResponse = setbucketreferantileech.Response

// 设置存储空间的防盗链模式
func (storage *Storage) SetBucketReferAntiLeech(ctx context.Context, request *SetBucketReferAntiLeechRequest, options *Options) (*SetBucketReferAntiLeechResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerSetBucketReferAntiLeechRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 1)
	pathSegments = append(pathSegments, "referAntiLeech")
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
	uplogInterceptor, err := uplog.NewRequestUplog("setBucketReferAntiLeech", bucketName, "", func() (string, error) {
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
		if options.OverwrittenBucketHosts != nil {
			req.Endpoints = options.OverwrittenBucketHosts
		} else {
			req.Endpoints = bucketHosts
		}
	}
	resp, err := storage.client.Do(ctx, &req)
	if err != nil {
		return nil, err
	}
	return &SetBucketReferAntiLeechResponse{}, resp.Body.Close()
}
