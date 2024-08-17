package util

import (
	"math/rand"
	"time"
)

// disini kita coba membuat until untuk bikin random string
// tapicara ini kurang efectif karena ada kemungkinana sama, walaupun kecil
// jadi jik punya cara lain juga boleh
// ini cuma simulasi saja

func RandomString(l int) string {
	const chartset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano()) // initnya dia bikin random berdasarkan waktu sekarang yg dibuat int
	// rand.NewSource(seed int64) rand.Source: Fungsi ini membuat instance baru dari rand.
	// Source yang menggunakan seed sebagai nilai awal untuk generator bilangan acak.
	// Objek rand.Source ini adalah antarmuka yang digunakan untuk menghasilkan angka acak. Dengan menggunakan seed yang berbeda,
	//  Anda bisa mendapatkan urutan angka acak yang berbeda.
	r := rand.New(seed) //kita masukan seednya yg suduah kita buat

	result := make([]byte, l) // kita bikin slice kosong
	for i := range result {   // kita lakukan perulangan sesuai jumplah panjang resultnya
		result[i] = chartset[r.Intn(len(chartset))] // kita ambil dari chartset indexnya rndom sesuai panjang chartset
	}
	return string(result)
}
