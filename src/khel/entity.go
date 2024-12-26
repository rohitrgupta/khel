package khel

type Entity struct {
	id         int
	Components Components
	tag        string
	alive      bool
}

func NewEntity(id int, tag string) *Entity {
	return &Entity{
		id:    id,
		tag:   tag,
		alive: true,
	}
}

func (e *Entity) ID() int {
	return e.id
}

func (e *Entity) Tag() string {
	return e.tag
}

func (e *Entity) Destroy() {
	e.alive = false
}
func (e *Entity) IsAlive() bool {
	return e.alive
}

// func (e *Entity) AddComponent(c Component) {
// 	switch component := c.(type) {
// 	case CTransform:
// 		e.Components.Transform = component
// 	case CCollision:
// 		e.Components.Collision = component
// 	case CScore:
// 		e.Components.Score = component
// 	case CLifespan:
// 		e.Components.Lifespan = component
// 	case CInput:
// 		e.Components.Input = component
// 	default:
// 		fmt.Printf("Unknown component type: %T\n", component)
// 	}
// }

// func (e *Entity) GetComponent(t string) Component {
// 	switch t {
// 	case "CTransform":
// 		return e.Components.Transform
// 	case "CShape":
// 		return e.Components.Shape
// 	case "CCollision":
// 		return e.Components.Collision
// 	case "CScore":
// 		return e.Components.Score
// 	case "CLifespan":
// 		return e.Components.Lifespan
// 	case "CInput":
// 		return e.Components.Input
// 	default:
// 		fmt.Printf("Unknown component type: %s\n", t)
// 		return nil
// 	}
// }
