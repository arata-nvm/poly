package poly

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	. "github.com/arata-nvm/poly/vecmath"
)

func LoadPly(filename string) *Mesh {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return parsePly(f)
}

func parsePly(r io.Reader) *Mesh {
	o := NewMesh()

	s := bufio.NewScanner(r)
	numVertex, numFace := parseHeader(s)

	vertices := make([]Vector3, 0, numVertex)
	o.Faces = make([]*Face, 0, numFace)

	for i := 0; i < numVertex; i++ {
		s.Scan()
		line := s.Text()
		cols := strings.Split(line, " ")
		vertex := parsePlyVertex(cols)
		vertices = append(vertices, vertex)
	}

	for i := 0; i < numFace; i++ {
		s.Scan()
		line := s.Text()
		cols := strings.Split(line, " ")
		face := parsePlyFace(cols)

		v1 := Vertex{
			Coordinates: vertices[face[0]],
		}
		v2 := Vertex{
			Coordinates: vertices[face[1]],
		}
		v3 := Vertex{
			Coordinates: vertices[face[2]],
		}

		o.Faces = append(o.Faces, &Face{v1, v2, v3})
	}

	return o
}

func parseHeader(s *bufio.Scanner) (int, int) {
	var numVertex int
	var numFace int
	var err error

	for s.Scan() {
		line := s.Text()
		if line == "end_header" {
			break
		}

		cols := strings.Split(line, " ")
		if cols[0] == "element" {
			switch cols[1] {
			case "vertex":
				numVertex, err = strconv.Atoi(cols[2])
				if err != nil {
					panic(err)
				}
			case "face":
				numFace, err = strconv.Atoi(cols[2])
				if err != nil {
					panic(err)
				}
			}
		}
	}

	return numVertex, numFace
}

func parsePlyVertex(cols []string) Vector3 {
	x, err := strconv.ParseFloat(cols[0], 32)
	if err != nil {
		panic(err)
	}

	y, err := strconv.ParseFloat(cols[1], 32)
	if err != nil {
		panic(err)
	}

	z, err := strconv.ParseFloat(cols[2], 32)
	if err != nil {
		panic(err)
	}

	return NewVector3(x, y, z)
}

func parsePlyFace(cols []string) []int {
	v1, err := strconv.Atoi(cols[1])
	if err != nil {
		panic(err)
	}

	v2, err := strconv.Atoi(cols[2])
	if err != nil {
		panic(err)
	}

	v3, err := strconv.Atoi(cols[3])
	if err != nil {
		panic(err)
	}

	return []int{v1, v2, v3}
}
