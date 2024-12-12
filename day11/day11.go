package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var stones []int
    var err error
    if stones, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Part One: %d\n", partOne(stones, 25))
    fmt.Printf("Part Two: %d\n", partTwo(stones, 75))
}


// reference: https://youtu.be/EOAFa8j-GVQ
func partTwo(stones []int, blinkTimes int) int {
    var result int = 0
    var cache map[[2]int]int = make(map[[2]int]int) // key: {number, blinkTimes}, value: length of end result
    
    var recv func(int, int) int
    recv =  func(stone, blinkLeft int) int {
        if (blinkLeft <= 0) {
            return 1
        }
        if _, isCached := cache[[2]int{stone, blinkLeft}]; !isCached {
            if (stone == 0) {
                cache[[2]int{stone, blinkLeft}] = recv(1, blinkLeft - 1)
            } else if numberOfDigits := countDigits(stone); (numberOfDigits % 2 == 0) {
                var left, right int = splitNumber(stone, numberOfDigits)
                cache[[2]int{stone, blinkLeft}] = recv(left, blinkLeft - 1) + recv(right, blinkLeft - 1)
            } else {
                cache[[2]int{stone, blinkLeft}] = recv(stone * 2024, blinkLeft - 1)
            }
        }
        return cache[[2]int{stone, blinkLeft}]
    }

    for _, stone := range(stones) {
        result += recv(stone, blinkTimes)
    }

    return result
}

// this solution works for part one when the number of blinks is small
// but will out of memory for part 2 when the number of blinks increased
func partOne(stones []int, blinkTimes int) int {
    for i := 0; i < blinkTimes; i += 1 {
        stones = blink(stones)
    }
    return len(stones)
}

func blink(inital []int) []int {
    var newArrangement []int = make([]int, len(inital))
    copy(newArrangement, inital)
    for i, val := range(inital) {
        if (val == 0) {
            newArrangement[i] = 1
        } else if numberOfDigits := countDigits(val); (numberOfDigits % 2 == 0) {
            var left, right int = splitNumber(val, numberOfDigits)
            newArrangement[i] = left
            newArrangement = append(newArrangement, right) // order of stones doesn't matter
        } else {
            newArrangement[i] *= 2024
        }
    }
    return newArrangement
}

func splitNumber(val, numberOfDigits int) (int, int) {
    var left int = val
    var right int = 0
    var multipleOfTen int = 1

    for i := 0; i < numberOfDigits / 2; i += 1 {
        right = (left % 10 * multipleOfTen) + right
        left /= 10
        multipleOfTen *= 10
    }

    return left, right
}

func countDigits(val int) int {
    var i int = 0
    for val > 0 {
        i += 1
        val /= 10
    }
    return i
}

func readInput() ([]int, error) {
    var stones []int
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    scanner.Scan()
    var line string = scanner.Text()

    var t int
    for _, sv := range(strings.Split(line, " ")) {
        if _, err = fmt.Sscanf(sv, "%d", &t); err != nil {
            return []int{}, err
        }
        stones = append(stones, t)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return stones, err
    }

    return stones, nil
}
