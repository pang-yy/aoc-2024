package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var rules [][2]int
    var updates [][]int
    var err error
    rules, updates, err = readInput()
    if (err != nil) {
        fmt.Printf("Error while reading input: %e\n", err)
        return
    }

    var validIndex []int = validUpdates(rules, updates)
    var totalBeforeReorder int = sumMiddle(updates, validIndex)
    
    var reorderedInvalidUpdates [][]int = processInvalidUpdates(updates, rules)
    var totalAfterReorder int
    for _, update := range(reorderedInvalidUpdates) {
        totalAfterReorder += update[len(update) / 2]
    }

    fmt.Printf("Total before reorder: %d\n", totalBeforeReorder)
    fmt.Printf("Total after reorder: %d\n", totalAfterReorder)
}

func processInvalidUpdates(original [][]int, rules [][2]int) [][]int {
    var reorderedInvalidUpdate [][]int = [][]int{}
    var formattedRules map[int][]int = constructRules(rules)
    for _, update := range(original) {
        if (!isValidUpdate(formattedRules, update)) {
            var copyOfInvalidUpdate []int
            copy(update, copyOfInvalidUpdate)
            
            // TODO: Topological Sort or Brute Force (seems like I choose to brute force)
            copyOfInvalidUpdate = reorderUpdateBF(copyOfInvalidUpdate, formattedRules)
            
            reorderedInvalidUpdate = append(reorderedInvalidUpdate, copyOfInvalidUpdate)
        }
    }
    return reorderedInvalidUpdate
}

func reorderUpdateBF(invalidUpdate []int, rules map[int][]int) []int {
    var x, y int
    for rightPtr:= 0; rightPtr < len(invalidUpdate); rightPtr += 1 {
        for leftPtr:= 0; leftPtr < rightPtr; leftPtr += 1 {
            x = invalidUpdate[leftPtr]
            y = invalidUpdate[rightPtr]
            cannotPrecede, hasRule := rules[y]
            if (!hasRule) {
                continue
            }
            for _, numNotAllow := range(cannotPrecede) {
                if (x == numNotAllow) {
                    invalidUpdate[leftPtr], invalidUpdate[rightPtr] = invalidUpdate[rightPtr], invalidUpdate[leftPtr]
                    break
                }
            }
        }
    }
    return invalidUpdate
}

func validUpdates(rules [][2]int, updates [][]int) []int { // return indexes of all valid updates
    var validIndex []int
    var formattedRules map[int][]int = constructRules(rules)
    for index, update := range(updates) {
        if (isValidUpdate(formattedRules, update)) {
            validIndex = append(validIndex, index)
        }
    }

    return validIndex
}

func isValidUpdate(rules map[int][]int, update []int) bool {
    if (len(update) <= 1) {
        return true
    }
    var freqMap map[int]bool = map[int]bool{}

    var beforeMe []int
    var hasRules bool
    for _, value := range(update) {
        freqMap[value] = true
        beforeMe, hasRules = rules[value]
        if (hasRules) {
            for _, precede := range(beforeMe) {
                if (freqMap[precede]) {
                    return false
                }
            }
        }
    }

    return true
}

func constructRules(rules [][2]int) map[int][]int { // key: number x, value: list of numbers CANNOT precede x
    var rulesMap map[int][]int = map[int][]int{}
    var beforeMe []int
    var exists bool

    for _, pair := range(rules) {
        beforeMe, exists = rulesMap[pair[0]]
        if (exists) {
            beforeMe = append(beforeMe, pair[1])
            rulesMap[pair[0]] = beforeMe
        } else {
            rulesMap[pair[0]] = []int{pair[1]}
        }
    }

    return rulesMap
}

func sumMiddle(arrays [][]int, indexes []int) int {
    var total int
    for _, i := range(indexes) {
        total += arrays[i][len(arrays[i]) / 2]
    }
    return total
}

func readInput() ([][2]int, [][]int, error) {
    var rules [][2]int
    var updates [][]int
    var err error
    var isRule bool = true
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0 && !isRule) {
            break
        } else if (len(line) == 0) {
            isRule = false
            continue
        }

        if (isRule) {
            var x int
            var y int
            var err error
            
            _, err = fmt.Sscanf(strings.Split(line, "|")[0], "%d", &x)
            if (err != nil) {
                return rules, updates, err
            }
            _, err = fmt.Sscanf(strings.Split(line, "|")[1], "%d", &y)
            if (err != nil) {
                return rules, updates, err
            }
            
            rules = append(rules, [2]int{x, y})
        } else {
            var temp []int
            var x int
            for _, value := range(strings.Split(line, ",")) {
                _, err = fmt.Sscanf(value, "%d", &x)
                if (err != nil) {
                    return rules, updates, err
                }
                temp = append(temp, x)
            }

            updates = append(updates, temp)
        }
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return rules, updates, err
    }

    return rules, updates, nil
}
