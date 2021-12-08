fn parse_input(input: &String) -> Vec<u8> {
    input
        .trim()
        .split(',')
        .map(|v| v.parse::<u8>().unwrap())
        .collect()
}
pub fn part1(input: String) {
    let mut report = parse_input(&input);
    let days = 80;

    for _ in 1..=days {
        for i in 0..report.len() {
            if report[i] >= 1 {
                report[i] -= 1
            } else {
                report[i] = 6;
                report.push(8);
            }
        }
    }

    println!("part1: {}", report.len() as u32);
}

pub fn part2(input: String) {
    let report = parse_input(&input);
    let days = 256;
    let mut fish = [0 as u64; 9];

    report.iter().for_each(|f| fish[*f as usize] += 1);

    for _ in 1..=days {
        let mut new_fish = [0; 9];
        let b = fish[0];
        new_fish[0] = fish[1];
        new_fish[1] = fish[2];
        new_fish[2] = fish[3];
        new_fish[3] = fish[4];
        new_fish[4] = fish[5];
        new_fish[5] = fish[6];
        new_fish[6] = fish[7] + b;
        new_fish[7] = fish[8];
        new_fish[8] = b;

        fish = new_fish;
    }

    println!("part2: {}", fish.iter().sum::<u64>());
}
