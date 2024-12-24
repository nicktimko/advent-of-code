import argparse
import importlib
import logging
import shutil
import sys
import time

# import timeit

from .io import REPO_ROOT


DAY_MODULE_PATH = REPO_ROOT / "aoc24" / "day"
LOG_LEVELS = {
    None: logging.WARNING,
    0: logging.WARNING,
    1: logging.INFO,
    2: logging.DEBUG,
    3: 0, # 'TRACE'
}


def template(day: int) -> None:
    template = DAY_MODULE_PATH / f"d00.py"
    target = DAY_MODULE_PATH / f"d{day:02d}.py"

    if target.exists():
        print(f"Module already exists for day {day}. Aborting init.", file=sys.stderr)
        return 1

    # shutil.copy(template, target)
    with template.open(mode="r") as f_in, target.open(mode="w") as f_out:
        f_out.write(
            f_in.read().replace(r"'''DAY'''", str(day)).replace(r'"""DAY"""', str(day))
        )


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("day", type=int)
    parser.add_argument("--init", action="store_true")
    parser.add_argument("-t", "--time", action="store_true")
    parser.add_argument("-v", "--verbose", action="count")
    # parser.add_argument("-T", "--timeit", action="store_true")
    args = parser.parse_args()

    logging.basicConfig(level=LOG_LEVELS[args.verbose])# format='%(relativeCreated)6d %(threadName)s %(message)s')

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

    # if args.timeit:
    #     timeit.Timer("day_module.run()")
    #     pass
    # else:
    start = time.monotonic_ns()
    results = list(day_module.run())
    elapsed = time.monotonic_ns() - start
    if args.time:
        print("Elapsed time: {:0.2f} ms".format(elapsed / 1e6))

    for n, r in enumerate(results, start=1):
        print(f"Day {args.day}, Part {n}: {r}")


if __name__ == "__main__":
    sys.exit(main())
