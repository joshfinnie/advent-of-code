use std::env;
use std::fs;
use std::io;

use aoc_2021::{get_day, noop};

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut day = String::new();
    let mut input = String::new();

    match args.len() {
        2 => day = args[1].clone(),
        _ => {
            println!("Enter day: ");
            io::stdin()
                .read_line(&mut day)
                .expect("Failed to read line.");
        }
    }

    day = day.trim().to_string();
    let day_num: u32 = match day.parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Invalid input for day: {}", day);
            return;
        }
    };

    let to_run = get_day(day_num);

    if to_run.0 != noop || to_run.1 != noop {
        let cwd = env::current_dir().unwrap();
        let filename = cwd.join("inputs").join(format!("day{:02}.txt", day_num));
        input = fs::read_to_string(filename).expect("Error while reading");
    } else {
        println!("There's no input for this day!")
    }

    if to_run.0 != noop {
        to_run.0(input.clone());
    }

    if to_run.1 != noop {
        to_run.1(input.clone());
    }
}
