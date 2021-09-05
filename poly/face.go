package poly

type Face struct {
	V1, V2, V3 Vertex
}

func (f *Face) CalcNormal() {
	d1 := f.V2.Coordinates.Sub(f.V1.Coordinates)
	d2 := f.V3.Coordinates.Sub(f.V1.Coordinates)
	n := d1.Cross(d2).Normalize()

	f.V1.Normal = n
	f.V2.Normal = n
	f.V3.Normal = n
}
