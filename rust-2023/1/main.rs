use std::fs;

fn part1() {
    let contents = fs::read_to_string("input.txt").unwrap();
    let sum: u32 = contents
        .lines()
        .filter_map(|line| {
            let first = line.chars().find_map(|c| c.to_digit(10))?;
            let last = line.chars().rev().find_map(|c| c.to_digit(10))?;
            Some(10 * first + last)
        })
        .sum();

    println!("Part 1: {}", sum);
}

fn part2() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let numbers = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let sum: u32 = contents
        .lines()
        .map(|line| {
            let mut mins: Vec<(usize, u32)> = vec![];
            let mut maxes: Vec<(usize, u32)> = vec![];

            let mindigit = line.chars().enumerate().find(|(_, c)| c.is_digit(10));
            match mindigit {
                Some(e) => mins.push((e.0, e.1.to_digit(10).unwrap())),
                _ => (),
            }

            let maxdigit = line
                .chars()
                .enumerate()
                .collect::<Vec<_>>()
                .into_iter()
                .rfind(|&(_, c)| c.is_digit(10));
            match maxdigit {
                Some(e) => maxes.push((e.0, e.1.to_digit(10).unwrap())),
                _ => (),
            }

            let minnumber = numbers
                .iter()
                .map(|n| line.find(n).unwrap_or(usize::MAX))
                .enumerate()
                .min_by(|(_, a), (_, b)| a.cmp(b));
            match minnumber {
                Some(e) => mins.push((e.1, 1 + e.0 as u32)),
                _ => (),
            }

            let maxnumber = numbers
                .iter()
                .enumerate()
                .filter_map(|(i, n)| {
                    let idx = line.rfind(n);
                    match idx {
                        Some(v) => Some((i, v)),
                        None => None,
                    }
                })
                .max_by(|(_, a), (_, b)| a.cmp(b));
            match maxnumber {
                Some(e) => maxes.push((e.1, 1 + e.0 as u32)),
                _ => (),
            }

            return 10 * mins.iter().min_by_key(|(i, _)| i).unwrap().1
                + maxes.iter().max_by_key(|(i, _)| i).unwrap().1;
        })
        .sum();

    println!("Part 2: {}", sum);
}

fn main() {
    println!("Day 1");
    println!("-----");
    part1();
    part2();
}
