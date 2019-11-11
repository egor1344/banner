package ucb1

import "testing"

func TestUCB(t *testing.T) {
	cases := []struct {
		object *Object
		result float64
	}{
		{&Object{0, 0, 0, 0}, 0},
		{&Object{1, -1, 1, 2}, 0},
		{&Object{2, 1, 1, 1}, 1},
		{&Object{3, 2, 3, 4}, 0.6666666666666666},
	}
	for _, c := range cases {
		//t.Log(i, c)
		usbResult, _ := UCB(c.object)
		if usbResult != c.result {
			t.Error("Error. usb_result = ", usbResult, ". Result must be ", c.result)
		}
	}
}

func TestUCBList(t *testing.T) {
	cases := []struct {
		objectsList *ListObject
		result      int
	}{
		//{
		//	&ListObject{
		//		[]*Object{
		//			{1, 0, 0, 0},
		//			{2, 0, 0, 0},
		//			{3, -1, 1, 2},
		//			{4, 1, 1, 1},
		//			{5, 2, 3, 4},
		//		},
		//	}, 4},
		{
			&ListObject{
				[]*Object{
					{1, 2, 3, 4},
					{2, 3, 4, 5},
					{3, 5, 7, 8},
					{4, 2, 1, 0},
					{5, 5, 3, 2},
				},
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
