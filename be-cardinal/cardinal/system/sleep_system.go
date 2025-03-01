package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "tamaland/component"
	"tamaland/msg"
	"tamaland/params"
)

// SleepSystem sets the player to the sleeping state.
func SleepSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.SleepMsg, msg.SleepMsgReply](
		world,
		func(sleep cardinal.TxData[msg.SleepMsg]) (msg.SleepMsgReply, error) {
			playerID, _, playerState, err := querySleepTargetPlayer(world, sleep.Msg.TargetNickname)
			if err != nil {
				return msg.SleepMsgReply{Success: false}, fmt.Errorf("failed to recover target: %w", err)
			}

			if !CanDoActions(*playerState) {
				return msg.SleepMsgReply{Success: false}, fmt.Errorf("failed to set player state: current state is %s", playerState.State)
			}

			playerState.State = params.StateSleeping
			playerState.EndStateTimestamp = world.Timestamp() + params.SleepDurationMs
			if err := cardinal.SetComponent[comp.State](world, playerID, playerState); err != nil {
				return msg.SleepMsgReply{Success: false}, fmt.Errorf("failed to set player state: %w", err)
			}

			return msg.SleepMsgReply{Success: false}, nil
		})
}
