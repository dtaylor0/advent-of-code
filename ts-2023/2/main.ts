import * as fs from "node:fs";

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
        const c = color.split(" ");
        const [count, colorName] = [+c[0], c[1]];
        return count <= limits[colorName];
    });
}

function part1() {
    let sum = 0;

    const lines = getLines("input.txt");
    for (const line of lines) {
        const gameId = +line.split(":")[0].split(" ")[1];
        const rounds = line.split(": ")[1].split("; ");
        const game = rounds.every((round: string) => valid(round));
        if (game) {
            sum += gameId;
        }
    }

    console.log("Part 1: ", sum);
}

function power(game: string): number {
    const maxes: { [key: string]: number } = {
        red: 0,
        green: 0,
        blue: 0,
    };
    const rounds = game.split("; ");
    for (const round of rounds) {
        const colors = round.split(", ");
        for (const color of colors) {
            const c = color.split(" ");
            const [count, colorName] = [+c[0], c[1]];
            if (count > maxes[colorName]) {
                maxes[colorName] = count;
            }
        }
    }
    return maxes.red * maxes.green * maxes.blue;
}

function part2() {
    let sum = 0;

    const lines = getLines("input.txt");
    for (const line of lines) {
        const game = line.split(": ")[1];
        sum += power(game);
    }

    console.log("Part 2: ", sum);
}

console.log("Day 2");
part1();
part2();
