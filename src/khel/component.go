package khel

type Component interface {
}

type Components struct {
	Transform CTransform
	Shape     CShape
	Collision CCollision
	Score     CScore
	Lifespan  CLifespan
	Input     CInput
}

type CTransform struct {
	Pos      Vec2
	Velocity Vec2
	Scale    Vec2
	Angle    float32
	Rotation float32
}

type CCollision struct {
	Radius float32
}

type CScore struct {
	Score int
}

type CShape struct {
	Sides  int32
	Radius float32
}

type CLifespan struct {
	Lifespan  int
	Remaining int
}

type CInput struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
	Shoot bool
}
