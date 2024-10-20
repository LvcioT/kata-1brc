package test

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
)

type Case struct {
	Name         string
	SourceFile   string
	ExpectedFile string
}

func GetCases() ([]Case, error) {
	_, b, _, _ := runtime.Caller(0)
	pattern := filepath.Join(b, strings.Repeat("../", 4)+"data/testcases/*.txt")

	testFiles, _ := filepath.Glob(pattern)

	tc := make([]Case, len(testFiles))

	if len(testFiles) == 0 {
		return []Case{}, errors.New("no test cases found")
	}

	for i, tf := range testFiles {
		tc[i] = Case{
			Name:         filepath.Base(tf),
			SourceFile:   tf,
			ExpectedFile: strings.Replace(tf, ".txt", ".out", 1),
		}
	}

	return tc, nil
}
