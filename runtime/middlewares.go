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

package zanzibar

import (
	"context"

	jsonschema "github.com/mcuadros/go-jsonschema-generator"
)

// MiddlewareStack is a stack of Middleware Handlers that can be invoked as an Handler.
// MiddlewareStack middlewares are evaluated for requests in the order that they are added to the stack
// followed by the underlying HandlerFn. The middleware responses are then executed in reverse.
type MiddlewareStack struct {
	middlewares []MiddlewareHandle
	handle      HandlerFn
}

// NewStack returns a new MiddlewareStack instance with no middleware preconfigured.
func NewStack(middlewares []MiddlewareHandle,
	handle HandlerFn) *MiddlewareStack {
	return &MiddlewareStack{
		handle:      handle,
		middlewares: middlewares,
	}
}

// Middlewares returns a list of all the handlers in the current MiddlewareStack.
func (m *MiddlewareStack) Middlewares() []MiddlewareHandle {
	return m.middlewares
}

// MiddlewareHandle used to define middleware
type MiddlewareHandle interface {
	// implement HandleRequest for your middleware.
	HandleRequest(
		req *ServerHTTPRequest,
		res *ServerHTTPResponse,
		shared SharedState) error
	// implement HandleResponse for your middleware
	HandleResponse(
		res *ServerHTTPResponse,
		shared SharedState) error
	// return any shared state for this middleware.
	OwnState() interface{}
	JSONSchema() *jsonschema.Document
	Name() string
}

// SharedState used to access other middlewares in the chain.
type SharedState struct {
	middlewareDict map[string]MiddlewareHandle
}

func newSharedState(middlewares []MiddlewareHandle) SharedState {
	sharedState := SharedState{}
	sharedState.middlewareDict = make(map[string]MiddlewareHandle)

	for i := 0; i < len(middlewares); i++ {
		sharedState.middlewareDict[middlewares[i].Name()] = middlewares[i]
	}
	return sharedState
}

// GetState returns the state from a different middleware
func (s SharedState) GetState(name string) interface{} {
	return s.middlewareDict[name].OwnState()
}

// Handle executes the middlewares in a stack and underlying handler.
func (m *MiddlewareStack) Handle(
	ctx context.Context,
	req *ServerHTTPRequest,
	res *ServerHTTPResponse) {

	shared := newSharedState(m.middlewares)

	for i := 0; i < len(m.middlewares); i++ {
		err := m.middlewares[i].HandleRequest(req, res, shared)
		if err != nil {
			// Decide whether to log and 500 or change the
			// zanzibar.HandleFn to return an error and let
			// router process those at that level.
			return
		}
	}

	m.handle(ctx, req, res)

	for i := len(m.middlewares) - 1; i >= 0; i-- {
		err := m.middlewares[i].HandleResponse(res, shared)
		if err != nil {
			// Decide whether to log and 500 or change the
			// zanzibar.HandleFn to return an error and let
			// router process those at that level.
			return
		}
	}
}
