import pathlib

PACKAGE_ROOT = pathlib.Path(__file__).parents[1]
INPUTS = PACKAGE_ROOT / "inputs"


def file_for_day(day_n):
    return INPUTS / "day{day_n:02d}.txt"
