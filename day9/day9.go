package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
    var diskStrMap string
    var diskMap []int
    var err error
    if diskStrMap, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }
    diskMap = strToInt(diskStrMap)

    // NOTE: blockCompressAndCheckSum will modify argument
    // NOTE: Pass in a copy instead if want answer for both parts at the same time
    //fmt.Printf("Part One: %d\n", blockCompressAndChecksum(diskMap))
    fmt.Printf("Part Two: %d\n", checksum(fileCompress(diskMap)))
}

func blockCompressAndChecksum(diskMap []int) int {
    var checksum int = 0
    var leftPtr int = 0
    var rightPtr int
    var currIdx int = 0
    var minID int = 0
    var maxID int
    var isFreeSpace bool = false
    if (len(diskMap) % 2 == 0) {
        maxID = len(diskMap) / 2 - 1
        rightPtr = len(diskMap) - 2
    } else {
        maxID = len(diskMap) / 2
        rightPtr = len(diskMap) - 1
    }
    for (leftPtr <= rightPtr) {
        if (isFreeSpace) {
            var freeSpace int = diskMap[leftPtr]
            for freeSpace > 0 && rightPtr > leftPtr {
                if (diskMap[rightPtr] <= 0) {
                    maxID -= 1
                    rightPtr -= 2
                } else {
                    checksum += (maxID * currIdx)
                    currIdx += 1
                    freeSpace -= 1
                    diskMap[rightPtr] -= 1
                }
            }
            leftPtr += 1
            if (diskMap[rightPtr] <= 0) {
                maxID -= 1
                rightPtr -= 2
            }
        } else {
            var count int = diskMap[leftPtr]
            for count > 0 { // INFO: Can use some formula instead
                checksum += (minID * currIdx)
                currIdx += 1
                count -= 1
            }
            minID += 1
            leftPtr += 1
        }
        isFreeSpace = !isFreeSpace
    }

    return checksum
}

func fileCompress(diskMap []int) []int {
    var formattedDiskMap []int = []int{}
    
    // key: free space size, value: list of startIndex with such size 
    var table map[int][]int = make(map[int][]int) 
    
    // set up hashtable and uncompress diskMap
    var isFreeSpace bool = false
    var id int = 0
    var temp int = 0
    var newDiskMapPtr int = 0
    for _, val := range(diskMap) {
        if (!isFreeSpace) {
            temp = id
            id += 1
        } else {
            temp = -1
            table[val] = append(table[val], newDiskMapPtr)
        }
        for count := 0; count < val; count += 1 {
            formattedDiskMap = append(formattedDiskMap, temp)
            newDiskMapPtr += 1
        }
        isFreeSpace = !isFreeSpace
    }

    // compress at file level and update hashtable
    newDiskMapPtr -= 1
    isFreeSpace = !isFreeSpace
    for i := len(diskMap) - 1; i >= 0; i -= 1 {
        if (!isFreeSpace) {
            var spaceNeeded int = diskMap[i]
            var using int
            var smallestStart int = newDiskMapPtr
            var foundSpace bool = false
            for f := 9; f >= spaceNeeded; f -= 1 {
                if ls, ok := table[f]; (ok && len(ls) > 0) {
                    if (ls[0] < smallestStart) {
                        using = f
                        smallestStart = ls[0]
                        foundSpace = true
                    }
                }
            }
            if (foundSpace) {
                var spaceExtra int = using - spaceNeeded
                var startIndex int = smallestStart
                if (len(table[using]) == 1) {
                    table[using] = []int{}
                } else {
                    table[using] = table[using][1:]
                }
                table[spaceExtra] = append(table[spaceExtra], startIndex + spaceNeeded)
                slices.SortFunc(table[spaceExtra], func(a, b int) int {
                    return a - b
                })
                    
                // update formmatted disk map
                var currID int = formattedDiskMap[newDiskMapPtr]
                for spaceNeeded > 0 {
                    formattedDiskMap[newDiskMapPtr] = -1
                    formattedDiskMap[startIndex] = currID
                    startIndex += 1
                    newDiskMapPtr -= 1
                    spaceNeeded -= 1
                }
            } else {
                newDiskMapPtr -= spaceNeeded
            }
        } else {
            newDiskMapPtr -= diskMap[i]
        }
        isFreeSpace = !isFreeSpace
    }

    return formattedDiskMap
}

func checksum(diskMap []int) int {
    var checksum int = 0
    for i, val := range(diskMap) {
        if (val >= 0) {
            checksum += (i * val)
        }
    }
    return checksum
}

func strToInt(diskStrMap string) []int {
    var diskMap []int
    for _, cn := range(diskStrMap) {
        diskMap = append(diskMap, int(cn - '0'))
    }
    return diskMap
}

func readInput() (string, error) {
    var diskStrMap string
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0) {
            break
        }

        diskStrMap = line
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return diskStrMap, err
    }

    return diskStrMap, nil
}
