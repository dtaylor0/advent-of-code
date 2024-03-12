import { readFileSync } from "fs";

console.log("Day 1");

function getValue(line: string): number {
    const reg = /(\d)/g;
    const matches = [...line.matchAll(reg)];
    if (!matches || matches.length === 0) {
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

part1();
