package vecmath

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func Zero() Vector3 {
	return Vector3{0, 0, 0}
}

func Unit() Vector3 {
	return Vector3{1, 1, 1}
}

func UnitX() Vector3 {
	return Vector3{1, 0, 0}
}

func UnitY() Vector3 {
	return Vector3{0, 1, 0}
}

func UnitZ() Vector3 {
	return Vector3{0, 0, 1}
}

func TransformCoordinate(v Vector3, transform Matrix4) Vector3 {
	w := 1 / (transform.M30*v.X + transform.M31*v.Y + transform.M32*v.Z + transform.M33)
	return transform.MulVector(v).MulScalar(w)
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

func (v1 Vector3) Mul(v2 Vector3) Vector3 {
	return Vector3{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

func (v1 Vector3) Div(v2 Vector3) Vector3 {
	return Vector3{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
}

func (v1 Vector3) AddScalar(f float64) Vector3 {
	return Vector3{v1.X + f, v1.Y + f, v1.Z + f}
}

func (v1 Vector3) SubScalar(f float64) Vector3 {
	return Vector3{v1.X - f, v1.Y - f, v1.Z - f}
}
func (v1 Vector3) MulScalar(f float64) Vector3 {
	return Vector3{v1.X * f, v1.Y * f, v1.Z * f}
}

func (v1 Vector3) DivScalar(f float64) Vector3 {
	invF := 1 / f
	return Vector3{v1.X * invF, v1.Y * invF, v1.Z * invF}
}

func (v1 Vector3) Dot(v2 Vector3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X,
	}
}

func (v1 Vector3) Length() float64 {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

func (v1 Vector3) LengthSq() float64 {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z
}

func (v1 Vector3) Normalize() Vector3 {
	invLen := 1 / v1.Length()
	return Vector3{v1.X * invLen, v1.Y * invLen, v1.Z * invLen}
}

func (v1 Vector3) Negate() Vector3 {
	return Vector3{-v1.X, -v1.Y, -v1.Z}
}

func (v1 Vector3) Clamp(min, max float64) Vector3 {
	return Vector3{
		Clamp(v1.X, min, max),
		Clamp(v1.Y, min, max),
		Clamp(v1.Z, min, max),
	}
}

func (v Vector3) Reflected(n Vector3) Vector3 {
	return n.MulScalar(2 * v.Dot(n)).Sub(v).Normalize()
}
