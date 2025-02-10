package array

import (
	"fmt"
)

// testing the slice type function
func SliceType() {

	dataMahasiswa := [...]string{"Fajar", "Aziz", "Kurniawan", "Semarang", "Indonesia"}
	// slice data mahasiswa
	dataDaerah := dataMahasiswa[3:5] // dataMahasiswa[3:5] artinya mengambil data dari index ke 3 sampai index ke sebelum 5 (4)
	// slice itu ada tiga komponen yaitu pointer, length, dan capacity
	fmt.Println("Data Mahasiswa =", dataDaerah)

	dataNama := dataMahasiswa[:4] // dataMahasiswa[:4] artinya mengambil data dari index ke 0 sampai index ke sebelum 4 (3) karna tidak ada angka di depan maka di anggap 0
	fmt.Println("Data Nama =", dataNama)

	dataAlamat := dataMahasiswa[3:] // dataMahasiswa[3:] artinya mengambil data dari index ke 3 sampai index terakhir
	fmt.Println("Data Alamat =", dataAlamat)

	dataSemua := dataMahasiswa[:] // dataMahasiswa[:] artinya mengambil semua data
	fmt.Println("Data Semua =", dataSemua)

	// beberapa function untuk slice itu sama seperti array
	// len, cap, append, copy
	fmt.Println("Panjang Data Mahasiswa =", len(dataNama))
	fmt.Println("Kapasitas Data Mahasiswa =", cap(dataNama))
	// Menambahkan data ke slice
	dataNama = append(dataNama, "Jakarta")
	fmt.Println("Data Nama =", dataNama)

	// jika ingin membuat slice baru
	dataSlice := make([]string, 2, 5)
	dataSlice[0] = "Fajar"
	// dataSlice[1] = "Aziz"
	// dataSlice[2] = "Kurniawan" // error karena kapasitasnya yang di definisikan hanya 2. Harus menggunakan append
	fmt.Println("Data Slice =", dataSlice)
	fmt.Println("Panjang Data Slice =", len(dataSlice))   // tetap akan dua meskipun dataSlice[1] tidak diisi
	fmt.Println("Kapasitas Data Slice =", cap(dataSlice)) // kapasitasnya 5

	dataSlice[1] = "Aziz"
	dataSlice1 := append(dataSlice, "Kurnaiwan") // dikarenakan kapasitasnya 5 maka dataSlice1 akan menambahkan data baru tanpa membuat array baru
	fmt.Println("Data Slice 1 =", dataSlice1)

	dataSlice1[0] = "Fajar Kurniawan"
	fmt.Println("Data Slice 1 =", dataSlice1)
	fmt.Println("Data Slice =", dataSlice)

	// Menyalin data dari slice
	dataNamaBaru := make([]string, len(dataNama), cap(dataNama))
	copy(dataNamaBaru, dataNama)
	fmt.Println("Data Nama Baru =", dataNamaBaru)

	dataMahasiswaSlice := dataMahasiswa[3:]
	fmt.Println("Data Mahasiswa Slice =", dataMahasiswaSlice)

	dataMahasiswaSlice[0] = "Bandung"
	dataMahasiswaSlice[1] = "Jawa Barat"
	fmt.Println("Data Mahasiswa Slice =", dataMahasiswa)
	// pada saat melakukan perubahan pada dataMahasiswaSlice maka data asli juga akan berubah, dikarenakan slice itu reference ke array aslinya

	dataMahasiswaAppend := append(dataMahasiswaSlice, "Indonesia")
	// berbeda dengan append pada array, pada slice append akan membuat data baru
	fmt.Println("Data Mahasiswa Append =", dataMahasiswaAppend)

	dataMahasiswaAppend[0] = "Jakarta Lama"
	fmt.Println("Data Mahasiswa Append =", dataMahasiswaAppend)
}
