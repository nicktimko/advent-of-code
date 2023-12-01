# 2023 Advent of Code with Rust

I worked with Rust before, but didn't do much. Maybe more this year? I'm coming back with not much real knowledge beyond "Rust has variable ownership", and mostly from Python and Golang experience, so I'll write from that perspective.

## Log

### Day 1

Looking through [Rust by Example](https://doc.rust-lang.org/rust-by-example/index.html) it's [file I/O section](https://doc.rust-lang.org/rust-by-example/std_misc/file.html) is buried behind "std misc"? Strange. Was going to work on creating basic file I/O readers, but will maybe do that later.

Strat to iterate over lines, and within the lines iterate from either side. Had vals `first|last_num` to store what I found, but the compiler was mad that they were maybe uninitalized because it didn't know that the loops would (should...given the input) set them. Just added 10x the first num found, and 1x the second.

Could do the second part by replacing the words and reusing the same logic, or a regex for the words/digits and get the first/last match. Word replace sounds best?