"""
Day 3:
https://adventofcode.com/2024/day/3
"""

import re

from ..io import read_lines


TODAY = 3

P2_REGEX = re.compile(
    r"""
        mul\(
            (\d{1,3})
            ,
            (\d{1,3})
        \)
        | don't\(\)
        | do\(\)
    """,
    flags=re.VERBOSE,
)


def run():
    data = "".join(read_lines(TODAY))

    part1_accum = 0
    for match in re.finditer(r"mul\((\d{1,3}),(\d{1,3})\)", data):
        n1, n2 = match.groups()
        part1_accum += int(n1) * int(n2)

    yield part1_accum

    part2_accum = 0
    enabled = True
    for match in P2_REGEX.finditer(data):
        if match.group() == "do()":
            enabled = True
        elif match.group() == "don't()":
            enabled = False
        elif enabled:
            n1, n2 = match.groups()
            part2_accum += int(n1) * int(n2)

    yield part2_accum
