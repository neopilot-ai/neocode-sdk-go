// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomneopilotaineocodesdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/neopilot-ai/neocode-sdk-go"
	"github.com/neopilot-ai/neocode-sdk-go/internal/testutil"
	"github.com/neopilot-ai/neocode-sdk-go/option"
)

func TestFileRead(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomneopilotaineocodesdkgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.File.Read(context.TODO(), githubcomneopilotaineocodesdkgo.FileReadParams{
		Path: githubcomneopilotaineocodesdkgo.F("path"),
	})
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomneopilotaineocodesdkgo.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.File.Status(context.TODO())
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
