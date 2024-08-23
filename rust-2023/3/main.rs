use regex::Regex;

const INPUT: &str = include_str!("input.txt");

#[derive(Debug)]
struct PartNumber {
    value: i32,
    row: usize,
    start: usize,
    end: usize,
}

fn part1() {
    let number_re = Regex::new("[0-9]+").unwrap();
    let sym_re = Regex::new("[^0-9.]").unwrap();
    let nums = INPUT.lines().enumerate().map(|(idx, line)| {
        number_re
            .find_iter(line)
            .map(|m| PartNumber {
                value: m.as_str().parse().expect("No number found"),
                row: idx,
                start: m.start(),
                end: m.end(),
            })
            .collect::<Vec<PartNumber>>()
    });
    let syms = INPUT.lines().enumerate().filter_map(|(idx, line)| {
        let matches = sym_re.find(line);
        match matches {
            Some(m) => Some((idx, m)),
            None => None,
        }
    });
    /*let res = nums.map(|(i, m)| {
        let curr_matches = m.unwrap();
        let cols = curr_matches.start()-1..curr_matches.end();
        let is_part = (i-1..=i+1).reduce(|acc, idx| {acc || (syms.nth(idx).unwrap().1.unwrap())});
    }).sum();
    */
    println!(
        "{:#?}\n\n\n",
        nums.collect::<Vec<_>>(),
    );
}

fn part2() {}

fn main() {
    println!("Day 3");
    part1();
    part2();
}
