package system

import (
	//"github.com/rs/zerolog/log"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	"github.com/rs/zerolog/log"

	comp "tamaland/component"
	"tamaland/params"
)

// LevelSystem increases the player's level during time.
func LevelSystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Level](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {
			log.Debug().Msg("LevelSystem")

			// go to the next one if not time to update.
			lastUpdate, err := cardinal.GetComponent[comp.LastUpdate](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error1: " + err.Error())
				return true
			}
			if lastUpdate.LevelUpdateTimestamp+params.LevelUpIntervalMs > world.Timestamp() {
				return true
			}

			level, err := cardinal.GetComponent[comp.Level](world, id)
			if err != nil {
				return true
			}
			state, err := cardinal.GetComponent[comp.State](world, id)
			if err != nil {
				log.Error().Msg("LevelSystem: error1: " + err.Error())
				return true
			}

			if state.State != params.StateDead {
				level.Lv += 1
				if err := cardinal.SetComponent[comp.Level](world, id, level); err != nil {
					log.Error().Msg("LevelSystem: error2: " + err.Error())
					return true
				}

				lastUpdate.LevelUpdateTimestamp = world.Timestamp()
				if err := cardinal.SetComponent[comp.LastUpdate](world, id, lastUpdate); err != nil {
					log.Error().Msg("LevelSystem: error2: " + err.Error())
					return true
				}
			}

			return true
		},
	)
	if err != nil {
		return err
	}
	return nil
}
