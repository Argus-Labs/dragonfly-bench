package system

import (
	comp "github.com/argus-labs/starter-game-template/cardinal/component"
	"pkg.world.dev/world-engine/cardinal"
)

// PlayerSpawnerSystem is a system that spawns 16000 players once for benchmarking
func PlayerSpawnerSystem(wCtx cardinal.WorldContext) error {
	q, err := wCtx.NewSearch(cardinal.Contains(comp.PlayerComponent{}))
	if err != nil {
		return err
	}
	c, err := q.Count(wCtx)
	if err != nil {
		return err
	}

	// Only spawn players if there are no players in the world
	if c == 0 {
		_, err := cardinal.CreateMany(wCtx, 15000,
			comp.PlayerComponent{Nickname: "CoolMage"},
			comp.HealthComponent{HP: 1},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
