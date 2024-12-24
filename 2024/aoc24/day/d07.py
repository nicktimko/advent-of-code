"""
Day 7:
https://adventofcode.com/2024/day/7
"""

import itertools
import logging
import operator

from ..io import read_lines


L = logging.getLogger(__name__)
TODAY = 7


def concat_number(a, b):
    return int(str(a) + str(b))


def try_operators(test_value, operands, operators):
    operator_combos = itertools.product(operators, repeat=len(operands) - 1)
    for operator_combo in operator_combos:
        operator_iter = iter(operator_combo)
        stack = list(reversed(operands))
        while len(stack) > 1:
            stack.append(next(operator_iter)(stack.pop(), stack.pop()))
        if stack[0] == test_value:
            return operator_combo
    return None

def run():
    L.info("starting day %d", TODAY)

    # ingest file
    data = read_lines(TODAY)
    equations = []
    for equation in data:
        test_value, operands = equation.split(":")
        test_value = int(test_value.strip())
        operands = [int(x) for x in operands.strip().split()]
        equations.append((test_value, operands))

    p1_ans = 0
    p2_ans = 0
    for test_value, operands in equations:
        operators = [operator.add, operator.mul]
        if try_operators(test_value, operands, operators) is not None:
            p1_ans += test_value
            p2_ans += test_value  # can skip if this works
        else:
            operators = [operator.add, operator.mul, concat_number]
            if try_operators(test_value, operands, operators) is not None:
                p2_ans += test_value

    yield p1_ans
    yield p2_ans
