package box

const (
	HIDDEN = iota
	VISIBLE
	FLAGGED

	EMPTY = iota
	BOMB
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
		State:    HIDDEN,
		Value:    EMPTY,
		Index:    index,
	}
}

func (b *Box) AsSymbol() string {
	switch b.State {
	case HIDDEN:
		return "?"

	case VISIBLE:
		if b.Value == BOMB {
			return "@"
		}

		return "0"

	case FLAGGED:
		return "X"

	default:
		return ""
	}
}
