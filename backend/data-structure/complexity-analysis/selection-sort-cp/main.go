// Selection Sort

// Algoritma Selection Sort didasarkan pada gagasan untuk menemukan elemen minimum atau maksimum
// dalam larik yang tidak diurutkan dan kemudian meletakkannya di posisi yang benar dalam larik yang diurutkan.

// Asumsikan bahwa array A[] = [7, 5, 4, 2] perlu diurutkan dalam urutan menaik.

// Elemen minimum dalam array yaitu 2 dicari dan kemudian ditukar dengan elemen yang saat ini berada di posisi
// pertama, yaitu 7. Sekarang elemen minimum di sisa array yang tidak disortir dicari dan diletakkan di posisi kedua,
// dan seterusnya.

// Contoh:
// Mari kita coba memahami kode semua dengan sebuah contoh: A[] = { 7, 5, 4, 2 } Pada i pertama dari iterasi, elemen dari posisi 0 ke i - 1 akan diurutkan.

// Image:
// Untuk memahami solusi sort diatas dengan Selection Sort, bisa diperhatikan gambar ini: https://he-s3.s3.amazonaws.com/media/uploads/2888f5b.png

// Visualisasi:
// Contoh Visualisasi Selection Sort dapat dilihat di sini: https://www.hackerearth.com/practice/algorithms/sorting/selection-sort/visualize/

// Time Complexity:
// Untuk menemukan elemen minimum dari array N elemen, N - 1 perbandingan diperlukan.
// Setelah menempatkan elemen minimum di posisi yang tepat, ukuran array yang tidak disortir berkurang
// menjadi N - 1 lalu N - 2 perbandingan diperlukan untuk menemukan minimum dalam array yang tidak disortir.

// Karena itu (N - 1)+(N - 2)+.......+ 1=(N.(N - 1))/2 perbandingan dan N swap menghasilkan kompleksitas keseluruhan yaitu O(n*n) / O(N2).

// Space Complexity:
// - Space Complexity dari Selection Sort adalah O(1)

package main

import "fmt"

func main() {
	sample := []int{3, 4, 5, 2, 1}
	arrSorted := SelectionSort(sample)

	fmt.Println("\nAfter SelectionSort")
	for _, val := range arrSorted {
		fmt.Println(val)
	}
}

func SelectionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			//TODO: answer here
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
		/*
			i= 1
			minIndex = 1
			j = 2
			arr[j] = 5
			arr[minIndex] = 2
			arr[i] = 5

			i= 2
			minIndex = 2
			j = 3
			arr[j] = 4
			arr[minIndex] = 2
			arr[i] = 4

			i= 3
			minIndex = 3
			j = 4
			arr[j] = 2
			arr[minIndex] = 2
			arr[i] = 2

			i= 4
			minIndex = 4
			j = 5
			arr[j] = 1
			arr[minIndex] = 2
			arr[i] = 1

			i= 5
			minIndex = 5
			j = 5
			arr[j] = 1
			arr[minIndex] = 1
			arr[i] = 1

		*/

	}

	return arr
}
