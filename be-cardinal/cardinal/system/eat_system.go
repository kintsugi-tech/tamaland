package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "tamaland/component"
	"tamaland/msg"
	"tamaland/params"
)

// EatSystem sets the player to the eating state.
func EatSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.EatMsg, msg.EatMsgReply](
		world,
		func(eat cardinal.TxData[msg.EatMsg]) (msg.EatMsgReply, error) {
			playerID, _, playerState, err := queryEatTargetPlayer(world, eat.Msg.TargetNickname)
			if err != nil {
				return msg.EatMsgReply{Success: false}, fmt.Errorf("failed to recover target: %w", err)
			}

			if !CanDoActions(*playerState) {
				return msg.EatMsgReply{Success: false}, fmt.Errorf("failed to set player state: current state is %s", playerState.State)
			}

			playerState.State = params.StateEating
			playerState.EndStateTimestamp = world.Timestamp() + params.EatingDurationMs
			if err := cardinal.SetComponent[comp.State](world, playerID, playerState); err != nil {
				return msg.EatMsgReply{Success: false}, fmt.Errorf("failed to set player state: %w", err)
			}

			return msg.EatMsgReply{Success: true}, nil
		})
}
