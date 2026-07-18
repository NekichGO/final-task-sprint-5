package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories вычисляет калории при ходьбе.
// Принимает количество шагов, вес в кг. и рост в метрах пользователя,
// а также продолжительность ходьбы.
// Возвращает потраченные калории при ходьбе, а также ошибку, при некорректных выходных данных.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive and greater then 0, got %v", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive and greater then 0, got %v", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be positive and greater then 0, got %v", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive and greater then 0, got %v", duration)
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	calories := (weight * avgSpeed * duration.Minutes()) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories вычисляет калории при беге.
// Принимает количество шагов, рост в метрах и вес в кг. пользователя,
// а также продолжительность бега.
// Возвращает потраченные калории при беге, а также ошибку, при некорректных входных данных.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 {
		return 0, fmt.Errorf("steps must be positive and greater than 0, got %v", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("weight must be positive and greater then 0, got %v", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("height must be positive and greater then 0, got %v", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("duration must be positive and greater than 0, got %v", duration)
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	return (weight * avgSpeed * duration.Minutes()) / minInH, nil
}

// MeanSpeed вычисляет средней скорости при беге или ходьбе.
// Возвращает среднюю скорость при беге или ходьбе.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 {
		return 0
	}

	if steps <= 0 {
		return 0
	}

	if height <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	avgSpeed := distance / duration.Hours()
	return avgSpeed
}

// Distance вычисляет дистанцию тренировки.
// Принимает количество шагов и рост пользователя в метрах.
// Возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	if height <= 0 {
		return 0
	}

	if steps <= 0 {
		return 0
	}

	stepLength := stepLengthCoefficient * height
	distance := (float64(steps) * stepLength) / mInKm
	return distance
}
