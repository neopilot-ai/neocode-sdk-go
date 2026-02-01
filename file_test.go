// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package neocode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/neopilot-ai/neocode-sdk-go"
	"github.com/neopilot-ai/neocode-sdk-go/internal/testutil"
	"github.com/neopilot-ai/neocode-sdk-go/option"
)

func TestFileListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := neocode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.File.List(context.TODO(), neocode.FileListParams{
		Path:      neocode.F("path"),
		Directory: neocode.F("directory"),
	})
	if err != nil {
		var apierr *neocode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileReadWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := neocode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.File.Read(context.TODO(), neocode.FileReadParams{
		Path:      neocode.F("path"),
		Directory: neocode.F("directory"),
	})
	if err != nil {
		var apierr *neocode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileStatusWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := neocode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.File.Status(context.TODO(), neocode.FileStatusParams{
		Directory: neocode.F("directory"),
	})
	if err != nil {
		var apierr *neocode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
