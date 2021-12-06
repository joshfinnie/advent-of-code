use std::cmp;
use std::collections::HashMap;

fn parse_input(input: &String) -> Vec<&str> {
    input.trim().split('\n').collect()
}

pub fn part1(input: String) {
    let mut vents: HashMap<(u16, u16), u16> = HashMap::new();

    let report = parse_input(&input);
    report.iter().for_each(|l| {
        let points: Vec<Vec<u16>> = l
            .split(" -> ")
            .map(|p| {
                p.split(',')
                    .map(|i| i.parse::<u16>().unwrap())
                    .collect::<Vec<u16>>()
            })
            .collect();
        let x1 = points[0][0];
        let x2 = points[1][0];
        let y1 = points[0][1];
        let y2 = points[1][1];

        if x1 == x2 {
            let ymin = cmp::min(y1, y2);
            let ymax = cmp::max(y1, y2);
            for y in ymin..=ymax {
                let counter = vents.entry((x1, y)).or_insert(0);
                *counter += 1;
            }
        }

        if y1 == y2 {
            let xmin = cmp::min(x1, x2);
            let xmax = cmp::max(x1, x2);
            for x in xmin..=xmax {
                let counter = vents.entry((x, y1)).or_insert(0);
                *counter += 1;
            }
        }
    });

    println!(
        "part1: {}",
        vents.into_values().filter(|&v| v > 1).count() as u16
    )
}

pub fn part2(input: String) {
    let mut vents: HashMap<(i16, i16), i16> = HashMap::new();

    let report = parse_input(&input);
    report.iter().for_each(|l| {
        let points: Vec<Vec<i16>> = l
            .split(" -> ")
            .map(|p| {
                p.split(',')
                    .map(|i| i.parse::<i16>().unwrap())
                    .collect::<Vec<i16>>()
            })
            .collect();
        let x1 = points[0][0];
        let x2 = points[1][0];
        let y1 = points[0][1];
        let y2 = points[1][1];

        if x1 == x2 {
            let ymin = cmp::min(y1, y2);
            let ymax = cmp::max(y1, y2);
            for y in ymin..=ymax {
                let counter = vents.entry((x1, y)).or_insert(0);
                *counter += 1;
            }
        }

        if y1 == y2 {
            let xmin = cmp::min(x1, x2);
            let xmax = cmp::max(x1, x2);
            for x in xmin..=xmax {
                let counter = vents.entry((x, y1)).or_insert(0);
                *counter += 1;
            }
        }

        // Diagonal...
        if i16::abs(x1 - x2) == i16::abs(y1 - y2) {
            let difference = i16::abs(x1 - x2);
            for i in 0..=difference {
                let x = x1 + (i * if x1 > x2 { -1 } else { 1 });
                let y = y1 + (i * if y1 > y2 { -1 } else { 1 });

                let counter = vents.entry((x, y)).or_insert(0);
                *counter += 1;
            }
        }
    });

    println!("part2: {}", vents.into_values().filter(|&v| v > 1).count())
}
