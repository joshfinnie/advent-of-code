use std::collections::HashMap;

pub fn part1(input: String) {
    let result_table = HashMap::from([(')', 3), (']', 57), ('}', 1197), ('>', 25137)]);
    let mut result = 0;

    input.lines().for_each(|l| {
        let mut stack: Vec<char> = vec![];
        l.chars().for_each(|c| {
            let last_char = stack.pop();
            match c {
                '(' | '[' | '{' | '<' => {
                    match last_char {
                        Some(c) => stack.push(c),
                        None => {}
                    }
                    stack.push(c);
                }
                ')' => {
                    if last_char.unwrap() != '(' {
                        result += result_table[&c];
                    }
                }
                ']' => {
                    if last_char.unwrap() != '[' {
                        result += result_table[&c];
                    }
                }
                '}' => {
                    if last_char.unwrap() != '{' {
                        result += result_table[&c];
                    }
                }
                '>' => {
                    if last_char.unwrap() != '<' {
                        result += result_table[&c];
                    }
                }
                _ => println!("yikes!"),
            }
        })
    });

    println!("part1: {}", result);
}

pub fn part2(input: String) {
    let result_table = HashMap::from([('(', 1), ('[', 2), ('{', 3), ('<', 4)]);
    let mut result = vec![];

    let _test = r"[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]";

    input.lines().for_each(|l| {
        let mut stack: Vec<char> = vec![];
        let mut corrupted = false;
        l.chars().for_each(|c| {
            let last_char = stack.pop();
            let is_corrupted = match c {
                '(' | '[' | '{' | '<' => {
                    match last_char {
                        Some(c) => stack.push(c),
                        None => {}
                    }
                    stack.push(c);
                    false
                }
                ')' => last_char.unwrap() != '(',
                ']' => last_char.unwrap() != '[',
                '}' => last_char.unwrap() != '{',
                '>' => last_char.unwrap() != '<',
                _ => true,
            };

            corrupted = corrupted || is_corrupted;
        });

        let mut score: u64 = 0;

        for c in stack.iter().rev() {
            score = score * 5 + result_table[&c];
        }

        if !corrupted {
            result.push(score);
        }
    });

    result.sort();
    println!("part2: {}", result[result.len() / 2])
}
