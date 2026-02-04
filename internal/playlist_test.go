package spotify

import (
	"testing"
)

func TestListTracks(t *testing.T) {
	mockTracks := []*Track{NewTrack("track one"), NewTrack("track two")}
	playlist := NewPlaylist("test", mockTracks)

	trackString := playlist.ListTracks()

	expected := "track one, track two"

	if trackString != expected {
		t.Errorf("Got '%s' but expected '%s'", trackString, expected)
	}
}
