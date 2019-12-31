package intcode

func opAdd(tape []int, ptr *int) bool {
	sum := tape[tape[*ptr+1]] + tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = sum
	// fmt.Println("ADD wrote", sum, "to index", targetIndex)

	*ptr += 4
	return false
}

func opMul(tape []int, ptr *int) bool {
	product := tape[tape[*ptr+1]] * tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = product
	// fmt.Println("MUL wrote", product, "to index", targetIndex)

	*ptr += 4
	return false
}

func opHalt(tape []int, ptr *int) bool {
	*ptr++
	return true
}

// Processor for Intcode tapes
func Processor(tape []int) {
	icOps := map[int](func([]int, *int) bool){
		1:  opAdd,
		2:  opMul,
		99: opHalt,
	}
	ptr := 0
	halt := false
	for !halt {
		halt = icOps[tape[ptr]](tape, &ptr)
	}
}
