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
    let currIsGear = false;
    let sum = 0;

    for (let i = 0; i < lines.length; i++) {
        for (let j = 0; j < lines[i].length; j++) {
            currIsNumber = isDigit(lines[i][j]);

            if (currIsNumber) {
                currNumber += lines[i][j];
                currIsGear = currIsGear || isGear(lines, i, j);
            } else {
                if (currIsGear) {
                    sum += +currNumber;
                }
                currNumber = "";
                currIsGear = false;
            }
        }

        if (currIsGear) {
            sum += +currNumber;
        }
        currNumber = "";
        currIsGear = false;
    }

    console.log("Part 1: ", sum);
}

function getGear(
    arr: string[],
    i: number,
    j: number,
): [number, number] | undefined {
    for (const [di, dj] of dirs) {
        let gear = get(arr, i + di, j + dj);
        let res = !isDot(gear) && !isDigit(gear);
        if (res) {
            return [i + di, j + dj];
        }
    }
}

function part2() {
    const lines = getLines("input.txt");
    let currNumber = "";
    let currIsNumber = false;
    let currGear: number[] | undefined;
    let sum = 0;
    let gears: number[][] = [];

    for (let i = 0; i < lines.length; i++) {
        for (let j = 0; j < lines[i].length; j++) {
            currIsNumber = isDigit(lines[i][j]);

            if (currIsNumber) {
                currNumber += lines[i][j];
                currGear = currGear || getGear(lines, i, j);
            } else {
                if (currGear) {
                    let gr = gears[currGear[0]];
                    if (gr) {
                        let gc = gr[currGear[1]];
                        if (gc) {
                            sum += +currNumber * gc;
                        }
                    }
                    if (!gears[currGear[0]]) {
                        gears[currGear[0]] = [];
                    }
                    gears[currGear[0]][currGear[1]] = +currNumber;
                }
                currNumber = "";
                currGear = undefined;
            }
        }

        if (currGear) {
            let gr = gears[currGear[0]];
            if (gr) {
                let gc = gr[currGear[1]];
                if (gc) {
                    sum += +currNumber * gc;
                }
            }
            if (!gears[currGear[0]]) {
                gears[currGear[0]] = [];
            }
            gears[currGear[0]][currGear[1]] = +currNumber;
        }
        currNumber = "";
        currGear = undefined;
    }

    console.log("Part 2: ", sum);
}

console.log("Day 3");
part1();
part2();
