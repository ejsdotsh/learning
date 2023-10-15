package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"

	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name   string
		col    int
		op     string
		exp    string
		files  []string
		expErr error
	}{
		{name: "RunAvg1File", col: 3, op: "avg", exp: "227.6\n",
			files:  []string{"./testdata/example.csv"},
			expErr: nil,
		},
		{name: "RunAvgMultiFile", col: 3, op: "avg", exp: "233.84\n",
			files:  []string{"./testdata/example.csv", "./testdata/example2.csv"},
			expErr: nil,
		},
		{name: "RunFailRead", col: 2, op: "avg", exp: "",
			files:  []string{"./testdata/example.csv", "./testdata/fakefile.csv"},
			expErr: os.ErrNotExist,
		},
		{name: "RunFailColumn", col: 0, op: "avg", exp: "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvalidColumn,
		},
		{name: "RunFailNoFiles", col: 2, op: "avg", exp: "",
			files:  []string{},
			expErr: ErrNoFiles,
		},
		{name: "RunFailOperation", col: 2, op: "invalid", exp: "",
			files:  []string{"./testdata/example.csv"},
			expErr: ErrInvalidOperation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var res bytes.Buffer
			err := run(tc.files, tc.op, tc.col, &res)

			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error, got nil instead")
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("expected error: %q, got %q instead", tc.expErr, err)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}
			if res.String() != tc.exp {
				t.Errorf("expected: %q, got %q instead", tc.exp, &res)
			}
		})
	}
}

func BenchmarkRun(b *testing.B) {
	filenames, err := filepath.Glob("./testdata/benchmark/*.csv")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	// benchmark -op avg
	for i := 0; i < b.N; i++ {
		if err := run(filenames, "avg", 2, io.Discard); err != nil {
			b.Error(err)
		}
	}
	/*
		b.ResetTimer()
		// benchmark -op sum
		for i := 0; i < b.N; i++ {
			if err := run(filenames, "sum", 2, io.Discard); err != nil {
				b.Error(err)
			}
		}

		b.ResetTimer()
		// benchmark -op min
		for i := 0; i < b.N; i++ {
			if err := run(filenames, "min", 2, io.Discard); err != nil {
				b.Error(err)
			}
		}

		b.ResetTimer()
		// benchmark -op max
		for i := 0; i < b.N; i++ {
			if err := run(filenames, "max", 2, io.Discard); err != nil {
				b.Error(err)
			}
		}
	*/
}
