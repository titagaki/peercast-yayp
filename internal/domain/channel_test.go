package domain

import "testing"

func TestMaskListeners(t *testing.T) {
	channels := []*Channel{
		{Name: "ch1", Listeners: 10, Relays: 5, HiddenListeners: false},
		{Name: "ch2", Listeners: 20, Relays: 8, HiddenListeners: true},
		{Name: "ch3", Listeners: 0, Relays: 0, HiddenListeners: false},
	}

	result := MaskListeners(channels)

	if len(result) != 3 {
		t.Fatalf("expected 3 channels, got %d", len(result))
	}

	if result[0].Listeners != 10 || result[0].Relays != 5 {
		t.Errorf("ch1: expected Listeners=10, Relays=5, got Listeners=%d, Relays=%d", result[0].Listeners, result[0].Relays)
	}

	if result[1].Listeners != -1 || result[1].Relays != -1 {
		t.Errorf("ch2: expected Listeners=-1, Relays=-1, got Listeners=%d, Relays=%d", result[1].Listeners, result[1].Relays)
	}

	if result[2].Listeners != 0 || result[2].Relays != 0 {
		t.Errorf("ch3: expected Listeners=0, Relays=0, got Listeners=%d, Relays=%d", result[2].Listeners, result[2].Relays)
	}
}

func TestMaskListeners_DoesNotMutateOriginal(t *testing.T) {
	original := &Channel{Listeners: 10, Relays: 5, HiddenListeners: true}
	channels := []*Channel{original}

	result := MaskListeners(channels)

	if result[0].Listeners != -1 {
		t.Errorf("masked copy: expected -1, got %d", result[0].Listeners)
	}
	if original.Listeners != 10 {
		t.Errorf("original mutated: expected 10, got %d", original.Listeners)
	}
}

func TestMaskListeners_Empty(t *testing.T) {
	result := MaskListeners(nil)
	if len(result) != 0 {
		t.Errorf("expected empty slice, got %d", len(result))
	}
}
