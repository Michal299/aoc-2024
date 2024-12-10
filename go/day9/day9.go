package day9

import (
	"container/list"
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
	block           *list.Element
	positionInBlock int
}

type filesystem struct {
	representation *list.List
}

func parseFilesystem(input string) (fs filesystem) {
	elems := strings.Split(input, "")
	representations := list.New()
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
		if b.length <= 0 {
			continue
		}
		representations.PushBack(b)
	}

	fs.representation = representations
	return
}

func splitBlock(size int, b block) (main, remaining block) {
	main = block{id: b.id, length: size, bt: b.bt}
	remaining = block{id: b.id, length: b.length - size, bt: b.bt}
	return
}

func (f *filesystem) fillBlockFromEnd(toFill *list.Element) {
	beforeToFill := toFill.Prev()
	f.representation.Remove(toFill)

	for toFill != nil {
		spaceDifference := toFill.Value.(block).length - f.representation.Back().Value.(block).length
		var filler block
		if spaceDifference == 0 {
			fillerElem := f.representation.Back()
			f.representation.Remove(fillerElem)
			filler = fillerElem.Value.(block)
			toFill = nil
			beforeToFill = f.representation.InsertAfter(filler, beforeToFill)
		} else if spaceDifference < 0 {
			end := f.representation.Back()
			f.representation.Remove(end)
			var remaining block
			filler, remaining = splitBlock(toFill.Value.(block).length, end.Value.(block))
			f.representation.PushBack(remaining)
			toFill = nil
			beforeToFill = f.representation.InsertAfter(filler, beforeToFill)
		} else if spaceDifference > 0 {
			fillerElem := f.representation.Back()
			f.representation.Remove(fillerElem)
			filler = fillerElem.Value.(block)
			remainingFreeBlock := block{id: toFill.Value.(block).id, bt: EMPTY, length: spaceDifference}
			beforeToFill = f.representation.InsertAfter(filler, beforeToFill)
			toFill = f.representation.InsertAfter(remainingFreeBlock, beforeToFill)
		}
	}
}

func (f *filesystem) fragmentation() {
	currentFront := f.representation.Front()

	for currentFront != nil {
		if currentFront.Value.(block).bt == FILE {
			currentFront = currentFront.Next()
			continue
		}

		toFill := currentFront
		currentFront = toFill.Next()
		f.fillBlockFromEnd(toFill)
	}
}

func (f *filesystem) fragmentation2() {

}

func (f *filesystem) checksum() (s int) {
	currentPosition := position{block: f.representation.Front(), positionInBlock: 0}
	for i := 0; true; i++ {
		if currentPosition.positionInBlock > currentPosition.block.Value.(block).length {
			nextBlock := currentPosition.block.Next()
			if nextBlock == nil {
				break
			}
			currentPosition.block = nextBlock
			currentPosition.positionInBlock = 0
		}
		s += currentPosition.block.Value.(block).id * i
	}
	return 0
}

func Part1(input []string) int {
	fs := parseFilesystem(input[0])
	fs.fragmentation()
	return fs.checksum()
}

func Part2(input []string) int {
	fs := parseFilesystem(input[0])
	fs.fragmentation2()
	return fs.checksum()
}
