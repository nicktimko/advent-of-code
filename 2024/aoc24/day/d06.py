"""
Day 6:
https://adventofcode.com/2024/day/6
"""

import logging

import numpy as np

from ..io import read_lines


L = logging.getLogger(__name__)
TODAY = 6


def run():
    L.info("starting day %d", TODAY)
    data = read_lines(TODAY)

    obstacles = []
    guard_loc = None
    guard_facing = np.array([0, -1])
    for n_row, row in enumerate(data):
        for n_col, char in enumerate(row):
            loc = np.array([n_col, n_row])
            if char == "^":
                assert guard_loc is None
                guard_loc = loc
            if char == "#":
                obstacles.append(tuple(loc))
    map_size = (n_col + 1, n_row + 1)
    assert guard_loc is not None

    L.info("map size: %r}", map_size)
    L.info("found %d obstacles", len(obstacles))
    L.debug("obstacles: %r", obstacles)
    L.info("starting guard loc: %r", guard_loc)

    visited = {tuple(guard_loc)}

    turn_right = np.array([[0, +1], [-1, 0]])

    while True:
        next_loc = guard_loc + guard_facing

        # if blocked, turn right
        if tuple(next_loc) in obstacles:
            guard_facing = guard_facing @ turn_right
            continue

        guard_loc = next_loc

        # check if we're outside the map
        if not (0 <= guard_loc[0] < map_size[0]):
            break
        if not (0 <= guard_loc[1] < map_size[1]):
            break

        visited.add(tuple(guard_loc))

    yield len(visited)

    ###
    yield "part 2 answer"
