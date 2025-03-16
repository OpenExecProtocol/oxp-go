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

func TestHealthCheck(t *testing.T) {
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
	err := client.Health.Check(context.TODO())
	if err != nil {
		var apierr *oxp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
