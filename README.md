# 📦 Dokumentasi Proyek

Proyek ini menggunakan berbagai konsep pemrograman Go, termasuk pemanfaatan interface, struct, penanganan error, manajemen file, serta pola desain seperti dependency injection dan constructor function.

# Struktur Program

Program ini dibagi menjadi beberapa paket, masing-masing dengan tanggung jawabnya sendiri sendiri

- cmdmanager: Menangani interaksi user melalui Command Line Interface (CLI).
- conversion: Berisi fungsi untuk konversi tipe data.
- filemanager: Mengelola operasi baca/tulis file.
- iomanager: Mendefinisikan antarmuka (interface) untuk abstraksi input/output.
- prices: Mengandung logika utama untuk perhitungan harga dan pajak.

# Konsep-Konsep Kunci yang Dipraktikkan

### Interface (Antarmuka)

- Interface digunakan untuk mendefinisikan kontrak fungsi, seperti `ReadFile` dan `WriteResult`.
- Membuat kode lebih fleksibel dalam penggunaan input/output, misalnya:
  - `CMDManager` untuk input dari command line
  - `FileManager` untuk input/output dari/ke file
- Interface ini didefinisikan di package `iomanager` dan diimplementasikan oleh `CMDManager` dan `FileManager`.

### Struct (Struktur Data)

- Struct digunakan untuk mengelompokkan data dan fungsi yang saling berkaitan, seperti:
  - `CMDManager`
  - `FileManager`
  - `TaxIncludedPricesJob`
- Method receiver digunakan untuk menempelkan fungsi ke struct, contoh:
  ```go
  func (fm FileManager) ReadFile() {...}
  ```

### Error Handling

- Setiap kali memanggil fungsi yang bisa gagal, error selalu dicek menggunakan pola standar berikut:
  ```go
  if err != nil {
    return err
  }
  ```
  Pola ini memastikan bahwa kesalahan bisa ditangani dengan baik tanpa menyebabkan crash atau hasil yang tidak diinginkan. Dalam Go, fungsi yang berpotensi gagal biasanya mengembalikan dua nilai: hasil dan error.
  Untuk membuat pesan error khusus, digunakan fungsi:

```go
errors.New("pesan error")
```

### Konversi Tipe Data

Untuk mengubah tipe data dari string ke float64, digunakan fungsi:

```go
strconv.ParseFloat(stringValue, 64)
Fungsi ini digunakan dalam package conversion.
```

Selalu penting untuk memeriksa error saat melakukan konversi, karena data string bisa saja tidak valid dan menyebabkan program gagal jika tidak ditangani.

### Manajemen File

Program ini menggunakan fungsi-fungsi dari package standar Go untuk membaca dan menulis file:

- os.Open → Membuka file untuk dibaca.
- os.Create → Membuat file baru untuk ditulis.
- bufio.NewScanner → Digunakan untuk membaca isi file baris per baris.
- json.NewEncoder → Menyimpan data ke dalam file dalam format JSON.

```go
file, err := os.Open("data.txt")
if err != nil {
  return err
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
  line := scanner.Text()
  // proses line
}
```

### Pola Desain yang Digunakan

- Dependency Injection
  Alih-alih membuat objek I/O secara langsung di dalam TaxIncludedPricesJob, objek IOManager disuntikkan (injected) melalui constructor. Hal ini membuat kode lebih fleksibel, mudah diuji, dan tidak tergantung pada implementasi tertentu. Contoh:

```go
func NewTaxIncludedPricesJob(io IOManager) *TaxIncludedPricesJob {
  return &TaxIncludedPricesJob{io: io}
}
```

- Constructor Function
  Untuk membuat instance dari sebuah struct dan memastikan semua dependensi tersedia, digunakan fungsi constructor seperti:
```go
func NewTaxIncludedPricesJob(io IOManager) *TaxIncludedPricesJob
```
Pendekatan ini membuat kode lebih bersih dan lebih mudah dirawat karena semua inisialisasi dilakukan di satu tempat.
