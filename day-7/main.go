package main

import (
	"fmt"
	"os"
	"strings"
)

type Node interface {
	Name() string
	Size() int
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name, size}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

type Dir struct {
	name  string
	nodes []Node
	sz    int
}

func NewDir(name string) *Dir {
	return &Dir{name: name, nodes: make([]Node, 0, 1), sz: -1}
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) Size() int {
	if d.sz == -1 { // so to not recompute it every time
		d.sz = 0
		for _, node := range d.nodes {
			d.sz += node.Size()
		}
	}
	return d.sz
}

func (d *Dir) AddNode(node Node) {
	d.sz = -1
	d.nodes = append(d.nodes, node)
}

func printFS(fs Node) string {
	var visit func(Node, string) []string
	visit = func(node Node, pref string) []string {
		res := make([]string, 0, 1)
		if dir, ok := node.(*Dir); ok {
			res = append(res, fmt.Sprintf("%s%s", pref+"└-", dir.Name()))
			for _, node := range dir.nodes {
				for _, sp := range visit(node, "  ") {
					res = append(res, pref+sp)
				}
			}
		} else {
			res = append(res, fmt.Sprintf("%s%s %d", pref+"└-", node.Name(), node.Size()))
		}
		return res
	}

	pp := visit(fs, "")
	return strings.Join(pp, "\n")
}

var (
	CD = NewDir("..")
)

func rebuildFS(cmds []string) Node {
	var recurse func(int) (Node, int)
	recurse = func(ix int) (Node, int) {
		_, ptr := readStr(cmds[ix], 0, "$ cd ")
		name := cmds[ix][ptr:]
		ix++
		if name == CD.name {
			return CD, ix
		}
		dir := NewDir(name)
		ix++ // read ls

		for ix < len(cmds) && cmds[ix][0] != '$' {
			if !startsWith(cmds[ix], "dir") {
				size, pp := readInt(cmds[ix], 0)
				file := NewFile(cmds[ix][pp+1:], size)
				dir.AddNode(file)
			}
			ix++
		}

		for ix < len(cmds) {
			var sdir Node
			sdir, ix = recurse(ix)
			debugf("visiting dir: %s at ix %d", sdir.Name(), ix)
			if sdir == CD {
				break
			}
			dir.AddNode(sdir)
		}

		return dir, ix
	}

	fs, _ := recurse(0)
	return fs
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	fs := rebuildFS(lines)

	println(printFS(fs))

	sumSize := 0
	var visit func(*Dir)
	visit = func(dir *Dir) {
		if ss := dir.Size(); ss < 100000 {
			sumSize += ss
		}
		for _, node := range dir.nodes {
			if sdir, ok := node.(*Dir); ok {
				visit(sdir)
			}
		}
	}

	visit(fs.(*Dir))

	printf("sum size: %d", sumSize)

	totalSize := 70000000
	requiredSize := 30000000
	rootSize := fs.Size()
	spaceLeft := totalSize - rootSize

	minDir := fs
	minSize := rootSize
	visit = func(dir *Dir) {
		if ss := dir.Size(); spaceLeft+ss >= requiredSize {
			if ss < minSize {
				minSize = ss
				minDir = dir
			}
		}
		for _, node := range dir.nodes {
			if sdir, ok := node.(*Dir); ok {
				visit(sdir)
			}
		}
	}
	visit(fs.(*Dir))
	printf("we need to remove dir %s of size %d", minDir.Name(), minSize)
}
