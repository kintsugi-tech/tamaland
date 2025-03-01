package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"
	"tamaland/params"
)

// TODO: remove
// queryTargetPlayer queries for the target player's entity ID and health component.
func queryAttackTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Health, error) {
	var playerID types.EntityID
	var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](),
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.Level](),
			filter.Component[comp.State]())).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if playerHealth == nil {
		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerHealth, err
}

func queryRespawnTargetEntity(world cardinal.WorldContext, targetNickname string) (types.EntityID, error) {
	var playerID types.EntityID
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Player](),
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.Level](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})

	if searchErr != nil {
		return 0, searchErr
	}
	if err != nil {
		return 0, err
	}

	return playerID, nil
}

func querySleepTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Energy, *comp.State, error) {
	var playerID types.EntityID
	var playerEnergy *comp.Energy
	var playerState *comp.State
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Player](),
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.Level](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerEnergy, err = cardinal.GetComponent[comp.Energy](world, id)
				if err != nil {
					return false
				}
				playerState, err = cardinal.GetComponent[comp.State](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, nil, err
	}
	if err != nil {
		return 0, nil, nil, err
	}
	if playerEnergy == nil {
		return 0, nil, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerEnergy, playerState, err
}

func queryEatTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Food, *comp.State, error) {
	var playerID types.EntityID
	var playerFood *comp.Food
	var playerState *comp.State
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Player](),
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.Level](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerFood, err = cardinal.GetComponent[comp.Food](world, id)
				if err != nil {
					return false
				}
				playerState, err = cardinal.GetComponent[comp.State](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, nil, err
	}
	if err != nil {
		return 0, nil, nil, err
	}
	if playerFood == nil {
		return 0, nil, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerFood, playerState, err
}

func queryAddLifeTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Health, error) {
	var playerID types.EntityID
	var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Player](),
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.Level](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		)).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if playerHealth == nil {
		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerHealth, err
}

func CanDoActions(currentState comp.State) bool {

	if currentState.State != params.StateDead &&
		currentState.State != params.StateEating &&
		currentState.State != params.StateSleeping {
		return true
	} else {
		return false
	}
}
