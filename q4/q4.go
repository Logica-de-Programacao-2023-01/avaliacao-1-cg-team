package q4

import "errors"

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	if len(prices) == 0 {
		return 0, errors.New("a lista de preços está vazia")
	}
	if len(prices) == 1 {
		return 3, nil
	}

	isAscending := true
	isDescending := true

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			isDescending = false
		} else if prices[i] < prices[i-1] {
			isAscending = false
		}
	}

	if isAscending {
		return 1, nil
	}
	if isDescending {
		return 2, nil
	}

	return 3, nil
}
