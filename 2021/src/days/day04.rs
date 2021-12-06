fn parse_numbers(input: &String) -> Vec<u8> {
    input.trim().split('\n').nth(0).unwrap().split(',').map(|s| s.parse::<u8>().unwrap()).collect()
}

#[derive(Debug)]
struct Board {
    data: Vec<Vec<(u8, bool)>>,
    has_bingo: bool,
}

impl Board {
    fn new(data: Vec<Vec<(u8, bool)>>) -> Board{
        Board {
            data,
            has_bingo: false,
        }
    }

    fn mark(&mut self, num: u8) {
        for row in self.data.iter_mut() {
            for d in row.iter_mut() {
                if d.0 == num {
                    d.1 = true;
                }
            }
        }
    }

    fn has_bingo(&self) -> bool {
        self.has_bingo
    }

    fn is_bingo(&mut self) -> bool {
        let mut row = 0;
        let mut col = 0;

        for x in 0..5 {
            for y in 0..5 {
                if self.data[x][y].1 {
                    row += 1;
                }
            }

            if row == 5 {
                self.has_bingo = true;
            }

            row = 0;
        }

        for y in 0..5 {
            for x in 0..5  {
                if self.data[x][y].1 {
                    col += 1;
                }
            }

            if col == 5 {
                self.has_bingo = true;
            }

            col = 0;
        }

        self.has_bingo
    }

    fn sum_unmarked(&mut self) -> u32 {
        let mut result: u32 = 0;
        for row in self.data.iter() {
            for data in row.iter() {
                if !data.1 {
                    result += data.0 as u32
                }
            }
        }

        result
    }
}

fn parse_boards(input: &String) -> Vec<Board> {
    let data: Vec<&str> = input.trim().split('\n').filter(|s| !s.is_empty()).skip(1).collect();
    let mut boards = vec![];

    for board in data.chunks(5) {
        let mut data: Vec<Vec<(u8, bool)>> = vec![];
        for row in board {
            data.push(row.split_whitespace().map(|s| s.parse::<u8>().unwrap()).map(|s| (s, false)).collect());
        }

        boards.push(Board::new(data));
    }

    boards
}

fn find_first_bingo(numbers: Vec<u8>, mut boards: Vec<Board>) -> u32 {
    for num in numbers {
        for board in &mut boards {
            board.mark(num);
            if board.is_bingo() {
                return board.sum_unmarked() * num as u32;
            }
        }
    }

    0
}

fn get_last_bingo(numbers: Vec<u8>, mut boards: Vec<Board>) -> u32 {
    let mut score = 0;

    for num in numbers {
        for board in &mut boards {
            board.mark(num);
            if !board.has_bingo() && board.is_bingo() {
                score = board.sum_unmarked() * num as u32;
            }
        }
    }


    score
}

pub fn part1(input: String) {
    let numbers = parse_numbers(&input);
    let boards: Vec<Board> = parse_boards(&input);

    println!("part1: {}", find_first_bingo(numbers, boards));
}

pub fn part2(input: String) {
    let numbers = parse_numbers(&input);
    let boards: Vec<Board> = parse_boards(&input);

    println!("part2: {}", get_last_bingo(numbers, boards));
}
