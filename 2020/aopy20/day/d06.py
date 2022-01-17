"""
Day 6: Custom Customs
https://adventofcode.com/2020/day/6
"""

from functools import reduce

from ..io import read_paragraphs


TODAY = 6


def run():
    anyone_total = 0
    everyone_total = 0
    for para in read_paragraphs(TODAY, coalesce=False):
        dec_forms = [set(f) for f in para.splitlines()]

        anyone_claims = reduce(set.union, dec_forms)
        everyone_claims = reduce(set.intersection, dec_forms)

        anyone_total += len(anyone_claims)
        everyone_total += len(everyone_claims)

    print("Part 1:", anyone_total)
    print("Part 2:", everyone_total)
