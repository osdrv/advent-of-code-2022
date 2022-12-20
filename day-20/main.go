package main

import (
	"bytes"
	"os"
	"strconv"
)

type Node struct {
	ix, num    int
	next, prev *Node
}

func NewNode(ix, num int) *Node {
	return &Node{ix: ix, num: num}
}

func numsToList(nums []int) *Node {
	var head, zero, prev *Node
	for ix, num := range nums {
		node := NewNode(ix, num)
		if head == nil {
			head = node
		}
		if num == 0 {
			zero = node
		}
		if prev != nil {
			node.prev = prev
			prev.next = node
		}
		prev = node
	}
	head.prev = prev
	prev.next = head

	return zero
}

func listToNums(list *Node) []int {
	ptr := list
	nums := make([]int, 0, 1)
	for {
		nums = append(nums, ptr.num)
		ptr = ptr.next
		if ptr == list {
			break
		}
	}

	return nums
}

func findNumByIx(list *Node, ix int) *Node {
	ptr := list
	for {
		if ptr.ix == ix {
			return ptr
		}
		ptr = ptr.next
		if ptr == list {
			break
		}
	}
	return nil
}

func moveNum(list *Node, ix int, cycle int) *Node {
	ptr := findNumByIx(list, ix)
	num := ptr.num
	if num == 0 {
		// list starts with a zero-node
		return list
	}
	prev, next := ptr.prev, ptr.next

	prev.next = next
	next.prev = prev

	for i := 0; i < abs(num)%cycle; i++ {
		if num > 0 {
			prev = prev.next
			next = next.next
		} else {
			prev = prev.prev
			next = next.prev
		}
	}

	prev.next = ptr
	ptr.prev = prev
	next.prev = ptr
	ptr.next = next

	return list
}

func printList(list *Node) string {
	var b bytes.Buffer
	ptr := list
	for {
		if b.Len() != 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(ptr.num))
		ptr = ptr.next
		if ptr == list {
			break
		}
	}
	return b.String()
}

const (
	KEY int = 811589153
)

func main() {
	f, err := os.Open("INPUT")
	noerr(err)
	defer f.Close()

	lines := readLines(f)
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		nums = append(nums, parseInt(line))
	}

	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	list := numsToList(nums)

	if DEBUG {
		debugf("initial nums: [%s]", printList(list))
	}

	cycle := len(nums) - 1

	for ix, num := range nums {
		debugf("moving num %d", num)
		list = moveNum(list, ix, cycle)
		if DEBUG {
			debugf("nums: [%s]", printList(list))
		}
	}

	printf("Part 1")

	nums = listToNums(list)

	p1, p2, p3 := (1000)%len(nums), (2000)%len(nums), (3000)%len(nums)

	printf("p1: %d, p2: %d, p3: %d, sum: %d", nums[p1], nums[p2], nums[p3], nums[p1]+nums[p2]+nums[p3])

	for ix := range nums2 {
		nums2[ix] *= KEY
	}

	debugf("nums2: %+v", nums2)

	list2 := numsToList(nums2)

	for j := 0; j < 10; j++ {
		for ix := range nums2 {
			list2 = moveNum(list2, ix, cycle)
		}
	}

	printf("Part 2")

	nums2 = listToNums(list2)
	printf("p1: %d, p2: %d, p3: %d, sum: %d", nums2[p1], nums2[p2], nums2[p3], nums2[p1]+nums2[p2]+nums2[p3])
}
