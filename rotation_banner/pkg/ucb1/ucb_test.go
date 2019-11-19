package ucb1

import "testing"

func TestUCB(t *testing.T) {
	cases := []struct {
		object *BannerStatistic
		result float64
	}{
		{&BannerStatistic{0, 0, 0}, 0},
		{&BannerStatistic{1, -1, 2}, 0},
		{&BannerStatistic{2, 1, 1}, 3},
		{&BannerStatistic{3, 2, 4}, 0.5},
	}
	for _, c := range cases {
		usbResult, _ := UCB(c.object, 7)
		if usbResult != c.result {
			t.Error("Error. usb_result = ", usbResult, ". Result must be ", c.result)
		}
	}
}

func TestUCBList(t *testing.T) {
	cases := []struct {
		objectsList *ListBannerStatistic
		result      int
	}{
		{
			&ListBannerStatistic{
				[]*BannerStatistic{
					{1, 0, 0},
					{2, 0, 0},
					{3, -1, 1},
					{4, 1, 1},
					{5, 2, 3},
				}, 7,
			}, 4},
		{
			&ListBannerStatistic{
				[]*BannerStatistic{
					{1, 2, 3},
					{2, 3, 4},
					{3, 5, 7},
					{4, 2, 1},
					{5, 5, 3},
				}, 18,
			}, 1},
	}
	for i, c := range cases {
		t.Log(i, c)
		IDObject, err := c.objectsList.GetRelevantObject()
		if err != nil {
			t.Error("Error.", err)
		}
		if IDObject != c.result {
			t.Error("Error. usb_result = ", IDObject, ". Result must be ", c.result)
		}
	}
}
