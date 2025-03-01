package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerStateRequest struct {
	Nickname string
}

type PlayerStateResponse struct {
	State comp.StateType
}

func PlayerState(world cardinal.WorldContext, req *PlayerStateRequest) (*PlayerStateResponse, error) {
	var playerState *comp.State
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.State]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
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
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if playerState == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerStateResponse{State: playerState.State}, nil
}
