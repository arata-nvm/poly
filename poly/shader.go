package poly

import (
	. "github.com/arata-nvm/poly/vecmath"
	"math"
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
	f := Clamp(v.Normal.Dot(s.Light), 0, 1)
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

type NormalShader struct{}

func NewNormalShader() *NormalShader {
	return &NormalShader{}
}

func (s *NormalShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.Coordinates = TransformCoordinate(v.Coordinates, m)
	v.Normal = TransformCoordinate(v.Normal, m).Normalize()
	return v
}

func (s *NormalShader) Fragment(v Vertex, _ Vector3) Color {
	n := v.Normal.Clamp(0, 1)
	return NewColorFromVec(n)
}

type PhongShader struct {
	Light Vector3
	Eye   Vector3
	Color Color
	Pow   float64
}

func NewPhongShader(light, eye Vector3, color Color, pow float64) *PhongShader {
	return &PhongShader{
		Light: light.Normalize(),
		Eye:   eye.Normalize(),
		Color: color,
		Pow:   pow,
	}
}

func (s *PhongShader) Vertex(v Vertex, m Matrix4) Vertex {
	v.Coordinates = TransformCoordinate(v.Coordinates, m)
	v.Normal = TransformCoordinate(v.Normal, m).Normalize()
	return v
}

func (s *PhongShader) Fragment(v Vertex, _ Vector3) Color {
	half := s.Light.Add(s.Eye).Normalize()
	diffuse := Clamp(v.Normal.Dot(s.Light), 0, 1)
	specular := math.Pow(Clamp(v.Normal.Dot(half), 0, 1), s.Pow)
	ambient := 0.05
	return NewColor(
		Clamp(s.Color.R*diffuse+specular+ambient, 0, 1),
		Clamp(s.Color.G*diffuse+specular+ambient, 0, 1),
		Clamp(s.Color.B*diffuse+specular+ambient, 0, 1),
		s.Color.A,
	)
}
