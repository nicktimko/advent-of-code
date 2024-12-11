from typing import Generator
import pathlib

REPO_ROOT = pathlib.Path(__file__).parents[1]

INPUT_DIR = REPO_ROOT / "inputs"


def input_file(day: int) -> pathlib.Path:
    """
    Get the path of the input file for *day*
    """
    return INPUT_DIR / f"day{day:02d}.txt"


def read_lines(day: int, skip_blanks: bool = True) -> Generator[str, None, None]:
    """
    Iteratively return all lines in *day* input file.
    """
    with open(input_file(day), mode="r") as f:
        lines = (line.strip() for line in f)
        for line in lines:
            if line or not skip_blanks:
                yield line


def read_ints(day: int) -> Generator[int, None, None]:
    """
    Iteratively return all non-blank lines in *day* input file as integers.
    """
    for line in read_lines(day):
        if line:
            yield int(line)


def read_int_series(day: int) -> Generator[list[int], None, None]:
    """
    Iteratively return all non-blank lines in *day* input file as series of integers.
    """
    for line in read_lines(day):
        if line:
            yield [int(x) for x in line.split()]


def read_paragraphs(day: int, coalesce: bool = True) -> Generator[str, None, None]:
    """
    Iteratively return lines separated by blank lines. If *coalesce* is true,
    the default, join the newlines with a space, otherwise the newline is kept.
    """
    joiner = " " if coalesce else "\n"
    para = []
    for line in read_lines(day, skip_blanks=False):
        if para and not line:
            yield joiner.join(para)
            para = []
        if line:
            para.append(line)
    if para:
        yield joiner.join(para)
