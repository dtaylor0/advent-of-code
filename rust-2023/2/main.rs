use std::fs;

#[derive(Debug, Default)]
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
        let mut parts = value.split(": ");
        let game_id = parts
            .next()
            .and_then(|s| s.split_whitespace().nth(1))
            .and_then(|s| s.parse().ok())
            .unwrap();

        let rounds = parts
            .next()
            .unwrap()
            .split("; ")
            .map(|round| {
                let mut r = Round::default();
                for color in round.split(", ") {
                    let mut color_count = color.split_whitespace();
                    if let (Some(count), Some(color)) = (color_count.next(), color_count.next()) {
                        match color {
                            "red" => r.red = count.parse().unwrap_or(0),
                            "green" => r.green = count.parse().unwrap_or(0),
                            "blue" => r.blue = count.parse().unwrap_or(0),
                            _ => {}
                        }
                    }
                }
                r
            })
            .collect();

        Game { game_id, rounds }
    }
}

fn part1(contents: &str) {
    let sum: u32 = contents
        .lines()
        .map(Game::from)
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
    let sum: u32 = contents
        .lines()
        .map(Game::from)
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
