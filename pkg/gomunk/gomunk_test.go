package gomunk

import "testing"

func TestFailedNewHandler(t *testing.T) {
	provider := "thisShouldFail"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(`GoMunk(%q) did not panic`, provider)
		}
	}()
	GoMunk()
}

func TestNewHandler(t *testing.T) {
	provider := "aws"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf(`GoMunk(%q) failed to panic`, provider)
		}
	}()
	GoMunk()
}
