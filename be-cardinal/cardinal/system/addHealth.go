package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "tamaland/component"
	"tamaland/msg"
)

const HealthRecover = 1

// SleepSystem recover energy to player's E based on `SleepPlayer` transactions.
func HealthDown(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.AddLifeMsg, msg.AddLifeMsgReply](
		world,
		func(addLife cardinal.TxData[msg.AddLifeMsg]) (msg.AddLifeMsgReply, error) {
			playerID, playerHealth, err := queryAddLifeTargetPlayer(world, addLife.Msg.TargetNickname)
			if err != nil {
				return msg.AddLifeMsgReply{}, fmt.Errorf("failed to recover energy: %w", err)
			}

			playerHealth.HP -= HealthRecover
			if err := cardinal.SetComponent[comp.Health](world, playerID, playerHealth); err != nil {
				return msg.AddLifeMsgReply{}, fmt.Errorf("failed to recover energy: %w", err)
			}

			return msg.AddLifeMsgReply{Recovery: HealthRecover}, nil
		})
}
