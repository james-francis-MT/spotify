package spotify

type Track struct {
	name string
}

func NewTrack(name string) *Track {
	return &Track{name: name}
}
