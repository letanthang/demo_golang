package fix

import (
	"testing"

	"go.uber.org/goleak"
)

// implement TestMain for every package that needed test leak
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSearch(t *testing.T) {
	actual := search()
	if len(actual) != 2 {
		t.Errorf("actual: %v, wanted: %v", len(actual), 2)
	}
}
