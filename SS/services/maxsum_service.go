// Service onde executa todas as funções

package services

import (
	"SS/domain/maxsum"
	"SS/utils/errors/rest_errors"
	"sort"
)

var (
	MaxSumService maxSumInterface = &maxSumService{}
)

type maxSumService struct {
}

type maxSumInterface interface {
	MaxSum(maxsum.Numbers) (*maxsum.Result, *rest_errors.RestErr)
}

func (s *maxSumService) MaxSum(numbers maxsum.Numbers) (*maxsum.Result, *rest_errors.RestErr) {
	var result maxsum.Result
	var resultArray maxsum.Result

	last_current_sum := 0

	// Algoritmo de Kadane - SOMA
	for _, element := range numbers.List {
		current_sum := max(last_current_sum+element, element)
		max_sum := max(current_sum, last_current_sum)
		last_current_sum = current_sum
		result.Sum = max_sum
		resultArray.Positions = append(resultArray.Positions, current_sum)
	}

	// Descobrir posições

	// Primeiro descobrimos a posição do último número a ser somado
	var positionMax int
	for i, element := range resultArray.Positions {
		if result.Sum == element {
			positionMax = i
			break
		}
	}

	// Adicionamos a posição do último número a ser somado ao array
	result.Positions = append(result.Positions, positionMax+1)

	// Depois adicionamos as posições dos números anteriores
	for i := positionMax; i > 0; i-- {
		if resultArray.Positions[i] == numbers.List[i]+resultArray.Positions[i-1] {
			result.Positions = append(result.Positions, i)
		} else {
			break
		}
	}

	sort.Ints(result.Positions)
	return &result, nil
}

// função para ver o valor máximo entre dois números
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
