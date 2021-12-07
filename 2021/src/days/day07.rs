fn parse_input(input: &String) -> Vec<i32> {
    input
        .trim()
        .split(',')
        .map(|v| v.parse::<i32>().unwrap())
        .collect()
}

pub fn part1(input: String) {
    let mut report = parse_input(&input);
    report.sort();

    let median = report[report.len() / 2];
    let mut fuel = 0;

    for x in &report {
        fuel += (x - median).abs()
    }

    println!("part1: {}", fuel);
}

pub fn part2(input: String) {
    let mut report = parse_input(&input);
    report.sort();

    let mean = report.iter().sum::<i32>() as f32 / input.len() as f32;
    let mut fuel = i32::MAX;

    for v in [mean.floor(), mean.ceil()] {
        let mut maybe_fuel = 0;
        for x in &report {
            let n = (x - v as i32).abs();
            maybe_fuel += n * (n + 1) / 2;
        }

        if maybe_fuel < fuel {
            fuel = maybe_fuel;
        }
    }

    println!("part2: {}", fuel);
}
