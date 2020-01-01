package intcode

// ParameterMode ...
type ParameterMode int

const (
	modePosition  ParameterMode = 0
	modeImmediate ParameterMode = 1
)

// OpCode ...
type OpCode struct {
	op int
	pm [3]ParameterMode
}

func decodeOp(opcode int) OpCode {
	var oc OpCode

	oc.op = opcode % 100
	opcode = opcode / 100

	for i := 0; i < 3; i++ {
		oc.pm[i] = ParameterMode(opcode % 10)
		opcode = opcode / 10
	}

	return oc
}

func getParam(pm [3]ParameterMode) {

}

func opAdd(tape []int, ptr *int, pm [3]ParameterMode) bool {
	sum := tape[tape[*ptr+1]] + tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = sum
	// fmt.Println("ADD wrote", sum, "to index", targetIndex)

	*ptr += 4
	return false
}

func opMul(tape []int, ptr *int, pm [3]ParameterMode) bool {
	product := tape[tape[*ptr+1]] * tape[tape[*ptr+2]]
	targetIndex := tape[*ptr+3]
	tape[targetIndex] = product
	// fmt.Println("MUL wrote", product, "to index", targetIndex)

	*ptr += 4
	return false
}

func opInput(tape []int, ptr *int, pm [3]ParameterMode) bool {
	return false
}

func opOutput(tape []int, ptr *int, pm [3]ParameterMode) bool {
	return false
}

func opHalt(tape []int, ptr *int, pm [3]ParameterMode) bool {
	*ptr++
	return true
}

var icOps = map[int](func([]int, *int, [3]ParameterMode) bool){
	1:  opAdd,
	2:  opMul,
	3:  opInput,
	4:  opOutput,
	99: opHalt,
}

// ProcessInstruction (single) for Intcode tapes
func ProcessInstruction(tape []int, ptr *int) bool {
	op := decodeOp(tape[*ptr])
	halt := icOps[op.op](tape, ptr, op.pm)
	return halt
}

// Processor for Intcode tapes
func Processor(tape []int) {
	ptr := 0
	halt := false
	for !halt {
		halt = ProcessInstruction(tape, &ptr)
	}
}
