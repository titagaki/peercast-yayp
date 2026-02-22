package worker

import (
	"testing"
	"time"

	"peercast-yayp/internal/domain"
)

func TestAggregateLogs_Empty(t *testing.T) {
	result := aggregateLogs(time.Now(), nil)
	if len(result) != 0 {
		t.Errorf("expected empty, got %d", len(result))
	}
}

func TestAggregateLogs_SingleChannel(t *testing.T) {
	logs := []*domain.ChannelLog{
		{Name: "ch1", Listeners: 10},
		{Name: "ch1", Listeners: 20},
		{Name: "ch1", Listeners: 30},
	}

	date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	result := aggregateLogs(date, logs)

	if len(result) != 1 {
		t.Fatalf("expected 1 summary, got %d", len(result))
	}

	s := result[0]
	if s.Name != "ch1" {
		t.Errorf("expected name 'ch1', got %q", s.Name)
	}
	if s.NumLogs != 3 {
		t.Errorf("expected 3 logs, got %d", s.NumLogs)
	}
	if s.MaxListeners != 30 {
		t.Errorf("expected max 30, got %d", s.MaxListeners)
	}
	if s.AverageListeners != 20.0 {
		t.Errorf("expected avg 20.0, got %f", s.AverageListeners)
	}
	if !s.LogDate.Equal(date) {
		t.Errorf("expected date %v, got %v", date, s.LogDate)
	}
}

func TestAggregateLogs_MultipleChannels(t *testing.T) {
	logs := []*domain.ChannelLog{
		{Name: "ch1", Listeners: 10},
		{Name: "ch2", Listeners: 100},
		{Name: "ch1", Listeners: 20},
		{Name: "ch2", Listeners: 200},
	}

	result := aggregateLogs(time.Now(), logs)

	if len(result) != 2 {
		t.Fatalf("expected 2 summaries, got %d", len(result))
	}

	if result[0].Name != "ch1" {
		t.Errorf("expected first 'ch1', got %q", result[0].Name)
	}
	if result[1].Name != "ch2" {
		t.Errorf("expected second 'ch2', got %q", result[1].Name)
	}

	if result[0].NumLogs != 2 || result[0].MaxListeners != 20 {
		t.Errorf("ch1: expected NumLogs=2 Max=20, got NumLogs=%d Max=%d", result[0].NumLogs, result[0].MaxListeners)
	}
	if result[0].AverageListeners != 15.0 {
		t.Errorf("ch1: expected avg 15.0, got %f", result[0].AverageListeners)
	}

	if result[1].NumLogs != 2 || result[1].MaxListeners != 200 {
		t.Errorf("ch2: expected NumLogs=2 Max=200, got NumLogs=%d Max=%d", result[1].NumLogs, result[1].MaxListeners)
	}
	if result[1].AverageListeners != 150.0 {
		t.Errorf("ch2: expected avg 150.0, got %f", result[1].AverageListeners)
	}
}
