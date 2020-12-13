package en

import "testing"

func TestErrNum(t *testing.T) {
	const fname = "ErrNum"
	var terr = New(6, "test", []string{
		"err -6",
		"err -5",
		"err -4",
		"err -3",
		"err -2",
		"err -1",
		"err 0",
		"err 1",
		"err 2",
		"err 3",
	})

	terr = terr.Set(-6)
	want := "test error: err -6"
	if terr.Error() != want {
		t.Errorf("%s: want %q got %q", fname, want, terr.Error())
	}
	terr = terr.Set(0)
	want = "test please contact package administration: err 0"
	if terr.Error() != want {
		t.Errorf("%s: want %q got %q", fname, want, terr.Error())
	}
	terr = terr.Set(3)
	want = "test warning: err 3"
	if terr.Error() != want {
		t.Errorf("%s: want %q got %q", fname, want, terr.Error())
	}
}

