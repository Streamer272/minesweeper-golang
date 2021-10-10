package box

const (
	Hidden = iota
	Visible

	Empty = iota
	Bomb
)

type Box struct {
	Selected bool
	State    int
	Value    int
	Index    int
}

func NewBox(index int, selected bool) Box {
	return Box{
		Selected: selected,
		State:    Hidden,
		Value:    Empty,
		Index:    index,
	}
}

func (b *Box) AsSymbol() string {
	switch b.State {
	case Hidden:
		return "?"

	case Visible:
		if b.Value == Bomb {
			return "X"
		}

		// TODO
		return "0"

	default:
		return ""
	}
}
