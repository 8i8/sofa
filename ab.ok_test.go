package sofa

import "testing"

//
//  - - - - - - -
//   T e s t A b
//  - - - - - - -
//
//  Test iauAb function.
//
//  Called:  Ab, vvd
//
//  This revision:  2013 October 1
//
func TestAb(t *testing.T) {
	const fname = "Ab"
	var pnat, v, ppr [3]float64
	var s, bm1 float64

	pnat[0] = -0.76321968546737951
	pnat[1] = -0.60869453983060384
	pnat[2] = -0.21676408580639883
	v[0] = 2.1044018893653786e-5
	v[1] = -8.9108923304429319e-5
	v[2] = -3.8633714797716569e-5
	s = 0.99980921395708788
	bm1 = 0.99999999506209258

	ppr = Ab(pnat, v, s, bm1)

	vvd(t, ppr[0], -0.7631631094219556269, 1e-12, fname, "1")
	vvd(t, ppr[1], -0.6087553082505590832, 1e-12, fname, "2")
	vvd(t, ppr[2], -0.2167926269368471279, 1e-12, fname, "3")
}

func TestGoAb(t *testing.T) {
	const fname = "Ab"
	var pnat, v, ppr [3]float64
	var s, bm1 float64

	pnat[0] = -0.76321968546737951
	pnat[1] = -0.60869453983060384
	pnat[2] = -0.21676408580639883
	v[0] = 2.1044018893653786e-5
	v[1] = -8.9108923304429319e-5
	v[2] = -3.8633714797716569e-5
	s = 0.99980921395708788
	bm1 = 0.99999999506209258

	ppr = goAb(pnat, v, s, bm1)

	vvd(t, ppr[0], -0.7631631094219556269, 1e-12, fname, "1")
	vvd(t, ppr[1], -0.6087553082505590832, 1e-12, fname, "2")
	vvd(t, ppr[2], -0.2167926269368471279, 1e-12, fname, "3")
}

func BenchmarkAb(b *testing.B) {
	var pnat, v [3]float64
	var s, bm1 float64

	pnat[0] = -0.76321968546737951
	pnat[1] = -0.60869453983060384
	pnat[2] = -0.21676408580639883
	v[0] = 2.1044018893653786e-5
	v[1] = -8.9108923304429319e-5
	v[2] = -3.8633714797716569e-5
	s = 0.99980921395708788
	bm1 = 0.99999999506209258

	for i := 0; i < b.N; i++ {
		_ = Ab(pnat, v, s, bm1)
	}
}

func BenchmarkGoAb(b *testing.B) {
	var pnat, v [3]float64
	var s, bm1 float64

	pnat[0] = -0.76321968546737951
	pnat[1] = -0.60869453983060384
	pnat[2] = -0.21676408580639883
	v[0] = 2.1044018893653786e-5
	v[1] = -8.9108923304429319e-5
	v[2] = -3.8633714797716569e-5
	s = 0.99980921395708788
	bm1 = 0.99999999506209258

	for i := 0; i < b.N; i++ {
		_ = goAb(pnat, v, s, bm1)
	}
}

func BenchmarkAbNoFuncCall(b *testing.B) {
	var pnat, v [3]float64
	var s, bm1 float64

	pnat[0] = -0.76321968546737951
	pnat[1] = -0.60869453983060384
	pnat[2] = -0.21676408580639883
	v[0] = 2.1044018893653786e-5
	v[1] = -8.9108923304429319e-5
	v[2] = -3.8633714797716569e-5
	s = 0.99980921395708788
	bm1 = 0.99999999506209258

	for i := 0; i < b.N; i++ {
		_ = AbNoFuncCall(pnat, v, s, bm1)
	}
}
