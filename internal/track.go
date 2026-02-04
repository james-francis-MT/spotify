package spotify

type Track struct {
	name string
}

func NewTrack(name string) *Track {
	return &Track{name: name}
}

func (t *Track) GetFormattedName() string {
	return t.name
}
