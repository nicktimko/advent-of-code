"""
Day 5:
https://adventofcode.com/2024/day/5
"""

import logging

import networkx as nx

from ..io import read_lines


L = logging.getLogger(__name__)
TODAY = 5


def update_ok_part1(orderings, update):
    for i1 in range(len(update)):
        for i2 in range(i1 + 1, len(update)):
            p1 = update[i1]
            p2 = update[i2]
            if (p2, p1) in orderings:
                L.debug("update %r violates ordering %r", update, (p2, p1))
                return False
    return True


def run():
    L.info("starting day %d", TODAY)
    data = read_lines(TODAY, skip_blanks=False)

    orderings = set()
    known_pages = set()

    for line in data:
        line = line.strip()
        if not line:
            break
        new_order = tuple(int(x) for x in line.split("|"))
        assert new_order not in orderings
        assert len(new_order) == 2
        orderings.add(new_order)
        known_pages.add(new_order[0])
        known_pages.add(new_order[1])

    known_pages = list(known_pages)

    L.info("read %d orderings", len(orderings))
    L.info("pages to order: %d ", len(known_pages))

    # Tried to pre-compute the correct page orders. This probably fails because
    # there's some ambiguity in the DAG? Unsure.
    #
    # ordered_pages = [known_pages[0]]
    # def check_insertion(pos, cur_pages, new_page):
    #     for prepage in cur_pages[:pos]:
    #         counterexample = (new_page, prepage)
    #         if counterexample in orderings:
    #             return False
    #     for postpage in cur_pages[pos:]:
    #         counterexample = (postpage, new_page)
    #         if counterexample in orderings:
    #             return False
    #     return True

    # for page in known_pages[1:]:
    #     for insertion_point in range(len(ordered_pages) + 1):
    #         if check_insertion(insertion_point, ordered_pages, page):
    #             ordered_pages.insert(insertion_point, page)
    #             break
    # L.info("page order: %r ", ordered_pages)
    # for x, y in orderings:
    #     assert ordered_pages.index(x) < ordered_pages.index(y)
    G = nx.DiGraph()
    for ordering in orderings:
        G.add_edge(*ordering)

    descendants = {}
    for page in known_pages:
        descendants[page] = list(nx.descendants(G, page))

    print(sorted(([pg, len(dcs)] for pg, dcs in descendants.items()), key=lambda x: x[1]))

    updates = []
    for line in data:
        new_update = [int(x) for x in line.split(",")]
        assert len(new_update) % 2 == 1  # must be odd
        updates.append(new_update)

    L.info("read %d updates", len(updates))

    p1_ans = 0
    bad_updates = []
    for update in updates:
        if update_ok_part1(orderings, update):
            L.debug("update OK: %r", update)
            p1_ans += update[len(update) // 2]
        else:
            bad_updates.append(update)

    yield p1_ans

    L.info("bad updates to fix: %d", len(bad_updates))


    yield "part 2 answer"
