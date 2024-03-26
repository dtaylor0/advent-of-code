use std::fs;

#[derive(Debug)]
struct Round {
    red: u32,
    green: u32,
    blue: u32,
}

#[derive(Debug)]
struct Game {
    game_id: u32,
    rounds: Vec<Round>,
}

impl From<&str> for Game {
    fn from(value: &str) -> Self {
        let mut game_and_rounds = value.split(": ");
        let game_id: u32 = game_and_rounds
            .next()
            .unwrap()
            .split_whitespace()
            .nth(1)
            .unwrap()
            .parse()
            .unwrap();
        let rounds = game_and_rounds
            .next()
            .unwrap()
            .split("; ")
            .map(|round| {
                let mut r = Round {
                    red: 0,
                    green: 0,
                    blue: 0,
                };
                for color in round.split(", ") {
                    let mut color_count = color.split_whitespace();
                    let count: u32 = color_count.next().unwrap().parse().unwrap();
                    match color_count.next().unwrap() {
                        "red" => r.red = count,
                        "green" => r.green = count,
                        "blue" => r.blue = count,
                        _ => {}
                    }
                }
                r
            })
            .collect();
        return Game { game_id, rounds };
    }
}

fn part1(contents: &str) {
    let games = contents.lines().map(|line| Game::from(line));
    let sum: u32 = games
        .filter(|game| {
            let (maxred, maxgreen, maxblue) =
                game.rounds.iter().fold((0, 0, 0), |(mr, mg, mb), r| {
                    (mr.max(r.red), mg.max(r.green), mb.max(r.blue))
                });
            return maxred <= 12 && maxgreen <= 13 && maxblue <= 14;
        })
        .map(|game| game.game_id)
        .sum();
    println!("Part 1: {}", sum);
}

fn part2(contents: &str) {
    let games = contents.lines().map(|line| Game::from(line));
    let sum: u32 = games
        .map(|game| {
            let (maxred, maxgreen, maxblue) =
                game.rounds.iter().fold((0, 0, 0), |(mr, mg, mb), r| {
                    (mr.max(r.red), mg.max(r.green), mb.max(r.blue))
                });
            maxred * maxgreen * maxblue
        })
        .sum();
    println!("Part 2: {}", sum);
}

fn main() {
    let contents = fs::read_to_string("input.txt").unwrap();
    println!("Day 2");
    part1(&contents);
    part2(&contents);
}
