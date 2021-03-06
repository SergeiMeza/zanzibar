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

package corgeclient

import (
	"context"

	"github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/wire"

	clientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/corge/corge"
)

// CorgeEchoStringFunc is the handler function for "echoString" method of thrift service "Corge".
type CorgeEchoStringFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsCorgeCorge.Corge_EchoString_Args,
) (string, map[string]string, error)

// NewCorgeEchoStringHandler wraps a handler function so it can be registered with a thrift server.
func NewCorgeEchoStringHandler(f CorgeEchoStringFunc) zanzibar.TChannelHandler {
	return &CorgeEchoStringHandler{f}
}

// CorgeEchoStringHandler handles the "echoString" method call of thrift service "Corge".
type CorgeEchoStringHandler struct {
	echostring CorgeEchoStringFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *CorgeEchoStringHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsCorgeCorge.Corge_EchoString_Args
	var res clientsCorgeCorge.Corge_EchoString_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echostring(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = &r

	return err == nil, &res, respHeaders, nil
}
