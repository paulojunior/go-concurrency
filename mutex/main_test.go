package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name   string
		result string
	}{
		{
			name:   "shouldReturnFullAmountWithoutRaceCondional",
			result: "$29640.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdOut := os.Stdout
			r, w, _ := os.Pipe()

			os.Stdout = w

			main()

			_ = w.Close()

			result, _ := io.ReadAll(r)
			output := string(result)

			os.Stdout = stdOut

			if !strings.Contains(output, tt.result) {
				t.Errorf("Wrong balance returned: %s", output)
			}
		})
	}
}
