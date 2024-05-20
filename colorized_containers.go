package colorizedcontainers

import (
	"fmt"
)

func swapColumns(matrix [][]uint64, col1, col2 int) [][]uint64 {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col1], matrix[i][col2] = matrix[i][col2], matrix[i][col1]
	}
	return matrix
}

func CheckSortAvailable(n int, containers [][]uint64) bool {
	if n == 1 {
		return true
	}

	// Подсчет суммы элементов нулевого контейнера
	var S uint64 = 0
	for i := 0; i < n; i++ {
		S += containers[0][i]
	}

	for i := 0; i < n; i++ {
		// Проверка, можно ли обменять в этом контейнере все шары разных цветова на один какой-нибудь цвет
		Si := S - containers[0][i]
		var Sj uint64 = 0
		for j := 1; j < n; j++ {
			Sj += containers[j][i]
		}
		// Если нельзя, то это условие не срабатывает и мы возвращаем false
		if Si == Sj {
			// Если можно, то:
			// Меняем местами этот цвет с 0-ым цветом
			containers = swapColumns(containers, 0, i)

			// Распределяем все остальные цвета по контейнерам
			con_i := 1
			size := containers[con_i][0]
			for j := 1; j < n; j++ {
				color_size := containers[0][j]
				for size != 0 {
					if size < color_size {
						containers[con_i][j] += size
						color_size -= size
						con_i++
						size = containers[con_i][0]
					} else {
						containers[con_i][j] += color_size
						size -= color_size
						break
					}
				}
			}

			// Вырезаем подмассив полученных контейнеров
			rowSlice := containers[1:]
			var colSlice [][]uint64
			for _, row := range rowSlice {
				colSlice = append(colSlice, row[1:])
			}

			// Рекурсивно повторяем эту функцию для меньшего количества контейнеров
			return CheckSortAvailable(n-1, colSlice)
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scanln(&n)
	if n < 1 && n > 100 {
		fmt.Println("Wrong n. Should: 1 <= n <= 100")
	}

	containers := make([][]uint64, n)
	for i := range containers {
		containers[i] = make([]uint64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&containers[i][j])
			if containers[i][j] > 1000*1000*1000 {
				fmt.Println("Wrong item. Should: 0 <= containers[i][j] <= 1'000'000'000")
			}
		}
	}
	if CheckSortAvailable(n, containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
