pub fn part1(input: String) {
    let report: Vec<u16> = input
        .trim()
        .split("\n")
        .map(|s| s.parse::<u16>().unwrap())
        .collect();
    let mut prev = report.first().unwrap();
    let mut result = 0;
    for depth in report.iter().skip(1) {
        if depth > prev {
            result += 1;
        }
        prev = depth;
    }
    println!("part1: {}", result);
}

pub fn part2(input: String) {
    let report: Vec<u16> = input
        .trim()
        .split("\n")
        .map(|s| s.parse::<u16>().unwrap())
        .collect();
    let mut prev: u16 = report.iter().take(3).sum();
    let mut result = 0;
    for (i, _) in report.iter().skip(1).enumerate() {
        if report.len() - i < 3 {
            break;
        }

        let a = report.iter().nth(i).unwrap();
        let b = report.iter().nth(i + 1).unwrap();
        let c = report.iter().nth(i + 2).unwrap();

        let depth = a + b + c;

        if depth > prev {
            result += 1;
        }

        prev = depth
    }
    println!("part2: {}", result);
}
