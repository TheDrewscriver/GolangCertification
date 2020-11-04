package test

func test() {
	//Underlying array
	arr := [...]string{"a", "b", "c", "d", "e"}
	//Slice 1
	s1 := arr[1:3]
	//Slice 2
	var s2 = arr[1:2]
	print(len(s1))
	print(cap(s2))
	//Slice Literals
	//Creates underlying array and slice and they will have same
	//len and cap
	sli := []int{1, 2, 3}
	print(cap(sli))

	/* VARIABLE SLICES  make() */
	//type, length=capacity
	sliMake2 := make([]int, 10)
	print(cap(sliMake2))
	// type, length, capacity
	sliMake3 := make([]int, 10, 15)
	print(cap(sliMake3))
	//Append can add to the slice
	sliMake2 = append(sliMake2, 100)
}
