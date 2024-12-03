package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    var memory string = readInput()
    var normalReuslt int = parseMemoryRegexp(memory, false)
    var conditionalReuslt int = parseMemoryRegexp(memory, true)

    fmt.Printf("Unconditional Result: %d\n", normalReuslt)
    fmt.Printf("Conditional Result: %d\n", conditionalReuslt)
}

func parseMemoryRegexp(memory string, conditional bool) int {
    var matches []string
    var x, y, result int = 0, 0, 0
    var dont bool = false
    var err error

    var re *regexp.Regexp
    if (!conditional) {
	    re = regexp.MustCompile(`mul\(([0-9]|[1-9][0-9]{1,2})[,]([0-9]|[1-9][0-9]{1,2})\)`)
    } else {
	    re = regexp.MustCompile(`mul\(([0-9]|[1-9][0-9]{1,2})[,]([0-9]|[1-9][0-9]{1,2})\)|do\(\)|don't\(\)`)
    }
	matches = re.FindAllString(memory, -1)

    if (!conditional) {
        for _, mul := range(matches) {
            x, y, err = parseMulInstruction(mul)
            if (err != nil) {
                fmt.Println(err)
                return -1
            }
            result += (x * y)
        }
    } else {
        for _, instruction := range(matches) {
            if (instruction[0:3] == "do(") {
                dont = false
            } else if (instruction[0:3] == "don") {
                dont = true
            } else if (!dont) {
                x, y, err = parseMulInstruction(instruction)
                if (err != nil) {
                    fmt.Println(err)
                    return -1
                }
                result += (x * y)
            }
        }
    }
    return result
}

func parseMemory(memory string, conditional bool) int {
    var x int
    var y int
    var result int
    var requirement int = 0

    var index int = 0
    for index < len(memory) {
        if (requirement == 0 && memory[index] == 'm') {
            requirement += 1
        } else if (requirement == 1 && memory[index] == 'u') {
            requirement += 1
        } else if (requirement == 2 && memory[index] == 'l') {
            requirement += 1
        } else if (requirement == 3 && memory[index] == '(') {
            requirement += 1
        } else if (requirement == 4) {
            var err error
            index, x, err = parseInt(memory, index)
            
            if (err != nil) {
                requirement = 0
            } else {
                requirement += 1
            }
        } else if (requirement == 5 && memory[index] == ',') {
            requirement += 1
        } else if (requirement == 6) {
            var err error
            index, y, err = parseInt(memory, index)
            
            if (err != nil) {
                requirement = 0
            } else {
                requirement += 1
            }
        } else if (requirement == 7 && memory[index] == ')') {
            requirement = 0
            result += (x * y)
        } else {
            requirement = 0
        }
        index += 1
    }

    return result
}

func parseMulInstruction(mulStr string) (int, int, error) {
    var left string = strings.Split(strings.Split(mulStr, ",")[0], "(")[1]
    var right string = strings.Split(strings.Split(mulStr, ",")[1], ")")[0]

    x, errX := strconv.Atoi(left)
    y, errY := strconv.Atoi(right)

    if (errX != nil) {
        return 0, 0, errX
    }
    if (errY != nil) {
        return 0, 0, errY
    }
    return x, y, nil
}

func parseInt(memory string, index int) (int, int, error) { // (new index, int, err)
    var temp string = ""
    var argInt int = 0
    var err error
    for index < len(memory) {
        if memory[index] >= 48 && memory[index] <= 57 {
            temp += string(memory[index])
            index += 1
        } else {
            break
        }
    }
    argInt, err = strconv.Atoi(temp)
    index -= 1
    if (err != nil) {
        return index, argInt, err
    }
    return index, argInt, nil
}

func readInput() string {
    var memory string
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()
        if (len(line) == 0) {
            break
        }

        memory += line
    }

    var err error = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return memory
    }

    return memory
}
