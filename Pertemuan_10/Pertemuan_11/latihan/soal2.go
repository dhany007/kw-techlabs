package latihan

func potongArray(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)

	return arr
}

func Pertarungan(petarung []int, kekuatan []int, self int) int {
	selesai := false
	powerAkhir := self
	tempPetarung := petarung
	tempKekuatan := kekuatan

	for !selesai {
		lenPetarung := len(tempPetarung)
		for i, v := range tempPetarung {
			if powerAkhir >= v {
				powerAkhir += tempKekuatan[i]
				tempPetarung = potongArray(tempPetarung, i)
				tempKekuatan = potongArray(tempKekuatan, i)
				lenPetarung += 1
				break
			}
		}

		if lenPetarung == len(tempPetarung) {
			selesai = true
		}
	}

	return powerAkhir
}
