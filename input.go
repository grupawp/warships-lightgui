package board

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

const (
	moveUp    = "\033[1A"
	moveDown  = "\033[1B"
	eraseLine = "\033[K"

	EOT = 0x3
	CR  = 0xD
)

/*
ReadLineWithTimer reads a line from stdin and displays a timer.
The function blocks until [ENTER] is pressed and returns typed in string and true.
If [CTRL+C] is pressed or the timer expires, the function blocks until any key is
pressed and returns an empty string and false.
*/
func ReadLineWithTimer(prompt string, timer time.Duration) (out string, ok bool) {
	fmt.Print("\n", prompt)
	defer fmt.Println()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var finished bool
	end := time.Now().Add(timer)
	ctx, cancel := context.WithDeadline(context.Background(), end)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()
		for {
			if finished {
				return
			}
			select {
			case <-ctx.Done():
				fmt.Print("\r"+eraseLine+moveUp+"\r", eraseLine)
				fmt.Printf("Time's up! Press any key to continue...")
				return
			case <-ticker.C:
				fmt.Print(moveUp + "\r" + eraseLine)
				fmt.Printf("Time left: %.0fs", time.Until(end).Seconds())
				fmt.Print(moveDown + "\r" + eraseLine)
				fmt.Print(prompt + out)
			}
		}
	}()

	reader := bufio.NewReader(os.Stdin)
outer:
	for {
		buf := make([]byte, 3)
		n, err := reader.Read(buf)
		if err != nil {
			panic(err)
		} else if n == 1 && buf[0] == EOT { // CTRL+C
			break outer
		}
		select {
		case <-ctx.Done():
			out = ""
			break outer
		default:
			if n == 1 && buf[0] == CR { // ENTER
				finished = true
				break outer
			}
			out += string(buf[:n])
		}
	}

	wg.Wait()
	return out, finished
}
