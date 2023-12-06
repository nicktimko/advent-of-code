use core::panic;
use std::collections::HashMap;
use std::cmp;
use std::fs::read_to_string;

use phf::phf_map;

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
    day3();
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

fn day3() {

    let raw_schematic = read_lines("inputs/day03.txt");

    let (rows, cols) = (raw_schematic.len(), raw_schematic[0].len());

    let mut schematic = vec![vec![0u8; cols]; rows];
    let mut legend = vec![vec![D3SchLegend::Empty; cols]; rows];

    for (i, row) in raw_schematic.iter().enumerate() {
        for (j, val) in row.as_bytes().iter().enumerate() {
            if *val == b'\n' && j != cols {
                // the newline should be at the end of the columns
                panic!();
            }
            schematic[i][j] = *val;
            legend[i][j] = SCH_CHARS[val];
        }
    }

    for (i, row) in schematic.iter().enumerate() {
        for (j, _val) in row.iter().enumerate() {
            if legend[i][j] == D3SchLegend::Symbol {
                for (ii, jj) in surrounding_coords(i, j, rows, cols) {
                    if legend[ii][jj] == D3SchLegend::RogueNumber {
                        legend[ii][jj] = D3SchLegend::PartNumber;

                        // flood fill along the line
                        for jback in (0..jj).rev() {
                            match legend[ii][jback] {
                                D3SchLegend::Empty => break,
                                D3SchLegend::Symbol => break,
                                _ => legend[ii][jback] = D3SchLegend::PartNumber
                            }
                        }
                        for jfwd in jj..cols {
                            match legend[ii][jfwd] {
                                D3SchLegend::Empty => break,
                                D3SchLegend::Symbol => break,
                                _ => legend[ii][jfwd] = D3SchLegend::PartNumber
                            }
                        }
                    }
                }
            }
        }
    }

    // debug view
    let sch_legend = HashMap::from([
        (D3SchLegend::Empty, '.'),
        (D3SchLegend::RogueNumber, '?'),
        (D3SchLegend::PartNumber, 'P'),
        (D3SchLegend::Symbol, 's'),
    ]);
    if false {
        for row in legend.iter() {
            for item in row.iter() {
                print!("{}", sch_legend[item]);
            }
            println!("");
        }
    }

    let mut accum_part1 = 0u32;
    for (leg_row, sch_row) in legend.iter().zip(schematic.iter()) {
        let mut scanning_number = false;
        let mut number = 0u32;

        println!("{}", std::str::from_utf8(sch_row).unwrap());
        for item in leg_row.iter() {
            print!("{}", sch_legend[item]);
        }
        println!("");

        for (n_col, (col_leg, col_char)) in leg_row.iter()
                .zip(sch_row.iter())
                // .chain([(D3SchLegend::Empty, b'.')].iter())
                // trying to run one more at the end ^^, didn't work
                // vv checking last col
                .enumerate()
                {
            let digit_value: u32 = (col_char.wrapping_sub(b'0')) as u32;
            // match (scanning_number, *col_leg == D3SchLegend::PartNumber) {
            match (scanning_number, *col_leg == D3SchLegend::PartNumber, n_col == cols-1) {
                (false, false, _) => continue,
                // start scanning number
                (false, true, _) => {
                    scanning_number = true;
                    number = digit_value;
                },
                // end of num
                (true, false, _) => {
                    scanning_number = false;
                    accum_part1 += number;
                    // println!("end of number: {}, running total {}", number, accum_part1);
                }
                // end of line WHILE scanning
                (true, true, true) => {
                    scanning_number = false;
                    accum_part1 += 10*number + digit_value;
                    // println!("end of number: {}, running total {}", number, accum_part1);
                }
                // middle of scanning number
                (true, true, _) => {
                    number = 10*number + digit_value;
                },
            }
        }
    }

    // println!("{:?}", schematic);

    println!("Day 3, Part 1: {}", accum_part1);
    println!("Day 3, Part 2: {}", 0);

}

fn surrounding_coords(x: usize, y: usize, xmax: usize, ymax: usize) -> Vec<(usize, usize)> {
    // in Python I'd write this as an iterator and yield back valid values, but
    // not sure how to do that in rustland, so will allocate vec and return that?

    // enumerate() gives `usize` and that's also what indexing wants, so keep with it
    let mut coords: Vec<(usize, usize)> = Vec::with_capacity(8);
    for dx in [-1, 0, 1].iter() {
        if (x == 0 && *dx == -1)
        || (x == xmax - 1 && *dx == 1) {
            continue;
        }
        for dy in [-1, 0, 1].iter() {
            if (y == 0 && *dy == -1)
            || (y == ymax - 1 && *dy == 1)
            || (*dx == 0 && *dy == 0) {
                continue;
            }
            coords.push((
                x.wrapping_add_signed(*dx),
                y.wrapping_add_signed(*dy)
            ))
        }
    }
    coords
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
enum D3SchLegend {
    Empty,
    Symbol,
    RogueNumber,
    PartNumber,
}

static SCH_CHARS: phf::Map<u8, D3SchLegend> = phf_map!{
    b'.' => D3SchLegend::Empty,
    b'0' => D3SchLegend::RogueNumber,
    b'1' => D3SchLegend::RogueNumber,
    b'2' => D3SchLegend::RogueNumber,
    b'3' => D3SchLegend::RogueNumber,
    b'4' => D3SchLegend::RogueNumber,
    b'5' => D3SchLegend::RogueNumber,
    b'6' => D3SchLegend::RogueNumber,
    b'7' => D3SchLegend::RogueNumber,
    b'8' => D3SchLegend::RogueNumber,
    b'9' => D3SchLegend::RogueNumber,
    b'#' => D3SchLegend::Symbol,
    b'$' => D3SchLegend::Symbol,
    b'%' => D3SchLegend::Symbol,
    b'&' => D3SchLegend::Symbol,
    b'*' => D3SchLegend::Symbol,
    b'+' => D3SchLegend::Symbol,
    b'-' => D3SchLegend::Symbol,
    b'/' => D3SchLegend::Symbol,
    b'=' => D3SchLegend::Symbol,
    b'@' => D3SchLegend::Symbol,
};

// the "naive" version as per Rust By Example, but simpler to reason about for now

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}