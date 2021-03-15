//Пакет для чтения данных из файла

package datafile

import (
	"bufio"
	"os"
	"strconv"
)

//GetFloats читает значение float64 из каждой строки файла
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64         //Объявление возвращаемого массива
	file, err := os.Open(fileName) //Открытие файла с переданным именем
	//Если при открытии файла произошла ошибка, сообщить о ней и завершить работу
	if err != nil {
		return numbers, err
	}
	i := 0                            //Переменная для хранения индекса, по которому должно выполняться присваивание
	scanner := bufio.NewScanner(file) //Для файла создается новое значение
	for scanner.Scan() {              //Читает строку из файла
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) //Выводит строку
		if err != nil {
			return numbers, err
		}

		i++ //Переход к следующему индексу массива

	} //Цикл выполняется до того, как будет достигнут конец файла, а scanner.Scan вернет false

	err = file.Close() //Закрывает файл для освобождения ресурсов

	if err != nil {
		return numbers, err
	} //Если при закрытии файла произошла ошибка, сообщить о ней и завершить работу

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	} //Если при сканировании файла произошла ошибка, сообщить о ней и завершить работу
	return numbers, nil //Если выполнение дошло до этой точки, значит, ошибок не было, поэтому программа возвращает массив чисел и значение ошибки nil
}
