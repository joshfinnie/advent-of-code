fn parse_input(input: &String) -> Vec<(&str, u32)>{
    let output = input.trim().split("\n").map(|s| {
        let data: Vec<&str> = s.split(" ").collect();
        (data[0], data[1].parse::<u32>().unwrap())
    }).collect();

    output
}

pub fn part1(input: String) {
    let report = parse_input(&input);
    let mut depth = 0;
    let mut forward = 0;
    for (dir, val) in report {
        match dir {
            "forward" => forward += val,
            "up" => depth -= val,
            "down" => depth += val,
            _ => println!("invalid direction!")
        }
    }

    println!("part1: {}", depth * forward);
}

struct Sub {
    depth: u32,
    forward: u32,
    aim: u32,
}

pub fn part2(input: String) {
    let report = parse_input(&input);
    let mut sub = Sub {
        depth: 0,
        forward: 0,
        aim: 0
    };
    for (dir, val) in report {
        match dir {
            "forward" => {
                sub.forward += val;
                sub.depth += sub.aim * val;
            },
            "up" => sub.aim -= val,
            "down" => sub.aim += val,
            _ => println!("invalid direction!")
        }
    }

    println!("part2: {}", sub.depth * sub.forward);
}
