package delivery

import (
	"be13/ca/features/mahasiswa"
)

type MahasiswaRequest struct {
	Nama   string `json:"nama" form:"nama"`
	Alamat string `json:"alamat" form:"alamat"`
}

func requestToCore(mahasiswaInput MahasiswaRequest) mahasiswa.Core {
	mahasiswaCoreData := mahasiswa.Core{
		Nama:   mahasiswaInput.Nama,
		Alamat: mahasiswaInput.Alamat,
	}
	return mahasiswaCoreData
}
