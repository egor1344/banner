package ucb1

import (
	"log"
	"testing"
)

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
		result      int64
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
	for _, c := range cases {
		IDObject, err := c.objectsList.GetRelevantObject()
		if err != nil {
			t.Error("Error.", err)
		}
		if IDObject != c.result {
			t.Error("Error. usb_result = ", IDObject, ". Result must be ", c.result)
		}
	}
}

// TestShowAll - Перебор всех
func TestShowAll(t *testing.T) {
	bannerList := &ListBannerStatistic{
		[]*BannerStatistic{
			{1, 1, 1},
			{2, 1, 1},
			{3, 1, 1},
			{4, 1, 1},
			{5, 1, 1},
		}, 5,
	}
	mapID := make(map[int64]int)
	for i := 0; i <= 100; i++ {
		IDObject, err := bannerList.GetRelevantObject()
		if err != nil {
			t.Error("Error.", err)
		}
		bannerList.Objects[IDObject-1].CountDisplay++
		bannerList.AllCountDisplay++
		//log.Println(bannerList.Objects[0], bannerList.Objects[1], bannerList.Objects[2], bannerList.Objects[3], bannerList.Objects[4])
		//log.Println(bannerList.AllCountDisplay)
		mapID[IDObject]++

	}
	for _, value := range mapID {
		if value <= 1 {
			t.Error("banner not show")
		}
	}
}

// TestShowPopular - Перебор всех
func TestShowPopular(t *testing.T) {
	bannerList := &ListBannerStatistic{
		[]*BannerStatistic{
			{1, 14, 15},
			{2, 18, 20},
			{3, 22, 25},
			{4, 23, 30},
			{5, 26, 35},
		}, 125,
	}
	mapID := make(map[int64]int)
	for i := 0; i <= 100; i++ {
		IDObject, err := bannerList.GetRelevantObject()
		if err != nil {
			t.Error("Error.", err)
		}
		bannerList.Objects[IDObject-1].CountDisplay++
		bannerList.AllCountDisplay++
		//log.Println(bannerList.Objects[0], bannerList.Objects[1], bannerList.Objects[2], bannerList.Objects[3], bannerList.Objects[4])
		//log.Println(bannerList.AllCountDisplay)
		mapID[IDObject]++
	}
	log.Println(mapID)
	var maxValue int
	var maxID int64
	for id, value := range mapID {
		if maxValue <= value {
			maxValue = value
			maxID = id
		}
	}
	if maxID != 1 && maxValue != 30 {
		t.Error("Wrong id or value show popular banner")
	}
}
