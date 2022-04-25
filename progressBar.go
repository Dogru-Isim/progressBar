package main

import(
        "fmt"
        "strings"
)

func progressBar(done, all float64) {

        barSize := 50

        var percent float64

        char := "█"

        percent = float64(barSize)*(float64(done) / all)

        bar := strings.Repeat(char, int(percent))

        gap := strings.Repeat(" ", (barSize-int(percent)) )

        if done == all {
                gap = ""
        }

        fmt.Printf("\r|%v%v|", bar, gap)
}

func main() {

                for i := 0; i <=1000000; i++ {
                        progressBar(float64(i), 1000000.)
                }

                fmt.Println() // jump to the next line
}
