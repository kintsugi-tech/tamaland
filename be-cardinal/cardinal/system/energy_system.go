package system

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"
	"tamaland/params"
)

// EnergySystem decrease the player's E by 1 every 10 ticks.
// This provides an example of a system that doesn't rely on a transaction to update a component.
func EnergySystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Energy](),
			filter.Component[comp.State](),
			filter.Component[comp.LastUpdate](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {
			log.Debug().Msg("EnergySystem")

			// go to the next one if not time to update.
			lastUpdate, err := cardinal.GetComponent[comp.LastUpdate](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error1: " + err.Error())
				return true
			}
			if lastUpdate.Timestamp+params.StatUpdateIntervalMs > world.Timestamp() {
				return true
			}

			energy, err := cardinal.GetComponent[comp.Energy](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error2: " + err.Error())
				return true
			}
			state, err := cardinal.GetComponent[comp.State](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error3: " + err.Error())
				return true
			}

			energyToAdd := 0
			if state.State == params.StateSleeping && energy.E < params.MaxEnergy {
				energyToAdd = 5
			} else if state.State != params.StateSleeping && energy.E > params.MinEnergy {
				energyToAdd = -1
			}

			// Debug purposes
			skipCause := ""
			if state.State == params.StateSleeping && energy.E >= params.MaxEnergy {
				skipCause = fmt.Sprintf("entity %d is sleeping but has full energy (%d)", id, energy.E)
			}
			if state.State != params.StateSleeping && energy.E <= params.MinEnergy {
				skipCause = fmt.Sprintf("entity %d is not sleeping but has minimum energy (%d)", id, energy.E)
			}

			energy.E += energyToAdd

			if energy.E > params.MaxEnergy {
				energy.E = params.MaxEnergy
			}
			if energy.E < params.MinEnergy {
				energy.E = params.MinEnergy
			}

			if err := cardinal.SetComponent[comp.Energy](world, id, energy); err != nil {
				log.Error().Msg("EnergySystem: error4: " + err.Error())
				return true
			}

			log.Debug().Msg(fmt.Sprintf("EnergySystem: skipped: %s", skipCause))

			return true
		},
	)
	if err != nil {
		log.Error().Msg("EnergySystem: error5: " + err.Error())
		return err
	}
	return nil
}
