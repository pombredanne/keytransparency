// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mutationclient

//
// This file contains Monitor'c grpc client logic: poll mutations from the
// kt-server mutations API and page if necessary.
//

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	gpb "github.com/google/keytransparency/core/proto/keytransparency_v1_grpc"
	pb "github.com/google/keytransparency/core/proto/keytransparency_v1_proto"
)

// Each page contains pageSize profiles. Each profile contains multiple
// keys. Assuming 2 keys per profile (each of size 2048-bit), a page of
// size 16 will contain about 8KB of data.
const pageSize = 16

// Client queries the server side mutations API.
type Client struct {
	client     gpb.MutationServiceClient
	pollPeriod time.Duration
}

// New initializes a new mutations API monitoring client.
func New(client gpb.MutationServiceClient, pollPeriod time.Duration) *Client {
	return &Client{
		client:     client,
		pollPeriod: pollPeriod,
	}
}

// StartPolling initiates polling the server side mutations API. It does not
// block returns a channel.
// The caller should listen on the channel to receiving the latest polled
// mutations response including all paged mutations. If anything went wrong
// while polling the response channel contains an error.
func (c *Client) StartPolling(domainID string, startEpoch int64) (<-chan *pb.GetMutationsResponse, <-chan error) {
	response := make(chan *pb.GetMutationsResponse)
	errChan := make(chan error)
	go func() {
		epoch := startEpoch
		t := time.NewTicker(c.pollPeriod)
		for now := range t.C {
			glog.Infof("Polling: %v", now)
			// time out if we exceed the poll period:
			ctx, cancel := context.WithTimeout(context.Background(), c.pollPeriod)
			monitorResp, err := c.PollMutations(ctx, domainID, epoch)
			cancel()
			if err != nil {
				glog.Infof("PollMutations(_): %v", err)
				errChan <- err
			} else {
				// only write to results channel and increment epoch
				// if there was no error:
				glog.Infof("Got response %v at %v", monitorResp.Epoch, now)
				response <- monitorResp
				epoch++
			}
		}
	}()
	return response, errChan
}

// PollMutations polls GetMutationsResponses from the configured
// key-transparency server's mutations API.
func (c *Client) PollMutations(ctx context.Context, domainID string, queryEpoch int64, opts ...grpc.CallOption) (*pb.GetMutationsResponse, error) {
	response, err := c.client.GetMutations(ctx, &pb.GetMutationsRequest{
		DomainId: domainID,
		PageSize: pageSize,
		Epoch:    queryEpoch,
	}, opts...)
	if err != nil {
		return nil, fmt.Errorf("GetMutations(%v, %v): %v", domainID, pageSize, err)
	}

	// Page if necessary: query all mutations in the current epoch
	for response.GetNextPageToken() != "" {
		req := &pb.GetMutationsRequest{
			DomainId: domainID,
			PageSize: pageSize,
		}
		resp, err := c.client.GetMutations(ctx, req, opts...)
		if err != nil {
			return nil, err
		}
		response.Mutations = append(response.Mutations, resp.GetMutations()...)
	}

	return response, nil
}
