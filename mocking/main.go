package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}


type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countDownStart = 3

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func CountDown(out io.Writer, sleeper Sleeper){
		for i := countDownStart; i > 0; i-- {
					sleeper.Sleep()
		}
		for i := countDownStart; i > 0; i-- {
				fmt.Fprintln(out, i)
			}
			fmt.Fprint(out, finalWord)
}


func main(){
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}

