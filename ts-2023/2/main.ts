import * as fs from "fs";

function getLines(fname: string): Array<string> {
    return fs.readFileSync(fname, "utf8").trim().split("\n");
}

function part1() {
    const limits: { [key: string]: number } = {
        red: 12,
        green: 13,
        blue: 14,
    };
    let sum = 0;

    const lines = getLines("input.txt");
    for (const line of lines) {
        const gameId = +line.split(":")[0].split(" ")[1];
        const rounds = line.split(": ")[1].split("; ");
        let game = rounds.reduce(
            (acc: boolean, round: string): boolean =>
                acc &&
                round
                    .split(", ")
                    .reduce(
                        (accColor: boolean, color: string): boolean =>
                            accColor &&
                            +color.split(" ")[0] <= limits[color.split(" ")[1]],
                        true,
                    ),
            true,
        );
        if (game) {
            sum += gameId;
        }
    }

    console.log("Part 1: ", sum);
}

function part2() {}

console.log("Day 2");
part1();
part2();
