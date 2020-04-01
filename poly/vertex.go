package poly

import . "github.com/arata-nvm/poly/vecmath"

type Vertex struct {
	Coordinates Vector3
	Uv          Vector3
	Normal      Vector3
}

func InterpolateVertex(v1, v2, v3 Vertex, w Vector3) Vertex {
	return Vertex{
		Coordinates: InterpolateVector(v1.Coordinates, v2.Coordinates, v3.Coordinates, w),
		Uv:          InterpolateVector(v1.Uv, v2.Uv, v3.Uv, w),
		Normal:      InterpolateVector(v1.Normal, v2.Normal, v3.Normal, w),
	}
}

func InterpolateVector(v1, v2, v3 Vector3, w Vector3) Vector3 {
	return NewVector3(
		w.X*v1.X+w.Y*v2.X+w.Z*v3.X,
		w.X*v1.Y+w.Y*v2.Y+w.Z*v3.Y,
		w.X*v1.Z+w.Y*v2.Z+w.Z*v3.Z,
	)
}
