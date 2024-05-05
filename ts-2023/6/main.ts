import * as fs from "node:fs";

function getLines(fname: string): string[] {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function distance(holdTime: number, raceTime: number): number {
    return (raceTime - holdTime) * holdTime;
}

function findMinHold(t: number, d: number, lo: number, hi: number): number {
    if (lo === hi - 1) {
        if (distance(lo, t) > d) {
            return lo;
        }
        return hi;
    }
    const mid = Math.floor((hi + lo) / 2);
    if (distance(mid, t) > d) {
        return findMinHold(t, d, lo, mid);
    }
    return findMinHold(t, d, mid, hi);
}

function part1() {
    const lines = getLines("input.txt");
    const [times, distances] = lines.map((line: string) =>
        line
            .split(/\s+/)
            .slice(1)
            .map((n) => +n),
    );

    let res = 1;
    for (let i = 0; i < times.length; i++) {
        const [t, d] = [times[i], distances[i]];
        const minHold = findMinHold(t, d, 0, t);
        const options = t + 1;
        res *= options - 2 * minHold;
    }

    console.log("Part 1: ", res);
}

function part2() {
    const lines = getLines("input.txt");
    const [t, d] = lines.map(
        (line: string) => +line.split(/\s+/).slice(1).join(""),
    );
    const minHold = findMinHold(t, d, 0, t);
    const options = t + 1;
    const res = options - 2 * minHold;
    console.log("Part 2: ", res);
}

console.log("Day 6: ");
part1();
part2();
