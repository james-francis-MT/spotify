package spotify

import (
	"testing"
)

type testClient struct{}

func (tc *testClient) search(_ string) []*Playlist {
	mockTracks := []*Track{NewTrack("track one"), NewTrack("track two")}
	return []*Playlist{NewPlaylist("test playlist", mockTracks)}
}

func TestGetPlaylist(t *testing.T) {
	service := NewPlaylistService(&testClient{})

	playlists := service.Search("search term")

	if len(playlists) == 0 {
		t.Error("Expected search to returns playlists")
	}

	expectedPlaylistName := "test playlist"
	if playlists[0].name != expectedPlaylistName {
		t.Errorf("Got '%s' but expected '%s'", playlists[0].name, expectedPlaylistName)
	}
}
