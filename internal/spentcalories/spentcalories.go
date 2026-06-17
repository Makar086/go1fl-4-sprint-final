package spentcalories

import (
	"time"
	"strings"
	"strconv"
	"errors"
	
	
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice)==3{
	step, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
			}
	if step<=0{
		return 0, "",0,  errors.New("Ошибка")
		}
	time, err := time.ParseDuration(slice[2])
		if err != nil {
		return 0, "", 0, err
		}
if time<=0{
		return 0, "",0, errors.New("Ошибка")
		}
	
	return step, slice[1], time, nil
	}
	return 0, "", 0, errors.New("Ошибка мало элементов")
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
if steps <=0||height <=0{
	return 0
	}
	lenStep:=height*stepLengthCoefficient
	return lenStep*float64(steps)/mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
if steps <=0||height <=0|| duration <=0{
	return 0
	}
	if duration>0{
	dist:= distance(steps , height)
	
	// Переводим в часы
	hours := duration.Hours()
	return dist/hours
	}
	return 0
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
if weight<=0||height <=0{
	return "", errors.New("Ошибка")
	}
	steps, vid, time, err:=parseTraining(data)
if err != nil {
			
			return  "", err
		}
	switch vid{
	case "Бег":
	calore, err:=RunningSpentCalories(steps, weight, height, time)
if err != nil {
			
			return  "", err
		}
	distan:=distance(steps, height)
	speed:=meanSpeed(steps, height, time)
	
	minutes := time.Minutes()
	time3:=minutes/minInH
	time1:=strconv.FormatFloat(time3, 'f', 2, 64)
	distan1:=strconv.FormatFloat(distan, 'f', 2, 64)
	speed1:=strconv.FormatFloat(speed, 'f', 2, 64)
	calore1:=strconv.FormatFloat(calore, 'f', 2, 64)
str:= "Тип тренировки: Бег\nДлительность: "+time1+" ч.\nДистанция: "+distan1+" км.\nСкорость: "+speed1+" км/ч\nСожгли калорий: "+calore1+"\n"
return str, err
	case "Ходьба":
	calore, err:=WalkingSpentCalories(steps, weight, height, time)
if err != nil {
			
			return  "", err
		}
	distan:=distance(steps, height)
	speed:=meanSpeed(steps, height, time)
	
	minutes := time.Minutes()
	time3:=minutes/minInH
	time1:=strconv.FormatFloat(time3, 'f', 2, 64)
	distan1:=strconv.FormatFloat(distan, 'f', 2, 64)
	speed1:=strconv.FormatFloat(speed, 'f', 2, 64)
	calore1:=strconv.FormatFloat(calore, 'f', 2, 64)
str:= "Тип тренировки: Ходьба\nДлительность: "+time1+" ч.\nДистанция: "+distan1+" км.\nСкорость: "+speed1+" км/ч\nСожгли калорий: "+calore1+"\n"
return str, err
	default:
return "", errors.New("неизвестный тип тренировки")
	}
	
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	
	speed:=meanSpeed(steps, height, duration)
	if steps <=0||weight<=0||height <=0|| duration <=0{
	return 0, errors.New("Ошибка")
	}
	

	// Переводим в минуты
	minute := duration.Minutes()
	return (weight*speed*minute)/minInH ,nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	speed:=meanSpeed(steps, height, duration)
	if steps <=0||weight<=0||height <=0|| duration <=0{
	return 0, errors.New("Ошибка")
	}
	

	// Переводим в минуты
	minute := duration.Minutes()
	calore:= (weight*speed*minute)/minInH
	return calore*walkingCaloriesCoefficient, nil
}
