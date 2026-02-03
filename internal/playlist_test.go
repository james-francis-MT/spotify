package spotify

import (
	"testing"
)

func TestListTracks(t *testing.T) {
	playlist := NewPlaylist("test", []*Track{NewTrack("track one"), NewTrack("track two")})

	trackString := playlist.ListTracks()

	expected := "track one, track two"

	if trackString != expected {
		t.Errorf("Got '%s' but expected '%s'", trackString, expected)
	}
}
