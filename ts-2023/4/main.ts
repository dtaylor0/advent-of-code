import * as fs from "fs";

function getLines(fname: string): string[] {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function part1() {
    const lines = getLines("input.txt");
    let res = lines
        .map((line: string): number => {
            const card = line.split(": ")[1].trim().split(" | ");
            const winners = card[0].split(/\s+/);
            const nums = card[1].split(/\s+/);

            let sum = nums.reduce(
                (acc, card) => acc + Number(winners.includes(card)),
                0,
            );
            return sum > 0 ? 1 * 2 ** (sum - 1) : 0;
        })
        .reduce((acc, item) => acc + item, 0);
    console.log("Part 1: ", res);
}

function part2() {
    const lines = getLines("input.txt");
    const counts = new Map<number, number>();
    let total = 0;
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i];
        const card = line.split(": ")[1].trim().split(" | ");
        const winners = card[0].split(/\s+/);
        const nums = card[1].split(/\s+/);

        let sum = nums.reduce(
            (acc, card) => acc + Number(winners.includes(card)),
            0,
        );
        const cardId = +line.split(": ")[0].trim().split(/\s+/)[1];
        let cardCount = counts.get(cardId) || 0;
        cardCount++;
        total += cardCount;

        if (sum > 0) {
            for (let j = 1; j <= sum; j++) {
                const currCount = counts.get(cardId + j) || 0;
                counts.set(cardId + j, cardCount + currCount);
            }
        }
    }
    console.log("Part 2: ", total);
}

console.log("Day 4");
part1();
part2();
