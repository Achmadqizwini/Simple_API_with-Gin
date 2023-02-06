package delivery

import (
	"be13/ca/features/mahasiswa"
)

type NilaiMhsResponse struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	NamaMatkul  string `json:"nama_matkul"`
	NilaiRerata uint   `json:"nilai_rata-rata"`
}

func fromCore(dataCore mahasiswa.NilaiMhs) NilaiMhsResponse {
	return NilaiMhsResponse{
		ID:          dataCore.ID,
		Nama:        dataCore.Nama,
		NamaMatkul:  dataCore.NamaMatkul,
		NilaiRerata: dataCore.NilaiRerata,
	}
}

func fromCoreList(dataCore []mahasiswa.NilaiMhs) []NilaiMhsResponse {
	var dataResponse []NilaiMhsResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
