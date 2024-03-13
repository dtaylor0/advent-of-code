import { readFileSync } from "fs";

function getValue(line: string): number {
    const reg = /(\d)/g;
    const matches = [...line.matchAll(reg)];
    if (!matches) {
        throw new Error("No digit");
    }

    const first = matches[0][0];
    const last = matches[matches.length - 1][0][0];

    return parseInt(first + last);
}

function part1() {
    const lines = readFileSync("input.txt").toString().trim().split("\n");
    let sum = lines.reduce((acc: number, line: string): number => {
        return acc + getValue(line);
    }, 0);
    console.log("Part 1: %d", sum);
}

type Value = {
    idx: number;
    val: string;
};

function getValuePart2(line: string): number {
    let first: Value | undefined;
    let last: Value | undefined;
    const reg = /(\d)/g;
    const matches = [...line.matchAll(reg)];
    if (matches) {
        let firstMatch = matches[0];
        let firstIdx = firstMatch.index;
        if (firstIdx === undefined) {
            firstIdx = -1;
        }
        if (firstIdx >= 0) {
            first = { idx: firstIdx, val: firstMatch[0] };
        }

        let lastMatch = matches[matches.length - 1];
        let lastIdx = lastMatch.index;
        if (lastIdx === undefined) {
            lastIdx = -1;
        }
        if (lastIdx >= 0) {
            last = { idx: lastIdx, val: lastMatch[0] };
        }
    }

    let digits = {
        one: "1",
        two: "2",
        three: "3",
        four: "4",
        five: "5",
        six: "6",
        seven: "7",
        eight: "8",
        nine: "9",
    };

    for (const [digit, v] of Object.entries(digits)) {
        let firstOcc = line.indexOf(digit);
        if (firstOcc >= 0) {
            if (!first || firstOcc < first.idx) {
                first = { idx: firstOcc, val: v };
            }
        }

        let lastOcc = line.lastIndexOf(digit);
        if (lastOcc >= 0) {
            if (!last || lastOcc > last.idx) {
                last = { idx: lastOcc, val: v };
            }
        }
    }

    if (first && last) {
        return parseInt(first.val + last.val);
    } else {
        throw new Error("No number(s) found");
    }
}

function part2() {
    const lines = readFileSync("input.txt").toString().trim().split("\n");
    let sum = lines.reduce((acc: number, line: string): number => {
        return acc + getValuePart2(line);
    }, 0);
    console.log("Part 2: %d", sum);
}

console.log("Day 1");
part1();
part2();
