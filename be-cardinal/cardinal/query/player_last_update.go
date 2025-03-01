package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerLastUpdateRequest struct {
	Nickname string
}

type PlayerLastUpdateResponse struct {
	Tick      uint64
	Timestamp uint64
}

func PlayerLastUpdate(world cardinal.WorldContext, req *PlayerLastUpdateRequest) (*PlayerLastUpdateResponse, error) {
	var playerLastUpdate *comp.LastUpdate
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.LastUpdate]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerLastUpdate, err = cardinal.GetComponent[comp.LastUpdate](world, id)
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

	if playerLastUpdate == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerLastUpdateResponse{Tick: playerLastUpdate.Tick, Timestamp: playerLastUpdate.Timestamp}, nil
}
