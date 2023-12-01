use std::fs::read_to_string;

// use regex::RegexBuilder;

struct WordNums {
    word: String,
    num: u8,
}

fn main() {
    let mut accum = 0u32;
    let mut accum2 = 0u32;

    for line in read_lines("inputs/day01.txt") {
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
        // println!("{} -> {}{}", line, first_num, last_num);
    }

    let wns: [WordNums; 9] = [
        WordNums{word: "one".to_string(), num: 1},
        WordNums{word: "two".to_string(), num: 2},
        WordNums{word: "three".to_string(), num: 3},
        WordNums{word: "four".to_string(), num: 4},
        WordNums{word: "five".to_string(), num: 5},
        WordNums{word: "six".to_string(), num: 6},
        WordNums{word: "seven".to_string(), num: 7},
        WordNums{word: "eight".to_string(), num: 8},
        WordNums{word: "nine".to_string(), num: 9},
    ];

    for line in read_lines("inputs/day01.txt") {
        for wn in wns {
        }
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
        // println!("{} -> {}{}", line, first_num, last_num);
    }

    // println!("Part 1: {}", accum);
    // let pat = r"
    //     (
    //         [0-9]  # single digit
    //         |one   # ...or 1-9 in english
    //         |two
    //         |three
    //         |four
    //         |five
    //         |six
    //         |seven
    //         |eight
    //         |nine
    //     )
    // ";
    // let re = RegexBuilder::new(pat)
    //     .ignore_whitespace(true)
    //     .build()
    //     .unwrap();

}


// the "naive" version as per Rust By Example, but simpler to reason about for now

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}