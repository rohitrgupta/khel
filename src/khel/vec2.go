package khel

import "math"

type Vec2 struct {
	X, Y float32
}

func (v *Vec2) Add(other Vec2) {
	v.X += other.X
	v.Y += other.Y
}

func (v *Vec2) Sub(other Vec2) {
	v.X -= other.X
	v.Y -= other.Y
}

func (v *Vec2) Mul(other float32) {
	v.X *= other
	v.Y *= other
}

func (v Vec2) Dot(other Vec2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Cross(other Vec2) float32 {
	return v.X*other.Y - v.Y*other.X
}

func (v Vec2) LengthSquared() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Length() float32 {
	return float32(math.Sqrt(float64(v.LengthSquared())))
}

func (v *Vec2) Normalize() {
	v.Mul(1 / v.Length())
}

func (v Vec2) Normalized() Vec2 {
	return Mul(v, float32(1/v.Length()))
}

func (v Vec2) Plus(v2 Vec2) Vec2 {
	return Add(v, v2)
}

func (v Vec2) Minus(v2 Vec2) Vec2 {
	return Sub(v, v2)
}

func (v Vec2) Times(r float32) Vec2 {
	return Mul(v, r)
}

func Add(v, u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

func Sub(v, u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

func Mul(v Vec2, r float32) Vec2 {
	return Vec2{v.X * r, v.Y * r}
}
