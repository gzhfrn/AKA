//// GRAFIK PERBANDINGAN RUNNING TIME //////

package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

// LoadData membaca file CSV dan mengembalikan data dalam bentuk slice
func LoadData(filePath string) ([][]string, error) {
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

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	filePath := "C:/Users/Ghaisani/Documents/Telkom/Semester 3/AKA/gofood_ratings.csv"
	data, err := LoadData(filePath)
	if err != nil {
		fmt.Println("Error membaca file:", err)
		return
	}

	// Ukuran data yang akan diuji
	sizes := []int{10, 50, 100, 200, 500, 800, 1000}
	var bubbleTimes []float64
	var insertionTimes []float64

	fmt.Println("Membandingkan running time Bubble Sort dan Recursive Insertion Sort...")
	for _, size := range sizes {
		if size > len(data) {
			fmt.Printf("Ukuran %d melebihi jumlah data. Melewatkan.\n", size)
			continue
		}

		sampleData := make([][]string, size)
		copy(sampleData, data[:size])

		// Ukur waktu untuk Bubble Sort
		startBubble := time.Now()
		BubbleSort(sampleData)
		elapsedBubble := time.Since(startBubble)
		bubbleTimes = append(bubbleTimes, elapsedBubble.Seconds())

		// Salin ulang data untuk Insertion Sort
		sampleData = make([][]string, size)
		copy(sampleData, data[:size])

		// Ukur waktu untuk Recursive Insertion Sort
		startInsertion := time.Now()
		RecursiveInsertionSort(sampleData, 1)
		elapsedInsertion := time.Since(startInsertion)
		insertionTimes = append(insertionTimes, elapsedInsertion.Seconds())

		// Cetak hasil untuk ukuran data ini
		fmt.Printf("Ukuran data: %d\n", size)
		fmt.Printf("Bubble Sort: %.4f detik | Insertion Sort: %.4f detik\n\n", elapsedBubble.Seconds(), elapsedInsertion.Seconds())
	}

	// Buat grafik
	p := plot.New()
	p.Title.Text = "Perbandingan Running Time"
	p.X.Label.Text = "Ukuran Data"
	p.Y.Label.Text = "Waktu (detik)"

	// Atur ticks pada sumbu X dan Y
	p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 10, Label: "10"},
		{Value: 50, Label: "50"},
		{Value: 100, Label: "100"},
		{Value: 200, Label: "200"},
		{Value: 500, Label: "500"},
		{Value: 800, Label: "800"},
		{Value: 1000, Label: "1000"},
	})

	// Pastikan data pada bubblePoints dan insertionPoints valid
	if len(bubbleTimes) == len(sizes) && len(insertionTimes) == len(sizes) {
		bubblePoints := make(plotter.XYs, len(sizes))
		insertionPoints := make(plotter.XYs, len(sizes))
		for i, size := range sizes {
			bubblePoints[i].X = float64(size)
			bubblePoints[i].Y = bubbleTimes[i]
			insertionPoints[i].X = float64(size)
			insertionPoints[i].Y = insertionTimes[i]
		}

		lineBubble, _ := plotter.NewLine(bubblePoints)
		lineBubble.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Merah
		lineInsertion, _ := plotter.NewLine(insertionPoints)
		lineInsertion.Color = color.RGBA{B: 255, A: 255} // Biru

		// Tambahkan garis ke plot
		p.Add(lineBubble, lineInsertion)
		p.Legend.Add("Bubble Sort", lineBubble)
		p.Legend.Add("Recursive Insertion Sort", lineInsertion)
	} else {
		fmt.Println("Error: Panjang data tidak sesuai dengan ukuran.")
	}

	// Simpan grafik
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "grafikrn2.png"); err != nil {
		fmt.Println("Error menyimpan grafik:", err)
		return
	}

	fmt.Println("Grafik berhasil disimpan sebagai grafikrn2.png")
}
