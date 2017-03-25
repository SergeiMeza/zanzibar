// Copyright (c) 2017 Uber Technologies, Inc.
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

package runtime_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	// TODO(sindelar): Refactor into a unit test and remove the
	// example middleware (which creates a cyclic dependency)
	"github.com/uber/zanzibar/examples/example-gateway/middlewares/example"
	"github.com/uber/zanzibar/examples/example-gateway/middlewares/example_reader"

	zanzibar "github.com/uber/zanzibar/runtime"
	"github.com/uber/zanzibar/test/lib/bench_gateway"
)

// Ensures that a middleware stack can correctly return all of its handlers.
func TestHandlers(t *testing.T) {
	ex := example.NewMiddleWare(
		nil, // *zanzibar.Gateway
		example.Options{
			Foo: "foo",
			Bar: 2,
		},
	)

	middles := []zanzibar.MiddlewareHandle{ex}
	middlewareStack := zanzibar.NewStack(middles, noopHandlerFn)

	// Verify the custom middleware has been added.
	middlewares := middlewareStack.Middlewares()
	assert.Equal(t, 1, len(middlewares))

	// Run the zanzibar.HandleFn of composed middlewares.
	// TODO(sindelar): Refactor. We some helpers to build zanzibar
	// request/responses without setting up a backend and register.
	// Currently they require endpoints to instantiate.
	gateway, err := benchGateway.CreateGateway(nil, nil)
	if !assert.NoError(t, err) {
		return
	}

	bgateway := gateway.(*benchGateway.BenchGateway)

	bgateway.ActualGateway.Router.Register(
		"GET", "/foo",
		zanzibar.NewEndpoint(
			bgateway.ActualGateway,
			"foo",
			"foo",
			middlewareStack.Handle,
		),
	)
	resp, err := gateway.MakeRequest("GET", "/foo", nil)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

// Ensures that a middleware can read state from a middeware earlier in the stack.
func TestMiddlewareSharedStates(t *testing.T) {
	ex := example.NewMiddleWare(
		nil, // nil Gateway
		example.Options{
			Foo: "test_state",
			Bar: 2,
		},
	)
	exReader := exampleReader.NewMiddleWare(
		nil, // *zanzibar.Gateway
		exampleReader.Options{
			Foo: "foo",
		},
	)

	middles := []zanzibar.MiddlewareHandle{ex, exReader}
	middlewareStack := zanzibar.NewStack(middles, noopHandlerFn)

	// Verify the custom middleware has been added.
	middlewares := middlewareStack.Middlewares()
	assert.Equal(t, 2, len(middlewares))

	// Run the zanzibar.HandleFn of composed middlewares.
	// TODO(sindelar): Refactor. We some helpers to build zanzibar
	// request/responses without setting up a backend and register.
	// Currently they require endpoints to instantiate.
	gateway, err := benchGateway.CreateGateway(nil, nil)
	if !assert.NoError(t, err) {
		return
	}

	bgateway := gateway.(*benchGateway.BenchGateway)

	bgateway.ActualGateway.Router.Register(
		"GET", "/foo",
		zanzibar.NewEndpoint(
			bgateway.ActualGateway,
			"foo",
			"foo",
			middlewareStack.Handle,
		),
	)
	resp, err := gateway.MakeRequest("GET", "/foo", nil)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func noopHandlerFn(ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse) {
	return
}
