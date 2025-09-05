package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) == 2 {
		steps, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, err
		}
		if steps <= 0 {
			return 0, 0, errors.New("cannot be equal to 0")
		}
		duration, err := time.ParseDuration(parts[1])
		if err != nil {
			return 0, 0, err
		}
		if duration <= 0 {
			return 0, 0, errors.New("cannot be equal to 0")
		}
		return steps, duration, nil
	} else {
		return 0, 0, errors.New("parts len is less then 2")
	}
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
	}
	if steps <= 0 {
		return ""
	}
	distation := float64(steps) * stepLength / mInKm

	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distation, calories)
}
