package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerFoodRequest struct {
	Nickname string
}

type PlayerFoodResponse struct {
	Fd int
}

func PlayerFood(world cardinal.WorldContext, req *PlayerFoodRequest) (*PlayerFoodResponse, error) {
	var PlayerFood *comp.Food
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.Food]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				PlayerFood, err = cardinal.GetComponent[comp.Food](world, id)
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

	if PlayerFood == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerFoodResponse{Fd: PlayerFood.Fd}, nil
}
