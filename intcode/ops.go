package intcode

func opAdd(c *intcodeComputerState, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = getParam(c, i, pm)
		if crashed {
			return
		}
	}
	sum := p[0] + p[1]
	targetIndex := c.tape[c.ptr+3]
	c.tape[targetIndex] = sum
	// fmt.Println("ADD wrote", sum, "to index", targetIndex)

	c.ptr += 4
}

func opMul(c *intcodeComputerState, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = getParam(c, i, pm)
		if crashed {
			return
		}
	}
	product := p[0] * p[1]
	targetIndex := c.tape[c.ptr+3]
	c.tape[targetIndex] = product
	// fmt.Println("MUL wrote", product, "to index", targetIndex)

	c.ptr += 4
}

func opInput(c *intcodeComputerState, pm [3]ParameterMode) {
	var input int
	input, c.inputs = c.inputs[0], c.inputs[1:]
	c.tape[c.tape[c.ptr+1]] = input
	c.ptr += 2
}

func opOutput(c *intcodeComputerState, pm [3]ParameterMode) {
	output, crashed := getParam(c, 0, pm)
	if crashed {
		return
	}

	c.outputs = append(c.outputs, output)
	c.ptr += 2
}

func subopJump(c *intcodeComputerState, pm [3]ParameterMode) (int, int, bool) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = getParam(c, i, pm)
		if crashed {
			return 0, 0, true
		}
	}
	return p[0], p[1], false
}

func opJumpIfTrue(c *intcodeComputerState, pm [3]ParameterMode) {
	predicate, jumpptr, crashed := subopJump(c, pm)
	if crashed {
		return
	}
	if predicate != 0 {
		c.ptr = jumpptr
	} else {
		c.ptr += 3
	}
}

func opJumpIfFalse(c *intcodeComputerState, pm [3]ParameterMode) {
	predicate, jumpptr, crashed := subopJump(c, pm)
	if crashed {
		return
	}
	if predicate == 0 {
		c.ptr = jumpptr
	} else {
		c.ptr += 3
	}
}

func opLT(c *intcodeComputerState, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = getParam(c, i, pm)
		if crashed {
			return
		}
	}
	targetIndex := c.tape[c.ptr+3]
	if p[0] < p[1] {
		c.tape[targetIndex] = 1
	} else {
		c.tape[targetIndex] = 0
	}
	c.ptr += 4
}

func opEQ(c *intcodeComputerState, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = getParam(c, i, pm)
		if crashed {
			return
		}
	}
	targetIndex := c.tape[c.ptr+3]
	if p[0] == p[1] {
		c.tape[targetIndex] = 1
	} else {
		c.tape[targetIndex] = 0
	}
	c.ptr += 4
}

func opHalt(c *intcodeComputerState, pm [3]ParameterMode) {
	c.ptr++
	c.status = Halted
}

var icOps = map[int](func(*intcodeComputerState, [3]ParameterMode)){
	1:  opAdd,
	2:  opMul,
	3:  opInput,
	4:  opOutput,
	5:  opJumpIfTrue,
	6:  opJumpIfFalse,
	7:  opLT,
	8:  opEQ,
	99: opHalt,
}
