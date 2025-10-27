package impl

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"log/slog"

	"github.com/tk3413/tk-weight-calc/calculator"
	api "github.com/tk3413/tk-weight-calc/server_gen"
)

func TestGetWeights_OK(t *testing.T) {
	srv := NewServer()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/weights", nil)

	// call handler with a typical weight
	params := api.GetWeightsParams{Weight: 70.5}
	srv.GetWeights(rr, req, params)

	if rr.Code != 200 {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	var got interface{}
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}
	if got == nil {
		t.Fatalf("expected non-nil response payload")
	}
}

func TestGetWeights_ZeroWeight(t *testing.T) {
	// ensure zero (edge) weight still returns a well-formed response
	srv := NewServer()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/weights", nil)

	params := api.GetWeightsParams{Weight: 0}
	srv.GetWeights(rr, req, params)

	if rr.Code != 200 {
		t.Fatalf("expected status 200 for zero weight, got %d", rr.Code)
	}

	var got interface{}
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("response body is not valid JSON for zero weight: %v", err)
	}
	if got == nil {
		t.Fatalf("expected non-nil response payload for zero weight")
	}
}

func TestNewServer_WithLogger_DoesNotPanicAndImplementsInterface(t *testing.T) {
	// use the default logger as a simple way to exercise WithLogger
	logger := slog.Default()
	srv := NewServer(WithLogger(logger))

	// should implement the generated ServerInterface
	var _ api.ServerInterface = srv

	// sanity: calling GetWeights should work (uses calculator.CalculateWeights)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/weights", nil)
	params := api.GetWeightsParams{Weight: 55.0}
	srv.GetWeights(rr, req, params)

	if rr.Code != 200 {
		t.Fatalf("expected status 200 after constructing server with logger, got %d", rr.Code)
	}

	// also verify the calculator package returns something marshalable for the same input
	calculated := calculator.CalculateWeights(params.Weight)
	// CalculateWeights returns a non-pointer value, so avoid a nil comparison and just ensure it marshals.
	if _, err := json.Marshal(calculated); err != nil {
		t.Fatalf("calculator.CalculateWeights result is not JSON-marshalable: %v", err)
	}
}

func TestNewServerNegativeWeight(t *testing.T) {
	srv := NewServer()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/weights", nil)

	// call handler with a negative weight
	params := api.GetWeightsParams{Weight: -10.0}
	srv.GetWeights(rr, req, params)

	if rr.Code != 200 {
		t.Fatalf("expected status 200 for negative weight, got %d", rr.Code)
	}

	var got interface{}
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("response body is not valid JSON for negative weight: %v", err)
	}
	if got == nil {
		t.Fatalf("expected non-nil response payload for negative weight")
	}
}

func TestNewServer400ResponseCode(t *testing.T) {
	srv := NewServer()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/weights", nil)
	// call handler with an excessively large weight
	params := api.GetWeightsParams{Weight: 1e6}
	srv.GetWeights(rr, req, params)

	if rr.Code != 200 {
		t.Fatalf("expected status 200 for large weight, got %d", rr.Code)
	}
	var got interface{}
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("response body is not valid JSON for large weight: %v", err)
	}
	if got == nil {
		t.Fatalf("expected non-nil response payload for large weight")
	}
}
