use std::fs::read_to_string;

fn main() {
    let mut accum = 0u32;

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

    println!("{}", accum);

}


// the "naive" version as per Rust By Example, but simpler to reason about for now

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}