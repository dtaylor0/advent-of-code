import * as fs from "fs";

function getLines(fname: string): string[] {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function part1() {
    const lines = getLines("input.txt");
    let res = lines
        .map((line: string): number => {
            const game = line.split(": ")[1];
            const winners = game
                .split(" | ")[0]
                .trim()
                .split(/[ ]+/)
                .map((num) => +num);
            const cards = game
                .split(" | ")[1]
                .trim()
                .split(/[ ]+/)
                .map((num) => +num);

            let sum = 0;
            for (const card of cards) {
                if (winners.includes(card)) {
                    sum += 1;
                }
            }
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
