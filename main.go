package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
)

var badContent = make(map[string]struct{})

func main() {
  var path string
  fmt.Scan(&path)
  readContent(path)
  for {
    censorship(getUserSentence())
  }
}

func readContent(path string) {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  wordScanner := bufio.NewScanner(file)
  wordScanner.Split(bufio.ScanWords)

  for wordScanner.Scan() {
    badContent[strings.ToLower(wordScanner.Text())] = struct{}{}
  }
  if err := wordScanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func getUserSentence() []string {
  var result []string
  wordScanner := bufio.NewScanner(os.Stdin)
  wordScanner.Scan()
  result = strings.Split(wordScanner.Text(), " ")
  return result
}

func censoredOutput(word []string) {
  fmt.Println(wordProcessing(word))
}

func wordProcessing(word []string) string {
  var b strings.Builder
  for index, word := range word {
    if isBadWord(word) {
      b.WriteString(strings.Repeat("*", len(word)))
    } else {
      b.WriteString(word)
    }
    if index != len(word)-1 {
      b.WriteString(" ")
    }
  }
  return b.String()
}

func isBadWord(word string) bool {
  _, ok := badContent[strings.ToLower(word)]
  return ok
}

func censorship(word []string) {
  if len(word) == 1 && word[0] == "exit" {
    fmt.Println("Bye!")
    os.Exit(0)
  }
  censoredOutput(word)
}
