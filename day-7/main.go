package main

import (
	"bytes"
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
}

func NewDir(name string) *Dir {
	return &Dir{name: name, nodes: make([]Node, 0, 1)}
}

func (d *Dir) Name() string {
	return d.name
}

func (d *Dir) recSize(visited map[*Dir]bool) int {
	if visited[d] {
		debugf("cycle detected on node %s", d.name)
		return 0
	}
	visited[d] = true

	size := 0
	for _, node := range d.nodes {
		if sdir, ok := node.(*Dir); ok {
			size += sdir.recSize(visited)
		} else {
			size += node.Size()
		}
	}

	delete(visited, d)

	return size
}

func (d *Dir) Size() int {
	return d.recSize(make(map[*Dir]bool))
}

func (d *Dir) AddNode(node Node) {
	d.nodes = append(d.nodes, node)
}

func fullPath(path []Node, dirName string) string {
	var b bytes.Buffer
	for _, n := range path {
		b.WriteString(n.Name())
		if n.Name() != "/" {
			b.WriteByte('/')
		}
	}
	b.WriteString(dirName)
	return b.String()
}

func printFS(fs Node) string {
	var visit func(Node, map[Node]bool, string) []string
	visit = func(node Node, visited map[Node]bool, pref string) []string {
		res := make([]string, 0, 1)
		if visited[node] {
			res = append(res, fmt.Sprintf("%s<symlink -> %s>", pref, node.Name()))
		} else {
			if dir, ok := node.(*Dir); ok {
				visited[node] = true
				res = append(res, fmt.Sprintf("%s%s", pref+"└-", dir.Name()))
				for _, node := range dir.nodes {
					for _, sp := range visit(node, visited, "  ") {
						res = append(res, pref+sp)
					}
				}
				delete(visited, node)
			} else {
				res = append(res, fmt.Sprintf("%s%s %d", pref+"└-", node.Name(), node.Size()))
			}
		}
		return res
	}

	pp := visit(fs, make(map[Node]bool), "")
	return strings.Join(pp, "\n")
}

func traverseFS(cmds []string) (Node, map[string]Node) {
	ix := 0
	dirs := make(map[string]Node)
	path := make([]Node, 0, 1)
Cmd:
	for ix < len(cmds) {
		name := cmds[ix][5:]
		ix++
		if name == ".." {
			path = path[:len(path)-1]
			continue Cmd
		}
		fp := fullPath(path, name)
		debugf("visiting dir: %s", fp)
		if _, ok := dirs[fp]; !ok {
			dirs[fp] = NewDir(fp)
		}
		dir := dirs[fp]
		path = append(path, dir)

		ix++ // read ls
		for ix < len(cmds) && cmds[ix][0] != '$' {
			if startsWith(cmds[ix], "dir") {
				sname := cmds[ix][4:]
				sfp := fullPath(path, sname)
				if _, ok := dirs[sfp]; !ok {
					dirs[sfp] = NewDir(sname)
				}
				sdir := dirs[sfp]
				(dir.(*Dir)).AddNode(sdir)
			} else {
				ptr := 0
				var size int
				size, ptr = readInt(cmds[ix], ptr)
				file := NewFile(cmds[ix][ptr+1:], size)
				debugf("visiting file %s of size: %d", cmds[ix][ptr+1:], size)
				(dir.(*Dir)).AddNode(file)
			}
			ix++
		}
	}

	return path[0], dirs
}

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)

	fs, dirs := traverseFS(lines)

	println(printFS(fs))

	sumSize := 0
	for name, dir := range dirs {
		if ss := dir.Size(); ss <= 100000 {
			debugf("dir %s is below than 100000 in size", name)
			sumSize += ss
		}
	}

	printf("sum size: %d", sumSize)

	totalSize := 70000000
	requiredSize := 30000000
	rootSize := fs.Size()
	spaceLeft := totalSize - rootSize

	debugf("space left: %d, required: %d", spaceLeft, requiredSize)

	minDir := fs
	minSize := rootSize
	for _, dir := range dirs {
		ds := dir.Size()
		if spaceLeft+ds >= requiredSize {
			if ds < minSize {
				minSize = ds
				minDir = dir
			}
		}
	}

	printf("we need to remove dir %s of size %d", minDir.Name(), minSize)
}
