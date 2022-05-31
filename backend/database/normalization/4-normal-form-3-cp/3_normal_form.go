package main

// Normalisasi database dalam bentuk 3NF bertujuan untuk menghilangkan seluruh atribut atau field yang tidak berhubungan dengan primary key.
// Dengan demikian tidak ada ketergantungan transitif pada setiap kandidat key
// fungsi normalisasi 3NF antara lain :
// 1. Memenuhi semua persyaratan dari bentuk normal kedua.
// 2. Menghapus kolom yang tidak tergantung pada primary key.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//the struct are irrelevant for the code, but hint for column
type Rekap struct {
	NoBon     int
	Discount  int
	Total     int
	Bayar     int
	Kembalian int
	KodeKasir string
	Tanggal   string
	Waktu     string
}

type DetailRekap struct {
	NoBon      int
	KodeBarang string
	Harga      int
	Jumlah     int
}

type Barang struct {
	KodeBarang string
	NamaBarang string
	Harga      int
}

type Kasir struct {
	KodeKasir string
	NamaKasir string
}

// Migrate digunakan untuk melakukan migrasi database dengan data yang dibutuhkan
// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi Migrate:
// Buatlah tabel dengan nama rekap, rekap_detail, barang, dan kasir
// Lalu insert data ke masing-masing tabel seperti pada contoh di bagian bawah file ini
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS rekap (
		NoBon INTEGER PRIMARY KEY,
		Discount INTEGER,
		Total INTEGER,
		Bayar INTEGER,
		Kembalian INTEGER,
		KodeKasir VARCHAR(10),
		Tanggal VARCHAR(10),
		Waktu VARCHAR(10)
	);` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO 
	rekap (NoBon, Discount, Total, Bayar, Kembalian, KodeKasir, Tanggal, Waktu)
	VALUES
	(1, 0, 0, 0, 0, "Kasir1", "2020-01-01", "12:00:00"),
	(2, 0, 0, 0, 0, "Kasir2", "2020-01-01", "12:00:00"),
	(3, 0, 0, 0, 0, "Kasir3", "2020-01-01", "12:00:00")
	  ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS rekap_detail (
		NoBon INTEGER,
		KodeBarang VARCHAR(10),
		Harga INTEGER,
		Jumlah INTEGER,
		FOREIGN KEY (NoBon) REFERENCES rekap_3nf(NoBon)
	);` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO
	rekap_detail (NoBon, KodeBarang, Harga, Jumlah)
	VALUES
	(1, "B001", 100, 1),
	(2, "B002", 200, 2),
	(2, "B002", 200, 2),
	(2, "B002", 200, 2),
	(3, "B003", 300, 3),
	(84, "B024", 400, 4),
	(85, "B015", 500, 5)
	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = ` CREATE TABLE IF NOT EXISTS barang (
		KodeBarang VARCHAR(10) PRIMARY KEY,
		NamaBarang VARCHAR(50),
		Harga INTEGER
	);` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	INSERT INTO
	barang (KodeBarang, NamaBarang, Harga)
	VALUES
	("B001", "Buku", 100),
	("B002", "Komputer", 200),
	("B003", "Laptop", 300),
	("B004", "Printer", 400),
	("B005", "Monitor", 500)
	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = ` CREATE TABLE IF NOT EXISTS kasir (
		KodeKasir VARCHAR(10) PRIMARY KEY,
		NamaKasir VARCHAR(50)
	);` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(` 
	INSERT INTO
	kasir (KodeKasir, NamaKasir)
	VALUES
	("Kasir1", "Nama Kasir 1"),
	("Kasir2", "Nama Kasir 2"),
	("Kasir3", "Nama Kasir 3")

	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkNoBonExists:
// checkNoBonExists digunakan untuk menghitung jumlah data yang ada berdasarkan no_bon
func checkNoBonExists(noBon string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM rekap WHERE NoBon = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi countRekapDetailByNoBon:
// countRekapDetailByNoBon digunakan untuk menghitung jumlah rekap detail yang ada berdasarkan no_bon
func countRekapDetailByNoBon(noBon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM rekap_detail WHERE NoBon = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, noBon)
	var countBon int
	err = row.Scan(&countBon)
	if err != nil {
		return 0, err
	}
	return countBon, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkBarangExists:
func checkBarangExists(kodeBarang string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM barang WHERE KodeBarang = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeBarang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkKasirExists:
func checkKasirExists(kodeKasir string) (bool, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM kasir WHERE KodeKasir = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, kodeKasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return false, err
	}
	return true, nil
}
