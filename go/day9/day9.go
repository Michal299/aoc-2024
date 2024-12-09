package day9

import (
	"strconv"
	"strings"
)

type blocktype int

const (
	FILE  blocktype = 0
	EMPTY blocktype = 1
)

type block struct {
	bt     blocktype
	length int
	id     int
}

type position struct {
	blockIndex      int
	positionInBlock int
}

type filesystem struct {
	front          position
	tail           position
	representation []block
	absoluteId     int
}

func parseFilesystem(input string) (fs filesystem) {
	elems := strings.Split(input, "")
	var representations []block
	blockId := 0
	for i, elem := range elems {
		num, _ := strconv.Atoi(elem)
		b := block{}
		b.length = num

		if i%2 == 0 {
			b.id = blockId
			blockId++
			b.bt = FILE
		} else {
			b.id = -1
			b.bt = EMPTY
		}
		representations = append(representations, b)

	}
	fs.representation = representations
	fs.front = position{blockIndex: 0, positionInBlock: 0}
	fs.tail = position{blockIndex: len(representations) - 1, positionInBlock: fs.representation[len(representations)-1].length - 1}
	fs.absoluteId = 0
	return
}

func (f *filesystem) isEof() bool {
	if f.front.blockIndex == f.tail.blockIndex {
		return f.front.positionInBlock > f.tail.positionInBlock
	}
	return f.front.blockIndex > f.tail.blockIndex
}

func (f *filesystem) getFrontBlock() block {
	frontBlock := f.representation[f.front.blockIndex]
	if f.front.positionInBlock >= frontBlock.length {
		f.front.blockIndex++
		for f.representation[f.front.blockIndex].bt == EMPTY && f.representation[f.front.blockIndex].length == 0 {
			f.front.blockIndex++
		}
		f.front.positionInBlock = 0
		frontBlock = f.representation[f.front.blockIndex]
	}
	return frontBlock
}

func (f *filesystem) getNextPosition() (pos, fileId int) {
	frontBlock := f.getFrontBlock()
	tailBlock := f.getTailBlock()

	pos = f.absoluteId
	f.absoluteId++

	switch frontBlock.bt {
	case FILE:
		fileId = frontBlock.id
	case EMPTY:
		fileId = tailBlock.id
		f.tail.positionInBlock--
	}
	f.front.positionInBlock++
	return
}

func (f *filesystem) getTailBlock() block {
	tailBlock := f.representation[f.tail.blockIndex]
	if f.tail.positionInBlock < 0 {
		f.tail.blockIndex--
		for f.representation[f.tail.blockIndex].bt != FILE {
			f.tail.blockIndex--
		}
		tailBlock = f.representation[f.tail.blockIndex]
		f.tail.positionInBlock = tailBlock.length - 1
	}
	return tailBlock
}

func (f *filesystem) checksum() (s int) {
	for !f.isEof() {
		pos, id := f.getNextPosition()
		s += pos * id
	}
	return
}

func Part1(input []string) int {
	fs := parseFilesystem(input[0])
	return fs.checksum()
}

func Part2(input []string) int {
	return 0
}
