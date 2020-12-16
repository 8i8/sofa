package sofa

import "testing"

//
//  - - - - - - - - - - -
//   T e s t E c e q 0 6
//  - - - - - - - - - - -
//
//  Test Eceq06 function.
//
//  Called:  Eceq06, vvd
//
//  This revision:  2016 March 12
//
func TestEceq06(t *testing.T) {
	const fname = "Eceq06"
	var date1, date2, dl, db, dr, dd float64

	date1 = 2456165.5
	date2 = 0.401182685
	dl = 5.1
	db = -0.9

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoEceq06},
		{"go", GoEceq06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		dr, dd = test.fn(date1, date2, dl, db)

		vvd(t, dr, 5.533459733613627767, 1e-14, tname, "dr")
		vvd(t, dd, -1.246542932554480576, 1e-14, tname, "dd")
	}
}

func BenchmarkEceq06(b *testing.B) {
	var date1, date2, dl, db float64

	date1 = 2456165.5
	date2 = 0.401182685
	dl = 5.1
	db = -0.9

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64) (c1, c2 float64)
	}{
		{"cgo", CgoEceq06},
		{"go", GoEceq06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = test.fn(date1, date2, dl, db)
			}
		})
	}
}
