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

fn max_value(line: &str, numbers: &[&str; 9]) -> u32 {
    let mut maxes: Vec<(usize, u32)> = vec![];
    let maxidx = line.rfind(char::is_numeric);
    match maxidx {
        Some(e) => maxes.push((e, line.chars().nth(e).unwrap().to_digit(10).unwrap())),
        _ => (),
    }

    let maxnumber = numbers
        .iter()
        .enumerate()
        .filter_map(|(i, n)| {
            let idx = line.rfind(n);
            match idx {
                Some(v) => Some((v, i as u32 + 1)),
                None => None,
            }
        })
        .max_by(|(a, _), (b, _)| a.cmp(b));
    match maxnumber {
        Some(e) => maxes.push(e),
        _ => (),
    }

    return maxes.iter().max().unwrap().1;
}

fn min_value(line: &str, numbers: &[&str; 9]) -> u32 {
    let mut mins: Vec<(usize, u32)> = vec![];
    let min_digit = line.chars().enumerate().find(|(_, c)| c.is_digit(10));

    match min_digit {
        Some(value) => mins.push((value.0, value.1.to_digit(10).unwrap())),
        _ => (),
    }

    let min_number = numbers
        .iter()
        .enumerate()
        .filter_map(|(i, n)| {
            let idx = line.find(n);
            match idx {
                Some(v) => Some((v, i as u32 + 1)),
                None => None,
            }
        })
        .min_by(|(a, _), (b, _)| a.cmp(b));

    match min_number {
        Some(value) => mins.push(value),
        _ => (),
    }

    mins.iter().min().unwrap().1
}

fn part2() {
    let contents = fs::read_to_string("input.txt").unwrap();

    let numbers = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let sum: u32 = contents
        .lines()
        .map(|line| 10 * min_value(line, &numbers) + max_value(line, &numbers))
        .sum();

    println!("Part 2: {}", sum);
}

fn main() {
    println!("Day 1");
    println!("-----");
    part1();
    part2();
}
