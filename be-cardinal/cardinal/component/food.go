package component

type Food struct {
	Fd int
}

func (Food) Name() string {
	return "Food"
}
