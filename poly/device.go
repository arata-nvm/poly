package poly

import (
	. "github.com/arata-nvm/poly/vecmath"
	"image"
	"math"
)

type Device struct {
	Width  int
	Height int

	camera      Camera
	shader      Shader
	colorBuffer *image.NRGBA
	depthBuffer []float64

	viewMatrix       Matrix4
	projectionMatrix Matrix4
}

func NewDevice(width, height int) *Device {
	d := &Device{
		Width:       width,
		Height:      height,
		colorBuffer: image.NewNRGBA(image.Rect(0, 0, width, height)),
		depthBuffer: make([]float64, width*height),
	}

	d.ClearDepthBuffer(math.MaxFloat64)

	return d
}

func (d *Device) ClearColorBuffer(c Color) {
	for y := 0; y < d.Height; y++ {
		for x := 0; x < d.Width; x++ {
			d.colorBuffer.SetNRGBA(x, y, c.NRGBA())
		}
	}
}

func (d *Device) ClearDepthBuffer(f float64) {
	for i := range d.depthBuffer {
		d.depthBuffer[i] = f
	}
}

func (d *Device) Image() image.Image {
	return d.colorBuffer
}

func (d *Device) DepthBuffer() []float64 {
	return d.depthBuffer
}

func (d *Device) SetCamera(c Camera) {
	d.camera = c
	d.viewMatrix = LookAt(c.Position, c.Target, c.Up)
}

func (d *Device) SetShader(s Shader) {
	d.shader = s
}

func (d *Device) Perspective(fovy, aspect, near, far float64) {
	d.projectionMatrix = Perspective(fovy, aspect, near, far)
}

func (d *Device) putPixel(x, y int, z float64, c Color) {
	if x < 0 || y < 0 || x >= d.Width || y >= d.Height {
		return
	}

	index := x + y*d.Width
	if d.depthBuffer[index] < z {
		return
	}

	d.depthBuffer[index] = z
	d.colorBuffer.Set(x, y, c.NRGBA())
}

func (d *Device) DrawPoint(v Vector3, c Color) {
	d.putPixel(int(v.X), int(v.Y), v.Y, c)
}

// TODO v1 > v2
func (d *Device) DrawLine(v1, v2 Vector3, c Color) {
	x1 := int(v1.X)
	y1 := int(v1.Y)
	x2 := int(v2.X)
	y2 := int(v2.Y)

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := sign(x2 - x1)
	sy := sign(y2 - y1)
	err := dx - dy

	gz := 1.0
	if dx > dy {
		gz /= float64(dx)
	} else {
		gz /= float64(dy)
	}

	for {
		var g float64
		if dx > dy {
			g = gz * float64(x2-x1)
		} else {
			g = gz * float64(y2-y1)
		}

		z := interpolate(v1.Z, v2.Z, 1-g)

		d.putPixel(x1, y1, z, c)

		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dx {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func (d *Device) DrawMesh(mesh *Mesh) {
	tm := Translate(mesh.Position)
	rm := RotateX(mesh.Rotation.X).Mul(RotateY(mesh.Rotation.Y)).Mul(RotateZ(mesh.Rotation.Z))
	sm := Scale(mesh.Scale)
	modelMatrix := tm.Mul(rm).Mul(sm)
	transformMatrix := d.projectionMatrix.Mul(d.viewMatrix).Mul(modelMatrix)

	for _, f := range mesh.Faces {
		v1 := d.transformVertex(f.V1, transformMatrix).WorldCoordinates
		v2 := d.transformVertex(f.V2, transformMatrix).WorldCoordinates
		v3 := d.transformVertex(f.V3, transformMatrix).WorldCoordinates
		c := d.shader.Fragment(f.V1) // TODO
		d.DrawTriangle(v1, v2, v3, c)
	}
}

func (d *Device) transformVertex(v Vertex, m Matrix4) Vertex {
	cx, cy := d.Width/2, d.Height/2
	scale := float64(d.Width) / 2

	v = d.shader.Vertex(v, m)
	v.WorldCoordinates.X = v.WorldCoordinates.X*scale + float64(cx)
	v.WorldCoordinates.Y = v.WorldCoordinates.Y*scale + float64(cy)

	return v
}

func (d *Device) DrawWiredTriangle(v1, v2, v3 Vector3, c Color) {
	d.DrawLine(v1, v2, c)
	d.DrawLine(v2, v3, c)
	d.DrawLine(v3, v1, c)
}

// TODO fix errors(float -> int)
func (d *Device) DrawTriangle(v1, v2, v3 Vector3, c Color) {
	v1, v2, v3 = sortVectorsWithY(v1, v2, v3)
	d12, d13 := 0.0, 0.0
	if v2.Y-v1.Y > 0 {
		d12 = (v2.X - v1.X) / (v2.Y - v1.Y)
	}
	if v3.Y-v1.Y > 0 {
		d13 = (v3.X - v1.X) / (v3.Y - v1.Y)
	}

	if d12 > d13 {
		for y := int(v1.Y); y <= int(v3.Y); y++ {
			if float64(y) < v2.Y {
				d.scanLine(y, v1, v3, v1, v2, c)
			} else {
				d.scanLine(y, v1, v3, v2, v3, c)
			}
		}
	} else {
		for y := int(v1.Y); y <= int(v3.Y); y++ {
			if float64(y) < v2.Y {
				d.scanLine(y, v1, v2, v1, v3, c)
			} else {
				d.scanLine(y, v2, v3, v1, v3, c)
			}
		}
	}
}

func (d *Device) scanLine(y int, va, vb, vc, vd Vector3, c Color) {
	g1 := (float64(y) - va.Y) / (vb.Y - va.Y)
	x1 := int(interpolate(va.X, vb.X, g1))
	z1 := interpolate(va.Z, vb.Z, g1)

	g2 := (float64(y) - vc.Y) / (vd.Y - vc.Y)
	x2 := int(interpolate(vc.X, vd.X, g2))
	z2 := interpolate(vc.Z, vd.Z, g2)

	if math.IsNaN(g1) || math.IsNaN(g2) {
		return
	}

	xs, xe := min(x1, x2), max(x1, x2)

	for x := xs; x <= xe; x++ {
		g := float64(x-xs) / float64(xe-xs)
		if math.IsNaN(g) {
			g = 0
		}
		z := interpolate(z1, z2, g)

		d.putPixel(x, y, z, c)
	}
}

func sortVectorsWithY(v1, v2, v3 Vector3) (Vector3, Vector3, Vector3) {
	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}
	if v2.Y > v3.Y {
		v2, v3 = v3, v2
	}
	if v1.Y > v2.Y {
		v1, v2 = v2, v1
	}
	return v1, v2, v3
}
