package assigment1

type Teman struct {
	Nama                 string
	Alamat               string
	Pekerjaan            string
	Alasan_memilih_kelas string
}

func GetAbsenceAttendece(absen int) Teman {
	dataTeman := map[int]Teman{
		1: {"Joko", "Jakarta", "Software Engineer", "Saya ingin memperluas pengetahuan saya tentang bahasa pemrograman dan Golang menarik perhatian saya"},
		2: {"Siti", "Bandung", "Data Scientist", "Golang menawarkan kemudahan dalam penulisan kode yang efisien, dan saya tertarik dengan kemampuannya dalam pemrograman konkurensi."},
		3: {"Sabil", "Surabaya", "UI/UX Designer", "Saya ingin mendalami aspek pengembangan aplikasi web dan Golang adalah pilihan yang tepat untuk itu."},
		4: {"firman", "Medan", "System Administrator", "Saya tertarik dengan kehandalan Golang dalam membangun sistem terdistribusi."},
		5: {"Nanang", "Semarang", "Full Stack Developer", "Saya ingin mempelajari Golang untuk meningkatkan keahlian saya dalam pengembangan perangkat lunak."},
	}

	return dataTeman[absen]
}
