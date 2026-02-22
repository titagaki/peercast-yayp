package handler

import (
	"testing"
	"time"
)

func TestParseDate_Valid(t *testing.T) {
	date, err := parseDate("20231225")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if date.Year() != 2023 || date.Month() != time.December || date.Day() != 25 {
		t.Errorf("expected 2023-12-25, got %v", date)
	}
}

func TestParseDate_InvalidLength(t *testing.T) {
	_, err := parseDate("2023")
	if err == nil {
		t.Error("expected error for short string")
	}
}

func TestParseDate_InvalidChars(t *testing.T) {
	_, err := parseDate("2023abcd")
	if err == nil {
		t.Error("expected error for non-numeric characters")
	}
}

func TestParseDate_Empty(t *testing.T) {
	_, err := parseDate("")
	if err == nil {
		t.Error("expected error for empty string")
	}
}

func TestParseDate_Boundary(t *testing.T) {
	date, err := parseDate("20240101")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if date.Year() != 2024 || date.Month() != time.January || date.Day() != 1 {
		t.Errorf("expected 2024-01-01, got %v", date)
	}
}
