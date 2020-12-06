// use rand::Rng;
// use std::cmp::Ordering;
// use std::io;
use std::collections::HashSet;
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    let day = &args[1];
    let day: u32 = day.parse().expect("non-integer day");

    let path = format!("inputs/day{:02}.txt", day);

    let contents = fs::read_to_string(path).expect("Failed reading file.");

    let lines = contents.lines();
    let mut costs: HashSet<i32> = HashSet::new();

    for cost in lines {
        let cost = cost.trim().parse().expect("not a number");
        costs.insert(cost);
    }

    let target_total = 2020;

    for cost in costs.iter() {
        let remainder = target_total - cost;
        if costs.contains(&remainder) {
            println!("Part 1: {}", cost * remainder);
            break;
        }
    }

    'loops: for cost1 in costs.iter() {
        for cost2 in costs.iter() {
            let remainder = target_total - cost1 - cost2;
            if costs.contains(&remainder) {
                println!("Part 2: {}", cost1 * cost2 * remainder);
                break 'loops;
            }
        }
    }
}
