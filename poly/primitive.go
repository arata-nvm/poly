package poly

import . "github.com/arata-nvm/poly/vecmath"

func NewPlane() *Mesh {
	v1 := NewVector3(-1, 0, 1)
	v2 := NewVector3(1, 0, 1)
	v3 := NewVector3(-1, 0, -1)
	v4 := NewVector3(1, 0, -1)

	t1 := NewVector3(1, 0, 0)
	t2 := NewVector3(0, 1, 0)
	t3 := NewVector3(0, 0, 0)
	t4 := NewVector3(1, 1, 0)

	m := &Mesh{
		Faces:    []*Face{
			{
				Vertex{
					Coordinates: v2,
					Uv:          t1,
				},
				Vertex{
					Coordinates: v3,
					Uv:          t2,
				},
				Vertex{
					Coordinates: v1,
					Uv:          t3,
				},
			},
			{
				Vertex{
					Coordinates: v2,
					Uv:          t1,
				},
				Vertex{
					Coordinates: v4,
					Uv:          t4,
				},
				Vertex{
					Coordinates: v3,
					Uv:          t2,
				},
			},
		},
	}

	m.CalcNormal()
	return m
}

//func NewBox() *Mesh {
//
//}
//
//func NewSphere() *Mesh {
//
//}
//
//func NewTrous() *Mesh {
//
//}
//
