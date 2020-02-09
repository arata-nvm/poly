package poly

import . "github.com/arata-nvm/poly/vecmath"

type Camera struct {
	Position Vector3
	Target   Vector3
	Up       Vector3
}

func NewCamera(position, target, up Vector3) Camera {
	return Camera{
		Position: position,
		Target:   target,
		Up:       up,
	}
}
