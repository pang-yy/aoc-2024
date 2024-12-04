package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "strings"
)

func main() {
    var leftArr, rightArr []int = readInput()
    
    var totalDistance int = calculateDistance(leftArr, rightArr)
    var totalSimilarity int = calculateSimilarity(leftArr, rightArr)

    fmt.Printf("Distance: %d\n", totalDistance)
    fmt.Printf("Distance: %d\n", totalSimilarity)
}

func calculateSimilarity(leftArr []int, rightArr []int) int {
    var result int = 0
    var freqMap map[int]int = make(map[int]int)

    for _, value := range rightArr {
        freqMap[value] += 1
    }

    for _, value := range leftArr {
        result += (value * freqMap[value])
    }
    
    return result
}

func calculateDistance(leftArr []int, rightArr []int) int {
    var result int = 0
    slices.Sort(leftArr)
    slices.Sort(rightArr)
    for index := range leftArr {
        if (leftArr[index] > rightArr[index]) {
            result += (leftArr[index] - rightArr[index])
        } else {
            result += (rightArr[index] - leftArr[index])
        }
    }
    return result
}

func readInput() ([]int, []int) {
    var leftArr, rightArr []int
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()
        if (len(line) == 0) {
            break
        }
        var t1, t2 int
        var err1, err2 error
        _, err1 = fmt.Sscanf(strings.Split(line, "   ")[0], "%d", &t1)
        _, err2 = fmt.Sscanf(strings.Split(line, "   ")[1], "%d", &t2)

        if (err1 == nil && err2 == nil) {
            leftArr = append(leftArr, t1)
            rightArr = append(rightArr, t2)
        } else {
            fmt.Println("Error occured when trying to convert string to integer!")
            return leftArr, rightArr
        }
    }

    var err error = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return leftArr, rightArr
    }

    return leftArr, rightArr
}
