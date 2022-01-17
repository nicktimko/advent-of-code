"""
"""

import argparse
import importlib
import shutil
import sys
import time
import timeit

from .io import REPO_ROOT


DAY_MODULE_PATH = REPO_ROOT / "aopy20" / "day"


def template(day: int) -> None:
    template = DAY_MODULE_PATH / f"d00.py"
    target = DAY_MODULE_PATH / f"d{day:02d}.py"

    if target.exists():
        print(f"Module already exists for day {day}. Aborting init.", file=sys.stderr)
        return 1

    shutil.copy(template, target)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("day", type=int)
    parser.add_argument("--init", action="store_true")
    parser.add_argument("-t", "--time", action="store_true")
    parser.add_argument("-T", "--timeit", action="store_true")
    args = parser.parse_args()

    if args.init:
        return template(args.day)

    try:
        day_module = importlib.import_module(
            f".day.d{args.day:02d}", package=__package__
        )
    except ImportError:
        print(
            f"No module defined for day {args.day}. Initialize with --init",
            file=sys.stderr,
        )
        return 1

    if args.timeit:
        # timeit.Timer("day_module.run()")
        pass
    else:
        start = time.monotonic_ns()
        day_module.run()
        elapsed = time.monotonic_ns() - start
        if args.time:
            print("Elapsed time: {:0.2f} ms".format(elapsed / 1e6))


if __name__ == "__main__":
    sys.exit(main())
