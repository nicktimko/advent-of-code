package intcode

func opAdd(c *State, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = c.getParam(i, pm)
		if crashed {
			return
		}
	}
	sum := p[0] + p[1]
	targetIndex := c.tape[c.ptr+3]
	c.setTapeIndex(targetIndex, sum)

	c.ptr += 4
}

func opMul(c *State, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = c.getParam(i, pm)
		if crashed {
			return
		}
	}
	product := p[0] * p[1]
	targetIndex := c.tape[c.ptr+3]
	c.setTapeIndex(targetIndex, product)

	c.ptr += 4
}

func opInput(c *State, pm [3]ParameterMode) {
	var input int
	input, c.inputs = c.inputs[0], c.inputs[1:]

	// input parameter is always immediate
	pm[0] = modeImmediate

	targetIndex, crashed := c.getParam(0, pm)
	if crashed {
		return
	}
	c.setTapeIndex(targetIndex, input)
	c.ptr += 2
}

func opOutput(c *State, pm [3]ParameterMode) {
	output, crashed := c.getParam(0, pm)
	if crashed {
		return
	}

	c.outputs = append(c.outputs, output)
	c.ptr += 2
}

func subopJump(c *State, pm [3]ParameterMode) (int, int, bool) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = c.getParam(i, pm)
		if crashed {
			return 0, 0, true
		}
	}
	return p[0], p[1], false
}

func opJumpIfTrue(c *State, pm [3]ParameterMode) {
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

func opJumpIfFalse(c *State, pm [3]ParameterMode) {
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

func opLT(c *State, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = c.getParam(i, pm)
		if crashed {
			return
		}
	}
	targetIndex := c.tape[c.ptr+3]
	if p[0] < p[1] {
		c.setTapeIndex(targetIndex, 1)
	} else {
		c.setTapeIndex(targetIndex, 0)
	}
	c.ptr += 4
}

func opEQ(c *State, pm [3]ParameterMode) {
	var p [2]int
	var crashed bool
	for i := range p {
		p[i], crashed = c.getParam(i, pm)
		if crashed {
			return
		}
	}
	targetIndex := c.tape[c.ptr+3]
	if p[0] == p[1] {
		c.setTapeIndex(targetIndex, 1)
	} else {
		c.setTapeIndex(targetIndex, 0)
	}
	c.ptr += 4
}

func opBaseAdjust(c *State, pm [3]ParameterMode) {
	adjustment, crashed := c.getParam(0, pm)
	if crashed {
		return
	}
	c.relativeBase += adjustment
	c.ptr += 2
}

func opHalt(c *State, pm [3]ParameterMode) {
	c.ptr++
	c.status = Halted
}

// opcodes
const opcAdd = 1
const opcMul = 2
const opcInput = 3
const opcOutput = 4
const opcJumpIfTrue = 5
const opcJumpIfFalse = 6
const opcLT = 7
const opcEQ = 8
const opcBaseAdjust = 9
const opcHalt = 99

var icOps = map[int](func(*State, [3]ParameterMode)){
	opcAdd:         opAdd,
	opcMul:         opMul,
	opcInput:       opInput,
	opcOutput:      opOutput,
	opcJumpIfTrue:  opJumpIfTrue,
	opcJumpIfFalse: opJumpIfFalse,
	opcLT:          opLT,
	opcEQ:          opEQ,
	opcBaseAdjust:  opBaseAdjust,
	opcHalt:        opHalt,
}
