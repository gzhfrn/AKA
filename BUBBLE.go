/////BUBBLE SORT ITERATIF//////

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BubbleSort untuk mengurutkan berdasarkan rating
func BubbleSort(data [][]string) [][]string {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			rating1, _ := strconv.ParseFloat(data[j][0], 64)
			rating2, _ := strconv.ParseFloat(data[j+1][0], 64)
			if rating1 < rating2 { // Tukar jika rating1 lebih kecil
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main() {
	// Buka file CSV
	file, err := os.Open("C:/Users/Ghaisani/Documents/Telkom/Semester 3/AKA/gofood_ratings.csv")
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()

	// Baca file
	scanner := bufio.NewScanner(file)
	data := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) > 2 {
			// Ambil hanya 'average' dari ratings
			rating := strings.Split(fields[0], ":")[1]
			rating = strings.TrimSpace(strings.Trim(rating, "'}"))
			data = append(data, []string{rating, fields[2]}) // Simpan rating dan brand.displayName
		}
	}

	// Urutkan menggunakan Bubble Sort
	fmt.Println("Mengurutkan dengan Bubble Sort...")
	sortedData := BubbleSort(data)

	// Tampilkan hasil
	fmt.Println("\nHasil Pengurutan:")
	for _, row := range sortedData {
		fmt.Println(row[1], "-", row[0]) // Tampilkan brand.displayName dan rating
	}
}
