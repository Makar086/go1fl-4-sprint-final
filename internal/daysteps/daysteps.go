package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"errors"
	
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice)==2{
	step, err := strconv.Atoi(slice[0])
	if err != nil {
			fmt.Println("Ошибка при конвертации строки в число:", err)
			return 0, 0, err
			}
	if step<=0{
		return 0, 0, err
		}
	time, err := time.ParseDuration(slice[1])
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
			return 0, 0, err
		}
	if time<=0{
		return 0, 0, err
		}
	return step, time, nil
	}
	return 0, 0, errors.New("ошибка")
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	step, time, err := parsePackage(data)
	if err != nil {
		// Выводим ошибку на экран
		return ""
	}
	if step<=0{
	return ""
	}
	dist:=float64(step)*stepLength 
	dist1:=float64(dist/mInKm)
	calorie, err:=spentcalories.WalkingSpentCalories(step, weight, height, time)
	step1:=strconv.Itoa(step)
	dist2:=strconv.FormatFloat(dist1, 'f', 2, 64)
	calorie1:=strconv.FormatFloat(calorie, 'f', 2, 64)

str:= "Количество шагов: "+step1+"/n.Дистанция составила "+dist2+ "/nкм. Вы сожгли: "+calorie1+"/n"
return str
}