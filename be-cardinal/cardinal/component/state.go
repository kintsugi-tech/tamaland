package component

type StateType string
type State struct {
	State             StateType
	EndStateTimestamp uint64
}

func (State) Name() string {
	return "State"
}
