// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package userplex

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/dqnamo/userplex-go/internal/apijson"
	"github.com/dqnamo/userplex-go/internal/requestconfig"
	"github.com/dqnamo/userplex-go/option"
	"github.com/dqnamo/userplex-go/packages/param"
	"github.com/dqnamo/userplex-go/packages/respjson"
)

// EventService contains methods and other services that help with interacting with
// the userplex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Creates or uses an existing event and records an event occurrence for an end
// user. Requires a valid API key for authentication.
func (r *EventService) New(ctx context.Context, body EventNewParams, opts ...option.RequestOption) (res *EventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/event"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EventNewResponse struct {
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
func (r EventNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EventNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventNewParams struct {
	Name string `json:"name,required"`
	// External user ID
	UserID string `json:"user_id,required"`
	// Event timestamp (ISO 8601)
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Additional event properties
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r EventNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EventNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
