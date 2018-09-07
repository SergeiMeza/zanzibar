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

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	"go.uber.org/zap"
)

// SimpleServiceTransHeadersWorkflow defines the interface for SimpleServiceTransHeaders workflow
type SimpleServiceTransHeadersWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsBazBaz.SimpleService_TransHeaders_Args,
	) (*endpointsBazBaz.TransHeader, zanzibar.Header, error)
}

// NewSimpleServiceTransHeadersWorkflow creates a workflow
func NewSimpleServiceTransHeadersWorkflow(deps *module.Dependencies) SimpleServiceTransHeadersWorkflow {
	return &simpleServiceTransHeadersWorkflow{
		Clients: deps.Client,
		Logger:  deps.Default.Logger,
	}
}

// simpleServiceTransHeadersWorkflow calls thrift client Baz.TransHeaders
type simpleServiceTransHeadersWorkflow struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
}

// Handle calls thrift client.
func (w simpleServiceTransHeadersWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_TransHeaders_Args,
) (*endpointsBazBaz.TransHeader, zanzibar.Header, error) {
	clientRequest := convertToTransHeadersClientRequest(r)
	clientRequest = propagateHeadersTransHeadersClientRequests(clientRequest, reqHeaders)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("Token")
	if ok {
		clientHeaders["Token"] = h
	}
	h, ok = reqHeaders.Get("Uuid")
	if ok {
		clientHeaders["Uuid"] = h
	}
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Shadow-Start-Trace-Id")
	if ok {
		clientHeaders["X-Shadow-Start-Trace-Id"] = h
	}
	h, ok = reqHeaders.Get("X-Zanzibar-Use-Staging")
	if ok {
		clientHeaders["X-Zanzibar-Use-Staging"] = h
	}

	clientRespBody, _, err := w.Clients.Baz.TransHeaders(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertTransHeadersAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		case *clientsBazBaz.OtherAuthErr:
			serverErr := convertTransHeadersOtherAuthErr(
				errValue,
			)

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "Baz"),
			)

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertSimpleServiceTransHeadersClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToTransHeadersClientRequest(in *endpointsBazBaz.SimpleService_TransHeaders_Args) *clientsBazBaz.SimpleService_TransHeaders_Args {
	out := &clientsBazBaz.SimpleService_TransHeaders_Args{}

	if in.Req != nil {
		out.Req = &clientsBazBase.TransHeaders{}
	} else {
		out.Req = nil
	}

	return out
}

func convertTransHeadersAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertTransHeadersOtherAuthErr(
	clientError *clientsBazBaz.OtherAuthErr,
) *endpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertSimpleServiceTransHeadersClientResponse(in *clientsBazBase.TransHeaders) *endpointsBazBaz.TransHeader {
	out := &endpointsBazBaz.TransHeader{}

	return out
}

func propagateHeadersTransHeadersClientRequests(in *clientsBazBaz.SimpleService_TransHeaders_Args, headers zanzibar.Header) *clientsBazBaz.SimpleService_TransHeaders_Args {
	if in == nil {
		in = &clientsBazBaz.SimpleService_TransHeaders_Args{}
	}
	if key, ok := headers.Get("x-token"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W1 == nil {
			in.Req.W1 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W1.N1 == nil {
			in.Req.W1.N1 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W1.N1.Token = &key

	}
	if key, ok := headers.Get("x-uuid"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W1 == nil {
			in.Req.W1 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W1.N1 == nil {
			in.Req.W1.N1 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W1.N1.UUID = key

	}
	if key, ok := headers.Get("x-token"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W1 == nil {
			in.Req.W1 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W1.N2 == nil {
			in.Req.W1.N2 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W1.N2.Token = &key

	}
	if key, ok := headers.Get("x-uuid"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W1 == nil {
			in.Req.W1 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W1.N2 == nil {
			in.Req.W1.N2 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W1.N2.UUID = key

	}
	if key, ok := headers.Get("x-token"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W2 == nil {
			in.Req.W2 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W2.N1 == nil {
			in.Req.W2.N1 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W2.N1.Token = &key

	}
	if key, ok := headers.Get("x-uuid"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W2 == nil {
			in.Req.W2 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W2.N1 == nil {
			in.Req.W2.N1 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W2.N1.UUID = key

	}
	if key, ok := headers.Get("x-token"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W2 == nil {
			in.Req.W2 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W2.N2 == nil {
			in.Req.W2.N2 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W2.N2.Token = &key

	}
	if key, ok := headers.Get("x-uuid"); ok {
		if in.Req == nil {
			in.Req = &clientsBazBase.TransHeaders{}
		}
		if in.Req.W2 == nil {
			in.Req.W2 = &clientsBazBase.Wrapped{}
		}
		if in.Req.W2.N2 == nil {
			in.Req.W2.N2 = &clientsBazBase.NestHeaders{}
		}
		in.Req.W2.N2.UUID = key

	}
	return in
}
