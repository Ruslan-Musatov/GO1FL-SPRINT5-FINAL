package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчётов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчёта длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчёта калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateTrainingParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	durationInMinutes := duration.Minutes()
	calories := (weight * MeanSpeed(steps, height, duration) * durationInMinutes) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validateTrainingParams(steps, weight, height, duration); err != nil {
		return 0, err
	}

	durationInMinutes := duration.Minutes()
	calories := (weight * MeanSpeed(steps, height, duration) * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}

func validateTrainingParams(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errors.New("количество шагов должно быть больше нуля")
	case weight <= 0:
		return errors.New("вес должен быть больше нуля")
	case height <= 0:
		return errors.New("рост должен быть больше нуля")
	case duration <= 0:
		return errors.New("продолжительность должна быть больше нуля")
	default:
		return nil
	}
}
