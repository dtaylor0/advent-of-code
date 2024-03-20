#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int BUFSIZE = 100;
char numbers[][9] = {"one", "two",   "three", "four", "five",
                     "six", "seven", "eight", "nine"};

// char **getlines(char *fname) {
//     FILE *fp;
//     char *line = malloc(BUFSIZE);
//     char **output;
//
//     fp = fopen("input.txt", "r");
//     if (fp == NULL) {
//         perror("Error opening file");
//         exit(1);
//     }
//
//     while (fgets(line, BUFSIZE, fp) != NULL) {
//     }
// }

int part1() {
    FILE *fp;
    char str[BUFSIZE];

    fp = fopen("input.txt", "r");
    if (fp == NULL) {
        perror("Error opening file");
        exit(1);
    }

    int sum = 0;
    while (fgets(str, BUFSIZE, fp) != NULL) {
        int leftIdx = -1, rightIdx = -1;
        for (int i = 0; i < BUFSIZE; i++) {
            if (str[i] == '\0') {
                break;
            }
            if (str[i] >= '0' && str[i] <= '9') {
                if (leftIdx == -1) {
                    leftIdx = i;
                }
                rightIdx = i;
            }
        }

        char res[3];
        res[0] = str[leftIdx], res[1] = str[rightIdx], res[2] = '\0';

        int value = atoi(res);
        sum += value;

        leftIdx = -1, rightIdx = -1;
    }
    fclose(fp);

    printf("Part 1: %d\n", sum);

    return 0;
}

int *findNumberRight(char *line) {
    int *res = malloc(sizeof(int) * 2);
    res[0] = -1;

    for (int i = 0; i < 9; i++) {
        char *idx = strstr(line, numbers[i]);
        if (idx != NULL) {
            char *validIdx;
            while (idx != NULL) {
                validIdx = idx;
                idx = strstr(validIdx+1, numbers[i]);
            }
            int currIdx = validIdx - &line[0];
            if (res[0] == -1 || currIdx > res[0]) {
                res[0] = currIdx;
                res[1] = i + 1;
            }
        }
    }
    return res;
}

int *findNumberLeft(char *line) {
    int *res = malloc(sizeof(int) * 2);
    res[0] = -1;

    for (int i = 0; i < 9; i++) {
        char *idx = strstr(line, numbers[i]);
        if (idx != NULL) {
            int currIdx = idx - &line[0];
            if (res[0] == -1 || currIdx < res[0]) {
                res[0] = currIdx;
                res[1] = i + 1;
            }
        }
    }
    return res;
}

int findDigitRight(char *line) {
    int rightIdx = -1;
    for (int i = 0; i < BUFSIZE; i++) {
        if (line[i] == '\0') {
            break;
        }
        if (line[i] >= '0' && line[i] <= '9') {
            rightIdx = i;
        }
    }
    return rightIdx;
}

int findDigitLeft(char *line) {
    int leftIdx = -1;
    for (int i = 0; i < BUFSIZE; i++) {
        if (line[i] == '\0') {
            break;
        }
        if (line[i] >= '0' && line[i] <= '9') {
            if (leftIdx == -1) {
                leftIdx = i;
            }
        }
    }
    return leftIdx;
}

int part2() {
    FILE *fp;
    char line[BUFSIZE];

    fp = fopen("input.txt", "r");
    if (fp == NULL) {
        perror("Error opening file");
        exit(1);
    }

    int sum = 0;
    while (fgets(line, BUFSIZE, fp) != NULL) {
        int leftIdx = findDigitLeft(line);
        char ldigit[2] = {line[leftIdx], '\0'};
        int lvalue = atoi(ldigit);

        int *leftNum = findNumberLeft(line);
        if (leftNum[0] < leftIdx && leftNum[0] != -1) {
            lvalue = leftNum[1];
        }

        int rightIdx = findDigitRight(line);
        char rdigit[2] = {line[rightIdx], '\0'};
        int rvalue = atoi(rdigit);

        int *rightNum = findNumberRight(line);
        if (rightNum[0] > rightIdx && rightNum[0] != -1) {
            rvalue = rightNum[1];
        }

        sum += 10 * lvalue + rvalue;
    }
    fclose(fp);

    printf("Part 2: %d\n", sum);

    return 0;
}

int main() {
    printf("Day 1\n");
    part1();
    part2();
    return 0;
}
