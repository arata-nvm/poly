package poly

import (
	"bufio"
	"fmt"
	. "github.com/arata-nvm/poly/vecmath"
	"io"
	"os"
	"strconv"
	"strings"
)

func LoadObj(filename string) *Mesh {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return parseObj(f)
}

func parseObj(r io.Reader) *Mesh {
	o := NewMesh()

	vertices := make([]Vector3, 0)
	uvs := make([]Vector3, 0)
	normals := make([]Vector3, 0)

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		cols := strings.Split(line, " ")
		switch cols[0] {
		case "v":
			vertices = append(vertices, parseVertex(cols))
		case "vt":
			uvs = append(uvs, parseUv(cols))
		case "vn":
			normals = append(normals, parseNormal(cols))
		case "f":
			col1 := parseFaceIndices(cols[1])
			col2 := parseFaceIndices(cols[2])
			col3 := parseFaceIndices(cols[3])

			v1 := Vertex{
				Coordinates: vertices[col1[0]],
				Uv:          uvs[col1[1]],
				Normal:      normals[col1[2]],
			}

			v2 := Vertex{
				Coordinates: vertices[col2[0]],
				Uv:          uvs[col2[1]],
				Normal:      normals[col2[2]],
			}

			v3 := Vertex{
				Coordinates: vertices[col3[0]],
				Uv:          uvs[col3[1]],
				Normal:      normals[col3[2]],
			}

			o.Faces = append(o.Faces, Face{v1, v2, v3})
		default:
			fmt.Printf("unexpected kind %s\n", cols[0])
			os.Exit(1)
		}
	}

	return o
}

func parseVertex(cols []string) Vector3 {
	x, err := strconv.ParseFloat(cols[1], 32)
	if err != nil {
		panic(err)
	}

	y, err := strconv.ParseFloat(cols[2], 32)
	if err != nil {
		panic(err)
	}

	z, err := strconv.ParseFloat(cols[3], 32)
	if err != nil {
		panic(err)
	}

	return NewVector3(x, y, z)
}

func parseUv(cols []string) Vector3 {
	u, err := strconv.ParseFloat(cols[1], 32)
	if err != nil {
		panic(err)
	}

	v, err := strconv.ParseFloat(cols[2], 32)
	if err != nil {
		panic(err)
	}

	return NewVector3(u, v, 0)
}

func parseNormal(cols []string) Vector3 {
	x, err := strconv.ParseFloat(cols[1], 32)
	if err != nil {
		panic(err)
	}

	y, err := strconv.ParseFloat(cols[2], 32)
	if err != nil {
		panic(err)
	}

	z, err := strconv.ParseFloat(cols[3], 32)
	if err != nil {
		panic(err)
	}

	return NewVector3(x, y, z)
}

func parseFaceIndices(col string) []int {
	cols := strings.Split(col, "/")
	v1, err := strconv.ParseInt(cols[0], 10, 32)
	if err != nil {
		panic(err)
	}

	v2, err := strconv.ParseInt(cols[1], 10, 32)
	if err != nil {
		panic(err)
	}

	v3, err := strconv.ParseInt(cols[2], 10, 32)
	if err != nil {
		panic(err)
	}

	return []int{int(v1) - 1, int(v2) - 1, int(v3) - 1}
}
