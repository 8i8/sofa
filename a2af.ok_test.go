package sofa

import (
	"log"
	"testing"
)

func TestA2af(t *testing.T) {
	const fname = "A2af"
	sign, idmsf := A2af(4, 2.345)
	if sign != '+' {
		t.Errorf("%s failed: want \"+\" got %q", fname, sign)
	} else if verbose {
		log.Printf("%s passed: want \"+\" got %q", fname, sign)
	}
	viv(t, idmsf[0], 134, fname, "0")
	viv(t, idmsf[1], 21, fname, "1")
	viv(t, idmsf[2], 30, fname, "2")
	viv(t, idmsf[3], 9706, fname, "3")
}
