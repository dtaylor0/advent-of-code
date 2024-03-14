import * as fs from "fs";

function getLines(fname: string): Array<string> {
    return fs.readFileSync(fname, "utf8").split("\n");
}

function isDot(char: string): boolean {
    return char === ".";
}

function isDigit(char: string): boolean {
    return char >= "0" && char <= "9";
}

const dirs = [
    [0, 1],
    [1, 1],
    [1, 0],
    [1, -1],
    [0, -1],
    [-1, -1],
    [-1, 0],
    [-1, 1],
];

function get(arr: string[], i: number, j: number): string {
    if (arr[i] === undefined || arr[i][j] === undefined) {
        return ".";
    }
    return arr[i][j];
}

function isGear(arr: string[], i: number, j: number): boolean {
    return dirs.some(([di, dj]): boolean => {
        let res =
            !isDot(get(arr, i + di, j + dj)) &&
            !isDigit(get(arr, i + di, j + dj));
        return res;
    });
}

function part1() {
    const lines = getLines("input.txt");
    let currNumber = "";
    let currIsNumber = false;
    let currIsPart = false;
    let sum = 0;

    for (let i = 0; i < lines.length; i++) {
        for (let j = 0; j < lines[i].length; j++) {
            currIsNumber = isDigit(lines[i][j]);

            if (currIsNumber) {
                currNumber += lines[i][j];
                currIsPart = currIsPart || isGear(lines, i, j);
            } else {
                if (currIsPart) {
                    sum += +currNumber;
                }
                currNumber = "";
                currIsPart = false;
            }
        }

        if (currIsPart) {
            sum += +currNumber;
        }
        currNumber = "";
        currIsPart = false;
    }

    console.log("Part 1: ", sum);
}

function part2() {}

console.log("Day 3");
part1();
part2();
