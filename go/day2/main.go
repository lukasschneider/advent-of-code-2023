package main

import (
    "fmt"
    "log"
    "strconv"
    "strings"
    "bufio"
    "os"
    "time"
)

type Game struct {
    id int
    subsets []subset
}

type subset struct {
    blue int
    red int
    green int
}

func parse(line string) (Game) {
    var subsets []subset

    gameString := strings.Split(line, ":")[0]

    // extract game id
    id, err := strconv.Atoi(strings.Split(gameString, " ")[1])
    if err != nil {
        log.Fatal(err)
    }

    // extract subsets
    subsetString := strings.Split(strings.Split(line, ":")[1], ";")
    for _, s := range subsetString {
        cubes := strings.Split(s, ",")
        for _, cube := range cubes {
            number, err := strconv.Atoi(strings.Split(cube, " ")[1])
            if err != nil {
                log.Fatal(err)
            }
            color := strings.Split(cube, " ")[2]
            subset := subset{}
            switch color {
            case "blue":
                subset.blue = number
            case "red":
                subset.red = number
            case "green":
                subset.green = number
            default:
                log.Fatal("Invalid color")
            }
            subsets = append(subsets, subset)
        }

    }

    // create game
    game := Game{id, subsets}

    return game
}

func part1(games []Game) int{
    start := time.Now()
    sum := 0
    for _, game := range games {
        if (func (game Game) bool {
            for _, subset := range game.subsets {
                if subset.blue > 14 || subset.red > 12 || subset.green > 13 {
                    return false
                }
            }
            return true
        })(game) {
            sum += game.id
        }
    }
    elapsed := time.Since(start)
    fmt.Printf("Part 1 took %s\n", elapsed)
    return sum
}

func part2(games []Game) int{
    start := time.Now()
    sum := 0
    for _, game := range games {
        max_blue, max_red, max_green := func (game Game) (int, int, int) {
            max_blue := 0
            max_red := 0
            max_green := 0
            for _, subset := range game.subsets {
                if subset.blue > max_blue {
                    max_blue = subset.blue
                }
                if subset.red > max_red {
                    max_red = subset.red
                }
                if subset.green > max_green {
                    max_green = subset.green
                }
            }
            return max_blue, max_red, max_green
        }(game)
        sum += max_blue * max_red * max_green
    }
    elapsed := time.Since(start)
    fmt.Printf("Part 2 took %s\n", elapsed)
    return sum
}


func main() {
    games := []Game {}
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        games = append(games, parse(scanner.Text()))
    }
    fmt.Printf("Part 1: %d\n", part1(games))
    fmt.Printf("Part 2: %d\n", part2(games))
}

