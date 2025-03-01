package component

type Player struct {
	Nickname    string `json:"nickname"`
	DisplayName string `json:"display_name"`
}

func (Player) Name() string {
	return "Player"
}
