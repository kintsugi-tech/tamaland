package system

import (
	"fmt"
	// "github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"

	comp "tamaland/component"
	"tamaland/msg"
	"tamaland/params"
)

// PlayerSpawnerSystem spawns players based on `CreatePlayer` transactions.
// This provides an example of a system that creates a new entity.
func PlayerSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](
		world,
		func(create cardinal.TxData[msg.CreatePlayerMsg]) (msg.CreatePlayerResult, error) {
			id, err := cardinal.Create(world,
				comp.Player{Nickname: create.Msg.Nickname, DisplayName: create.Msg.DisplayName},
				comp.Level{Lv: params.InitialLevel},
				comp.Health{HP: params.InitialHealth},
				comp.Energy{E: params.InitialEnergy},
				comp.Food{Fd: params.InitialFood},
				comp.State{State: params.InitialState, EndStateTimestamp: params.InitialEndStateTimestamp},
				comp.LastUpdate{Tick: world.CurrentTick(), Timestamp: world.Timestamp(), LevelUpdateTimestamp: world.Timestamp()},
			)
			if err != nil {
				return msg.CreatePlayerResult{}, fmt.Errorf("error creating player: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_player",
				"id":    id,
			})
			if err != nil {
				return msg.CreatePlayerResult{}, err
			}
			return msg.CreatePlayerResult{Success: true}, nil
		})
}

// PlayerRespawnerSystem spawns players based on `CreatePlayer` transactions.
// This deletes the old player entity and creates a new one.
func PlayerRespawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.RespawnPlayerMsg, msg.RespawnPlayerResult](
		world,
		func(create cardinal.TxData[msg.RespawnPlayerMsg]) (msg.RespawnPlayerResult, error) {

			entityId, err := queryRespawnTargetEntity(world, create.Msg.Nickname)
			//log.Debug().Msg(fmt.Sprintf("entityId: %d, err:%s", entityId, err))
			if err != nil {
				return msg.RespawnPlayerResult{Success: false}, fmt.Errorf("error creating player: %w", err)
			}

			err = cardinal.Remove(world, entityId)
			if err != nil {
				return msg.RespawnPlayerResult{Success: false}, fmt.Errorf("error creating player: %w", err)
			}
			id, err := cardinal.Create(world,
				comp.Player{Nickname: create.Msg.Nickname, DisplayName: create.Msg.DisplayName},
				comp.Level{Lv: params.InitialLevel},
				comp.Health{HP: params.InitialHealth},
				comp.Energy{E: params.InitialEnergy},
				comp.Food{Fd: params.InitialFood},
				comp.State{State: params.InitialState, EndStateTimestamp: params.InitialEndStateTimestamp},
				comp.LastUpdate{Tick: world.CurrentTick(), Timestamp: world.Timestamp(), LevelUpdateTimestamp: world.Timestamp()},
			)
			if err != nil {
				return msg.RespawnPlayerResult{Success: false}, fmt.Errorf("error creating player: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_player",
				"id":    id,
			})
			if err != nil {
				return msg.RespawnPlayerResult{Success: false}, err
			}
			return msg.RespawnPlayerResult{Success: true}, nil
		})
}
