package vecmath

import (
	"math"
)

type Matrix4 struct {
	M00, M01, M02, M03 float64
	M10, M11, M12, M13 float64
	M20, M21, M22, M23 float64
	M30, M31, M32, M33 float64
}

func Identity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotateX(theta float64) Matrix4 {
	s := math.Sin(theta)
	c := math.Cos(theta)
	return Matrix4{
		1, 0, 0, 0,
		0, c, -s, 0,
		0, s, c, 0,
		0, 0, 0, 1,
	}
}

func RotateY(theta float64) Matrix4 {
	s := math.Sin(theta)
	c := math.Cos(theta)
	return Matrix4{
		c, 0, s, 0,
		0, 1, 0, 0,
		-s, 0, c, 0,
		0, 0, 0, 1,
	}
}

func RotateZ(theta float64) Matrix4 {
	s := math.Sin(theta)
	c := math.Cos(theta)
	return Matrix4{
		c, -s, 0, 0,
		s, c, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func RotateAxis(axis Vector3, theta float64) Matrix4 {
	v := axis.Normalize()
	s := math.Sin(theta)
	c := math.Cos(theta)
	c2 := 1.0 - c
	return Matrix4{
		v.X*v.X*c2 + c,
		v.X*v.Y*c2 - v.X*s,
		v.X*v.Y*c2 + v.X*s,
		0,
		v.Y*v.X*c2 + v.X*s,
		v.Y*v.Y*c2 + c,
		v.Y*v.X*c2 - v.X*s,
		0,
		v.Z*v.X*c2 - v.Y*s,
		v.Z*v.Y*c2 + v.X*s,
		v.Z*v.Z*c2 + c,
		0,
		0, 0, 0, 1,
	}
}

func Translate(v Vector3) Matrix4 {
	return Matrix4{
		1, 0, 0, v.X,
		0, 1, 0, v.Y,
		0, 0, 1, v.Z,
		0, 0, 0, 1,
	}
}

func Scale(v Vector3) Matrix4 {
	return Matrix4{
		v.X, 0, 0, 0,
		0, v.Y, 0, 0,
		0, 0, v.Z, 0,
		0, 0, 0, 1,
	}
}

func LookAt(from, to, upward Vector3) Matrix4 {
	z := from.Sub(to).Normalize()
	x := upward.Cross(z).Normalize()
	y := z.Cross(x)

	return Matrix4{
		x.X, x.Y, x.Z, -x.Dot(from),
		y.X, y.Y, y.Z, -y.Dot(from),
		z.X, z.Y, z.Z, -z.Dot(from),
		0, 0, 0, 1,
	}
}

func Orthographic(left, right, bottom, top, near, far float64) Matrix4 {
	return Matrix4{
		2 / (right - left), 0, 0, -(right + left) / (right - left),
		0, 2 / (top - bottom), 0, -(top + bottom) / (top - bottom),
		0, 0, -2 / (far - near), -(far + near) / (far - near),
		0, 0, 0, 1,
	}
}

func Frustum(left, right, bottom, top, near, far float64) Matrix4 {
	return Matrix4{
		(2 * near) / (right - left), 0, (right + left) / (right - left), 0,
		0, (2 * near) / (top - bottom), (top + bottom) / (top - bottom), 0,
		0, 0, -(far + near) / (far - near), (-2 * near * far) / (far - near),
		0, 0, -1, 0,
	}
}

func Perspective(fovy, aspect, near, far float64) Matrix4 {
	top := math.Tan(fovy*math.Pi/360) * near
	bottom := -top
	left := bottom * aspect
	right := top * aspect
	return Frustum(left, right, bottom, top, near, far)
}

func (m1 Matrix4) Add(m2 Matrix4) Matrix4 {
	return Matrix4{
		m1.M00 + m2.M00, m1.M01 + m2.M01, m1.M02 + m2.M02, m1.M03 + m2.M03,
		m1.M10 + m2.M10, m1.M11 + m2.M11, m1.M12 + m2.M12, m1.M13 + m2.M13,
		m1.M20 + m2.M20, m1.M21 + m2.M21, m1.M22 + m2.M22, m1.M23 + m2.M23,
		m1.M30 + m2.M30, m1.M31 + m2.M31, m1.M32 + m2.M32, m1.M33 + m2.M33,
	}
}

func (m1 Matrix4) Sub(m2 Matrix4) Matrix4 {
	return Matrix4{
		m1.M00 - m2.M00, m1.M01 - m2.M01, m1.M02 - m2.M02, m1.M03 - m2.M03,
		m1.M10 - m2.M10, m1.M11 - m2.M11, m1.M12 - m2.M12, m1.M13 - m2.M13,
		m1.M20 - m2.M20, m1.M21 - m2.M21, m1.M22 - m2.M22, m1.M23 - m2.M23,
		m1.M30 - m2.M30, m1.M31 - m2.M31, m1.M32 - m2.M32, m1.M33 - m2.M33,
	}
}

func (m1 Matrix4) Mul(m2 Matrix4) Matrix4 {
	return Matrix4{
		m1.M00*m2.M00 + m1.M01*m2.M10 + m1.M02*m2.M20 + m1.M03*m2.M30,
		m1.M00*m2.M01 + m1.M01*m2.M11 + m1.M02*m2.M21 + m1.M03*m2.M31,
		m1.M00*m2.M02 + m1.M01*m2.M12 + m1.M02*m2.M22 + m1.M03*m2.M32,
		m1.M00*m2.M03 + m1.M01*m2.M13 + m1.M02*m2.M23 + m1.M03*m2.M33,
		m1.M10*m2.M00 + m1.M11*m2.M10 + m1.M12*m2.M20 + m1.M13*m2.M30,
		m1.M10*m2.M01 + m1.M11*m2.M11 + m1.M12*m2.M21 + m1.M13*m2.M31,
		m1.M10*m2.M02 + m1.M11*m2.M12 + m1.M12*m2.M22 + m1.M13*m2.M32,
		m1.M10*m2.M03 + m1.M11*m2.M13 + m1.M12*m2.M23 + m1.M13*m2.M33,
		m1.M20*m2.M00 + m1.M21*m2.M10 + m1.M22*m2.M20 + m1.M23*m2.M30,
		m1.M20*m2.M01 + m1.M21*m2.M11 + m1.M22*m2.M21 + m1.M23*m2.M31,
		m1.M20*m2.M02 + m1.M21*m2.M12 + m1.M22*m2.M22 + m1.M23*m2.M32,
		m1.M20*m2.M03 + m1.M21*m2.M13 + m1.M22*m2.M23 + m1.M23*m2.M33,
		m1.M30*m2.M00 + m1.M31*m2.M10 + m1.M32*m2.M20 + m1.M33*m2.M30,
		m1.M30*m2.M01 + m1.M31*m2.M11 + m1.M32*m2.M21 + m1.M33*m2.M31,
		m1.M30*m2.M02 + m1.M31*m2.M12 + m1.M32*m2.M22 + m1.M33*m2.M32,
		m1.M30*m2.M03 + m1.M31*m2.M13 + m1.M32*m2.M23 + m1.M33*m2.M33,
	}
}

func (m1 Matrix4) MulScalar(f float64) Matrix4 {
	return Matrix4{
		m1.M00 * f, m1.M01 * f, m1.M02 * f, m1.M03 * f,
		m1.M10 * f, m1.M11 * f, m1.M12 * f, m1.M13 * f,
		m1.M20 * f, m1.M21 * f, m1.M22 * f, m1.M23 * f,
		m1.M30 * f, m1.M31 * f, m1.M32 * f, m1.M33 * f,
	}
}

func (m1 Matrix4) MulVector(v Vector3) Vector3 {
	return Vector3{
		m1.M00*v.X + m1.M01*v.Y + m1.M02*v.Z + m1.M03,
		m1.M10*v.X + m1.M11*v.Y + m1.M12*v.Z + m1.M13,
		m1.M20*v.X + m1.M21*v.Y + m1.M22*v.Z + m1.M23,
	}
}
