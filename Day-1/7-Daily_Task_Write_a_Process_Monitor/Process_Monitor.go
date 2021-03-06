package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type callBackChan chan struct {}

// Trigger the callBack channel every d duration units
func checkEvery(ctx context.Context, d time.Duration, cb callBackChan) {
	for {
		select {
		case <-ctx.Done():
			// ctx is canceled
			return
		case <-time.After(d):
			// wait for the duration
			if cb != nil {
				cb <-struct {}{}
			}
		}
	}
}

func PrintProcessList() {
	psCommand := exec.Command("ps", "a")
	resp, err := psCommand.CombinedOutput()
	if err != nil {
		log.Fatal("error: ps command failed")
	}

	out := string(resp)
	lines := strings.Split(out, "\n\n")
	fmt.Println("\nSystem Monitor Sequence Initiated, Updating Every 5 Seconds...\n--------------------------------------------------------------")

	for _, line := range lines {
		if line != "" {
			fmt.Println(line)
		}
	}
}

func main() {
	ctx := context.Background()
	PrintProcessList()

	callBack := make(callBackChan)
	go checkEvery(ctx, 5 * time.Second, callBack)
	go func() {
		for {
			select {
			case <-callBack:
				PrintProcessList()
			}
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}