// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package workflow

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"

	clientsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/googlenow/googlenow"
	endpointsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/googlenow/googlenow"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	"go.uber.org/zap"
)

// GoogleNowAddCredentialsWorkflow defines the interface for GoogleNowAddCredentials workflow
type GoogleNowAddCredentialsWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args,
	) (zanzibar.Header, error)
}

// NewGoogleNowAddCredentialsWorkflow creates a workflow
func NewGoogleNowAddCredentialsWorkflow(deps *module.Dependencies) GoogleNowAddCredentialsWorkflow {
	return &googleNowAddCredentialsWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// googleNowAddCredentialsWorkflow calls thrift client GoogleNow.AddCredentials
type googleNowAddCredentialsWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w googleNowAddCredentialsWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args,
) (zanzibar.Header, error) {
	clientRequest := convertToAddCredentialsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Shadow-Start-Trace-Id")
	if ok {
		clientHeaders["X-Shadow-Start-Trace-Id"] = h
	}
	h, ok = reqHeaders.Get("X-Token")
	if ok {
		clientHeaders["X-Token"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}
	h, ok = reqHeaders.Get("X-Zanzibar-Use-Staging")
	if ok {
		clientHeaders["X-Zanzibar-Use-Staging"] = h
	}

	cliRespHeaders, err := w.Clients.GoogleNow.AddCredentials(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "GoogleNow"),
			)

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}
	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	return resHeaders, nil
}

func convertToAddCredentialsClientRequest(in *endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args) *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args {
	out := &clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args{}

	out.AuthCode = string(in.AuthCode)

	return out
}
