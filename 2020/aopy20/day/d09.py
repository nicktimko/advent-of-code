"""
Day 9: Encoding Error
https://adventofcode.com/2020/day/9
"""

from ..io import read_ints


TODAY = 9

SPAN_SIZE = 25


def find_addends(xmas_tape, index, span_size=SPAN_SIZE):
    niq = xmas_tape[index]
    span = xmas_tape[index - span_size : index]
    for n in span:
        if n * 2 >= niq:
            continue
        diff = niq - n
        if diff in span:
            return n, diff
    return None


def run():
    xmas = list(read_ints(TODAY))

    ptr = SPAN_SIZE
    while True:
        if find_addends(xmas, ptr) is None:
            break
        ptr += 1

    print("Part 1:", xmas[ptr])

    target_sum = xmas[ptr]
    acc = 0
    i_lead = 0
    i_follow = 0
    while acc != target_sum:
        if acc < target_sum:
            acc += xmas[i_lead]
            i_lead += 1
        elif acc > target_sum:
            acc -= xmas[i_follow]
            i_follow += 1

    print("Part 2:", max(xmas[i_follow:i_lead]) + min(xmas[i_follow:i_lead]))
    # amazingly not off by one? maybe just the test case didn't rely on
    # it that tightly
