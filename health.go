// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oxp

import (
	"context"
	"net/http"

	"github.com/OpenExecProtocol/oxp-go/internal/requestconfig"
	"github.com/OpenExecProtocol/oxp-go/option"
)

// HealthService contains methods and other services that help with interacting
// with the oxp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r *HealthService) {
	r = &HealthService{}
	r.Options = opts
	return
}

// Checks the health status of the server.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}
