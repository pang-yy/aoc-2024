package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
    "time"
)

func main() {
    var rawInputs []string
    var inputs [][]int
    var err error
    if rawInputs, err = readInput(); err == nil {
        if inputs, err = parseInput(rawInputs); err != nil {
            fmt.Println(err)
            return
        }
    } else {
        fmt.Println(err)
        return
    }

    start := time.Now()
    var ans1, ans2 int
    ans1, ans2 = sumPossible(inputs)
    fmt.Printf("Part One: %d\n", ans1)
    fmt.Printf("Part Two: %d\n", ans2)
    fmt.Printf("Time taken: %s\n", time.Since(start))
}

func sumPossible(inputs [][]int) (int, int) {
    var partOneTotal int = 0
    var partTwoTotal int = 0

    for _, in := range(inputs) {
        var isPossible bool = false
        isValidRecurse(in, 2, in[1], &isPossible)
        if (isPossible) {
            partOneTotal += in[0]
            partTwoTotal += in[0]
        } else {
            if isValid2Recurse(in, 2, in[1], &isPossible); isPossible {
                partTwoTotal += in[0]
            }
        }
    }

    return partOneTotal, partTwoTotal
}

func isValidRecurse(in []int, idx, soFar int, isPossible *bool) {
    if ((*isPossible) || in[0] < soFar) {
        return
    }
    if (idx == int(len(in))) {
        if (soFar == in[0]) {
            *isPossible = true
        }
        return
    }
    isValidRecurse(in, idx + 1, soFar + in[idx], isPossible)
    isValidRecurse(in, idx + 1, soFar * in[idx], isPossible)
}

func isValid2Recurse(in []int, idx, soFar int, isPossible *bool) {
    if ((*isPossible) || in[0] < soFar) {
        return
    }
    if (idx == int(len(in))) {
        if (soFar == in[0]) {
            *isPossible = true
        }
        return
    }
    isValid2Recurse(in, idx + 1, soFar + in[idx], isPossible)
    isValid2Recurse(in, idx + 1, soFar * in[idx], isPossible)
    isValid2Recurse(in, idx + 1, concateInteger(soFar, in[idx]), isPossible)
}

func concateInteger(x, y int) int { // asume no overflow
    if (y == 0) { return x * 10 }
    var cpY, lenY int = y, 1
    for cpY > 0 {
        cpY /= 10
        lenY *= 10
    }
    return (x * lenY) + y
}

func parseInput(rawInputs []string) ([][]int, error) {
    var parseResult [][]int = [][]int{}
    var err error
    for _, line := range(rawInputs) {
        var in []int = []int{}
        var x int
        var temp []string = strings.Split(line, ": ")

        _, err = fmt.Sscanf(temp[0], "%d", &x)
        if (err != nil) {
            return [][]int{}, err
        }
        in = append(in, x)
        
        for _, num := range(strings.Split(temp[1], " ")) {
            _, err = fmt.Sscanf(num, "%d", &x)
            if (err != nil) {
                return [][]int{}, err
            }
            in = append(in, x)
        }
        parseResult = append(parseResult, in)
    }
    return parseResult, nil
}

func readInput() ([]string, error) {
    var input []string
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0) {
            break
        }

        input = append(input, line)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return input, err
    }

    return input, nil
}
