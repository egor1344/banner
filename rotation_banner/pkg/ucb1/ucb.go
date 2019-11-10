package ucb1

import (
	"errors"
	log "github.com/egor1344/banner/rotation_banner/pkg/logger"
	"math"
)

// Расчет показа алгоритма
func UCB(countClick, countDisplay, allCountDisplay int) (result float64, err error) {
	/*
		count_click       - количество кликов по баннеру
		count_display     - количество показов баннера
		all_count_display - сумма показов всех баннеров
	*/
	log.Logger.Info(countClick, countDisplay, allCountDisplay)
	if countClick == 0 || countDisplay == 0 || allCountDisplay == 0 {
		return 0, errors.New("variables equals 0")
	}
	bannerActual := float64(countClick / countDisplay)
	result = bannerActual + math.Sqrt((float64(2)*math.Logb(float64(allCountDisplay)))/float64(countDisplay))
	return
}
