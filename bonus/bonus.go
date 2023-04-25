package bonus

import "sort"

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	if len(barLengths) < 1 {
		return 0, 0
	}
	sort.Ints(barLengths)
	maxHeight := 0
	numTowers := 1
	for i := 1; i < len(barLengths); i++ {
		if barLengths[i] != barLengths[i-1] {

			numTowers++
		}
		if barLengths[i] > maxHeight {
			maxHeight = barLengths[i]
		}
	}
	numTowers++
	return numTowers, maxHeight
}
