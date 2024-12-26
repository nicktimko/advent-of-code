"""
Day 8:
https://adventofcode.com/2024/day/8
"""

import itertools
import logging
import string

import numpy as np

from ..io import read_lines


L = logging.getLogger(__name__)
TODAY = 8


def find_antinodes(locs):
    for l1, l2 in (np.array(l) for l in itertools.combinations(locs, r=2)):
        delta = l2 - l1
        yield tuple(int(x) for x in l1 - delta)
        yield tuple(int(x) for x in l2 + delta)


def filter_locs(x_dim, y_dim, locs):
    for x, y in locs:
        if x not in range(x_dim):
            continue
        if y not in range(y_dim):
            continue
        yield x, y


def run():
    L.info("starting day %d", TODAY)
    data = read_lines(TODAY)

    locations = {}
    freqs = string.ascii_letters + string.digits
    for n_row, row in enumerate(data):
        for n_col, cell in enumerate(row):
            if cell in freqs:
                locations[(n_col, n_row)] = cell

    field_dims = n_col + 1, n_row + 1

    L.info("loaded %d antennas", len(locations))
    L.info("field size: %r", field_dims)
    L.info("signals: %r", "".join(sorted(set(locations.values()))))

    L.debug("locations: %r", locations)

    freqs = {}
    for loc, freq in locations.items():
        freqs.setdefault(freq, [])
        freqs[freq].append(loc)

    L.info("freq counts: %r", sorted([len(locs), freq] for freq, locs in freqs.items()))

    antinodes = []
    for freq, locs in freqs.items():
        L.debug("freq %s has antennas @ %r", freq, locs)
        antinodes.extend((l, freq) for l in find_antinodes(locs))

    L.info("antinodes locations (may be outside area, not unique): %d", len(antinodes))

    L.debug("antinodes: %r", antinodes)

    p1_ans = 0
    antinode_locs = {}
    for loc, freq in antinodes:
        if loc[0] not in range(field_dims[0]) or loc[1] not in range(field_dims[1]):
            L.debug("antinode for freq %s outside field @ %r", freq, loc)
            continue
        # if loc in locations:
        #     L.debug("antinode for freq %s occupied by antenna @ %r", freq, loc)
        #     continue
        if loc in antinode_locs:
            L.debug("antinode for freq %s repeated @ %r", freq, loc)
            antinode_locs[loc].append(freq)
            continue

        p1_ans += 1
        L.debug("new antinode for freq %s @ %r", freq, loc)
        antinode_locs[loc] = [freq]

    yield p1_ans
    yield "part 2 answer"
