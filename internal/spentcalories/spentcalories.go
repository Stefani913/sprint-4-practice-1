package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
	walking                    = "Ходьба"
	running                    = "Бег"
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")

	if len(parts) == 3 {
		steps, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, parts[1], 0, err
		}
		duration, err := time.ParseDuration(parts[2])
		if err != nil {
			return 0, parts[1], 0, err
		}
		return steps, parts[1], duration, nil
	}
	return 0, "", 0, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return height * stepLengthCoefficient * float64(steps) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)

	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}

	switch activity {
	case walking:
		{
			calories, _ := WalkingSpentCalories(steps, weight, height, duration)
			distance := distance(steps, height)
			meanSpeed := meanSpeed(steps, height, duration)

			return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
				activity, duration.Hours(), distance, meanSpeed, calories), nil

		}
	case running:
		{
			calories, _ := RunningSpentCalories(steps, weight, height, duration)
			distance := distance(steps, height)
			meanSpeed := meanSpeed(steps, height, duration)

			return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
				activity, duration.Hours(), distance, meanSpeed, calories), nil

		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("cannot be equal to 0")
	}
	average := meanSpeed(steps, height, duration)

	return weight * average * duration.Minutes() / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("cannot be equal to 0")
	}
	average := meanSpeed(steps, height, duration)

	return weight * average * duration.Minutes() / minInH * walkingCaloriesCoefficient, nil
}
