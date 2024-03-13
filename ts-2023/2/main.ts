import * as fs from "fs";

function getLines(fname: string): Array<string> {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function valid(round: string): boolean {
    const limits: { [key: string]: number } = {
        red: 12,
        green: 13,
        blue: 14,
    };
    const colors = round.split(", ");
    return colors.every((color: string) => {
        let c = color.split(" ");
        let [count, colorName] = [+c[0], c[1]];
        return count <= limits[colorName];
    });
}

function part1() {
    let sum = 0;

    const lines = getLines("input.txt");
    for (const line of lines) {
        const gameId = +line.split(":")[0].split(" ")[1];
        const rounds = line.split(": ")[1].split("; ");
        let game = rounds.every((round: string) => valid(round));
        if (game) {
            sum += gameId;
        }
    }

    console.log("Part 1: ", sum);
}

function part2() {}

console.log("Day 2");
part1();
part2();
