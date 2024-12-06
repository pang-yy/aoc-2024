package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var patrolMap []string
    var err error
    patrolMap, err = readInput()
    if (err != nil) {
        fmt.Println(err)
        return
    }

    var partOne int
    var partTwo int
    partOne, _ = guardPaths(patrolMap)
    partTwo = findAllCycle(patrolMap)

    fmt.Printf("Part One: %d\n", partOne)
    fmt.Printf("Part Two: %d\n", partTwo)
}

func findAllCycle(patrolMap []string) int {
    var result int = 0
    for y, path := range(patrolMap) {
        for x, object := range(path) {
            if (object == '.') {
                patrolMap[y] = replacePath(path, x, '#')
                _, isCycle := guardPaths(patrolMap)
                if (isCycle) {
                    result += 1
                }
                patrolMap[y] = replacePath(path, x, '.')
            }
        }
    }
    return result
}

func guardPaths(patrolMap []string) (int, bool) {
    var directions [][2]int = [][2]int{
        {0, -1}, // up
        {1, 0}, // right
        {0, 1}, // down
        {-1, 0}, // left
    }
    var dirPtr int = 0
    var isExit bool = false
    var totalDistinctMoves int = 1
    var currentMoves int = 0
    var xNow, yNow int = findGuard(patrolMap)
    var beenTo [][]bool = [][]bool{}

    // for path two
    var MAX_MOVES int = len(patrolMap) * len(patrolMap[0]) * 4
    var totalMoves int = 0
    var actualMoves int = 0

    // initialise beenTo
    for y := 0; y < len(patrolMap); y += 1 {
        beenTo = append(beenTo, []bool{})
        for x := 0; x < len(patrolMap[0]); x += 1 {
            beenTo[y] = append(beenTo[y], false)
        }
    }
    beenTo[yNow][xNow] = true

    for !isExit {
        beenTo, currentMoves, xNow, yNow, isExit, actualMoves = moveUtil(patrolMap, beenTo, xNow, yNow, 
            directions[dirPtr][0], directions[dirPtr][1])
        totalDistinctMoves += currentMoves
        
        totalMoves += actualMoves
        if (totalMoves >= MAX_MOVES) {
            return totalMoves, true
        }
        
        dirPtr = (dirPtr + 1) % len(directions)
    }

    return totalMoves, false
}

func moveUtil(patrolMap []string, beenTo [][]bool, x, y, dx, dy int) (
    [][]bool, int, int, int, bool, int) { // beenTo, distinctMovesCount, newX, newY, isExit, actualMovesCount
    var distinctMoves int = 0

    var totalMoves int = 0

    for (y + dy) < len(patrolMap) && (y + dy) >= 0 && 
        (x + dx) < len(patrolMap[0]) && (x + dx) >= 0 {
        x += dx
        y += dy
        if (patrolMap[y][x] == '#') {
            return beenTo, distinctMoves, x - dx, y - dy, false, totalMoves
        }
        totalMoves += 1
        if (!beenTo[y][x]) {
            distinctMoves += 1
            beenTo[y][x] = true
        }
    }
    return beenTo, distinctMoves, x, y, true, totalMoves
}

func findGuard(patrolMap []string) (int, int) {
    for y, path := range(patrolMap) {
        for x, object := range(path) {
            if (object == '^') {
                return x, y
            }
        }
    }
    return -1, -1
}

func replacePath(ori string, idx int, to rune) string {
    var newPath string = ""
    for i, char := range(ori) {
        if (i == idx) {
            newPath += string(to)
        } else {
            newPath += string(char)
        }
    }
    return newPath
}

func readInput() ([]string, error) {
    var patrolMap []string
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0) {
            break
        }

        patrolMap = append(patrolMap, line)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return patrolMap, err
    }

    return patrolMap, nil
}
