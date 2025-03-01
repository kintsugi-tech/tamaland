package msg

type SleepMsg struct {
	TargetNickname string `json:"target"`
}

type SleepMsgReply struct {
	Success bool `json:"success"`
}
