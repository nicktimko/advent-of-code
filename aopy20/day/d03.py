"""
Day 3: Toboggan Trajectory
https://adventofcode.com/2020/day/3
"""

from ..io import read_lines


class TreeFilledSlope:
    def __init__(self, data):
        self.rows = list(data)
        self.repeat_period = len(self.rows[0])

    def has_tree(self, x, y):
        return self.rows[y][x % self.repeat_period] == "#"

    def hits(self, slope_x, slope_y, start_x=0):
        hits = 0
        x = start_x
        y = 0
        while y < len(self.rows):
            hits += 1 if self.has_tree(x, y) else 0
            x += slope_x
            y += slope_y
        return hits

def run():
    slope = TreeFilledSlope(read_lines(3))
    print("Part 1:", slope.hits(3, 1))

    paths = [
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    ]
    product = 1
    for p in paths:
        product *= slope.hits(*p)
    print("Part 2:", product)
