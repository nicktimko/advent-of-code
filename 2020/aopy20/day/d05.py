"""
Day 5: Binary Boarding
https://adventofcode.com/2020/day/5
"""

from ..io import read_lines


TODAY = 5


def decode_seat(code):
    code = code.translate({ord(x): y for x, y in zip("FBLR", "0101")})
    return int(code, base=2)


assert decode_seat("BFFFBBFRRR") == 567


def run():
    max_seat = 0
    seats = []
    for line in read_lines(TODAY):
        seat_id = decode_seat(line)
        # row, col = divmod(seat_id, 8)
        max_seat = max(seat_id, max_seat)
        seats.append(seat_id)

    print("Part 1:", max_seat)

    seats.sort()
    iseats = iter(seats)
    last_seat = next(iseats)
    for seat_id in iseats:
        if seat_id - last_seat != 1:
            missing_seat = seat_id - 1
            break
        last_seat = seat_id
    else:
        raise AssertionError

    print("Part 2:", missing_seat)
