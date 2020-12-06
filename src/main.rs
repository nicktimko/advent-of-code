// use rand::Rng;
// use std::cmp::Ordering;
// use std::io;
use std::env;
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();

    let day = &args[1];
    let day: u32 = day.parse().expect("non-integer day");

    let path = format!("inputs/day{:02}.txt", day);

    let contents = fs::read_to_string(path).expect(
        "Failed reading file."
    );

    println!("text: {}", contents);

}
