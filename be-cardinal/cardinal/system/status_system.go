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

// StatusSystem handles status transitions.
func StatusSystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.State](),
			filter.Component[comp.Health](),
			filter.Component[comp.Food](),
			filter.Component[comp.Energy](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {

			log.Debug().Msg("StatusSystem")

			state, err := cardinal.GetComponent[comp.State](world, id)
			if err != nil {
				log.Error().Msg("StatusSystem: error1: " + err.Error())
				return true
			}

			health, err := cardinal.GetComponent[comp.Health](world, id)
			if err != nil {
				log.Error().Msg("StatusSystem: error2: " + err.Error())
				return true
			}

			if health.HP <= params.MinHeath {
				state.State = params.StateDead
				state.EndStateTimestamp = 0

				food, err := cardinal.GetComponent[comp.Food](world, id)
				if err != nil {
					log.Error().Msg("StatusSystem: error3: " + err.Error())
					return true
				}
				energy, err := cardinal.GetComponent[comp.Energy](world, id)
				if err != nil {
					log.Error().Msg("StatusSystem: error4: " + err.Error())
					return true
				}

				food.Fd = params.MinFood
				energy.E = params.MinEnergy

				if err := cardinal.SetComponent[comp.Food](world, id, food); err != nil {
					log.Error().Msg("StatusSystem: error5: " + err.Error())
					return true
				}
				if err := cardinal.SetComponent[comp.Energy](world, id, energy); err != nil {
					log.Error().Msg("StatusSystem: error6: " + err.Error())
					return true
				}

			}

			if state.State != params.StateDead && !CanDoActions(*state) && world.Timestamp() >= state.EndStateTimestamp {
				state.State = params.StateNormal
				state.EndStateTimestamp = 0
			}

			if err := cardinal.SetComponent[comp.State](world, id, state); err != nil {
				log.Error().Msg("StatusSystem: error7: " + err.Error())
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
