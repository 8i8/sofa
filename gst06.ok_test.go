package sofa

import "testing"

//
//  - - - - - - - - - -
//   T e s t G s t 0 6
//  - - - - - - - - - -
//
//  Test Gst06 function.
//
//  Called:  Gst06, vvd
//
//  This revision:  2013 August 7
//
func TestGst06(t *testing.T) {
	const fname = "Gst06"
	var theta float64
	var rnpb [3][3]float64

	rnpb[0][0] = 0.9999989440476103608
	rnpb[0][1] = -0.1332881761240011518e-2
	rnpb[0][2] = -0.5790767434730085097e-3

	rnpb[1][0] = 0.1332858254308954453e-2
	rnpb[1][1] = 0.9999991109044505944
	rnpb[1][2] = -0.4097782710401555759e-4

	rnpb[2][0] = 0.5791308472168153320e-3
	rnpb[2][1] = 0.4020595661593994396e-4
	rnpb[2][2] = 0.9999998314954572365

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64,
			a5 [3][3]float64) float64
	}{
		{"cgo", CgoGst06},
		{"go", GoGst06},
	}

	for _, test := range tests {
		tname := fname + " " + test.ref

		theta = test.fn(
			2400000.5, 53736.0, 2400000.5, 53736.0, rnpb)

		vvd(t, theta, 1.754166138018167568, 1e-12, tname, "")
	}
}

func BenchmarkGst06(b *testing.B) {
	var rnpb [3][3]float64

	rnpb[0][0] = 0.9999989440476103608
	rnpb[0][1] = -0.1332881761240011518e-2
	rnpb[0][2] = -0.5790767434730085097e-3

	rnpb[1][0] = 0.1332858254308954453e-2
	rnpb[1][1] = 0.9999991109044505944
	rnpb[1][2] = -0.4097782710401555759e-4

	rnpb[2][0] = 0.5791308472168153320e-3
	rnpb[2][1] = 0.4020595661593994396e-4
	rnpb[2][2] = 0.9999998314954572365

	tests := []struct {
		ref string
		fn  func(a1, a2, a3, a4 float64,
			a5 [3][3]float64) float64
	}{
		{"cgo", CgoGst06},
		{"go", GoGst06},
	}

	for _, test := range tests {
		b.Run(test.ref, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = test.fn(2400000.5, 53736.0,
					2400000.5, 53736.0, rnpb)
			}
		})
	}
}
