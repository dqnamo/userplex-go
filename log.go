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

// LogService contains methods and other services that help with interacting with
// the userplex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r LogService) {
	r = LogService{}
	r.Options = opts
	return
}

// Records multiple log occurrences in a single request. Requires a valid API key
// for authentication.
func (r *LogService) Batch(ctx context.Context, body LogBatchParams, opts ...option.RequestOption) (res *LogBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/logs/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates or uses an existing log and records a log occurrence for an end user.
// Requires a valid API key for authentication.
func (r *LogService) New(ctx context.Context, body LogNewParams, opts ...option.RequestOption) (res *LogNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/log"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LogBatchResponse struct {
	// Number of logs processed
	Count float64 `json:"count,required"`
	// Operation success status
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LogBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *LogBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogNewResponse struct {
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
func (r LogNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LogNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogBatchParams struct {
	// List of logs to track
	Logs []LogBatchParamsLog `json:"logs,omitzero,required"`
	paramObj
}

func (r LogBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow LogBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, UserID are required.
type LogBatchParamsLog struct {
	Name string `json:"name,required"`
	// External user ID
	UserID string `json:"user_id,required"`
	// Log timestamp (ISO 8601)
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Additional log data
	Data map[string]any `json:"data,omitzero"`
	// Alias for data, for compatibility
	Properties  map[string]any `json:"properties,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r LogBatchParamsLog) MarshalJSON() (data []byte, err error) {
	type shadow LogBatchParamsLog
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *LogBatchParamsLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LogNewParams struct {
	Name string `json:"name,required"`
	// External user ID
	UserID string `json:"user_id,required"`
	// Log timestamp (ISO 8601)
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// Additional log data
	Data map[string]any `json:"data,omitzero"`
	// Alias for data, for compatibility
	Properties map[string]any `json:"properties,omitzero"`
	paramObj
}

func (r LogNewParams) MarshalJSON() (data []byte, err error) {
	type shadow LogNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LogNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
