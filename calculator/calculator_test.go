package calculator

import (
	"encoding/json"
	"testing"
)

func TestCalculateWeights_60(t *testing.T) {
	// 60 - 45 = 15 lbs remaining => 1x5 lb and 1x2.5 lb per side
	got := CalculateWeights(65)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got10 := m["10"]; got10 != 1 {
		t.Fatalf("expected '10' == 1, got %d; full: %v", got10, m)
	}
}

func TestCalculateWeights_45(t *testing.T) {
	// 45 - 45 = 0 lbs remaining => no plates needed
	got := CalculateWeights(45)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got45 := m["45"]; got45 != 0 {
		t.Fatalf("expected '45' == 0, got %d; full: %v", got45, m)
	}
}

func TestCalculateWeights_135(t *testing.T) {
	// 135 - 45 = 90 lbs remaining => 1x45 lb per side
	got := CalculateWeights(135)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got45 := m["45"]; got45 != 1 {
		t.Fatalf("expected '45' == 1, got %d; full: %v", got45, m)
	}
}

func TestCalculateWeights_90(t *testing.T) {
	// 90 - 45 = 45 lbs remaining => 2x20 lb per side => 1x2.5 lb per side
	got := CalculateWeights(90)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got10 := m["10"]; got10 != 2 {
		t.Fatalf("expected '10' == 2, got %d; full: %v", got10, m)
	}
	if got2pt5 := m["2pt5"]; got2pt5 != 1 {
		t.Fatalf("expected '2pt5' == 1, got %d; full: %v", got2pt5, m)
	}
}

func TestCalculateWeights_225(t *testing.T) {
	// 225 - 45 = 180 lbs remaining => 2x45 lb per side
	got := CalculateWeights(225)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got45 := m["45"]; got45 != 2 {
		t.Fatalf("expected '45' == 2, got %d; full: %v", got45, m)
	}
}

func TestCalculateWeights_10(t *testing.T) {
	// 10 - 45 = -35 lbs remaining => no plates needed (weight too low)
	got := CalculateWeights(10)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got45 := m["45"]; got45 != 0 {
		t.Fatalf("expected '45' == 0, got %d; full: %v", got45, m)
	}
}

func TestCalculateWeights_neg10(t *testing.T) {
	// -10 - 45 = -55 lbs remaining => no plates needed (weight too low)
	got := CalculateWeights(-10)

	b, err := json.Marshal(got)
	if err != nil {
		t.Fatalf("failed to marshal result: %v", err)
	}

	var m map[string]int
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("failed to unmarshal result into map: %v", err)
	}

	if got45 := m["45"]; got45 != 0 {
		t.Fatalf("expected '45' == 0, got %d; full: %v", got45, m)
	}
}
