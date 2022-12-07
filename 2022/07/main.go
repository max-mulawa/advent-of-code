package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Dir struct {
	parent   *Dir
	name     string
	children []*Dir
	files    []*File
}

type File struct {
	parent *Dir
	name   string
	size   int
}

const (
	parentDir = ".."
)

type cdAction struct {
	destDir string
	pwd     *Dir
}

type lsAction struct {
	pwd *Dir
}

type Action interface {
	execute(string) *Dir
}

func NewAction(pwd *Dir, line string) Action {
	tokens := strings.Split(line, " ")
	actionType := tokens[1]

	switch actionType {
	case "ls":
		return &lsAction{pwd: pwd}
	case "cd":
		return &cdAction{pwd: pwd, destDir: tokens[2]}
	}

	panic(line)
}

func (a *lsAction) execute(line string) *Dir {
	if strings.HasPrefix(line, "$") {
		return a.pwd
	}

	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "dir":
		a.pwd.AddDir(&Dir{name: tokens[1]})
	default:
		fileSize, _ := strconv.Atoi(tokens[0])
		a.pwd.AddFile(&File{name: tokens[1], size: fileSize})
	}

	return a.pwd
}

func (a *cdAction) execute(line string) *Dir {
	if a.destDir == parentDir {
		return a.pwd.parent
	}

	if a.pwd == nil {
		return &Dir{name: a.destDir}
	}

	for _, d := range a.pwd.children {
		if d.name == a.destDir {
			return d
		}
	}
	panic(line)
}

func (d *Dir) AddFile(f *File) {
	f.parent = d
	d.files = append(d.files, f)
}

func (d *Dir) AddDir(dir *Dir) {
	dir.parent = d
	d.children = append(d.children, dir)
}

func (d *Dir) GetTotalSize() int {
	total := 0
	for _, f := range d.files {
		total += f.size
	}
	for _, c := range d.children {
		total += c.GetTotalSize()
	}
	return total
}

func main() {
	payload, err := os.ReadFile("input.txt")
	// payload, err := os.ReadFile("input-test.txt")
	if err != nil {
		panic(err)
	}

	var pwd *Dir

	r := bufio.NewReader(bytes.NewReader(payload))
	var action Action
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		if strings.HasPrefix(line, "$") {
			action = NewAction(pwd, line)
		}

		pwd = action.execute(line)
	}

	// go to root level
	for {
		if pwd.parent == nil {
			break
		}
		pwd = pwd.parent
	}

	root := pwd
	atMostTotal := calcTotal(root, 100000)
	fmt.Printf("at most total is: %d\n", atMostTotal)
	fmt.Printf("size %s: %d\n", root.name, root.GetTotalSize())

	diskSpace := 70000000
	expectedAtLeastFreeSpace := 30000000
	usedSpace := root.GetTotalSize()
	atLeastSpaceToFreeUp := expectedAtLeastFreeSpace - (diskSpace - usedSpace)

	fmt.Printf("disk space to free up: %d\n", atLeastSpaceToFreeUp)

	calcDirSizeToFree(root, atLeastSpaceToFreeUp)

	sort.Ints(folderSizes)
	fmt.Printf("folder with total size %d need to be deleted\n", folderSizes[0])
}

var folderSizes = []int{}

func calcDirSizeToFree(dir *Dir, atMost int) {
	total := dir.GetTotalSize()
	if total >= atMost {
		folderSizes = append(folderSizes, total)
	}
	for _, c := range dir.children {
		calcDirSizeToFree(c, atMost)
	}
}

func calcTotal(dir *Dir, atMost int) int {
	atMostTotal := 0
	total := dir.GetTotalSize()
	if total <= atMost {
		atMostTotal += total
	}
	for _, c := range dir.children {
		atMostTotal += calcTotal(c, atMost)
	}
	return atMostTotal
}
