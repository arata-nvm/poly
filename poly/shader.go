package poly

import . "github.com/arata-nvm/poly/vecmath"

type Shader interface {
	Vertex(Vertex, Matrix4) Vertex
	Fragment(Vertex) Color
}

type SolidShader struct {
	Color Color
}

func NewSolidShader(color Color) *SolidShader {
	return &SolidShader{Color: color}
}

func (s *SolidShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.WorldCoordinates = TransformCoordinate(v.Coordinates, m)
	return v
}

func (s *SolidShader) Fragment(v Vertex) Color {
	return s.Color
}
