package ucb1

import (
	"errors"
	"math"
	"sync"
)

// Объект для расчетов
type Object struct {
	id, countClick, countDisplay, allCountDisplay int
}

// Список объектов для расчетов
type ListObject struct {
	objects []*Object
}

// Алгоритм расчета для показа баннера
func UCB(object *Object) (result float64, err error) {
	/*
		count_click       - количество кликов по баннеру
		count_display     - количество показов баннера
		all_count_display - сумма показов всех баннеров
	*/
	if object.countClick <= 0 || object.countDisplay <= 0 || object.allCountDisplay <= 0 {
		return 0, errors.New("variables equals 0")
	}
	if object.allCountDisplay < object.countDisplay || object.countDisplay < object.countClick {
		return 0, errors.New("error logic")
	}
	bannerActual := float64(object.countClick / object.countDisplay) // Актуальность баннера
	up := float64(2) * math.Logb(float64(object.allCountDisplay))    // Врехняя часть уравнения
	result = bannerActual + (math.Sqrt(up) / float64(object.countDisplay))
	return
}

// Получение одного релевантного объекта из списка
func (lo *ListObject) GetRelevantObject() (id int, err error) {
	if len(lo.objects) == 0 {
		return 0, errors.New("blank object list")
	}
	if len(lo.objects) == 1 {
		return lo.objects[0].id, nil
	}
	var maxResult float64
	var m sync.Mutex
	var wg sync.WaitGroup
	for _, obj := range lo.objects {
		wg.Add(1)
		go func(c *Object, m *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			result, _ := UCB(c)
			if result >= maxResult {
				m.Lock()
				maxResult = result
				id = c.id
				m.Unlock()
			}
		}(obj, &m, &wg)
	}
	wg.Wait()
	return
}
