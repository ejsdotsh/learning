package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing/iotest"

	"testing"
)

func TestOperations(t *testing.T) {
	data := [][]float64{
		{15, 20, 15, 33, 85, 10, 99, 30},
		{2.5, 8, 5.5, 7.75, 8.35, 3, 3.5, 11.15, 4.65, 1.5, 7.65, 13.25, 5.45},
		{-15, 0, -25},
		{12, 87, 54, 13, 69, 119},
	}

	testCases := []struct {
		name string
		op   statsFunc
		exp  []float64
	}{
		{"sum", sum, []float64{307, 82.25, -40, 354}},
		{"avg", avg, []float64{38.375, 6.326923076923077, -13.333333333333334, 59.0}},
		{"max", max, []float64{99, 13.25, 0, 119}},
		{"min", min, []float64{10, 1.5, -25, 12}},
	}

	for _, tc := range testCases {
		for k, exp := range tc.exp {
			name := fmt.Sprintf("%sdata%d", tc.name, k)
			t.Run(name, func(t *testing.T) {
				res := tc.op(data[k])

				if res != exp {
					t.Errorf("expected: %g, got %g instead", exp, res)
				}
			})
		}
	}
}

func TestCSV2Float(t *testing.T) {
	csvData := `IP address,requests,response time
192.168.0.199,2056,236
192.168.0.88,899,220
192.168.0.199,3054,226
192.168.0.100,4133,218
192.168.0.199,950,238
`
	testCases := []struct {
		name   string
		col    int
		exp    []float64
		expErr error
		r      io.Reader
	}{
		{name: "column2", col: 2,
			exp:    []float64{2056, 899, 3054, 4133, 950},
			expErr: nil,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "column3", col: 3,
			exp:    []float64{236, 220, 226, 218, 238},
			expErr: nil,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "FailRead", col: 1,
			exp:    nil,
			expErr: iotest.ErrTimeout,
			r:      iotest.TimeoutReader(bytes.NewReader([]byte{0})),
		},
		{name: "FailedNotNumber", col: 1,
			exp:    nil,
			expErr: ErrNotNumber,
			r:      bytes.NewBufferString(csvData),
		},
		{name: "FailedInvalidColumn", col: 4,
			exp:    nil,
			expErr: ErrInvalidColumn,
			r:      bytes.NewBufferString(csvData),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := csv2float(tc.r, tc.col)
			if tc.expErr != nil {
				if err == nil {
					t.Errorf("expected error, got nil instead")
				}
				if !errors.Is(err, tc.expErr) {
					t.Errorf("expected error %q, got %q instead", tc.expErr, err)
				}

				return
			}

			if err != nil {
				t.Errorf("unexpected error: %q", err)
			}

			for i, exp := range tc.exp {
				if res[i] != exp {
					t.Errorf("expected %g, got %g instead", exp, res[i])
				}
			}
		})
	}

}
