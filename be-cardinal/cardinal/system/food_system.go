package system

import (
	//"github.com/rs/zerolog/log"
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/rs/zerolog/log"

	comp "tamaland/component"
	"tamaland/params"
)

// FoodSystem decrease the player's Fd by 1 every 10 ticks.
// This provides an example of a system that doesn't rely on a transaction to update a component.
func FoodSystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Food](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {
			log.Debug().Msg("FoodSystem")

			// go to the next one if not time to update.
			lastUpdate, err := cardinal.GetComponent[comp.LastUpdate](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error1: " + err.Error())
				return true
			}
			if lastUpdate.Timestamp+params.StatUpdateIntervalMs > world.Timestamp() {
				return true
			}

			food, err := cardinal.GetComponent[comp.Food](world, id)
			if err != nil {
				return true
			}
			state, err := cardinal.GetComponent[comp.State](world, id)
			if err != nil {
				log.Error().Msg("FoodSystem: error2: " + err.Error())
				return true
			}

			foodToAdd := 0
			if state.State == params.StateEating && food.Fd < params.MaxFood {
				foodToAdd = 5
			} else if state.State != params.StateEating && food.Fd > params.MinFood {
				foodToAdd = -1
			}

			// Debug purposes
			skipCause := ""
			if state.State == params.StateEating && food.Fd >= params.MaxFood {
				skipCause = fmt.Sprintf("entity %d is eating but has full food (%d)", id, food.Fd)
			}
			if state.State != params.StateEating && food.Fd <= params.MinFood {
				skipCause = fmt.Sprintf("entity %d is not eating but has minimum food (%d)", id, food.Fd)
			}

			food.Fd += foodToAdd

			if food.Fd > params.MaxFood {
				food.Fd = params.MaxFood
			}
			if food.Fd < params.MinFood {
				food.Fd = params.MinFood
			}

			if err := cardinal.SetComponent[comp.Food](world, id, food); err != nil {
				log.Error().Msg("FoodSystem: error3: " + err.Error())
				return true
			}

			log.Debug().Msg(fmt.Sprintf("FoodSystem: skipped: %s", skipCause))
			return true
		},
	)
	if err != nil {
		return err
	}
	return nil
}
