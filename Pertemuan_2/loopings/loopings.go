package loopings

import "fmt"

func Loopings() {
	fmt.Println("loopings first way")

	for i := 0; i < 3; i++ {
		fmt.Println("Angka ke -", i)
	}

	fmt.Println("loopings second way")
	j := 0
	for j < 3 {
		fmt.Println("Angka ke -", j)
		j++
	}

	fmt.Println("loopings third way")

	j = 0
	for {
		fmt.Println("Angka ke -", j)
		j++
		if j == 3 {
			break
		}
	}
}

func LoopingsBreakContinue() {
	fmt.Println("looping break and continue keyword")
	// break => menghentikan perulangan
	// continue => melanjutkan perulangan tanpa mengecek baris kode selanjutnya

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}

		fmt.Println("Angka ke -", i)
	}
}

func NestedLoopings() {
	fmt.Println("nested loopings")

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func LoopingsLabel() {
	fmt.Println("looping label")
	// fungsi break hanya bisa menghentikan perulangan pada scope nya
	// jika dibuat break label, maka akan menghentikan perulangan dimana  label nya ditempatkan

outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke -", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}

func LatihanLoopingsDeretPrima() {
	fmt.Println("latihan menampilkan deret bilangan prima")

	max := 50

	for i := 2; i <= max; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i, " ")
		}
	}
}
