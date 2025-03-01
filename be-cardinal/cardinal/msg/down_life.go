package msg

// TODO: remove

type AddLifeMsg struct {
	TargetNickname string `json:"target"`
}

type AddLifeMsgReply struct {
	Recovery int `json:"recovery"`
}
