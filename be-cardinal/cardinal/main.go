package main

import (
	"errors"

	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"

	"tamaland/component"
	"tamaland/msg"
	"tamaland/query"
	"tamaland/system"
)

func main() {
	w, err := cardinal.NewWorld(cardinal.WithDisableSignatureVerification())
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	MustInitWorld(w)

	Must(w.StartGame())
}

// MustInitWorld registers all components, messages, queries, and systems. This initialization happens in a helper
// function so that this can be used directly in tests.
func MustInitWorld(w *cardinal.World) {
	// Register components
	// NOTE: You must register your components here for it to be accessible.
	Must(
		cardinal.RegisterComponent[component.Player](w),
		cardinal.RegisterComponent[component.Level](w),
		cardinal.RegisterComponent[component.State](w),
		cardinal.RegisterComponent[component.Health](w),
		cardinal.RegisterComponent[component.Energy](w),
		cardinal.RegisterComponent[component.Food](w),
		cardinal.RegisterComponent[component.LastUpdate](w),
	)

	// Register messages (user action)
	// NOTE: You must register your transactions here for it to be executed.
	Must(
		cardinal.RegisterMessage[msg.CreatePlayerMsg, msg.CreatePlayerResult](w, "create-player"),
		cardinal.RegisterMessage[msg.RespawnPlayerMsg, msg.RespawnPlayerResult](w, "respawn-player"),
		cardinal.RegisterMessage[msg.SleepMsg, msg.SleepMsgReply](w, "sleep"),
		cardinal.RegisterMessage[msg.EatMsg, msg.EatMsgReply](w, "eat"),
		cardinal.RegisterMessage[msg.AddLifeMsg, msg.AddLifeMsgReply](w, "down-life"),
	)

	// Register queries
	// NOTE: You must register your queries here for it to be accessible.
	Must(
		cardinal.RegisterQuery[query.PlayerHealthRequest, query.PlayerHealthResponse](w, "player-health", query.PlayerHealth),
		cardinal.RegisterQuery[query.PlayerEnergyRequest, query.PlayerEnergyResponse](w, "player-energy", query.PlayerEnergy),
		cardinal.RegisterQuery[query.PlayerFoodRequest, query.PlayerFoodResponse](w, "player-hunger", query.PlayerFood),
		cardinal.RegisterQuery[query.PlayerLevelRequest, query.PlayerLevelResponse](w, "player-level", query.PlayerLevel),
		cardinal.RegisterQuery[query.PlayerStateRequest, query.PlayerStateResponse](w, "player-state", query.PlayerState),
		cardinal.RegisterQuery[query.PlayerNicknameRequest, query.PlayerNicknameResponse](w, "player-nickname", query.PlayerNickname),
		cardinal.RegisterQuery[query.PlayerStatsRequest, query.PlayerStatsResponse](w, "player-stats", query.PlayerStats),
	)

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	Must(cardinal.RegisterSystems(w,
		system.StatusSystem,
		system.LevelSystem,
		system.HealthDown,
		system.EatSystem,
		system.SleepSystem,
		system.FoodSystem,
		system.EnergySystem,
		system.HealthSystem,
		system.PlayerSpawnerSystem,
		system.PlayerRespawnerSystem,
		system.LastUpdateSystem,
	))

	Must(cardinal.RegisterInitSystems(w,
		system.SpawnDefaultPlayersSystem,
	))
}

func Must(err ...error) {
	e := errors.Join(err...)
	if e != nil {
		log.Fatal().Err(e).Msg("")
	}
}
