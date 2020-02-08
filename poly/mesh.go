package poly

type Mesh struct {
	Vertices []Vertex
	Faces    []Face
}

func NewMesh() Mesh {
	return Mesh{}
}
