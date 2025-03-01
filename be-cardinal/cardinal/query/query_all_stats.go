package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerStatsRequest struct {
	Nickname string
}

type PlayerStatsResponse struct {
	E  int
	F  int
	H  int
	L  int
	S  comp.State
	LU comp.LastUpdate
}

func PlayerStats(world cardinal.WorldContext, req *PlayerStatsRequest) (*PlayerStatsResponse, error) {
	var playerEnergy *comp.Energy
	var playerLastUpdate *comp.LastUpdate
	var playerState *comp.State
	var playerFood *comp.Food
	var playerHealth *comp.Health
	var playerLevel *comp.Level
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerEnergy, err = cardinal.GetComponent[comp.Energy](world, id)
				if err != nil {
					return false
				}
				playerLastUpdate, err = cardinal.GetComponent[comp.LastUpdate](world, id)
				if err != nil {
					return false
				}
				playerState, err = cardinal.GetComponent[comp.State](world, id)
				if err != nil {
					return false
				}
				playerFood, err = cardinal.GetComponent[comp.Food](world, id)
				if err != nil {
					return false
				}
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				playerLevel, err = cardinal.GetComponent[comp.Level](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if playerEnergy == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerStatsResponse{E: playerEnergy.E, F: playerFood.Fd, H: playerHealth.HP, S: *playerState, LU: *playerLastUpdate, L: *&playerLevel.Lv}, nil
}
