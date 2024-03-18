use std::fs;

fn file_to_string(fname: &str) -> Result<String, Box<dyn std::error::Error>> {
    let contents = fs::read_to_string(fname)?;
    Ok(contents)
}

fn main() {
    let fname = "input.txt";
    let parsed = file_to_string(fname).unwrap_or_default();
    let trimmed = parsed.trim();
    let s = trimmed.split('\n');

    let mut sum = 0;
    for line in s {
        let first = line.chars().find(|&l| l.is_digit(10));
        let last = line.chars().rfind(|&l| l.is_digit(10));

        let mut res = String::new();
        if first.is_some() {
            res.push(first.unwrap());
        }
        if last.is_some() {
            res.push(last.unwrap());
        }

        let v: i32 = res.parse::<i32>().unwrap();
        sum += v;
    }

    println!("Part 1: {}", sum);
}
