mod days;
use days::{day01, day02, day03, day04};

pub fn noop(_: String) {}

type DayFn = fn(String);

pub fn get_day(day: u32) -> (DayFn, DayFn) {
    return match day {
        1 => (day01::part1, day01::part2),
        2 => (day02::part1, day02::part2),
        3 => (day03::part1, day03::part2),
        4 => (day04::part1, day04::part2),
        5 => (noop, noop),
        6 => (noop, noop),
        7 => (noop, noop),
        8 => (noop, noop),
        9 => (noop, noop),
        10 => (noop, noop),
        11 => (noop, noop),
        12 => (noop, noop),
        13 => (noop, noop),
        14 => (noop, noop),
        15 => (noop, noop),
        16 => (noop, noop),
        17 => (noop, noop),
        18 => (noop, noop),
        19 => (noop, noop),
        20 => (noop, noop),
        21 => (noop, noop),
        22 => (noop, noop),
        23 => (noop, noop),
        24 => (noop, noop),
        25 => (noop, noop),
        _ => {
            println!("Not a day that's part of AOC!");
            (noop, noop)
        }
    }
}