package ucb1

import (
	"errors"
	"math"
)

// BannerStatistic -Объект для расчетов
type BannerStatistic struct {
	ID, CountClick, CountDisplay int64
}

// ListBannerStatistic - Список объектов для расчетов
type ListBannerStatistic struct {
	Objects         []*BannerStatistic
	AllCountDisplay int64
}

// UCB - Алгоритм расчета для показа баннера
func UCB(object *BannerStatistic, allCountDisplay int64) (result float64, err error) {
	/*
		object           - статистика баннера
		allCountDisplay  - сумма показов всех баннеров
	*/
	if object.CountClick <= 0 || object.CountDisplay <= 0 || allCountDisplay <= 0 {
		return 0, errors.New("variables equals 0")
	}
	if allCountDisplay < object.CountDisplay || object.CountDisplay < object.CountClick {
		return 0, errors.New("error logic")
	}
	bannerActual := float64(object.CountClick / object.CountDisplay) // Актуальность баннера
	up := float64(2) * math.Logb(float64(allCountDisplay))           // Врехняя часть уравнения
	result = bannerActual + (math.Sqrt(up) / float64(object.CountDisplay))
	return
}

// GetRelevantObject - Получение одного релевантного объекта из списка
func (lo *ListBannerStatistic) GetRelevantObject() (id int64, err error) {
	if len(lo.Objects) == 0 {
		return 0, errors.New("blank object list")
	}
	if len(lo.Objects) == 1 {
		return lo.Objects[0].ID, nil
	}
	var maxResult float64
	for _, obj := range lo.Objects {
		result, _ := UCB(obj, lo.AllCountDisplay)
		//log.Println(i, result, maxResult, id, obj.ID)
		if result >= maxResult {
			maxResult = result
			id = obj.ID
		}
	}
	return
}
