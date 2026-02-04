package spotify

type searchPort interface {
	search(string) []*Playlist
}

type PlaylistService struct {
	searchClient searchPort
}

func NewPlaylistService(client searchPort) *PlaylistService {
	return &PlaylistService{searchClient: client}
}

func (ps *PlaylistService) Search(searchTerm string) []*Playlist {
	return ps.searchClient.search(searchTerm)
}
