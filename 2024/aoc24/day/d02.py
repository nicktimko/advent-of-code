from ..io import read_int_series

DAY = 2


def diff(nums):
    nums = iter(nums)
    prev = next(nums)
    for n in nums:
        yield n - prev
        prev = n


def report_safe(report):
    d_report = list(diff(report))
    if max(d_report) <= 3 and min(d_report) >= 1:
        return True
    if min(d_report) >= -3 and max(d_report) <= -1:
        return True
    return False


def run():
    # yield "previously solved"

    n_safe = n_modsafe = 0

    for report in read_int_series(DAY):
        basic_safe = mod_safe = report_safe(report)
        if not mod_safe:
            # for idx, d in enumerate(d_report):
            #     if ((-3 <= d <= -1) or (1 <= d <= 3)):
            #         continue
            for drop_idx in range(len(report)):
                mod_report = report[:drop_idx] + report[drop_idx + 1 :]
                if report_safe(mod_report):
                    mod_safe = True
                    break

        if basic_safe:
            n_safe += 1
        if mod_safe:
            n_modsafe += 1

    yield n_safe
    yield n_modsafe
