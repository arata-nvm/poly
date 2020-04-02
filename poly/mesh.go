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

func (m *Mesh) SmoothNormals() {
	// 同じ位置にある頂点の法線ベクトルの総和をとり、正規化
	normals := make(map[Vector3]Vector3)
	for _, f := range m.Faces {
		normals[f.V1.Coordinates] = normals[f.V1.Coordinates].Add(f.V1.Normal)
		normals[f.V2.Coordinates] = normals[f.V2.Coordinates].Add(f.V2.Normal)
		normals[f.V3.Coordinates] = normals[f.V3.Coordinates].Add(f.V3.Normal)
	}

	for coord, normal := range normals {
		normals[coord] = normal.Normalize()
	}

	for _, f := range m.Faces {
		f.V1.Normal = normals[f.V1.Coordinates]
		f.V2.Normal = normals[f.V2.Coordinates]
		f.V3.Normal = normals[f.V3.Coordinates]
	}
}
