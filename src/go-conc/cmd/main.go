package main

import (
	"github.com/schollz/progressbar/v3"
	"lvciot/go-conc/internal/tool"
	"lvciot/go-conc/shared/concurrency"
	"path/filepath"
	"runtime"
	"time"
)

const (
	MaxRows = 1_000_000_000
	SrcFile = "../../../../data/measurements.txt"
	DstFile = "../../measurements.out"
)

func main() {
	_, b, _, _ := runtime.Caller(0)
	srcFile := filepath.Join(b, SrcFile)
	dstFile := filepath.Join(b, DstFile)

	bar := progressbar.Default(MaxRows)
	ticker := time.NewTicker(time.Second)

	caw := concurrency.CountAndWait{}

	go func() {
		for {
			select {
			case <-ticker.C:
				_ = bar.Set(caw.Counter())
			}
		}
	}()

	tool.Parser(srcFile, dstFile, &caw)

	caw.WaitAllRoutinesDone()
	_ = bar.Set(MaxRows)
	println()
}
