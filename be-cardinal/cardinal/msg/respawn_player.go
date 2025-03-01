package msg

type RespawnPlayerMsg struct {
	Nickname    string `json:"nickname"`
	DisplayName string `json:"display_name"`
}

type RespawnPlayerResult struct {
	Success bool `json:"success"`
}
