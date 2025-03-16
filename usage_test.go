// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package oxp_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/oxp-go"
	"github.com/stainless-sdks/oxp-go/internal/testutil"
	"github.com/stainless-sdks/oxp-go/option"
)

func TestUsage(t *testing.T) {
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
	tool, err := client.Tools.List(context.TODO(), oxp.ToolListParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", tool.Items)
}
