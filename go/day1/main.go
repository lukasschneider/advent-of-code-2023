package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
    "time"
)

func findFirstAndLastNumber(s string) int {
    var loopUpTable = map[string]string {
        "one": "one1e",
        "two": "two2o",
        "three": "three3e",
        "four": "four4r",
        "five": "five5e",
        "six": "six6x",
        "seven": "seven7n",
        "eight": "eight8t",
        "nine": "nine9e",
    }
    for key, value := range loopUpTable {
        if strings.Contains(s, key) {
            s = strings.Replace(s, key, value, -1)
        }
    }
    number := findFirstAndLastDigit(s)

    return number

}

func findFirstAndLastDigit(s string) int {
    if len(s) == 0 {
        return 0
    }
    re := regexp.MustCompile("[1-9]")
    arr := re.FindAllString(s, -1)
    first := arr[0]
    last := arr[len(arr)-1]
    concat := first + last
    number, err := strconv.ParseInt(concat, 10, 32)
    if err != nil {
        log.Fatal(err)
    }
    return int(number)
}

func part1(data []string) int {
    start := time.Now()
    value := 0
    for _, line := range data {
        value += findFirstAndLastDigit(line)
    }
    elapsed := time.Since(start)
    fmt.Println("Part 1 took: ", elapsed)
    return value
}   

func part2(data []string) int {
    start := time.Now()
    value := 0
    for _, line := range data {
        value += findFirstAndLastNumber(line)
    }
    elapsed := time.Since(start)
    fmt.Println("Part 2 took: ", elapsed)
    return value
}

func main() {
    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var data []string 

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        data = append(data, line)
    }

    fmt.Println(part1(data))
    fmt.Println(part2(data))

}
