package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strings"
  "time"
  "regexp"
)

const (
  star = "Pythia"
  venue = "Delphi"
  prompt = "> "
)

func main() {
  fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
  fmt.Println("Your questions will be answered in due time.")

  questions := Oracle()
  reader := bufio.NewReader(os.Stdin)
  for {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
	  continue
	}
	fmt.Printf("%s heard: %s\n", star, line)
	questions <- line
  }
}

func Oracle() chan<- string {
  questions := make(chan string)
  answers := make(chan string)

  // Listen routine
  go func() {
    for q := range questions {
      question := strings.TrimRight(q, "\n.!")
      question := strings.ToLower(q)
      go prophecy(q, answers)
    }
  }()

  // Speak routine
  go func() {
    for a := range answers {
      fmt.Printf("%s: %s\n> ", star, a)
    }
  }()

  // Shitpost routine
  go func(answer chan<- string) {
    shitposts := make([]string, 0)
    shitposts = append(shitposts,
      "Russia government go reeeee",
      "Xin Winn Poong",
      "monke",
      "pappata pizza",
      "oompa loompa",
    )
    rand.Seed(time.Now().Unix())

    for {
      time.Sleep(time.Duration(10 + rand.Intn(20)) * time.Second)
      answer <- shitposts[rand.Intn(len(shitposts))]
    }
  }(answers)

  return questions
}

var responses = map[string][]string {
  `(i need|i want) (.*)`: {
    "why tho?",
    "trust me, you dont want %s.",
  },
  `i am (.*)`: {
    "cool, but i didnt ask.",
    "hello %s, i'm dad.",
  },
  // Can add without rewriting code
}

func prophecy(question string, answer chan<- string) {
  time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
  longestWord := ""
  words := strings.Fields(question) // Fields extracts the words into a slice.
  for _, w := range words {
	if len(w) > len(longestWord) {
      longestWord = w
	}
  }

  for pattern, response := range responses {
    rgx := regexp.MustCompile(pattern)
    match := rgx.FindStringSubmatch(question)
    if len(match) > 0 {
      rand.Seed(time.Now().Unix())
      answer <- response[rand.Intn(len(response))]
      return
    }
  }

  nonsense := []string{
	"The moon is dark.",
	"The sun is bright.",
  }
  answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() {
  rand.Seed(time.Now().Unix())
}
