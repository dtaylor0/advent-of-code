import * as fs from "fs";

const getLines = (fname: string) => {
    return fs.readFileSync(fname, "utf8").trim().split("\n").filter(Boolean);
};

type Convert = {
    src: number;
    dest: number;
    range: number;
};

function contains(conv: Convert, seed: number): boolean {
    return seed >= conv.src && seed < conv.src + conv.range;
}

type CMap = {
    name: string;
    converts: Convert[];
};

function getMap(lines: string[]): CMap {
    let nameLine = lines.shift()!;
    let name = nameLine.slice(0, -5);
    let converts: Convert[] = [];
    while (lines.length && /\d/.test(lines[0][0])) {
        let conv = lines
            .shift()!
            .split(/\s/)
            .map((n) => +n);
        let c = { src: conv[1], dest: conv[0], range: conv[2] };
        converts.unshift(c);
    }
    return { name: name, converts: converts };
}

function part1() {
    const lines = getLines("input.txt");
    const seedsline = lines.shift()!;
    let seeds = seedsline
        .split(/\s+/)
        .slice(1)
        .map((n) => +n);
    while (lines.length) {
        const currMap = getMap(lines);
        let negatives: number[] = [];
        let next: number[] = [];
        for (const conv of currMap.converts) {
            for (const seed of seeds) {
                if (contains(conv, seed)) {
                    let translation = conv.dest - conv.src;
                    next.unshift(seed + translation);
                } else {
                    negatives.unshift(seed);
                }
            }

            seeds = negatives;
            negatives = [];
        }
        seeds = next.concat(negatives);
    }
    console.log("Part 1: ", Math.min(...seeds));
}

function part2() { }

console.log("Day 5");
part1();
part2();
