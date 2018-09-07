package panic

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic/module"
	"github.com/uber/zanzibar/examples/example-gateway/build/endpoints/panic/workflow"
	"github.com/uber/zanzibar/runtime"
)

// NewServiceCFrontHelloWorkflow ...
func NewServiceCFrontHelloWorkflow(
	deps *module.Dependencies,
) workflow.ServiceCFrontHelloWorkflow {
	return &EndpointPanicEndpoint{
		Clients: deps.Client,
	}
}

// EndpointPanicEndpoint ...
type EndpointPanicEndpoint struct {
	Clients *module.ClientDependencies
}

// Handle ...
func (w EndpointPanicEndpoint) Handle(
	ctx context.Context,
	headers zanzibar.Header,
) (string, zanzibar.Header, error) {
	panic("panic at user's code ...")
}
