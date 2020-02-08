package poly

import . "github.com/arata-nvm/poly/vecmath"

type Mesh struct {
	Vertices []Vector3
	Faces    []Face
}

func NewMesh() Mesh {
	return Mesh{}
}
