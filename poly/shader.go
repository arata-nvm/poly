package poly

import (
	. "github.com/arata-nvm/poly/vecmath"
)

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

type FlatShader struct {
	Color Color
	Light Vector3
}

func NewFlatShader(color Color, light Vector3) *FlatShader {
	return &FlatShader{
		Color: color,
		Light: light.Normalize(),
	}
}

func (s *FlatShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.Coordinates = TransformCoordinate(v.Coordinates, m)
	v.Normal = TransformCoordinate(v.Normal, m).Normalize()
	return v
}

func (s *FlatShader) Fragment(v Vertex, _ Vector3) Color {
	f := clamp(v.Normal.Dot(s.Light), 0, 1)
	return NewColor(s.Color.R*f, s.Color.G*f, s.Color.B*f, s.Color.A)
}

type TextureShader struct {
	Texture *Texture
}

func NewTextureShader(texture *Texture) *TextureShader {
	return &TextureShader{
		Texture: texture,
	}
}

func (s *TextureShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.Coordinates = TransformCoordinate(v.Coordinates, m)
	v.Normal = TransformCoordinate(v.Normal, m).Normalize()
	return v
}

func (s *TextureShader) Fragment(v Vertex, _ Vector3) Color {
	return s.Texture.Map(v.Uv.X, v.Uv.Y)
}
