package behaviortree

// ActionType ...
type ActionType int

// List of all possible actions an entity can perform.
const (
	_                  ActionType = iota
	ActionMove                    // Move in a certain direction.
	ActionFace                    // Face a certain direction.
	ActionSetTarget               // Set a target.
	ActionAttack                  // Attack the current target, if any.
	ActionUseSkill                // Use skill by ID (AI input only).
	ActionUseSkillSlot            // Use skill at slot (player input only).
)

type ActionData struct {
	Type ActionType
	Args interface{}
}

type MoveAction struct {
	x, y float64
}

func NewMoveAction(x, y float64) *MoveAction {
	return &MoveAction{x, y}
}

func (a *MoveAction) Tick(ctx *Context) Status {
	ctx.board.Write(ctx.owner.String()+".x", a.x)
	ctx.board.Write(ctx.owner.String()+".y", a.y)
	return StatusSuccess
}

func SetTarget(ctx *Context) {

}

type Action struct {
	Leaf
}

// import (
// 	"github.com/alexanderskafte/lair/ai"
// )

// type Action struct {
// 	Leaf

// 	SelfID int // Entity ID of self
// 	Input  interface{}
// 	Output interface{}
// }

// // ActionFn is a function emitted by the behavior tree.
// type ActionFn func(*Context, interface{}, interface{}) Status

// var actionFns = map[string]ActionFn{
// 	"MoveTo": MoveTo,
// }

// type Condition struct {
// 	Leaf
// }

// func MoveTo(ctx *Context, in, out interface{}) Status {
// 	target := in.(ai.Target).Position()
// 	return StatusFailure
// }

// var targetMap

// func TargetSelect(ctx *Context) Status {

// }

// type ActionCommand interface {
// 	OldActor() OldActor
// 	Output() interface{}
// }

// The behavior tree should emit ActionCommands: put them on
// a channel owned by the same entity that owns the behavior
// tree instance.
