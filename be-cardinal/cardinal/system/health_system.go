package system

import (
	//"github.com/rs/zerolog/log"
	"fmt"
	"math"

	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "tamaland/component"
	"tamaland/params"
)

// HealthSystem decrease the player's HP by 1 every 10 ticks.
// This provides an example of a system that doesn't rely on a transaction to update a component.
func HealthSystem(world cardinal.WorldContext) error {
	err := cardinal.NewSearch().Entity(
		filter.Contains(
			filter.Component[comp.Health](),
			filter.Component[comp.Energy](),
			filter.Component[comp.Food](),
			filter.Component[comp.LastUpdate](),
		),
	).Each(
		world,
		func(id types.EntityID) bool {
			log.Debug().Msg("HealthSystem")

			// go to the next one if not time to update.
			lastUpdate, err := cardinal.GetComponent[comp.LastUpdate](world, id)
			if err != nil {
				log.Error().Msg("EnergySystem: error1: " + err.Error())
				return true
			}
			if lastUpdate.Timestamp+params.StatUpdateIntervalMs > world.Timestamp() {
				return true
			}

			health, err := cardinal.GetComponent[comp.Health](world, id)
			if err != nil {
				return true
			}
			energy, err := cardinal.GetComponent[comp.Energy](world, id)
			if err != nil {
				return true
			}
			food, err := cardinal.GetComponent[comp.Food](world, id)
			if err != nil {
				return true
			}
			if health.HP > params.MinHeath {

				currentEnergyPercentage := float64(energy.E) / params.MaxEnergy * 100
				currentFoodPercentage := float64(food.Fd) / params.MaxFood * 100

				totalHealthLoss := 0

				if currentEnergyPercentage <= 80 && currentEnergyPercentage > 60 {
					totalHealthLoss += 1
				} else if currentEnergyPercentage <= 60 && currentEnergyPercentage > 40 {
					totalHealthLoss += 2
				} else if currentEnergyPercentage <= 40 && currentEnergyPercentage > 20 {
					totalHealthLoss += 3
				} else if currentEnergyPercentage <= 20 {
					totalHealthLoss += 4
				}

				if currentFoodPercentage <= 80 && currentFoodPercentage > 60 {
					totalHealthLoss += 1
				} else if currentFoodPercentage <= 60 && currentFoodPercentage > 40 {
					totalHealthLoss += 2
				} else if currentFoodPercentage <= 40 && currentFoodPercentage > 20 {
					totalHealthLoss += 3
				} else if currentFoodPercentage <= 20 {
					totalHealthLoss += 4
				}

				totalHealthLossScaled := int(math.Floor(float64(totalHealthLoss) / 2))

				log.Debug().Msg(fmt.Sprintf("HealthSystem: entity %d has energyPercentage=%d, foodPercentage=%d ; totalLoss=%d scaled to %d", id, energy.E, food.Fd, totalHealthLoss, totalHealthLossScaled))

				health.HP -= totalHealthLossScaled
				if health.HP < params.MinHeath {
					health.HP = params.MinHeath
				}
				if health.HP > params.MaxHealth {
					health.HP = params.MaxHealth
				}

				if err := cardinal.SetComponent[comp.Health](world, id, health); err != nil {
					return true
				}
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
