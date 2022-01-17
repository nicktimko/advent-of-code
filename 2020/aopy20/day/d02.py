"""
Day 2: Password Philosophy
https://adventofcode.com/2020/day/2
"""

import re
from ..io import read_lines

PASSWORD_POLICY_PATTERN = re.compile(
    r"""
    ^
    (?P<count_min>\d+)-(?P<count_max>\d+)
    [\ ]
    (?P<letter>\w):
    [\ ]
    (?P<password>\w+)
    $
""",
    flags=re.VERBOSE,
)


def password_policy(entry: str):
    match = PASSWORD_POLICY_PATTERN.fullmatch(entry.strip())
    policy = match.groupdict()

    policy["count_min"] = int(policy["count_min"])
    policy["count_max"] = int(policy["count_max"])
    return policy


def valid_password(policy, password: str) -> bool:
    count = password.count(policy["letter"])
    return policy["count_min"] <= count <= policy["count_max"]


def officially_valid_password(policy, password: str) -> bool:
    pos1 = password[policy["count_min"] - 1] == policy["letter"]
    pos2 = password[policy["count_max"] - 1] == policy["letter"]
    return pos1 ^ pos2  # XOR


def run():
    entries = [password_policy(line) for line in read_lines(2)]
    valid_passwords = sum(1 for p in entries if valid_password(p, p["password"]))
    print("Part 1:", valid_passwords)

    official_passwords = sum(
        1 for p in entries if officially_valid_password(p, p["password"])
    )
    print("Part 2:", official_passwords)
