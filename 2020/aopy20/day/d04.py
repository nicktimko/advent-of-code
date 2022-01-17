"""
Day 4: Passport Processing
https://adventofcode.com/2020/day/4
"""
import string
from ..io import read_paragraphs

TODAY = 4

REQUIRED_FIELDS = {
    "byr",  # Birth Year
    "iyr",  # Issue Year
    "eyr",  # Expiration Year
    "hgt",  # Height
    "hcl",  # Hair Color
    "ecl",  # Eye Color
    "pid",  # Passport ID
}
OPTIONAL_FIELDS = {
    "cid",  # Country ID
}

VALIDATORS = {}


def validator(field):
    def _validator(f):
        global VALIDATORS
        VALIDATORS[field] = f

    return _validator


@validator("byr")
def _(value):
    try:
        return 1920 <= int(value) <= 2002
    except ValueError:
        return False


@validator("iyr")
def _(value):
    try:
        return 2010 <= int(value) <= 2020
    except ValueError:
        return False


@validator("eyr")
def _(value):
    try:
        return 2020 <= int(value) <= 2030
    except ValueError:
        return False


@validator("hgt")
def _(value):
    try:
        if value.endswith("in"):
            return 59 <= int(value[:-2]) <= 76
        elif value.endswith("cm"):
            return 150 <= int(value[:-2]) <= 193
        else:
            return False
    except ValueError:
        # e.g "xin" "32mcm"
        return False


@validator("hcl")
def _(value):
    try:
        return value[0] == "#" and all(
            value[n] in string.hexdigits for n in range(1, 7)
        )
    except IndexError:
        # e.g. "", "#abc"
        return False


@validator("ecl")
def _(value):
    return value in {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}


@validator("pid")
def _(value):
    return len(value) == 9 and value.isnumeric()


def run():
    well_formed = 0
    valid = 0

    for para in read_paragraphs(TODAY):
        data = dict(kv.split(":") for kv in para.split())

        if not (REQUIRED_FIELDS - set(data)):
            well_formed += 1

            # if all(VALIDATORS[field](value) for field, value in data.items()):
            if all(
                VALIDATORS.get(field, lambda x: True)(value)
                for field, value in data.items()
            ):
                valid += 1

    print("Part 1:", well_formed)
    print("Part 2:", valid)
