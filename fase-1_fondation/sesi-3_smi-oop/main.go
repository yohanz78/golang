package main

import "fmt"

// 1. Interface: Mendefinisikan "kontrak" bahwa objek apapun yang
// memiliki method Describe() bisa dianggap sebagai tipe "Activity"
type Activity interface {
	Describe() string
}

// 2. Struct: Mendefinisikan cetak biru (blueprint) untuk objek Task
type Task struct {
	Name     string
	Duration int
	Priority string
}

// Struct lain untuk contoh variasi data
type BreakTime struct {
	Type     string
	Duration int
}

// 3. Method untuk Task: Fungsi ini menempel pada struct Task (menggunakan receiver 't Task')
func (t Task) Describe() string {
	return fmt.Sprintf("[TUGAS] %s - Prioritas %s (%d menit)", t.Name, t.Priority, t.Duration)
}

// Method untuk BreakTime: Fungsi ini menempel pada struct BreakTime
func (b BreakTime) Describe() string {
	return fmt.Sprintf("[ISTIRAHAT] Waktunya %s selama %d menit", b.Type, b.Duration)
}

// 4. Fungsi umum yang menerima Interface sebagai parameter
// Fungsi ini tidak peduli apakah objeknya Task atau BreakTime,
// asalkan objek tersebut punya method Describe()
func PrintActivityInfo(a Activity) {
	fmt.Println(a.Describe())
}

func main() {
	fmt.Println("=== Routine Integrator (Struct & Interface) ===")

	// Membuat instance (objek) dari struct Task
	codingTask := Task{
		Name:     "Learn Golang Structs",
		Duration: 120,
		Priority: "High",
	}

	readingTask := Task{
		Name:     "Read Clean Code Book",
		Duration: 45,
		Priority: "Medium",
	}

	// Membuat instance dari struct BreakTime
	coffeeBreak := BreakTime{
		Type:     "Coffee & Stretching",
		Duration: 15,
	}

	// Memanggil fungsi PrintActivityInfo
	// Walaupun codingTask, readingTask, dan coffeeBreak adalah tipe struct yang berbeda,
	// semuanya bisa dimasukkan ke parameter Interface 'Activity'
	PrintActivityInfo(codingTask)
	PrintActivityInfo(readingTask)
	PrintActivityInfo(coffeeBreak)

	// Bonus: Mengakses field struct secara langsung
	fmt.Printf("\nTarget utama hari ini adalah '%s' selama %d menit.\n", codingTask.Name, codingTask.Duration)
}
