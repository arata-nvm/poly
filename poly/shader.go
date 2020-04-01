package poly

import . "github.com/arata-nvm/poly/vecmath"

type Shader interface {
	Vertex(Vertex, Matrix4) Vertex
	Fragment(Vertex, Vector3) Color
}

type SolidShader struct {
	Color Color
}

func NewSolidShader(color Color) *SolidShader {
	return &SolidShader{Color: color}
}

func (s *SolidShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.Coordinates = TransformCoordinate(v.Coordinates, m)
	v.Normal = TransformCoordinate(v.Normal, m).Normalize()
	return v
}

func (s *SolidShader) Fragment(_ Vertex, _ Vector3) Color {
	return s.Color
}
