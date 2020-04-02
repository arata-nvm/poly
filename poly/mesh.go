package poly

import . "github.com/arata-nvm/poly/vecmath"

type Mesh struct {
	Faces []*Face

	Position Vector3
	Rotation Vector3
	Scale    Vector3
}

func NewMesh() *Mesh {
	return &Mesh{
		Position: Zero(),
		Rotation: Zero(),
		Scale:    Unit(),
	}
}
