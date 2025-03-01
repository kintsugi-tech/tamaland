package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerNicknameRequest struct {
	Nickname string
}

type PlayerNicknameResponse struct {
	Nickname string
}

func PlayerNickname(world cardinal.WorldContext, req *PlayerNicknameRequest) (*PlayerNicknameResponse, error) {
	var playerNickname *comp.Player
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
				playerNickname, err = cardinal.GetComponent[comp.Player](world, id)
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

	if playerNickname == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerNicknameResponse{Nickname: playerNickname.Nickname}, nil
}
