"""
Day 11:
https://adventofcode.com/2020/day/11
"""
import enum
import itertools
from typing import List, Tuple

from ..io import read_lines


TODAY = 11

class State(enum.Enum):
    TAKEN = "#"
    FREE = "L"
    FLOOR = "."


def seqxy(shape, c):
    y, x = divmod(c, shape[0])
    return x, y


def xyseq(shape, xy):
    return shape[0] * xy[0] + xy[1]


def nearby(seats: List[State], shape: Tuple[int, int], coords: Tuple[int, int]) -> int:
    print(coords, shape)
    box = itertools.product([-1, 0, 1], [-1, 0, 1])
    consider = [
        (coords[0] + dx, coords[1] + dy)
        for dx, dy
        in box
        if dx or dy
    ]
    assert len(consider) == 8
    print(consider)
    consider = [
        y * shape[0] + x
        for x, y
        in consider
        if 0 <= x < shape[0]
        and 0 <= y < shape[1]
    ]
    print(consider, shape)
    assert 3 <= len(consider) <= 8, (coords, len(consider), consider)
    return sum(1 for c in consider if seats[c] is State.TAKEN)


def taken(seats: List[State]) -> int:
    return sum(1 for s in seats if s is State.TAKEN)


def advance(seats, shape, coords):
    seq = xyseq(shape, coords)
    seat = seats[seq]
    if seat is State.FLOOR:
        return State.FLOOR
    neighbors = nearby(seats, shape, coords)
    if seat is State.FREE and neighbors == 0:
        return State.TAKEN
    if seat is State.TAKEN and neighbors >= 4:
        return State.FREE
    return seat


def viz(seats, shape):
    for i in range(0, shape[1]):
        print(''.join(s.value for s in seats[shape[0] * i:shape[0] * (i + 1)]))



def test():
    seats = [
        ".L..#L",
        "LLL.#L",
        "LLL.#L",
    ]
    shape = len(seats[0]), len(seats)
    assert shape == (6, 3)

    seats = [State(c) for c in itertools.chain(*seats)]
    for n in range(len(seats)):
        print(nearby(seats, shape, seqxy(shape, n)))
    assert nearby(seats, shape, (0, 4)) == 1


def run():
    test()

    seats = list(read_lines(TODAY))
    shape = len(seats[0]), len(seats)
    seats = [State(c) for c in itertools.chain(*seats)]

    assert taken(seats) == 0
    assert nearby(seats, shape, (0, 0)) == 0

    # next_seats = seats.copy()
    # assert next_seats == seats

    while True:
        next_seats = [advance(seats, shape, seqxy(shape, n)) for n in range(len(seats))]
        viz(seats, shape)
        input()
        if seats == next_seats:
            break
        seats = next_seats

    print("Part 1:", taken(next_seats))
