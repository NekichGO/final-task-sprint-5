package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training описывает параметры тренировки,
// а также параметры пользователя.
type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse парсит строку и записывает данные в соответствующие поля структуры Training.
// Возвращает ошибку в случае возникновения проблем с парсингом строки.
func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	sliceStr := strings.Split(datastring, ",")

	if len(sliceStr) != 3 {
		return fmt.Errorf("invalid length of slice, got: %d, expected 3", len(sliceStr))
	}

	steps, err := strconv.Atoi(sliceStr[0])
	if err != nil {
		return fmt.Errorf("invalid conversion string to int %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("steps must be positive and greater then 0, got %d", steps)
	}
	t.Steps = steps

	t.TrainingType = sliceStr[1]

	duration, err := time.ParseDuration(sliceStr[2])
	if err != nil {
		return fmt.Errorf("invalid conversion string to duration %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive and greater then 0, got: %v", duration)
	}
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с данными о тренировке,
// исходя из того, какой тип тренировки был передан.
// Возвращает строку с данными о тренировке и ошибку, при ее возникновении.
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distance := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		runCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error calculating burned calories while running: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, avgSpeed, runCalories), nil

	case "Ходьба":
		walkCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("error calculating burned calories while walking: %w", err)
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, avgSpeed, walkCalories), nil

	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

}
