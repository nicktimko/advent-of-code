"""
Day 1: Expense Report
https://adventofcode.com/2020/day/1
"""

import itertools
import logging

from ..io import read_ints

L = logging.getLogger(__name__)


def run():
    entries = list(read_ints(1))
    entries.sort()

    i_lo = 0
    i_hi = len(entries) - 1

    while True:
        sum_ = entries[i_lo] + entries[i_hi]

        if sum_ == 2020:
            break
        elif sum_ < 2020:
            i_lo += 1
        else:  # >2020
            i_hi -= 1

    lo = entries[i_lo]
    hi = entries[i_hi]

    L.debug(f"{lo} + {hi} = {lo+hi}")
    L.debug(f"{lo} * {hi} = {lo*hi}")
    print("Part 1:", lo * hi)

    # optimize search space by ignoring values that can't work, i.e. those
    # with values so high that when added to the _lowest_ two values, are
    # always going to be too much.
    lowest_pair_sum = entries[0] + entries[1]
    viable_entries = [i for i in entries if i + lowest_pair_sum <= 2020]

    viable_entry_set = set(viable_entries)
    for a, b in itertools.combinations(viable_entries, r=2):
        c = 2020 - a - b
        if c in viable_entry_set:
            print("Part 2:", a * b * c)
            break
    else:
        raise RuntimeError("failed to find the triple")
