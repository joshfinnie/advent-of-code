fn parse_input(input: &String) -> Vec<Vec<u32>> {
    input
        .lines()
        .map(|l| l.chars().map(|n| n.to_digit(10).unwrap()).collect())
        .collect()
}

pub fn part1(input: String) {
    let report = parse_input(&input);
    let mut low_points = vec![];

    for x in 0..report.len() {
        for y in 0..report[x].len() {
            let d = report[x][y];
            if x > 0 && d >= report[x - 1][y]
                || x < report.len() - 1 && d >= report[x + 1][y]
                || y > 0 && d >= report[x][y - 1]
                || y < report[y].len() - 1 && d >= report[x][y + 1]
            {
                continue;
            }

            low_points.push(d);
        }
    }

    let result = low_points.iter().map(|d| d + 1).sum::<u32>();

    println!("part1: {}", result);
}

fn get_basins(
    report: &Vec<Vec<u32>>,
    pts_visited: &mut Vec<Vec<bool>>,
    low_pt: (usize, usize),
) -> u32 {
    let (x, y) = low_pt;
    let mut basin_size = 1;
    let point = report[y][x];

    pts_visited[y][x] = true;

    if x > 0 && !pts_visited[y][x - 1] && point < report[y][x - 1] && report[y][x - 1] != 9 {
        basin_size += get_basins(&report, pts_visited, (x - 1, y));
    }

    if x < report[y].len() - 1
        && !pts_visited[y][x + 1]
        && point < report[y][x + 1]
        && report[y][x + 1] != 9
    {
        basin_size += get_basins(&report, pts_visited, (x + 1, y));
    }

    if y > 0 && !pts_visited[y - 1][x] && point < report[y - 1][x] && report[y - 1][x] != 9 {
        basin_size += get_basins(&report, pts_visited, (x, y - 1));
    }

    if y < report.len() - 1
        && !pts_visited[y + 1][x]
        && point < report[y + 1][x]
        && report[y + 1][x] != 9
    {
        basin_size += get_basins(&report, pts_visited, (x, y + 1));
    }

    basin_size
}

pub fn part2(input: String) {
    let report = parse_input(&input);
    let mut low_points = vec![];

    for x in 0..report.len() {
        for y in 0..report[x].len() {
            let d = report[x][y];
            if x > 0 && d >= report[x - 1][y]
                || x < report.len() - 1 && d >= report[x + 1][y]
                || y > 0 && d >= report[x][y - 1]
                || y < report[y].len() - 1 && d >= report[x][y + 1]
            {
                continue;
            }

            // y is really x and x is really y here...
            low_points.push((y, x));
        }
    }

    let mut basins = vec![];
    let mut pts_visited = vec![vec![false; report.len()]; report[0].len()];
    for pts in low_points {
        basins.push(get_basins(&report, &mut pts_visited, pts));
    }

    basins.sort();

    let result = basins
        .iter_mut()
        .rev()
        .take(3)
        .fold(1, |acc, &mut x| acc * x);

    println!("part2: {}", result);
}
