package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"errors"
	"log"
	
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
			
			return 0, 0, err
			}
	if step<=0{
		return 0, 0, errors.New("ошибка")
		}
	time, err := time.ParseDuration(slice[1])
		if err != nil {
			
			return 0, 0, err
		}
	if time<=0{
		return 0, 0, errors.New("ошибка")
		}
	return step, time, nil
	}
	return 0, 0, errors.New("ошибка")
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	step, time, err := parsePackage(data)
	if err != nil {

log.Println("Ошибка при разборе данных:", err)
		return ""
	}
	
	dist:=float64(step)*stepLength 
	dist1:=float64(dist/mInKm)
	calorie, err:=spentcalories.WalkingSpentCalories(step, weight, height, time)
	if err != nil {
		
		log.Println("Ошибка при разборе данных:", err)
		return ""
	}
	if calorie==0{
fmt.Println("Ошибка ошибка 2 при разборе данных:", err)
	return ""
	}
	step1:=strconv.Itoa(step)
	dist2:=strconv.FormatFloat(dist1, 'f', 2, 64)
	calorie1:=strconv.FormatFloat(calorie, 'f', 2, 64)

str:= "Количество шагов: "+step1+".\nДистанция составила "+dist2+ " км.\nВы сожгли "+calorie1+" ккал.\n"
return str
}