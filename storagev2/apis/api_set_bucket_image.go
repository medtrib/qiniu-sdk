// THIS FILE IS GENERATED BY api-generator, DO NOT EDIT DIRECTLY!

package apis

import (
	"context"
	"encoding/base64"
	auth "github.com/medtrib/qiniu-sdk/v7/auth"
	uplog "github.com/medtrib/qiniu-sdk/v7/internal/uplog"
	setbucketimage "github.com/medtrib/qiniu-sdk/v7/storagev2/apis/set_bucket_image"
	errors "github.com/medtrib/qiniu-sdk/v7/storagev2/errors"
	httpclient "github.com/medtrib/qiniu-sdk/v7/storagev2/http_client"
	region "github.com/medtrib/qiniu-sdk/v7/storagev2/region"
	uptoken "github.com/medtrib/qiniu-sdk/v7/storagev2/uptoken"
	"strings"
	"time"
)

type innerSetBucketImageRequest setbucketimage.Request

func (pp *innerSetBucketImageRequest) getBucketName(ctx context.Context) (string, error) {
	return pp.Bucket, nil
}
func (path *innerSetBucketImageRequest) buildPath() ([]string, error) {
	allSegments := make([]string, 0, 5)
	if path.Bucket != "" {
		allSegments = append(allSegments, path.Bucket)
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Bucket"}
	}
	if path.Url != "" {
		allSegments = append(allSegments, "from", base64.URLEncoding.EncodeToString([]byte(path.Url)))
	} else {
		return nil, errors.MissingRequiredFieldError{Name: "Url"}
	}
	if path.Host != "" {
		allSegments = append(allSegments, "host", path.Host)
	}
	return allSegments, nil
}

type SetBucketImageRequest = setbucketimage.Request
type SetBucketImageResponse = setbucketimage.Response

// 设置源站镜像回源
func (storage *Storage) SetBucketImage(ctx context.Context, request *SetBucketImageRequest, options *Options) (*SetBucketImageResponse, error) {
	if options == nil {
		options = &Options{}
	}
	innerRequest := (*innerSetBucketImageRequest)(request)
	serviceNames := []region.ServiceName{region.ServiceBucket}
	if innerRequest.Credentials == nil && storage.client.GetCredentials() == nil {
		return nil, errors.MissingRequiredFieldError{Name: "Credentials"}
	}
	pathSegments := make([]string, 0, 6)
	pathSegments = append(pathSegments, "image")
	if segments, err := innerRequest.buildPath(); err != nil {
		return nil, err
	} else {
		pathSegments = append(pathSegments, segments...)
	}
	path := "/" + strings.Join(pathSegments, "/")
	var rawQuery string
	bucketName := options.OverwrittenBucketName
	if bucketName == "" {
		var err error
		if bucketName, err = innerRequest.getBucketName(ctx); err != nil {
			return nil, err
		}
	}
	uplogInterceptor, err := uplog.NewRequestUplog("setBucketImage", bucketName, "", func() (string, error) {
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
	return &SetBucketImageResponse{}, resp.Body.Close()
}
