package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"

	comp "tamaland/component"
	params "tamaland/params"
)

// SpawnDefaultPlayersSystem creates 3 players with nicknames "Test-[0-9]". This System is registered as an
// Init system, meaning it will be executed exactly one time on tick 0.
func SpawnDefaultPlayersSystem(world cardinal.WorldContext) error {

	namesMap := make(map[int]string)
	namesMap[0] = "alice"
	namesMap[1] = "bob"
	namesMap[2] = "carol"

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("Test-%d", i)
		displayName := namesMap[i]
		_, err := cardinal.Create(world,
			comp.Player{Nickname: name, DisplayName: displayName},
			comp.Level{Lv: params.InitialLevel},
			comp.Health{HP: params.InitialHealth},
			comp.Energy{E: params.InitialEnergy},
			comp.Food{Fd: params.InitialFood},
			comp.State{State: params.InitialState, EndStateTimestamp: params.InitialEndStateTimestamp},
			comp.LastUpdate{Tick: world.CurrentTick(), Timestamp: world.Timestamp()},
		)
		if err != nil {
			return err
		}
	}
	return nil
}
