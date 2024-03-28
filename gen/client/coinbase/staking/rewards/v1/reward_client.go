// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.


package v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	rewardpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/rewards/v1"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newRewardClientHook clientHook

// RewardCallOptions contains the retry settings for each method of RewardClient.
type RewardCallOptions struct {
	ListRewards []gax.CallOption
	ListStakes []gax.CallOption
}

func defaultRewardRESTCallOptions() *RewardCallOptions {
	return &RewardCallOptions{
		ListRewards: []gax.CallOption{
		},
		ListStakes: []gax.CallOption{
		},
	}
}

// internalRewardClient is an interface that defines the methods available from .
type internalRewardClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListRewards(context.Context, *rewardpb.ListRewardsRequest, ...gax.CallOption) *RewardIterator
	ListStakes(context.Context, *rewardpb.ListStakesRequest, ...gax.CallOption) *StakeIterator
}

// RewardClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// RewardService exposes publicly available onchain staking-related rewards data across all supported protocols.
type RewardClient struct {
	// The internal transport-dependent client.
	internalClient internalRewardClient

	// The call options for this service.
	CallOptions *RewardCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *RewardClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *RewardClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *RewardClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListRewards list rewards for a given protocol.
func (c *RewardClient) ListRewards(ctx context.Context, req *rewardpb.ListRewardsRequest, opts ...gax.CallOption) *RewardIterator {
	return c.internalClient.ListRewards(ctx, req, opts...)
}

// ListStakes list staking activities for a given protocol.
func (c *RewardClient) ListStakes(ctx context.Context, req *rewardpb.ListStakesRequest, opts ...gax.CallOption) *StakeIterator {
	return c.internalClient.ListStakes(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type rewardRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing RewardClient
	CallOptions **RewardCallOptions
}

// NewRewardRESTClient creates a new reward service rest client.
//
// RewardService exposes publicly available onchain staking-related rewards data across all supported protocols.
func NewRewardRESTClient(ctx context.Context, opts ...option.ClientOption) (*RewardClient, error) {
	clientOpts := append(defaultRewardRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRewardRESTCallOptions()
	c := &rewardRESTClient{
		endpoint: endpoint,
		httpClient: httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &RewardClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRewardRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://api.developer.coinbase.com"),
		internaloption.WithDefaultMTLSEndpoint("https://api.developer.coinbase.com"),
		internaloption.WithDefaultAudience("https://api.developer.coinbase.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}
// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *rewardRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *rewardRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *rewardRESTClient) Connection() *grpc.ClientConn {
	return nil
}
// ListRewards list rewards for a given protocol.
func (c *rewardRESTClient) ListRewards(ctx context.Context, req *rewardpb.ListRewardsRequest, opts ...gax.CallOption) *RewardIterator {
	it := &RewardIterator{}
	req = proto.Clone(req).(*rewardpb.ListRewardsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*rewardpb.Reward, string, error) {
		resp := &rewardpb.ListRewardsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/rewards", req.GetParent())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil{
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetRewards(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
// ListStakes list staking activities for a given protocol.
func (c *rewardRESTClient) ListStakes(ctx context.Context, req *rewardpb.ListStakesRequest, opts ...gax.CallOption) *StakeIterator {
	it := &StakeIterator{}
	req = proto.Clone(req).(*rewardpb.ListStakesRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*rewardpb.Stake, string, error) {
		resp := &rewardpb.ListStakesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1/%v/stakes", req.GetParent())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		headers := buildHeaders(ctx, c.xGoogMetadata, metadata.Pairs("Content-Type", "application/json"))
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			httpRsp, err := c.httpClient.Do(httpReq)
			if err != nil{
				return err
			}
			defer httpRsp.Body.Close()

			if err = googleapi.CheckResponse(httpRsp); err != nil {
				return err
			}

			buf, err := ioutil.ReadAll(httpRsp.Body)
			if err != nil {
				return err
			}

			if err := unm.Unmarshal(buf, resp); err != nil {
				return maybeUnknownEnum(err)
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetStakes(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
// RewardIterator manages a stream of *rewardpb.Reward.
type RewardIterator struct {
	items    []*rewardpb.Reward
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*rewardpb.Reward, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RewardIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RewardIterator) Next() (*rewardpb.Reward, error) {
	var item *rewardpb.Reward
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RewardIterator) bufLen() int {
	return len(it.items)
}

func (it *RewardIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// StakeIterator manages a stream of *rewardpb.Stake.
type StakeIterator struct {
	items    []*rewardpb.Stake
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*rewardpb.Stake, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *StakeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *StakeIterator) Next() (*rewardpb.Stake, error) {
	var item *rewardpb.Stake
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *StakeIterator) bufLen() int {
	return len(it.items)
}

func (it *StakeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
