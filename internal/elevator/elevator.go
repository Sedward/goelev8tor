package elevator

type DirectionType string

const (
	Up        DirectionType = "Up"
	Down                    = "Down"
	Staionary               = "Stationary"
)

type Elevator struct {
	Direction DirectionType
}
