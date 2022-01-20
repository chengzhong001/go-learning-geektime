package condition_test

import "testing"

func someFunc() (int, error) {
	return 1, nil
}

func TestIfMultiSec(t *testing.T) {
	if v, err := someFunc(); err == nil {
		t.Log("no error", v)

	} else {
		t.Log("error")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("it is not 0-3", i)
		}

	}
}
