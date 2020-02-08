package poly

import . "github.com/arata-nvm/poly/vecmath"

type Mesh struct {
	Vertices []Vertex
	Faces    []Face

	Rotation Vector3
}

func NewMesh() Mesh {
	return Mesh{
		Rotation: Zero(),
	}
}
