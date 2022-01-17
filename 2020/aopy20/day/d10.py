"""
Day 10:
https://adventofcode.com/2020/day/10
"""

import collections

from ..io import read_ints


TODAY = 10


def run():
    adapters = list(read_ints(TODAY))
    adapters.sort()

    # start at 0, and final device +3 from max
    adapters = [0] + adapters + [adapters[-1] + 3]

    diffs = collections.Counter()
    prev = adapters[0]
    for adapter in adapters[1:]:
        diff = adapter - prev
        diffs[diff] += 1
        prev = adapter

    print(diffs)

    print("Part 1:", diffs[1] * diffs[3])

    # graph = {}
    nic = {0: 1}  # node_inbound_counts
    for n, adapter in enumerate(adapters[1:], start=1):
        possible_sources = [a for a in adapters[max(0, n-3):n] if a >= adapter - 3]
        # print(possible_sources, adapter)
        # if n > 10: break

        nic[adapter] = sum(nic[s] for s in possible_sources)

    print("Part 2:", print(nic[adapters[-1]]))
