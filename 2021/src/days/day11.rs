fn parse_input(input: &String) -> Vec<Vec<u32>> {
    input
        .lines()
        .map(|l| l.chars().map(|n| n.to_digit(10).unwrap()).collect())
        .collect()
}

fn count_flashes(report: &mut Vec<Vec<u32>>, size: usize) -> u32 {
    let mut flashes = 0;

    for y in 0..size {
        for x in 0..size {
            if report[x][y] >= 10 {
                report[x][y] = 0;
                flashes += 1;

                if x > 0 && report[x - 1][y] != 0 {
                    report[x - 1][y] += 1;
                }

                if x < size - 1 && report[x + 1][y] != 0 {
                    report[x + 1][y] += 1;
                }

                if y > 0 && report[x][y - 1] != 0 {
                    report[x][y - 1] += 1;
                }

                if y < size - 1 && report[x][y + 1] != 0 {
                    report[x][y + 1] += 1;
                }

                if x > 0 && y > 0 && report[x - 1][y - 1] != 0 {
                    report[x - 1][y - 1] += 1;
                }

                if x > 0 && y < size - 1 && report[x - 1][y + 1] != 0 {
                    report[x - 1][y + 1] += 1;
                }

                if x < size - 1 && y > 0 && report[x + 1][y - 1] != 0 {
                    report[x + 1][y - 1] += 1;
                }

                if x < size - 1 && y < size - 1 && report[x + 1][y + 1] != 0 {
                    report[x + 1][y + 1] += 1;
                }

                flashes += count_flashes(report, size);
            }
        }
    }

    flashes
}

pub fn part1(input: String) {
    let mut report = parse_input(&input);
    let mut flashes = 0;
    let size = 10;

    for _steps in 1..=100 {
        for y in 0..size {
            for x in 0..size {
                report[x][y] += 1;
            }
        }

        flashes += count_flashes(&mut report, size);
    }

    println!("part1: {}", flashes);
}

pub fn part2(input: String) {
    let mut report = parse_input(&input);
    let size = 10;
    let total = (size * size) as u32;

    for steps in 1..=250 {
        for y in 0..size {
            for x in 0..size {
                report[x][y] += 1;
            }
        }

        if count_flashes(&mut report, size) == total {
            println!("part2: {}", steps);
        }
    }

    println!("didn't find a solution");
}
