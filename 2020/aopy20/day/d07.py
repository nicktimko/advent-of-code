"""
Day 7:
https://adventofcode.com/2020/day/7
"""

from typing import Dict, Generator, Set, NamedTuple

from ..io import read_lines


TODAY = 7
TARGET_BAG = "shiny gold"


class N_COLORS(NamedTuple):
    n: int
    color: str


def process_inner_bags(inner_rules: str) -> Generator[N_COLORS, None, None]:
    if inner_rules.strip() == "no other bags.":
        return

    for bag in inner_rules.split(","):
        bag = bag.strip().rstrip(".")

        if bag.endswith(" bag"):
            bag = bag[:-4]
        elif bag.endswith(" bags"):
            bag = bag[:-5]
        else:
            raise AssertionError()

        n, color = bag.split(" ", 1)
        n = int(n)

        # bag_can_be_in.setdefault(color, set())
        yield N_COLORS(n, color)


def n_bags_in(contain_tree: Dict[str, N_COLORS], color: str) -> int:
    """Finally, a good use of recursion"""
    return sum(
        x.n * (1 + n_bags_in(contain_tree, x.color)) for x in contain_tree[color]
    )


def run() -> None:
    bag_can_be_in: Dict[str, Set[str]] = {}
    bag_contains: Dict[str, N_COLORS] = {}
    for line in read_lines(TODAY):
        outer_bag, inner_bags = (x.strip() for x in line.split("contain", 1))

        assert outer_bag.endswith(" bags")
        outer_color = outer_bag[:-5]

        bag_contains.setdefault(outer_color, set())

        for inner in process_inner_bags(inner_bags):
            bag_can_be_in.setdefault(inner.color, set())
            bag_can_be_in[inner.color].add(outer_color)

            bag_contains[outer_color].add(inner)

    possible_containers = set(bag_can_be_in[TARGET_BAG])

    added_any = True
    while added_any:
        # print("LOOP")
        added_any = False
        for color in list(possible_containers):
            if color not in bag_can_be_in:
                continue  # top level bag
            new_containers = bag_can_be_in[color]
            # print(color, "can be in", new_containers)
            added_any = added_any or not (new_containers <= possible_containers)
            possible_containers |= new_containers

    # print(possible_containers)
    print("Part 1:", len(possible_containers))

    print("Part 2:", n_bags_in(bag_contains, TARGET_BAG))
