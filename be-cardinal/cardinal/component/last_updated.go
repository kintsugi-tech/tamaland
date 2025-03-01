package component

type LastUpdate struct {
	Tick                 uint64 //unused
	Timestamp            uint64
	LevelUpdateTimestamp uint64
}

func (LastUpdate) Name() string {
	return "LastUpdate"
}
