package peercast

import "testing"

func TestParseGenre_Basic(t *testing.T) {
	opts, ok := ParseGenre("apゲーム", "ap")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Prefix != "ap" {
		t.Errorf("expected prefix 'ap', got %q", opts.Prefix)
	}
	if opts.Genre != "ゲーム" {
		t.Errorf("expected genre 'ゲーム', got %q", opts.Genre)
	}
	if opts.HiddenListeners {
		t.Error("expected HiddenListeners=false")
	}
	if opts.Restriction != RestrictionNone {
		t.Errorf("expected restriction None, got %d", opts.Restriction)
	}
}

func TestParseGenre_WrongPrefix(t *testing.T) {
	_, ok := ParseGenre("spゲーム", "ap")
	if ok {
		t.Error("expected ok=false for wrong prefix")
	}
}

func TestParseGenre_EmptyString(t *testing.T) {
	_, ok := ParseGenre("", "ap")
	if ok {
		t.Error("expected ok=false for empty string")
	}
}

func TestParseGenre_HiddenListeners(t *testing.T) {
	opts, ok := ParseGenre("ap?ゲーム", "ap")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if !opts.HiddenListeners {
		t.Error("expected HiddenListeners=true")
	}
	if opts.Genre != "ゲーム" {
		t.Errorf("expected genre 'ゲーム', got %q", opts.Genre)
	}
}

func TestParseGenre_Restriction(t *testing.T) {
	tests := []struct {
		input       string
		restriction int
	}{
		{"ap@music", RestrictionFirewalled},
		{"ap@@music", RestrictionBandwidth},
		{"ap@@@music", RestrictionHighBandwidth},
		{"ap@@@@music", RestrictionHighBandwidth},
	}

	for _, tt := range tests {
		opts, ok := ParseGenre(tt.input, "ap")
		if !ok {
			t.Fatalf("input %q: expected ok=true", tt.input)
		}
		if opts.Restriction != tt.restriction {
			t.Errorf("input %q: expected restriction %d, got %d", tt.input, tt.restriction, opts.Restriction)
		}
	}
}

func TestParseGenre_Namespace(t *testing.T) {
	opts, ok := ParseGenre("apns:ゲーム", "ap")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Namespace != "ns" {
		t.Errorf("expected namespace 'ns', got %q", opts.Namespace)
	}
	if opts.Genre != "ゲーム" {
		t.Errorf("expected genre 'ゲーム', got %q", opts.Genre)
	}
}

func TestParseGenre_HiddenAndRestriction(t *testing.T) {
	opts, ok := ParseGenre("ap?@ゲーム", "ap")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if !opts.HiddenListeners {
		t.Error("expected HiddenListeners=true")
	}
	if opts.Restriction != RestrictionFirewalled {
		t.Errorf("expected restriction Firewalled, got %d", opts.Restriction)
	}
	if opts.Genre != "ゲーム" {
		t.Errorf("expected genre 'ゲーム', got %q", opts.Genre)
	}
}

func TestParseGenre_PrefixOnly(t *testing.T) {
	opts, ok := ParseGenre("ap", "ap")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if opts.Genre != "" {
		t.Errorf("expected empty genre, got %q", opts.Genre)
	}
}
