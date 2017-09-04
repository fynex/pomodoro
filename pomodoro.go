package main

import (
    "time"
    "os/exec"
    "flag"
    "fmt"
)


var HEADLINE = "Pomodoro"

var w = flag.Int("w", 25, "Duration of a working session")
var s = flag.Int("s", 5, "Duration of a short break")
var l = flag.Int("l", 15, "Duration of a long break")
var p = flag.String("p", "wswswl", "Pattern to  follow (for example wswswl)")


type Pomodoro struct { 
    curr_state string
}


func (p *Pomodoro) Notify(summary string, body string, isUrgent bool) {
    if isUrgent {
        exec.Command("notify-send", summary, body, "-u", "critical").Run()
    } else {
        exec.Command("notify-send", summary, body).Run()
    }

}


func (p *Pomodoro) Work(duration *int, num int) {
    p.Notify(HEADLINE, "Start Working!", false)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
    fmt.Printf("    ☑ %d. Working phase\n", num)
}


func (p *Pomodoro) ShortBreak(duration *int) {
    p.Notify(HEADLINE, "Take a BREAK", true)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
}


func (p *Pomodoro) LongBreak (duration *int) {
    p.Notify(HEADLINE, "Take a long BREAK", true)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
}

func (pom *Pomodoro) PrintHeader() {
    fmt.Println(".===========================.")
    fmt.Println("|          Pomodoro         |")
    fmt.Println("'==========================='")
    fmt.Printf("  * Working    : %d minutes\n", *w)
    fmt.Printf("  * Break      : %d minutes\n", *s)
    fmt.Printf("  * Long Break : %d minutes\n", *l)
    fmt.Printf("  * Working Set: %s \n", *p)
    fmt.Println("\n  [*] Status ")
}

func (pom *Pomodoro) Run() {
    i := 1

    for _,state := range *p {
        pom.curr_state = string(state)

        switch pom.curr_state {
        case "w":
            pom.Work(w, i)
            i += 1
        case "s":
            pom.ShortBreak(s)
        case "l":
            pom.LongBreak(l)
        }

    }
}

func main() {
    flag.Parse()

    p := &Pomodoro{}

    p.PrintHeader()
    p.Run()

    p.Notify(HEADLINE, "Great job!", false)
}
