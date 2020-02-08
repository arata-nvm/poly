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

func LoadObj(filename string) Mesh {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return parseObj(f)
}

func parseObj(r io.Reader) Mesh {
	o := NewMesh()

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		switch line[0] {
		case 'v':
			o.Vertices = append(o.Vertices, parseVertex(line))
		case 'f':
			o.Faces = append(o.Faces, parseFace(line))
		default:
			fmt.Printf("unexpected kind %c\n", line[0])
			os.Exit(1)
		}
	}

	return o
}

func parseVertex(line string) Vertex {
	cols := strings.Split(line, " ")
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

	return Vertex{Coordinates:NewVector3(x, y, z)}
}

func parseFace(line string) Face {
	cols := strings.Split(line, " ")
	v1, err := strconv.ParseInt(cols[1], 10, 32)
	if err != nil {
		panic(err)
	}

	v2, err := strconv.ParseInt(cols[2], 10, 32)
	if err != nil {
		panic(err)
	}

	v3, err := strconv.ParseInt(cols[3], 10, 32)
	if err != nil {
		panic(err)
	}

	return Face{int(v1) - 1, int(v2) - 1, int(v3) - 1}
}
