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

func TestAppGet(t *testing.T) {
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
	_, err := client.App.Get(context.TODO())
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppInit(t *testing.T) {
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
	_, err := client.App.Init(context.TODO())
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppLogWithOptionalParams(t *testing.T) {
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
	_, err := client.App.Log(context.TODO(), githubcomneopilotaineocodesdkgo.AppLogParams{
		Level:   githubcomneopilotaineocodesdkgo.F(githubcomneopilotaineocodesdkgo.AppLogParamsLevelDebug),
		Message: githubcomneopilotaineocodesdkgo.F("message"),
		Service: githubcomneopilotaineocodesdkgo.F("service"),
		Extra: githubcomneopilotaineocodesdkgo.F(map[string]interface{}{
			"foo": "bar",
		}),
	})
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppModes(t *testing.T) {
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
	_, err := client.App.Modes(context.TODO())
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAppProviders(t *testing.T) {
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
	_, err := client.App.Providers(context.TODO())
	if err != nil {
		var apierr *githubcomneopilotaineocodesdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
