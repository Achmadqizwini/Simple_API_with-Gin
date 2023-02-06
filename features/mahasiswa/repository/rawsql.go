package repository

import (
	"be13/ca/features/mahasiswa"
	"database/sql"
	"log"
)

type mahasiswaRepository struct {
	db *sql.DB
}

func NewRaw(db *sql.DB) mahasiswa.RepositoryInterface {
	return &mahasiswaRepository{
		db: db,
	}
}

// Create implements user.RepositoryInterface
func (repo *mahasiswaRepository) Create(input mahasiswa.Core) (err error) {
	var query = "Insert into mahasiswa (Nama, Alamat) Values (?, ?)"
	statement, errPrepare := repo.db.Prepare(query)
	if errPrepare != nil {
		return errPrepare
	}

	_, errExec := statement.Exec(input.Nama, input.Alamat)
	if errExec != nil {
		return errExec
	}
	return nil
}

// DeleteUser implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Delete(id int) (err error) {
	query := ("Delete from mahasiswa where id in (?)")
	statement, errPrepare := repo.db.Prepare(query)
	if errPrepare != nil {
		return err
	}

	_, errExec := statement.Exec(id)
	if errExec != nil {
		return errExec
	}
	return nil
}

// UpdateUser implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Update(input mahasiswa.Core, id int) (err error) {
	var query = ("Update users set nama = ?, alamat = ? where id = ?")
	statement, errPrepare := repo.db.Prepare(query)
	if errPrepare != nil {
		return errPrepare
	}
	_, errExec := statement.Exec(input.Nama, input.Alamat, id)
	if errExec != nil {
		return errExec
	}
	return nil
}

// Read implements mahasiswa.RepositoryInterface
func (repo *mahasiswaRepository) Read(id int) (data []mahasiswa.NilaiMhs, err error) {
	result, errSelect := repo.db.Query("select u.nama, v.nama_matkul, avg(nilai) from nilai inner join mahasiswa u on nilai.mahasiswa_id = u.id inner join matkul v on nilai.matkul_id = v.id where nilai.mahasiswa_id = (?)", id)
	if errSelect != nil {
		log.Fatal("error select", errSelect.Error())
	}

	var mhs []mahasiswa.NilaiMhs
	for result.Next() {
		var mhsRow mahasiswa.NilaiMhs
		errScan := result.Scan(&mhsRow.Nama, &mhsRow.NamaMatkul, &mhsRow.NilaiRerata)
		if errScan != nil {
			log.Fatal("error scan", errScan.Error())
		}
		mhs = append(mhs, mhsRow)
	}
	return mhs, nil
}
