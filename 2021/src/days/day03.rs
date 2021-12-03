fn parse_input(input: &String) -> Vec<&str> {
   input.trim().lines().collect()
}

pub fn part1(input: String) {
    let report = parse_input(&input);
    let mut nth_bit_num = [0; 12];
    let length = &report.len();

    for l in report {
        for i in 0..12 {
            if l.chars().nth(i).unwrap() == '1' {
                nth_bit_num[i] += 1
            };
        }
    }

    let bits: String = nth_bit_num.map(|i| if i > length / 2 {
        '1'
    } else {
        '0'
    }).iter().collect();

    let gamma = u32::from_str_radix(&bits, 2).unwrap();
    let epsilon = !gamma & 0b0000111111111111;

    println!("part1: {}", gamma * epsilon)
}

fn get_rating(report: &Vec<&str>, most_common: bool) -> u32 {
    let mut left = report.clone();
    for i in 0..12 {
        let mut bit_num = 0;
        let mut input_length = 0;

        for data in &left {
            input_length += 1;
            if data.chars().nth(i).unwrap() == '1' {
                bit_num += 1
            };
        }

        let bit = if bit_num >= (input_length - bit_num) {
            '1'
        } else {
            '0'
        };

        left = left.iter().cloned().filter(|data| {
            if most_common && bit == '1'{
                '1' == data.chars().nth(i).unwrap()
            } else if most_common && bit == '0' {
                '0' == data.chars().nth(i).unwrap()
            } else if !most_common && bit == '1' {
                '0' == data.chars().nth(i).unwrap()
            } else {
                '1' == data.chars().nth(i).unwrap()
            }
        }).collect();

        if left.iter().count() == 1 {
            break
        }
    }

    u32::from_str_radix(left.first().unwrap(), 2).unwrap()
}

pub fn part2(input: String) {
    let report = parse_input(&input);
    let oxygen = get_rating(&report, true);
    let co2 = get_rating(&report, false);

    println!("part2: {}", oxygen * co2);
}
