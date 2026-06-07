package actioninfo

import "fmt"

// DataParser описывает типы, поддерживающие парсинг строки с данными о тренировке
// и вывод информации о ней.
type DataParser interface {
	// TODO: добавить методы

	Parse(string) error
	ActionInfo() (string, error)
}

// Info принимает слайс строк с данными о тренировке или прогулке.
// Формирует и выводи строку с информацией об активности.
func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию

	for _, info := range dataset {
		err := dp.Parse(info)
		if err != nil {
			fmt.Printf("error parsing dataset %s: %v\n", info, err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("string formation error: %s\n ", err)
			continue
		}
		fmt.Println(result)
	}
}
