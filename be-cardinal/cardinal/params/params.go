package params

import comp "tamaland/component"

const (
	// Interval between statistics-update checks
	StatUpdateIntervalMs = 2 * Seconds2milliseconds

	// Time durations for the altered states
	SleepDurationMs  uint64 = 20 * Seconds2milliseconds
	EatingDurationMs uint64 = 5 * Seconds2milliseconds

	// Interval between levelups
	LevelUpIntervalMs = 60 * Seconds2milliseconds
)
const (
	StateHealthy   comp.StateType = "Healthy"   //unused
	StateUnhealthy comp.StateType = "Unhealthy" //unused
	StateNormal    comp.StateType = "Normal"
	StateSleeping  comp.StateType = "Sleeping"
	StateEating    comp.StateType = "Eating"
	StateHungry    comp.StateType = "Hungry" //unused
	StateTired     comp.StateType = "Tired"  //unused
	StateDead      comp.StateType = "Dead"

	Seconds2milliseconds uint64 = 1000
)

const (
	InitialHealth            = 100
	InitialEnergy            = 100
	InitialFood              = 100
	InitialLevel             = 1
	InitialState             = StateNormal
	InitialEndStateTimestamp = 0

	MaxHealth = 100
	MaxEnergy = 100
	MaxFood   = 100

	MinHeath  = 0
	MinEnergy = 0
	MinFood   = 0
)
