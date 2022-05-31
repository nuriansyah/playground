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
func Migrate() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `CREATE TABLE IF NOT EXISTS rekap_3nf (
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
	rekap_3nf (NoBon, Discount, Total, Bayar, Kembalian, KodeKasir, Tanggal, Waktu)
	VALUES
	(1, 0, 0, 0, 0, "Kasir1", "2020-01-01", "12:00:00"),
	(2, 0, 0, 0, 0, "Kasir2", "2020-01-01", "12:00:00") 
	  ;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = `CREATE TABLE IF NOT EXISTS rekap_detail_3nf (
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
	rekap_detail_3nf (NoBon, KodeBarang, Harga, Jumlah)
	VALUES
	(1, "B001", 100, 1),
	(1, "B002", 200, 2)
	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = ` CREATE TABLE IF NOT EXISTS barang_3nf (
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
	barang_3nf (KodeBarang, NamaBarang, Harga)
	VALUES
	("B001", "Buku", 100),
	("B002", "Komputer", 200)
	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	sqlStmt = ` CREATE TABLE IF NOT EXISTS kasir_3nf (
		KodeKasir VARCHAR(10) PRIMARY KEY,
		NamaKasir VARCHAR(50)
	);` // TODO: replace this

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(` 
	INSERT INTO
	kasir_3nf (KodeKasir, NamaKasir)
	VALUES
	("Kasir1", "Nama Kasir 1"),
	("Kasir2", "Nama Kasir 2")
	;`) // TODO: replace this

	if err != nil {
		panic(err)
	}

	return db, nil
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoBon:
func checkLatestNoBon(no_bon string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM rekap_3nf WHERE NoBon = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, no_bon)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoBonDetail:
func checkLatestNoBonDetail(no_bon_detail string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM rekap_detail_3nf WHERE NoBon = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, no_bon_detail)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoBarang:
func checkLatestNoBarang(kode_barang string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM barang_3nf WHERE KodeBarang = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, kode_barang)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}

// Tugas: Replace tanda ... dengan Query yang tepat pada fungsi checkLatestNoKasir:
func checkLatestNoKasir(kode_kasir string) (int, error) {
	db, err := sql.Open("sqlite3", "./normalize-cp.db")
	if err != nil {
		panic(err)
	}

	sqlStmt := `SELECT COUNT(*) FROM kasir_3nf WHERE KodeKasir = ?` // TODO: replace this

	row := db.QueryRow(sqlStmt, kode_kasir)
	var latestId int
	err = row.Scan(&latestId)
	if err != nil {
		return 0, err
	} else {
		return 1, nil
	}
}
