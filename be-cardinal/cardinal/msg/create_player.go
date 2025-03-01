package msg

type StateStruct struct {
	State             string `json:"state"`
	EndStateTimestamp int    `json:"end_state"`
}

type CreatePlayerMsg struct {
	Nickname    string `json:"nickname"`
	DisplayName string `json:"display_name"`
}

type CreatePlayerResult struct {
	Success bool `json:"success"`
}
