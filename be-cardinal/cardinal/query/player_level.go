package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerLevelRequest struct {
	Nickname string
}

type PlayerLevelResponse struct {
	Lv int
}

func PlayerLevel(world cardinal.WorldContext, req *PlayerLevelRequest) (*PlayerLevelResponse, error) {
	var playerLevel *comp.Level
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Contains(filter.Component[comp.Player](), filter.Component[comp.Level]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
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

	if playerLevel == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerLevelResponse{Lv: playerLevel.Lv}, nil
}
