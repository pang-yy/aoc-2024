package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var lines []string = readInput()
    var xmasCount int = findXmasPartOne(lines)
    var XmasCount int = findXmasPartTwo(lines)

    fmt.Printf("XMAS Count: %d\n", xmasCount)
    fmt.Printf("X-MAS Count: %d\n", XmasCount)
}

func findXmasPartOne(lines []string) int {
    var result int = 0
    var directions [8][2]int = [8][2]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
        {-1, -1},
        {-1, 1},
        {1, -1},
        {1, 1},
    }

    for i, line := range(lines) {
        for j := range(line) {
            if (line[j] == 'X') {
                for _, direction := range(directions) {
                    if (checkBoundary(j, i, direction[0], direction[1], len(line), len(lines)) &&
                        checkWord(lines[i + direction[1]][j + direction[0]], 
                                    lines[i + direction[1] * 2][j + direction[0] * 2], 
                                    lines[i + direction[1] * 3][j + direction[0] * 3])) {
                        result += 1
                    }
                }
            }
        }
    }

    return result
}

func findXmasPartTwo(lines []string) int {
    var result int = 0

    var i, j int = 1, 1
    for i < len(lines) {
        j = 1
        for j < len(lines[i]) {
            if (lines[i][j] == 'A' && (i >= 1 && i < len(lines) - 1) && (j >= 1 && j < len(lines[i]) - 1)) {
                if (((lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S') || 
                        (lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M')) && 
                    ((lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S') || 
                        (lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M'))) {
                    result += 1
                }
            }
            j += 1
        }
        i += 1
    }

    return result
}

func checkBoundary(x, y, dx, dy, maxX, maxY int) bool {
    return (x + dx * 3 >= 0) && (x + dx * 3 < maxX) && (y + dy * 3 >= 0) && (y + dy * 3 < maxY)
}

func checkWord(m,a,s byte) bool {
    return (m == 77) && (a == 65) && (s == 83)
}

func readInput() []string {
    var lines []string
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()
        if (len(line) == 0) {
            break
        }

        lines = append(lines, line)
    }

    var err error = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return lines
    }

    return lines
}
