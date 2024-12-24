"""
Day '''DAY''':
https://adventofcode.com/2024/day/'''DAY'''
"""

import logging

from ..io import read_lines


L = logging.getLogger(__name__)
TODAY = """DAY"""


def run():
    L.info("starting day %d", TODAY)
    data = read_lines(TODAY)
    yield "part 1 answer"
    yield "part 2 answer"
