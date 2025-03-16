// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oxp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/OpenExecProtocol/oxp-go"
	"github.com/OpenExecProtocol/oxp-go/internal/testutil"
	"github.com/OpenExecProtocol/oxp-go/option"
)

func TestToolList(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oxp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Tools.List(context.TODO(), oxp.ToolListParams{})
	if err != nil {
		var apierr *oxp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolCallWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := oxp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Tools.Call(context.TODO(), oxp.ToolCallParams{
		Request: oxp.F(oxp.ToolCallParamsRequest{
			ToolID: oxp.F("sqFnKL1N_jr_.U0_2_jv__a"),
			CallID: oxp.F("call_id"),
			Context: oxp.F(oxp.ToolCallParamsRequestContext{
				Authorization: oxp.F([]oxp.ToolCallParamsRequestContextAuthorization{{
					ID:    oxp.F("id"),
					Token: oxp.F("token"),
				}}),
				Secrets: oxp.F([]oxp.ToolCallParamsRequestContextSecret{{
					ID:    oxp.F("id"),
					Value: oxp.F("value"),
				}}),
				UserID: oxp.F("user_id"),
			}),
			Input: oxp.F(map[string]interface{}{
				"foo": "bar",
			}),
			TraceID: oxp.F("trace_id"),
		}),
		Schema: oxp.F("https://example.com"),
	})
	if err != nil {
		var apierr *oxp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
