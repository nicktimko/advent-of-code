"""
Day 4:
https://adventofcode.com/2024/day/4
"""

import logging
# import numpy as np
# from scipy import ndimage

from ..io import read_lines


L = logging.getLogger(__name__)

TODAY = 4
DIRS = ["N", "S", "E", "W", "NE", "NW", "SE", "SW"]


def check_pos_dir(matrix, x: int, y: int, direction: str, word="XMAS") -> bool:
    dx = dy = 0
    if "N" in direction:
        dy = -1
    elif "S" in direction:
        dy = 1
    if "W" in direction:
        dx = -1
    elif "E" in direction:
        dx = 1

    focus = matrix[y][x]
    if focus != word[0]:
        return False

    L.debug(f"checking dir: {direction} ({dx}, {dy})")
    for letter in word[1:]:
        x += dx
        y += dy
        if x < 0 or y < 0:
            # guard against indexes too small
            L.debug(f"({x}, {y}) exceeded bounds")
            return False

        try:
            focus = matrix[y][x]
        except IndexError:
            # indexes too big
            L.debug(f"({x}, {y}) exceeded bounds")
            return False

        if focus != letter:
            L.debug(f"({x}, {y}) is {focus} (needed {letter})")
            return False
        L.debug(f"({x}, {y}) is {focus}")

    return True


def check_pos(matrix, x: int, y: int) -> int:
    if matrix[y][x] != "X":
        return 0
    return sum(1 for d in DIRS if check_pos_dir(matrix, x, y, d))


def check_ex_mas(matrix, x, y):
    if matrix[y][x] != "A":
        return 0
    if not (0 < y < len(matrix) - 1):
        return 0
    if not (0 < x < len(matrix[0]) - 1):
        return 0

    nwse = {matrix[y-1][x-1], matrix[y+1][x+1]}
    nesw = {matrix[y-1][x+1], matrix[y+1][x-1]}

    if nwse == nesw == {"M", "S"}:
        return 1
    return 0


def test():
    inp = """\
    MMMSXXMASM
    MSAMXMSMSA
    AMXSXMAAMM
    MSAMASMSMX
    XMASAMXAMM
    XXAMMXXAMA
    SMSMSASXSS
    SAXAMASAAA
    MAMMMXMMMM
    MXMXAXMASX"""
    matrix = [line for line in (line.strip() for line in inp.splitlines()) if line]
    for d in DIRS:
        if check_pos_dir(matrix, 4, 0, d):
            L.info(d)
    L.info(check_pos(matrix, 9, 4))
    L.info(check_pos(matrix, 4, 0))

    matrix2 = [
        "S..S..S",
        ".A.A.A.",
        "..MMM..",
        "SAMXMAS",
        "..MMM..",
        ".A.A.A.",
        "S..S..S",
    ]
    L.info(check_pos(matrix2, 3, 3))

    L.info(check_ex_mas(matrix, 1, 2))


def run():

    # test()
    # return
    # return
    # data = read_lines(TODAY)

    # nums = [1, 10, 100, 1000]
    # target = sum(n**2 for n in nums)

    # lut = dict(zip("XMAS", nums))
    # g = np.array([[lut[c] for c in line] for line in data])

    # dir_e = np.array([[nums[0],],[nums[1],],[nums[2],],[nums[3],]])
    # dir_w = np.array([[nums[3],],[nums[2],],[nums[1],],[nums[0],]])
    # dir_s = np.array([nums])
    # dir_n = np.array([nums[::-1]])

    # dir_se = np.array([[nums[0], 0, 0, 0],[0, nums[1], 0, 0],[0, 0, nums[2], 0],[0, 0, 0, nums[3]]])
    # dir_ne = np.array([[0, 0, 0, nums[0]],[0, 0, nums[1], 0],[0, nums[2], 0, 0],[nums[3], 0, 0, 0]])
    # dir_nw = np.array([[nums[3], 0, 0, 0],[0, nums[2], 0, 0],[0, 0, nums[1], 0],[0, 0, 0, nums[0]]])
    # dir_sw = np.array([[0, 0, 0, nums[3]],[0, 0, nums[2], 0],[0, nums[1], 0, 0],[nums[0], 0, 0, 0]])

    # kernels = [dir_e, dir_w, dir_n, dir_s, dir_se, dir_sw, dir_nw, dir_ne]

    # yield sum(np.count_nonzero(ndimage.convolve(g, k, mode="constant", cval=0) == target) for k in kernels)

    data = [list(line) for line in read_lines(TODAY)]
    total = 0
    pt2 = 0

    permutations = set()
    for iy in range(len(data)):
        for ix in range(len(data[0])):
            more = check_pos(data, ix, iy)
            if more not in permutations:
                permutations.add(more)
                L.info(f"pos ({ix}, {iy}) has {more} xmases")
            total += more
            pt2 += check_ex_mas(data, ix, iy)
    yield total


    yield pt2
