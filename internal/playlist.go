package spotify

import "strings"

type Playlist struct {
	id          string
	name        string
	description string
	owner       string
	tracks      []*Track
}

func NewPlaylist(name string, tracks []*Track) *Playlist {
	return &Playlist{name: name, tracks: tracks}
}

func (p *Playlist) GetName() string {
	return p.name
}

func (p *Playlist) ListTracks() string {
	tracks := make([]string, 0, len(p.tracks))
	for _, t := range p.tracks {
		tracks = append(tracks, t.name)
	}

	return strings.Join(tracks, ", ")
}
