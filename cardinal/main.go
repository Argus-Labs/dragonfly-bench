package main

import (
	"errors"
	"github.com/argus-labs/starter-game-template/cardinal/component"
	"github.com/argus-labs/starter-game-template/cardinal/system"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg := GetConfig()
	world := NewWorld(cfg, cardinal.WithDisableSignatureVerification())

	// Register components
	// NOTE: You must register your components here,
	// otherwise it will show an error when you try to use them in a system.
	err := errors.Join(
		cardinal.RegisterComponent[component.PlayerComponent](world),
		cardinal.RegisterComponent[component.HealthComponent](world))
	if err != nil {
		log.Fatal().Err(err)
	}

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	cardinal.RegisterSystems(world,
		system.RegenSystem,
		system.PlayerSpawnerSystem,
	)

	err = world.StartGame()
	if err != nil {
		panic(err)
	}

}
