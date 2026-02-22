package handler

import "testing"

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		seconds uint
		want    string
	}{
		{0, "0:00"},
		{60, "0:01"},
		{3600, "1:00"},
		{3661, "1:01"},
		{7200, "2:00"},
		{86400, "24:00"},
		{90061, "25:01"},
	}

	for _, tt := range tests {
		got := formatDuration(tt.seconds)
		if got != tt.want {
			t.Errorf("formatDuration(%d) = %q, want %q", tt.seconds, got, tt.want)
		}
	}
}

func TestBoolToDigit(t *testing.T) {
	if boolToDigit(true) != "1" {
		t.Error("expected '1' for true")
	}
	if boolToDigit(false) != "0" {
		t.Error("expected '0' for false")
	}
}
