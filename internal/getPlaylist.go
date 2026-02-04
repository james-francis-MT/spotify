package spotify

type PlaylistSearcher interface {
	Search(string) []*Playlist
}

type PlaylistService struct {
	searchClient PlaylistSearcher
}

func NewPlaylistService(client PlaylistSearcher) *PlaylistService {
	return &PlaylistService{searchClient: client}
}

func (ps *PlaylistService) Search(searchTerm string) []*Playlist {
	return ps.searchClient.Search(searchTerm)
}
