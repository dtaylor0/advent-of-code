import * as fs from "fs";

function getLines(fname: string): string[] {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function part1() {
    const lines = getLines("input.txt");
    let res = lines
        .map((line: string): number => {
            const game = line.split(": ")[1].trim();
            const winners = game
                .split(" | ")[0]
                .split(/[ ]+/)
            const cards = game
                .split(" | ")[1]
                .split(/[ ]+/)

            let sum = cards.reduce(
                (acc, card) => acc + (winners.includes(card) ? 1 : 0),
                0,
            );
            if (sum > 0) {
                return 1 * 2 ** (sum - 1);
            }
            return 0;
        })
        .reduce((acc, item) => acc + item, 0);
    console.log("Part 1: ", res);
}

function part2() { }

console.log("Day 4");
part1();
part2();
