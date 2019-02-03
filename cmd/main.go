package main

import "fmt"

const (
	sparta = "sparta"
	_      = 1 << iota
	t1
	t2
	t3
	t4
	t5
)

func main() {
	fmt.Printf("This is %s\n", sparta)
	fmt.Printf("The constants are %v,%v,%v,%v,%v of type %T,%T,%T,%T,%T\n", t1, t2, t3, t4, t5, t1, t2, t3, t4, t5)
	original := []int{1: 10, 2: 20, 3: 30, 10: 40}
	fmt.Printf("The slice before chopping %v\n", original)
	chopped := original[0:5]
	fmt.Printf("The original slice after chopping %v\n", original)
	fmt.Printf("The chopped slice %v\n", chopped)
	chopped = append(chopped, 1)
	fmt.Printf("The original slice after appending to chopped %v\n", original)
	fmt.Printf("The chopped slice after appending to chopped %v\n", chopped)
	original[5] = 2
	fmt.Printf("The original slice after setting 5th index to original %v\n", original)
	fmt.Printf("The chopped slice after setting 5th index to original %v\n", chopped)
	chopped = append(chopped, []int{3, 4, 5, 6, 7, 8, 9}...)
	fmt.Printf("The original slice after appending an array to chopped %v\n", original)
	fmt.Printf("The chopped  slice after appending an array to chopped %v\n", chopped)
	newSlice := original[1:2:2]
	fmt.Printf("The original slice after 3 index chopping %v\n", original)
	fmt.Printf("The new slice after 3 index chopping %v\n", newSlice)
	newSlice = append(newSlice, 2)
	fmt.Printf("The original slice after 3 index chopping and appending %v\n", original)
	fmt.Printf("The new slice after 3 index chopping and appending %v\n", newSlice)
	anotherSlice := original[1:2:3]
	fmt.Printf("The original slice after 3 index chopping %v\n", original)
	fmt.Printf("anotherSlice after 3 index chopping %v\n", anotherSlice)
	anotherSlice = append(anotherSlice, 2)
	fmt.Printf("The original slice after 3 index chopping and appending %v\n", original)
	fmt.Printf("anotherSlice after 3 index chopping and appending %v\n", anotherSlice)
	tt := test(1)
	adherence(&tt)
	tt1 := test1(2)
	adherence(tt1)

}

type tester interface {
	testonce() bool
	testntime(n int) bool
}

type test int

func (t *test) testonce() bool {
	te := test(1)
	t = &te
	return false
}

func (t *test) testntime(n int) bool {
	te := test(1)
	t = &te
	return false
}

func adherence(t tester) bool {
	return false
}

type test1 int

func (t test1) testonce() bool {
	return false
}

func (t test1) testntime(n int) bool {
	return false
}
