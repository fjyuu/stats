package lib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		inputs []float64
		result *Result
	}{
		{
			[]float64{1, 2, 3},
			&Result{
				Count: 3,
				Mean:  2,
				Std:   1,
				Min:   1,
				Max:   3,
				Sum:   6,
			},
		},
		{
			[]float64{-2, 0, 2},
			&Result{
				Count: 3,
				Mean:  0,
				Std:   2,
				Min:   -2,
				Max:   2,
				Sum:   0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("inputs: %v", tc.inputs), func(t *testing.T) {
			calc := NewCalculator()
			for _, v := range tc.inputs {
				calc.Input(v)
			}

			r, err := calc.GetResult()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(r, tc.result) {
				t.Fatalf("got: %#v, want: %#v", r, tc.result)
			}
		})
	}
}
