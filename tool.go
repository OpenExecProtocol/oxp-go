// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oxp

import (
	"context"
	"net/http"
	"reflect"

	"github.com/OpenExecProtocol/oxp-go/internal/apijson"
	"github.com/OpenExecProtocol/oxp-go/internal/param"
	"github.com/OpenExecProtocol/oxp-go/internal/requestconfig"
	"github.com/OpenExecProtocol/oxp-go/option"
	"github.com/OpenExecProtocol/oxp-go/shared"
	"github.com/tidwall/gjson"
)

// ToolService contains methods and other services that help with interacting with
// the oxp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	return
}

// Returns a list of tool definitions.
func (r *ToolService) List(ctx context.Context, query ToolListParams, opts ...option.RequestOption) (res *ToolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Calls a tool with the given parameters.
func (r *ToolService) Call(ctx context.Context, body ToolCallParams, opts ...option.RequestOption) (res *ToolCallResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "tools/call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ToolListResponse struct {
	Items  []ToolListResponseItem `json:"items,required"`
	Schema string                 `json:"$schema" format:"uri"`
	JSON   toolListResponseJSON   `json:"-"`
}

// toolListResponseJSON contains the JSON metadata for the struct
// [ToolListResponse]
type toolListResponseJSON struct {
	Items       apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseJSON) RawJSON() string {
	return r.raw
}

type ToolListResponseItem struct {
	// A tool's unique identifier in the format 'Toolkit.Tool[@version]', where
	// @version is optional.
	ID string `json:"id,required"`
	// A plain language explanation of when and how the tool should be used.
	Description string `json:"description,required"`
	// The tool's name. Only allows alphanumeric characters, underscores, and dashes.
	Name string `json:"name,required"`
	// A tool's semantic version in the format 'x.y.z', where x, y, and z are integers.
	Version string `json:"version,required"`
	// JSON Schema describing the input parameters for the tool. Supports standard JSON
	// Schema validation but excludes $ref and definitions/schemas for simplicity.
	InputSchema map[string]interface{} `json:"input_schema"`
	// JSON Schema describing the output parameters for the tool. Supports standard
	// JSON Schema validation but excludes $ref and definitions/schemas for simplicity.
	OutputSchema map[string]interface{}            `json:"output_schema,nullable"`
	Requirements ToolListResponseItemsRequirements `json:"requirements"`
	JSON         toolListResponseItemJSON          `json:"-"`
}

// toolListResponseItemJSON contains the JSON metadata for the struct
// [ToolListResponseItem]
type toolListResponseItemJSON struct {
	ID           apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Version      apijson.Field
	InputSchema  apijson.Field
	OutputSchema apijson.Field
	Requirements apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ToolListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseItemJSON) RawJSON() string {
	return r.raw
}

type ToolListResponseItemsRequirements struct {
	Authorization []ToolListResponseItemsRequirementsAuthorization `json:"authorization"`
	Secrets       []ToolListResponseItemsRequirementsSecret        `json:"secrets"`
	// Whether the tool requires a user ID.
	UserID      bool                                  `json:"user_id"`
	ExtraFields map[string]interface{}                `json:"-,extras"`
	JSON        toolListResponseItemsRequirementsJSON `json:"-"`
}

// toolListResponseItemsRequirementsJSON contains the JSON metadata for the struct
// [ToolListResponseItemsRequirements]
type toolListResponseItemsRequirementsJSON struct {
	Authorization apijson.Field
	Secrets       apijson.Field
	UserID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolListResponseItemsRequirements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseItemsRequirementsJSON) RawJSON() string {
	return r.raw
}

type ToolListResponseItemsRequirementsAuthorization struct {
	// A provider's unique identifier, allowing the tool to specify a specific
	// authorization provider.
	ID string `json:"id"`
	// OAuth 2.0-specific authorization details.
	Oauth2 ToolListResponseItemsRequirementsAuthorizationOauth2 `json:"oauth2"`
	JSON   toolListResponseItemsRequirementsAuthorizationJSON   `json:"-"`
}

// toolListResponseItemsRequirementsAuthorizationJSON contains the JSON metadata
// for the struct [ToolListResponseItemsRequirementsAuthorization]
type toolListResponseItemsRequirementsAuthorizationJSON struct {
	ID          apijson.Field
	Oauth2      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListResponseItemsRequirementsAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseItemsRequirementsAuthorizationJSON) RawJSON() string {
	return r.raw
}

// OAuth 2.0-specific authorization details.
type ToolListResponseItemsRequirementsAuthorizationOauth2 struct {
	Scopes      []string                                                 `json:"scopes"`
	ExtraFields map[string]interface{}                                   `json:"-,extras"`
	JSON        toolListResponseItemsRequirementsAuthorizationOauth2JSON `json:"-"`
}

// toolListResponseItemsRequirementsAuthorizationOauth2JSON contains the JSON
// metadata for the struct [ToolListResponseItemsRequirementsAuthorizationOauth2]
type toolListResponseItemsRequirementsAuthorizationOauth2JSON struct {
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListResponseItemsRequirementsAuthorizationOauth2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseItemsRequirementsAuthorizationOauth2JSON) RawJSON() string {
	return r.raw
}

type ToolListResponseItemsRequirementsSecret struct {
	// The secret's unique identifier.
	ID          string                                      `json:"id,required"`
	ExtraFields map[string]interface{}                      `json:"-,extras"`
	JSON        toolListResponseItemsRequirementsSecretJSON `json:"-"`
}

// toolListResponseItemsRequirementsSecretJSON contains the JSON metadata for the
// struct [ToolListResponseItemsRequirementsSecret]
type toolListResponseItemsRequirementsSecretJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListResponseItemsRequirementsSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseItemsRequirementsSecretJSON) RawJSON() string {
	return r.raw
}

type ToolCallResponse struct {
	Result ToolCallResponseResult `json:"result,required"`
	Schema string                 `json:"$schema" format:"uri"`
	JSON   toolCallResponseJSON   `json:"-"`
}

// toolCallResponseJSON contains the JSON metadata for the struct
// [ToolCallResponse]
type toolCallResponseJSON struct {
	Result      apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolCallResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolCallResponseJSON) RawJSON() string {
	return r.raw
}

type ToolCallResponseResult struct {
	// The unique identifier (e.g. UUID) for this tool call. If an ID is not provided
	// by the client, the server will generate one.
	CallID string `json:"call_id"`
	// The runtime duration of the tool call, in milliseconds
	Duration float64 `json:"duration"`
	// This field can have the runtime type of [ToolCallResponseResultObjectError].
	Error   interface{}                   `json:"error"`
	Success ToolCallResponseResultSuccess `json:"success"`
	// This field can have the runtime type of
	// [ToolCallResponseResultObjectValueUnion].
	Value interface{}                `json:"value"`
	JSON  toolCallResponseResultJSON `json:"-"`
	union ToolCallResponseResultUnion
}

// toolCallResponseResultJSON contains the JSON metadata for the struct
// [ToolCallResponseResult]
type toolCallResponseResultJSON struct {
	CallID      apijson.Field
	Duration    apijson.Field
	Error       apijson.Field
	Success     apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r toolCallResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ToolCallResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ToolCallResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ToolCallResponseResultUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [ToolCallResponseResultObject],
// [ToolCallResponseResultObject].
func (r ToolCallResponseResult) AsUnion() ToolCallResponseResultUnion {
	return r.union
}

// Union satisfied by [ToolCallResponseResultObject] or
// [ToolCallResponseResultObject].
type ToolCallResponseResultUnion interface {
	implementsToolCallResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolCallResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ToolCallResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ToolCallResponseResultObject{}),
		},
	)
}

type ToolCallResponseResultObject struct {
	// The unique identifier (e.g. UUID) for this tool call. If an ID is not provided
	// by the client, the server will generate one.
	CallID string `json:"call_id"`
	// The runtime duration of the tool call, in milliseconds
	Duration float64                             `json:"duration"`
	Success  ToolCallResponseResultObjectSuccess `json:"success"`
	// The value returned from the tool.
	Value ToolCallResponseResultObjectValueUnion `json:"value,nullable"`
	JSON  toolCallResponseResultObjectJSON       `json:"-"`
}

// toolCallResponseResultObjectJSON contains the JSON metadata for the struct
// [ToolCallResponseResultObject]
type toolCallResponseResultObjectJSON struct {
	CallID      apijson.Field
	Duration    apijson.Field
	Success     apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolCallResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolCallResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r ToolCallResponseResultObject) implementsToolCallResponseResult() {}

type ToolCallResponseResultObjectSuccess bool

const (
	ToolCallResponseResultObjectSuccessTrue ToolCallResponseResultObjectSuccess = true
)

func (r ToolCallResponseResultObjectSuccess) IsKnown() bool {
	switch r {
	case ToolCallResponseResultObjectSuccessTrue:
		return true
	}
	return false
}

// The value returned from the tool.
//
// Union satisfied by [ToolCallResponseResultObjectValueMap],
// [ToolCallResponseResultObjectValueArray], [shared.UnionString],
// [shared.UnionFloat] or [shared.UnionBool].
type ToolCallResponseResultObjectValueUnion interface {
	ImplementsToolCallResponseResultObjectValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ToolCallResponseResultObjectValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ToolCallResponseResultObjectValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ToolCallResponseResultObjectValueMap map[string]interface{}

func (r ToolCallResponseResultObjectValueMap) ImplementsToolCallResponseResultObjectValueUnion() {}

type ToolCallResponseResultObjectValueArray []interface{}

func (r ToolCallResponseResultObjectValueArray) ImplementsToolCallResponseResultObjectValueUnion() {}

type ToolCallResponseResultSuccess bool

const (
	ToolCallResponseResultSuccessTrue  ToolCallResponseResultSuccess = true
	ToolCallResponseResultSuccessFalse ToolCallResponseResultSuccess = false
)

func (r ToolCallResponseResultSuccess) IsKnown() bool {
	switch r {
	case ToolCallResponseResultSuccessTrue, ToolCallResponseResultSuccessFalse:
		return true
	}
	return false
}

type ToolListParams struct {
}

type ToolCallParams struct {
	Request param.Field[ToolCallParamsRequest] `json:"request,required"`
	Schema  param.Field[string]                `json:"$schema" format:"uri"`
}

func (r ToolCallParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolCallParamsRequest struct {
	// A tool's unique identifier in the format 'Toolkit.Tool[@version]', where
	// @version is optional.
	ToolID param.Field[string] `json:"tool_id,required"`
	// A unique identifier (e.g. UUID) for this tool call. Used as an idempotency key.
	// If omitted, the server will generate an ID.
	CallID  param.Field[string]                       `json:"call_id"`
	Context param.Field[ToolCallParamsRequestContext] `json:"context"`
	// The input parameters for the tool call.
	Input param.Field[map[string]interface{}] `json:"input"`
	// An optional trace identifier for the tool call.
	TraceID param.Field[string] `json:"trace_id"`
}

func (r ToolCallParamsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolCallParamsRequestContext struct {
	// The authorization information for the tool call.
	Authorization param.Field[[]ToolCallParamsRequestContextAuthorization] `json:"authorization"`
	// The secrets for the tool call.
	Secrets param.Field[[]ToolCallParamsRequestContextSecret] `json:"secrets"`
	// A unique ID that identifies the user, if required by the tool.
	UserID      param.Field[string]    `json:"user_id"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ToolCallParamsRequestContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolCallParamsRequestContextAuthorization struct {
	// The unique identifier for the authorization method or authorization provider.
	ID param.Field[string] `json:"id,required"`
	// The token for the tool call.
	Token       param.Field[string]    `json:"token,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ToolCallParamsRequestContextAuthorization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolCallParamsRequestContextSecret struct {
	// The secret's unique identifier.
	ID param.Field[string] `json:"id,required"`
	// The secret's value.
	Value       param.Field[string]    `json:"value,required"`
	ExtraFields map[string]interface{} `json:"-,extras"`
}

func (r ToolCallParamsRequestContextSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
