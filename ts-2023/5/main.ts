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

function cut(
    conv: Convert,
    seed: Seed,
): [Seed | undefined, Seed[] | undefined] {
    const { src, range, dest } = conv;
    const { start, length } = seed;
    const translation = dest - src;

    if (start >= src + range || start + length <= src) {
        return [undefined, [seed]];
    }

    const cutStart = Math.max(start, src);
    const cutEnd = Math.min(start + length, src + range);
    const cutLength = cutEnd - cutStart;

    const cutSeed: Seed = {
        start: cutStart + translation,
        length: cutLength,
    };

    const newSeeds: Seed[] = [];
    if (start < cutStart) {
        newSeeds.push({ start, length: cutStart - start });
    }
    if (start + length > cutEnd) {
        newSeeds.push({ start: cutEnd, length: start + length - cutEnd });
    }

    return [cutSeed, newSeeds.length ? newSeeds : undefined];
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
        converts.push(c);
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

type Seed = {
    start: number;
    length: number;
};

function part2() {
    const lines = getLines("input.txt");
    const seedsline = lines.shift()!;
    let seedsValues = seedsline.split(/\s+/).slice(1);
    let seeds: Seed[] = [];
    for (let i = 0; i < seedsValues.length; i += 2) {
        seeds.push({
            start: +seedsValues[i],
            length: +seedsValues[i + 1],
        });
    }
    while (lines.length) {
        const currMap = getMap(lines);
        let negatives: Seed[] = [];
        let next: Seed[] = [];
        for (const conv of currMap.converts) {
            for (const seed of seeds) {
                let [cutSeed, newSeeds] = cut(conv, seed);
                if (cutSeed) {
                    next.push(cutSeed);
                }
                if (newSeeds) {
                    negatives.push(
                        ...newSeeds,
                    );
                }
            }

            seeds = negatives;
            negatives = [];
        }
        seeds = seeds.concat(next);
    }
    let m = seeds[0].start;
    for (const seed of seeds) {
        if (seed.start < m) {
            m = seed.start;
        }
    }
    console.log("Part 2: ", m);
}

console.log("Day 5");
part1();
part2();
