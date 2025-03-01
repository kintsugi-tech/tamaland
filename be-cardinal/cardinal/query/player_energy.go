package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerEnergyRequest struct {
	Nickname string
}

type PlayerEnergyResponse struct {
	E int
}

func PlayerEnergy(world cardinal.WorldContext, req *PlayerEnergyRequest) (*PlayerEnergyResponse, error) {
	var playerEnergy *comp.Energy
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.Energy]())).
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

	return &PlayerEnergyResponse{E: playerEnergy.E}, nil
}
