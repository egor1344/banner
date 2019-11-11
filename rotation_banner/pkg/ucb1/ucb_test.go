package ucb1

import "testing"

func TestUSB(t *testing.T) {
	cases := []struct {
		countClick, countDisplay, allCountDisplay int
		result                                    float64
	}{
		{0, 0, 0, 0},
		{-1, 1, 2, 0},
		{1, 1, 1, 1},
		{2, 3, 4, 1.1547005383792515},
	}
	for i, c := range cases {
		t.Log(i, c)
		usbResult, _ := UCB(c.countClick, c.countDisplay, c.allCountDisplay)
		if usbResult != c.result {
			t.Error("Error. usb_result = ", usbResult, ". Result must be ", c.result)
		}
	}
}
