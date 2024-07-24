package api_test

import (
	"testing"

	"rayjay.com/go/calc/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
