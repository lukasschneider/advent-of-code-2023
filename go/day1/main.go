package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "sort"
    // "strings"
)

var intToWord = map[int]string{
    0: "zero",
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
}

func findFirstAndLastNumber(s string) int {
    // sort map by key
    keys := make([]int, 0, len(intToWord))
    for k := range intToWord {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    // replace words with numbers
    for _, k := range keys {
        s = regexp.MustCompile(intToWord[k]).ReplaceAllString(s, strconv.Itoa(k))
    }
    re := regexp.MustCompile("[0-9]")
    arr := re.FindAllString(s, -1)
    if len(arr) == 0 {
        return 0
    }
    if len(arr) == 1 {
        number, err := strconv.ParseInt(arr[0], 10, 32)
        if err != nil {
            log.Fatal(err)
        }
        return int(number)
    }
    first := arr[0]
    last := arr[len(arr)-1]
    concat := first + last
    number, err := strconv.ParseInt(concat, 10, 32)
    if err != nil {
        log.Fatal(err)
    }
    return int(number)
}

func findFirstAndLastDigit(s string) int {
    re := regexp.MustCompile("[0-9]")
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
    value := 0
    for _, line := range data {
        value += findFirstAndLastDigit(line)
    }
    return value
}   

func part2(data []string) int {
    value := 0
    for _, line := range data {
        value += findFirstAndLastNumber(line)
    }
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
