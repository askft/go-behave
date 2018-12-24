package condition

import (
	"fmt"
	"sort"

	"github.com/alexanderskafte/behaviortree/core"

	"github.com/alexanderskafte/lair/core/vec"
	"github.com/alexanderskafte/lair/ecs/entities"
)

// TargetNearby ...
type TargetNearby struct {
	*core.Action
}

// Start ...
func (a *TargetNearby) Start(ctx *core.Context) {}

// Tick ...
func (a *TargetNearby) Tick(ctx *core.Context) core.Status {
	self := ctx.Owner.(*entities.Actor)

	data, ok := ctx.Store.Read("target_positions")
	if !ok {
		fmt.Println("could not read target_positions")
	}
	ps := data.([]vec.Vec)
	if len(ps) == 0 {
		return core.StatusFailure
	}
	pos := self.Position
	sort.Slice(ps, func(i, j int) bool {
		return pos.SquaredDistanceTo(ps[i]) < pos.SquaredDistanceTo(ps[i])
	})
	if pos.SquaredDistanceTo(ps[0]) > 100*100 {
		return core.StatusFailure
	}
	return core.StatusSuccess
}

// Stop ...
func (a *TargetNearby) Stop(ctx *core.Context) {}
