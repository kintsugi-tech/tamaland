package msg

type EatMsg struct {
	TargetNickname string `json:"target"`
}

type EatMsgReply struct {
	Success bool `json:"success"`
}
