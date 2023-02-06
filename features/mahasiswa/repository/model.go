package repository

type Mahasiswa struct {
	ID     uint
	Nama   string
	Alamat string
	Nilai  []Nilai
}
type Matkul struct {
	ID         uint
	NamaMatkul string
	Nilai      []Nilai
}

type Nilai struct {
	ID          uint
	MahasiswaID uint
	MatkulID    uint
	Nilai       uint
	Mahasiswa   Mahasiswa
	Matkul      Matkul
}
