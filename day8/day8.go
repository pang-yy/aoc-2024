package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var antennaStrMap []string
    var err error
    if antennaStrMap, err = readInput(); err != nil {
        fmt.Println(err)
        return
    }

    var antennaMap map[byte][][2]int = parseInput(antennaStrMap)
    
    fmt.Printf("Part One: %d\n", countAntinode1(antennaMap, len(antennaStrMap[0]), len(antennaStrMap)))
    fmt.Printf("Part Two: %d\n", countAntinode2(antennaMap, len(antennaStrMap[0]), len(antennaStrMap)))
}

func countAntinode1(antennaMap map[byte][][2]int, maxX, maxY int) int {
    var count int = 0
    var distinctMap []bool = make([]bool, maxX * maxY)
    for _, coors := range(antennaMap) {
        var i, j int = 0, 1
        for i < len(coors) - 1 {
            j = i + 1
            for j < len(coors) {
                var dx int = coors[i][0] - coors[j][0]
                var dy int = coors[i][1] - coors[j][1]
                
                var newC1 [2]int = [2]int{coors[i][0] + dx, coors[i][1] + dy}
                var newC1a [2]int = [2]int{coors[i][0] - dx, coors[i][1] - dy}
                var newC2 [2]int = [2]int{coors[j][0] + dx, coors[j][1] + dy}
                var newC2a [2]int = [2]int{coors[j][0] - dx, coors[j][1] - dy}
                if (isValidCoor1(newC1, coors[j], maxX, maxY) && 
                    !distinctMap[newC1[1] * maxY + newC1[0]]) {
                    count += 1
                    distinctMap[newC1[1] * maxY + newC1[0]] = true
                }
                if (isValidCoor1(newC1a, coors[j], maxX, maxY) &&
                    !distinctMap[newC1a[1] * maxY + newC1a[0]]) {
                    count += 1
                    distinctMap[newC1a[1] * maxY + newC1a[0]] = true
                }
                if (isValidCoor1(newC2, coors[i], maxX, maxY) &&
                    !distinctMap[newC2[1] * maxY + newC2[0]]) {
                    count += 1
                    distinctMap[newC2[1] * maxY + newC2[0]] = true
                }
                if (isValidCoor1(newC2a, coors[i], maxX, maxY) &&
                    !distinctMap[newC2a[1] * maxY + newC2a[0]]) {
                    count += 1
                    distinctMap[newC2a[1] * maxY + newC2a[0]] = true
                }
                j += 1
            }
            i += 1
        }
    }
    return count
}

func countAntinode2(antennaMap map[byte][][2]int, maxX, maxY int) int {
    var count int = 0
    var distinctMap []bool = make([]bool, maxX * maxY)
    for _, coors := range(antennaMap) {
        var i, j int = 0, 1
        for i < len(coors) - 1 {
            j = i + 1
            for j < len(coors) {
                var dx int = coors[i][0] - coors[j][0]
                var dy int = coors[i][1] - coors[j][1]
                
                var newC1 [2]int = [2]int{coors[i][0] + dx, coors[i][1] + dy}
                var newC1a [2]int = [2]int{coors[i][0] - dx, coors[i][1] - dy}
                var newC2 [2]int = [2]int{coors[j][0] + dx, coors[j][1] + dy}
                var newC2a [2]int = [2]int{coors[j][0] - dx, coors[j][1] - dy}
                for isValidCoor2(newC1, maxX, maxY) {
                    if (!distinctMap[newC1[1] * maxY + newC1[0]]) {
                        count += 1
                        distinctMap[newC1[1] * maxY + newC1[0]] = true
                    }
                    newC1[0] += dx
                    newC1[1] += dy
                }
                for isValidCoor2(newC1a, maxX, maxY) {
                    if (!distinctMap[newC1a[1] * maxY + newC1a[0]]) {
                        count += 1
                        distinctMap[newC1a[1] * maxY + newC1a[0]] = true
                    }
                    newC1a[0] -= dx
                    newC1a[1] -= dy
                }
                for isValidCoor2(newC2, maxX, maxY) {
                    if (!distinctMap[newC2[1] * maxY + newC2[0]]) {
                        count += 1
                        distinctMap[newC2[1] * maxY + newC2[0]] = true
                    }
                    newC2[0] += dx
                    newC2[1] += dy
                }
                for isValidCoor2(newC2a, maxX, maxY) {
                    if (!distinctMap[newC2a[1] * maxY + newC2a[0]]) {
                        count += 1
                        distinctMap[newC2a[1] * maxY + newC2a[0]] = true
                    }
                    newC2a[0] -= dx
                    newC2a[1] -= dy
                }
                j += 1
            }
            i += 1
        }
    }
    return count
}

func isValidCoor1(coor [2]int, other [2]int, maxX, maxY int) bool {
    return (coor[0] < maxX) && (coor[1] < maxY) && 
            (coor[0] >= 0) && (coor[1] >= 0) &&
            (coor[0] != other[0]) && (coor[1] != other[1])
}

func isValidCoor2(coor [2]int, maxX, maxY int) bool {
    return (coor[0] < maxX) && (coor[1] < maxY) && 
            (coor[0] >= 0) && (coor[1] >= 0)
}

func parseInput(antennaStrMap []string) map[byte][][2]int {
    var antennaMap map[byte][][2]int = make(map[byte][][2]int)

    var exists bool = false
    for y, line := range(antennaStrMap) {
        for x, antenna := range(line) {
            if (antenna != '.') {
                if _, exists = antennaMap[byte(antenna)]; exists {
                    antennaMap[byte(antenna)] = append(antennaMap[byte(antenna)], [2]int{x, y})
                } else {
                    antennaMap[byte(antenna)] = [][2]int{{x, y}}
                }
            }
        }
    }

    return antennaMap
}

func readInput() ([]string, error) {
    var antennaMap []string
    var err error
    var scanner bufio.Scanner = *bufio.NewScanner(os.Stdin)

    for {
        scanner.Scan()
        var line string = scanner.Text()

        if (len(line) == 0) {
            break
        }
    
        antennaMap = append(antennaMap, line)
    }

    err = scanner.Err()
    if (err != nil) {
        fmt.Println("Error occured when trying to read user input!")
        return antennaMap, err
    }

    return antennaMap, nil
}
