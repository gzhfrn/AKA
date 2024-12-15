package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// BubbleSort untuk mengurutkan berdasarkan rating
func BubbleSort(data [][]string) [][]string {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			rating1, _ := strconv.ParseFloat(data[j][0], 64)
			rating2, _ := strconv.ParseFloat(data[j+1][0], 64)
			if rating1 < rating2 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// Rekursif insertion sort helper untuk menyisipkan elemen ke posisi yang tepat
func insert(data [][]string, i int) {
	if i <= 0 {
		return
	}

	key := data[i]
	j := i - 1

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
	if i == len(data) {
		return data
	}

	insert(data, i)
	return RecursiveInsertionSort(data, i+1)
}

// Fungsi untuk membaca data dari file CSV
func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) > 2 {
			rating := strings.Split(fields[0], ":")[1]
			rating = strings.TrimSpace(strings.Trim(rating, "'}"))
			data = append(data, []string{rating, fields[2]})
		}
	}

	return data, scanner.Err()
}

// Fungsi utama
func main() {
	filePath := "C:/Users/Ghaisani/Documents/Telkom/Semester 3/AKA/gofood_ratings.csv"
	data, err := readCSV(filePath)
	if err != nil {
		fmt.Println("Error membaca file:", err)
		return
	}

	// Salin data untuk digunakan pada kedua algoritma
	dataForBubble := make([][]string, len(data))
	copy(dataForBubble, data)

	dataForInsertion := make([][]string, len(data))
	copy(dataForInsertion, data)

	// Hitung waktu untuk Bubble Sort
	start := time.Now()
	BubbleSort(dataForBubble)
	bubbleTime := time.Since(start).Milliseconds()

	// Hitung waktu untuk Recursive Insertion Sort
	start = time.Now()
	RecursiveInsertionSort(dataForInsertion, 1)
	insertionTime := time.Since(start).Milliseconds()

	// Cetak hasil waktu eksekusi
	fmt.Printf("Waktu eksekusi Bubble Sort: %d ms\n", bubbleTime)
	fmt.Printf("Waktu eksekusi Recursive Insertion Sort: %d ms\n", insertionTime)

	// Buat grafik perbandingan waktu eksekusi
	graph := charts.NewBar()
	graph.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Perbandingan Running Time Sorting"}))

	graph.SetXAxis([]string{"Bubble Sort", "Insertion Sort"}).
		AddSeries("Running Time (ms)", []opts.BarData{
			{Name: "Bubble Sort", Value: bubbleTime},
			{Name: "Insertion Sort", Value: insertionTime},
		})

	// Simpan grafik sebagai file HTML
	f, err := os.Create("sorting_runtime_comparison.html")
	if err != nil {
		fmt.Println("Error membuat file grafik:", err)
		return
	}
	defer f.Close()
	graph.Render(f)

	fmt.Println("Grafik perbandingan running time disimpan sebagai 'sorting_runtime_comparison.html'.")
}
