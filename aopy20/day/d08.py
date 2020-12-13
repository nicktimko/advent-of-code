"""
Day 8: Handheld Halting
https://adventofcode.com/2020/day/8
"""

import enum
from typing import NamedTuple

from ..io import read_lines


TODAY = 8


class Operation(enum.Enum):
    NOP = "nop"
    ACC = "acc"
    JMP = "jmp"


class Instruction(NamedTuple):
    operation: Operation
    argument: int

    def __str__(self):
        return "<{} {:+d}>".format(self.operation.value, self.argument)


def process_line(line: str) -> Instruction:
    op, arg = line.split()
    op = Operation(op)
    arg = int(arg)
    return Instruction(op, arg)


class Debugger:
    def __init__(self, lines):
        self.mem = [process_line(l) for l in lines]
        self.ic = 0
        self.acc = 0

        self.exec_count = [0 for _ in self.mem]

    def __repr__(self):
        try:
            inst = self.mem[self.ic]
        except KeyError:
            inst = "SEGFAULT"
        return "<{} acc={!r} ic={!r} inst={}>".format(
            self.__class__.__name__, self.acc, self.ic, inst
        )

    def step(self):
        instruction = self.mem[self.ic]
        self.exec_count[self.ic] += 1

        if instruction.operation is Operation.NOP:
            self.ic += 1
        elif instruction.operation is Operation.ACC:
            self.acc += instruction.argument
            self.ic += 1
        elif instruction.operation is Operation.JMP:
            self.ic += instruction.argument
        else:
            raise RuntimeError("crash, unhandled op", instruction.operation)


def run():
    dbg = Debugger(read_lines(TODAY))

    while dbg.exec_count[dbg.ic] != 1:
        # print(dbg)
        dbg.step()

    print("Part 1:", dbg.acc)

    # dbg.mem[dbg.ic] = Instruction
    # while True:
    #     dbg.step()
