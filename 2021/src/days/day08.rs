pub fn part1(input: String) {
    let mut result = 0;
    input.lines().for_each(|l| {
        let line: Vec<&str> = l.split("|").collect();
        result += line[1]
            .split_ascii_whitespace()
            .filter(|o| matches!(o.len(), 2 | 3 | 4 | 7))
            .count();
    });

    println!("part1: {}", result);
}

pub fn part2(input: String) {
    let mut result = 0;
    input.lines().for_each(|l| {
        let mut final_num = 0;
        let line: Vec<&str> = l.split("|").collect();
        let examples: Vec<&str> = line[0].split_ascii_whitespace().collect();
        let nums: Vec<&str> = line[1].split_ascii_whitespace().collect();
        let mut one = "";
        let mut four = "";
        let mut seven = "";
        for num in examples.iter() {
            match num.len() {
                2 => one = num,
                4 => four = num,
                3 => seven = num,
                _ => {}
            }
        }

        /*
        1, 4, 7, 8

         dddd
        e    a
        e    a
         ffff
        g    b
        g    b
         cccc

        ab: 1
        eafb: 4
        dab: 7
        acedgfb: 8
        cagedb: 0
        cdfgeb: 6
        cefabd: 9
        cdfbe: 5
        fbcad: 3
        gcdfa: 2
                 */

        let four_vec: Vec<&str> = four.split("").collect();
        let one_vec: Vec<&str> = one.split("").collect();
        let seven_vec: Vec<&str> = seven.split("").collect();
        let ef_vec = four_vec
            .into_iter()
            .filter(|ch| !one_vec.contains(ch))
            .collect();

        fn test_include(s: &str, v: &Vec<&str>) -> bool {
            let len_includes: Vec<&str> = s.split("").filter(|ch| v.contains(ch)).collect();
            v.len() == len_includes.len()
        }

        for (i, num) in nums.iter().enumerate() {
            let val = match num.len() {
                2 => 1,
                3 => 7,
                4 => 4,
                7 => 8,
                5 if test_include(num, &one_vec) => 3,
                5 if test_include(num, &ef_vec) => 5,
                5 => 2,
                6 if !test_include(num, &seven_vec) => 6,
                6 if test_include(num, &ef_vec) => 9,
                _ => 0,
            };

            final_num += val * 10_u32.pow(3 - i as u32);
        }

        result += final_num;
    });

    println!("part2: {}", result);
}
