package tool

import (
	"bufio"
	"lvciot/go-conc/internal/model"
	"lvciot/go-conc/shared/concurrency"
	"os"
	"strings"
)

func Parser(sf string, df string, caw *concurrency.CountAndWait) {
	//numCPU := runtime.NumCPU()
	//sem := semaphore.NewWeighted(int64(numCPU))
	//ctx := context.Background()

	mutexAggregates := NewMutexAggregates()

	srcFile, _ := os.Open(sf)
	defer srcFile.Close()
	srcScanner := bufio.NewScanner(srcFile)

	dstFile, _ := os.Create(df)
	defer dstFile.Close()
	dstWriter := bufio.NewWriter(dstFile)

	for srcScanner.Scan() {
		//_ = sem.Acquire(ctx, 1)
		caw.RoutineStart()

		go func(row string) {
			d := model.NewDetectionFromRow(row)

			mutexAggregates.AddDetection(d)

			//sem.Release(1)
			caw.RoutineDone()
		}(srcScanner.Text())
	}

	aggregateRows := mutexAggregates.SortedRows()

	_, _ = dstWriter.WriteString("{")
	_, _ = dstWriter.WriteString(strings.Join(aggregateRows, ", "))
	_, _ = dstWriter.WriteString("}\n")
	_ = dstWriter.Flush()

	return
}
