// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package userplex

import (
	"context"
	"net/http"
	"slices"

	"github.com/dqnamo/userplex-go/internal/apijson"
	"github.com/dqnamo/userplex-go/internal/requestconfig"
	"github.com/dqnamo/userplex-go/option"
	"github.com/dqnamo/userplex-go/packages/param"
	"github.com/dqnamo/userplex-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the userplex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Creates or updates an end user in InstantDB with the provided information.
// Requires a valid API key for authentication.
func (r *UserService) Identify(ctx context.Context, body UserIdentifyParams, opts ...option.RequestOption) (res *UserIdentifyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/identify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UserIdentifyResponse struct {
	// Operation success status
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserIdentifyResponse) RawJSON() string { return r.JSON.raw }
func (r *UserIdentifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserIdentifyParams struct {
	// Unique identifier for the user
	UserID string `json:"user_id,required"`
	// User email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// User full name
	Name param.Opt[string] `json:"name,omitzero"`
	// Additional user properties
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r UserIdentifyParams) MarshalJSON() (data []byte, err error) {
	type shadow UserIdentifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserIdentifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
