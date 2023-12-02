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

        println!("{} -> {}", line, alt_line);

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

    println!("Part 1: {}", accum);
    println!("Part 2: {}", accum2);
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