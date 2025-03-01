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

// LastUpdateSystem handles last-update update.
// It is the last-one called.
// Only the levelup sets its update in the level_system.
func LastUpdateSystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.LastUpdate](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {

			log.Debug().Msg("LastupdateSystem")

			// go to the next one if not time to update.
			lastUpdate, err := cardinal.GetComponent[comp.LastUpdate](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error1: " + err.Error())
				return true
			}
			if lastUpdate.Timestamp+params.StatUpdateIntervalMs > world.Timestamp() {
				return true
			}

			lastUpdate.Timestamp = world.Timestamp()
			lastUpdate.Tick = world.CurrentTick()
			if err := cardinal.SetComponent[comp.LastUpdate](world, id, lastUpdate); err != nil {
				log.Error().Msg("EnergySystem: error3: " + err.Error())
				return true
			}

			return true
		},
	)
	if err != nil {
		return err
	}
	return nil
}
