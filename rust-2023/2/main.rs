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
        let game_and_rounds: Vec<&str> = value.split(": ").collect();
        let game_id: u32 = game_and_rounds[0].split_whitespace().collect::<Vec<&str>>()[1]
            .parse()
            .unwrap();
        let rounds = game_and_rounds[1]
            .split("; ")
            .map(|round| {
                let mut r = Round {
                    red: 0,
                    green: 0,
                    blue: 0,
                };
                for color in round.split(", ") {
                    let color_count = color.split_whitespace().collect::<Vec<&str>>();
                    let count: u32 = color_count[0].parse().unwrap();
                    let color = color_count[1];
                    if color == "red" {
                        r.red = count;
                    } else if color == "green" {
                        r.green = count;
                    } else if color == "blue" {
                        r.blue = count;
                    }
                }
                r
            })
            .collect();
        return Game { game_id, rounds };
    }
}

fn part1() {
    let contents = fs::read_to_string("input.txt").unwrap();
    let games = contents.lines().map(|line| Game::from(line));

    let sum: u32 = games
        .map(|game| {
            let maxred = game.rounds.iter().max_by_key(|r| r.red).unwrap().red;
            let maxgreen = game.rounds.iter().max_by_key(|r| r.green).unwrap().green;
            let maxblue = game.rounds.iter().max_by_key(|r| r.blue).unwrap().blue;
            maxred * maxgreen * maxblue
        })
        .sum();
    println!("Part 1: {}", sum);
}

fn main() {
    println!("Day 2");
    part1();
}
