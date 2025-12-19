// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package userplex_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/dqnamo/userplex-go"
	"github.com/dqnamo/userplex-go/internal/testutil"
	"github.com/dqnamo/userplex-go/option"
)

func TestLogBatch(t *testing.T) {
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
	_, err := client.Logs.Batch(context.TODO(), userplex.LogBatchParams{
		Logs: []userplex.LogBatchParamsLog{{
			Name:   "name",
			UserID: "user_id",
			Data: map[string]any{
				"foo": "bar",
			},
			Properties: map[string]any{
				"foo": "bar",
			},
			Timestamp: userplex.Time(time.Now()),
		}},
	})
	if err != nil {
		var apierr *userplex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.New(context.TODO(), userplex.LogNewParams{
		Name:   "name",
		UserID: "user_id",
		Data: map[string]any{
			"foo": "bar",
		},
		Properties: map[string]any{
			"foo": "bar",
		},
		Timestamp: userplex.Time(time.Now()),
	})
	if err != nil {
		var apierr *userplex.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
