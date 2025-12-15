// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package userplex_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/userplex-go"
	"github.com/stainless-sdks/userplex-go/internal/testutil"
	"github.com/stainless-sdks/userplex-go/option"
)

func TestUserIdentifyWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := userplex.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Users.Identify(context.TODO(), userplex.UserIdentifyParams{
		UserID: "user_id",
		Email:  userplex.String("dev@stainless.com"),
		Name:   userplex.String("name"),
		Properties: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *userplex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
