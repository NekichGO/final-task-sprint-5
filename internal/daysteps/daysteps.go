package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps описывает дневную прогулку.
// Также содержит в себе параметры пользователя.
type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку и записывает данные в соответствующие поля структуры DaySteps.
// Возвращает ошибку в случае возникновения проблем с парсингом строки.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	sliceStr := strings.Split(datastring, ",")

	if len(sliceStr) != 2 {
		return fmt.Errorf("invalid length of slice, got %d, expected 2", len(sliceStr))
	}

	steps, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return fmt.Errorf("error conversion steps string into int: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive and greater than 0, got %v", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(sliceStr[1])
	if err != nil {
		return fmt.Errorf("error conversion duration string into duration: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive and greater than 0, got %v", duration)
	}
	ds.Duration = duration
	return nil
}

// ActionInfo формирует строку с данными о прогулке.
// Возвращает строку с информацией о прогулке, а также, ошибку,
// при ее возникновении.
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	burnedCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("error calculating burned calories: %w", err)
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, burnedCalories), nil
}
