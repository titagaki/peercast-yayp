package worker

import (
	"testing"

	"peercast-yayp/internal/domain"
	"peercast-yayp/internal/peercast"
)

func TestMakeChannelMap(t *testing.T) {
	channels := []*domain.Channel{
		{ID: 1, CID: "aaa"},
		{ID: 2, CID: "bbb"},
	}

	m := makeChannelMap(channels)

	if len(m) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(m))
	}
	if m["aaa"].ID != 1 {
		t.Errorf("expected ID=1 for 'aaa', got %d", m["aaa"].ID)
	}
	if m["bbb"].ID != 2 {
		t.Errorf("expected ID=2 for 'bbb', got %d", m["bbb"].ID)
	}
}

func TestMakeChannelMap_Empty(t *testing.T) {
	m := makeChannelMap(nil)
	if len(m) != 0 {
		t.Errorf("expected empty map, got %d", len(m))
	}
}

func TestFindTracker_TrackerHost(t *testing.T) {
	hosts := []peercast.XMLHost{
		{IP: "1.2.3.4:7144", Tracker: false, Push: false},
		{IP: "5.6.7.8:7144", Tracker: true, Push: false},
	}

	result := findTracker(hosts)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.IP != "5.6.7.8:7144" {
		t.Errorf("expected tracker IP '5.6.7.8:7144', got %q", result.IP)
	}
}

func TestFindTracker_NoTracker_ReturnFirstNonPush(t *testing.T) {
	hosts := []peercast.XMLHost{
		{IP: "1.2.3.4:7144", Tracker: false, Push: true},
		{IP: "5.6.7.8:7144", Tracker: false, Push: false},
		{IP: "9.10.11.12:7144", Tracker: false, Push: false},
	}

	result := findTracker(hosts)
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.IP != "5.6.7.8:7144" {
		t.Errorf("expected IP '5.6.7.8:7144', got %q", result.IP)
	}
}

func TestFindTracker_AllPush(t *testing.T) {
	hosts := []peercast.XMLHost{
		{IP: "1.2.3.4:7144", Push: true},
		{IP: "5.6.7.8:7144", Push: true},
	}

	result := findTracker(hosts)
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestFindTracker_Empty(t *testing.T) {
	result := findTracker(nil)
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestApplyXMLToChannel(t *testing.T) {
	ch := &domain.Channel{}
	xml := peercast.XMLChannel{
		Name:    "test channel",
		ID:      "abc123",
		Bitrate: 1000,
		Type:    "FLV",
		Desc:    "description",
		URL:     "http://example.com",
		Comment: "hello",
		Age:     3600,
	}
	xml.Track.Artist = "artist"
	xml.Track.Title = "title"
	xml.Track.Album = "album"
	xml.Track.Genre = "rock"
	xml.Track.Contact = "http://track.example.com"
	xml.Hits.Listeners = 42
	xml.Hits.Relays = 5

	opts := peercast.StreamOptions{
		Genre:           "ゲーム",
		HiddenListeners: true,
	}
	tracker := &peercast.XMLHost{
		IP:     "1.2.3.4:7144",
		Direct: true,
	}

	applyXMLToChannel(ch, xml, opts, tracker)

	if ch.CID != "abc123" {
		t.Errorf("CID: expected 'abc123', got %q", ch.CID)
	}
	if ch.Name != "test channel" {
		t.Errorf("Name: expected 'test channel', got %q", ch.Name)
	}
	if ch.Bitrate != 1000 {
		t.Errorf("Bitrate: expected 1000, got %d", ch.Bitrate)
	}
	if ch.Listeners != 42 {
		t.Errorf("Listeners: expected 42, got %d", ch.Listeners)
	}
	if ch.Genre != "ゲーム" {
		t.Errorf("Genre: expected 'ゲーム', got %q", ch.Genre)
	}
	if !ch.HiddenListeners {
		t.Error("HiddenListeners: expected true")
	}
	if ch.TrackerIP != "1.2.3.4:7144" {
		t.Errorf("TrackerIP: expected '1.2.3.4:7144', got %q", ch.TrackerIP)
	}
	if !ch.TrackerDirect {
		t.Error("TrackerDirect: expected true")
	}
	if !ch.IsPlaying {
		t.Error("IsPlaying: expected true")
	}
}
