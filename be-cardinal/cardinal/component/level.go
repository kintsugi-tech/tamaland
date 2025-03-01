package component

type Level struct {
	Lv int
}

func (Level) Name() string {
	return "Level"
}
