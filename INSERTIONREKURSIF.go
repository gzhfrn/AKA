////INSERTION SORT REKURSIF/////
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rekursif insertion sort helper untuk menyisipkan elemen ke posisi yang tepat
func insert(data [][]string, i int) {
	if i <= 0 {
		return
	}

	key := data[i]
	j := i - 1

	// Geser elemen ke kanan jika lebih kecil dari elemen sebelumnya
	for j >= 0 {
		rating1, _ := strconv.ParseFloat(data[j][0], 64)
		ratingKey, _ := strconv.ParseFloat(key[0], 64)
		if rating1 < ratingKey {
			data[j+1] = data[j]
			j--
		} else {
			break
		}
	}
	data[j+1] = key
}

// Rekursif insertion sort untuk mengurutkan data
func RecursiveInsertionSort(data [][]string, i int) [][]string {
	if i == len(data) { // Basis kasus
		return data
	}

	insert(data, i)                          // Sisipkan elemen ke tempat yang sesuai
	return RecursiveInsertionSort(data, i+1) // Rekursi ke elemen berikutnya
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

	// Urutkan menggunakan Insertion Sort rekursif
	fmt.Println("Mengurutkan dengan Insertion Sort rekursif...")
	sortedData := RecursiveInsertionSort(data, 1)

	// Tampilkan hasil
	fmt.Println("\nHasil Pengurutan:")
	for _, row := range sortedData {
		fmt.Println(row[1], "-", row[0]) // Tampilkan brand.displayName dan rating
	}
}
