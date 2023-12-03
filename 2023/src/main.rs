use core::panic;
use std::collections::HashMap;
use std::cmp;
use std::fs::read_to_string;

// use regex::RegexBuilder;

struct WordNums<'a> {
    word: &'a str,
    num: u8,
}

const WNS: [WordNums; 9] = [
    WordNums{word: "one", num: 1},
    WordNums{word: "two", num: 2},
    WordNums{word: "three", num: 3},
    WordNums{word: "four", num: 4},
    WordNums{word: "five", num: 5},
    WordNums{word: "six", num: 6},
    WordNums{word: "seven", num: 7},
    WordNums{word: "eight", num: 8},
    WordNums{word: "nine", num: 9},
];

fn main() {
    day1();
    day2();
}

fn day1() {

    let mut accum = 0u32;
    let mut accum2 = 0u32;

    for line in read_lines("inputs/day01.txt") {
        let mut alt_line = line.clone();
        for wn in WNS {
            alt_line = alt_line.replace(wn.word, &*wn.num.to_string());
        }

        // let first_num: u32;
        // let last_num: u32;
        for c in line.chars() {
            let digit = c.to_digit(10);
            match digit {
                None => continue,
                Some(val) => {
                    // first_num = val;
                    accum += 10*val;
                    break
                }
            }
        }

        for c in line.chars().rev() {
            let digit = c.to_digit(10);
            match digit {
                None => continue,
                Some(val) => {
                    // last_num = val;
                    accum += val;
                    break
                }
            }
        }

        // println!("{} -> {}", line, alt_line);

        for c in alt_line.chars() {
            let digit = c.to_digit(10);
            match digit {
                None => continue,
                Some(val) => {
                    // first_num = val;
                    accum2 += 10*val;
                    break
                }
            }
        }
        for c in alt_line.chars().rev() {
            let digit = c.to_digit(10);
            match digit {
                None => continue,
                Some(val) => {
                    // last_num = val;
                    accum2 += val;
                    break
                }
            }
        }
    }

    println!("Day 1, Part 1: {}", accum);
    println!("Day 1, Part 2: {}", accum2);
}

fn day2() {
    let mut id_accum = 0;
    let mut accum_pt2 = 0;
    let (bag_red, bag_green, bag_blue) = (12, 13, 14);
    for line in read_lines("inputs/day02.txt") {
        // println!("{:?}", parse_day2_game(&line).draws);
        let game = parse_day2_game(&line);
        let mut invalid_draw_found = false;
        let (mut min_red, mut min_green, mut min_blue) = (0, 0, 0);
        for draw in game.draws {
            if draw.0 > bag_red || draw.1 > bag_green || draw.2 > bag_blue {
                invalid_draw_found = true;
            }
            min_red = cmp::max(min_red, draw.0);
            min_green = cmp::max(min_green, draw.1);
            min_blue = cmp::max(min_blue, draw.2);
        }
        if !invalid_draw_found {
            id_accum += game.id
        }

        accum_pt2 += min_red * min_green * min_blue;
    }
    println!("Day 2, Part 1: {}", id_accum);
    println!("Day 2, Part 2: {}", accum_pt2);
}

struct Day2Game {
    id: u32,
    draws: Vec<(i32, i32, i32)>
}

fn parse_day2_game(line: &str) -> Day2Game {
    let x: Vec<&str> = line.split(":").collect();
    if x.len() != 2 {
        // Day2Game { id: 0, reveals: vec![] }
        panic!();
    }
    let id: u32 = x[0].split(" ").nth(1).unwrap().parse().unwrap();
    let mut draws: Vec<(i32, i32, i32)> = vec![];
    for step in x[1].split(";") {
        let mut g = HashMap::new();
        for group in step.split(",") {
            // println!("{}", group);
            let parts: Vec<&str> = group.trim().split(" ").collect();
            if parts.len() != 2 {
                continue;
            }
            let count: i32 = parts[0].parse().unwrap();
            let color: &str = parts[1].trim();
            g.insert(color, count);
        }
        draws.push((
            *g.entry("red").or_insert(0),
            *g.entry("green").or_insert(0),
            *g.entry("blue").or_insert(0),
        ))
    }
    Day2Game { id: id, draws: draws }
}

// the "naive" version as per Rust By Example, but simpler to reason about for now

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}