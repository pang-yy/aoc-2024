package main

import (
    "bufio"
    "fmt"
    "os"
    "slices"
    "strings"
)

func main() {
    var reports [][]int = readInput()
    var unsafeReports [][]int

    var minSafe int = 0
    var maxSafe int = 0

    for _, levels := range(reports) { // O(n^2)
        if (partOneReport(levels, levels[0] < levels[1])) {
            minSafe += 1
        } else {
            unsafeReports = append(unsafeReports, levels)
        }
    }

    maxSafe = minSafe
    for _, levels := range(unsafeReports) { // O(n^3) by using brute force
        if (partTwoReportBF(levels)) {
            maxSafe += 1
        }
    }

    fmt.Printf("Part 1 saves: %d\n", minSafe)
    fmt.Printf("Part 2 saves: %d\n", maxSafe)
}

func partOneReport(report []int, isInc bool) bool { // O(n) for one report
    if (isInc) {
        for index := 0; index < len(report) - 1; index++ {
            if (report[index + 1] - report[index] <= 0 ||
                report[index + 1] - report[index] > 3) {
                return false
            }
        }
    } else {
        for index := 0; index < len(report) - 1; index++ {
            if (report[index] - report[index + 1] <= 0 ||
                report[index] - report[index + 1] > 3) {
                return false
            }
        }
    }
    return true
}

func partTwoReportBF(report []int) bool { // O(n^2) for one report
    var temp []int
    for index := range(report) {
        temp = slices.Clone(report)
        temp = slices.Delete(temp, index, index + 1)
        if (partOneReport(temp, temp[0] < temp[1])) {
            return true
        }
    }
    return false
}

func partTwoReport(report []int) bool {
    panic("not yet implemented")
}

func readInput() [][]int {
    var reports [][]int
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()
        if (len(line) == 0) {
            break
        }
        var levels []int
        var temp int
        var err1 error

        for _, value := range(strings.Split(line, " ")) {
            _, err1 = fmt.Sscanf(value, "%d", &temp)
            if (err1 != nil) {
                fmt.Println("Error occured when trying to convert string to integer!")
                return reports
            }
            levels = append(levels, temp)
        }
        reports = append(reports, levels)
    }

    var err error = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return reports
    }

    return reports
}
