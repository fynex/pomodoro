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


func Notify(summary string, body string, isUrgent bool) {
    if isUrgent {
        exec.Command("notify-send", summary, body, "-u", "critical").Run()
    } else {
        exec.Command("notify-send", summary, body).Run()
    }

}


func Work(duration *int) {
    Notify(HEADLINE, "Start Working!", false)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
}


func ShortBreak(duration *int) {
    Notify(HEADLINE, "Take a BREAK", true)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
}


func LongBreak (duration *int) {
    Notify(HEADLINE, "Take a long BREAK", true)

    timer1 := time.NewTimer(time.Minute * time.Duration(*duration))

    <-timer1.C
}


func main() {
    flag.Parse()

    fmt.Println(".===========================.")
    fmt.Println("|          Pomodoro         |")
    fmt.Println("'==========================='")
    fmt.Printf("  * Working    : %d minutes\n", *w)
    fmt.Printf("  * Break      : %d minutes\n", *s)
    fmt.Printf("  * Long Break : %d minutes\n", *l)
    fmt.Printf("  * Working Set: %s \n", *p)

    for _,state := range *p {
        curr_state := string(state)

        switch curr_state {
        case "w":
            Work(w)
        case "s":
            ShortBreak(s)
        case "l":
            LongBreak(l)
        }
    }

    Notify(HEADLINE, "DONE", false)
}
