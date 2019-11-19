package ucb1

import (
	"errors"
	"math"
)

// Объект для расчетов
type BannerStatistic struct {
	id, countClick, countDisplay int
}

// Список объектов для расчетов
type ListBannerStatistic struct {
	objects         []*BannerStatistic
	allCountDisplay int
}

// Алгоритм расчета для показа баннера
func UCB(object *BannerStatistic, allCountDisplay int) (result float64, err error) {
	/*
		countClick       - количество кликов по баннеру
		countDisplay     - количество показов баннера
	*/
	if object.countClick <= 0 || object.countDisplay <= 0 || allCountDisplay <= 0 {
		return 0, errors.New("variables equals 0")
	}
	if allCountDisplay < object.countDisplay || object.countDisplay < object.countClick {
		return 0, errors.New("error logic")
	}
	bannerActual := float64(object.countClick / object.countDisplay) // Актуальность баннера
	up := float64(2) * math.Logb(float64(allCountDisplay))           // Врехняя часть уравнения
	result = bannerActual + (math.Sqrt(up) / float64(object.countDisplay))
	return
}

// Получение одного релевантного объекта из списка
func (lo *ListBannerStatistic) GetRelevantObject() (id int, err error) {
	if len(lo.objects) == 0 {
		return 0, errors.New("blank object list")
	}
	if len(lo.objects) == 1 {
		return lo.objects[0].id, nil
	}
	var maxResult float64
	for _, obj := range lo.objects {
		result, _ := UCB(obj, lo.allCountDisplay)
		if result >= maxResult {
			maxResult = result
			id = obj.id
		}
	}
	return
}
