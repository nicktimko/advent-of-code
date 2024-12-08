use std::fs::read_to_string;

fn main() {
    day01();
    day02();
}

/** given the sorted slice, return how many instances of the target number are
 * present in the slice
 */
fn count_occurrences(sorted: &[i32], target: i32) -> i32 {
    let lower = sorted.partition_point(|&x| x < target);
    let upper = sorted.partition_point(|&x| x < target + 1);
    // println!("{}-{}", lower, upper);
    return (upper - lower).try_into().unwrap();
}

fn day01() {
    let mut l1: Vec<i32> = Vec::new();
    let mut l2: Vec<i32> = Vec::new();

    for line in read_lines("inputs/day01.txt") {
        let nums: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        if nums.len() != 2 {
            println!("unexpected/malformed input");
            return;
        }
        l1.push(nums[0]);
        l2.push(nums[1]);
    }
    l1.sort();
    l2.sort();

    let mut diff: i32 = 0;
    for (n1, n2) in l1.iter().zip(l2.iter()) {
        diff += (n2 - n1).abs();
    }
    println!("Day 1, Part 1: {}", diff);

    let mut sim_score: i32 = 0;
    for n1 in l1.iter() {
        let occ = count_occurrences(&l2, *n1);
        // println!("n: {}, occ: {}", n1, occ);
        sim_score += n1 * occ;
    }
    println!("Day 1, Part 2: {}", sim_score);
}

fn discrete_diff(nn: &[i32]) -> Vec<i32> {
    nn.windows(2).map(|w| w[1] - w[0]).collect()
}

fn day02() {
    let mut safe_reports: i32 = 0;
    for report in read_lines("inputs/day02.txt") {
        let levels: Vec<i32> = report
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        let trend = discrete_diff(&levels);
        let trend_min = trend.iter().min().unwrap();
        let trend_max = trend.iter().max().unwrap();

        match (trend_min, trend_max) {
            (tmin, tmax) if *tmin >= -3 && *tmax <= -1 => {
                safe_reports += 1;
            }
            (tmin, tmax) if *tmin >= 1 && *tmax <= 3 => {
                safe_reports += 1;
            }
            _ => {}
        }
    }
    println!("Day 2, Part 1: {}", safe_reports);
}

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap() // panic on possible file-reading errors
        .lines() // split the string into an iterator of string slices
        .map(String::from) // make each slice into a string
        .collect() // gather them together into a vector
}
