package main

import (
	. "../util"
	"fmt"
	"sort"
	"strings"
)

type Food struct {
	ingredients StringSet
	allergens   StringSet
}

type DataType struct {
	foods        []Food
	allAllergens StringSet
}

func parseData() DataType {
	data := FetchInputData(21)
	dataSplit := strings.Split(data, "\n")

	foods := make([]Food, len(dataSplit))
	for i, line := range dataSplit {
		part := strings.Split(line[:len(line)-1], " (contains ")

		ingredients := NewStringSet(strings.Split(part[0], " "))
		allergens := NewStringSet(strings.Split(part[1], ", "))
		foods[i] = Food{ingredients, allergens}
	}

	allAllergens := NewStringSet([]string{})
	for _, food := range foods {
		for allergen := range food.allergens {
			allAllergens.Add(allergen)
		}
	}

	return DataType{foods, allAllergens}
}

func validate(data map[string]string, foods []Food) bool {
	for _, food := range foods {
		for allergen := range food.allergens {
			if _, ok := data[allergen]; ok {
				if !food.ingredients.Contains(data[allergen]) {
					return false
				}
			}
		}
	}

	return true
}

func solvePart(foods []Food, initTmp map[string]string) []map[string]string {
	initTmpValues := MapValuesAsStringSet(initTmp)
	initTmpKeys := MapKeysAsStringSet(initTmp)

	newFoods := make([]Food, len(foods))
	for i, food := range foods {
		newIngredients := food.ingredients.Difference(&initTmpValues)
		newAllergens := food.allergens.Difference(&initTmpKeys)
		newFoods[i] = Food{newIngredients, newAllergens}
	}

	foodWithMaxAllergens := newFoods[0]
	for _, food := range newFoods {
		if len(foodWithMaxAllergens.allergens) < len(food.allergens) {
			foodWithMaxAllergens = food
		}
	}

	ruleWithMaxAllergensIngredientsKeys := foodWithMaxAllergens.ingredients.AsSlice()
	ruleWithMaxAllergensAllergensKeys := foodWithMaxAllergens.allergens.AsSlice()

	result := make([]map[string]string, 0)
	for _, p := range StringPermutations(ruleWithMaxAllergensIngredientsKeys, len(ruleWithMaxAllergensAllergensKeys)) {
		tmp := make(map[string]string, len(p))
		for i := 0; i < len(p); i++ {
			tmp[ruleWithMaxAllergensAllergensKeys[i]] = p[i]
		}

		if validate(tmp, newFoods) {
			for initTmpKey, initTmpValue := range initTmp {
				tmp[initTmpKey] = initTmpValue
			}
			result = append(result, tmp)
		}
	}
	return result
}

func solve(data DataType) map[string]string {
	result := make([]map[string]string, 1)
	result[0] = make(map[string]string)

	for len(result[0]) < len(data.allAllergens) {
		tmp := make([]map[string]string, 0)
		for _, r := range result {
			for _, el := range solvePart(data.foods, r) {
				tmp = append(tmp, el)
			}
		}
		result = tmp
	}

	return result[0]
}

func solvePart1(data DataType) (rc int) {
	myList := solve(data)
	myListValues := MapValuesAsStringSet(myList)

	for _, food := range data.foods {
		rc += len(food.ingredients.Difference(&myListValues))
	}

	return
}

func solvePart2(data DataType) (rc string) {
	myList := solve(data)
	myListKeys := MapKeysAsStringSet(myList)
	myListKeysAsSlice := myListKeys.AsSlice()
	sort.Strings(myListKeysAsSlice)

	sortedValues := make([]string, len(myListKeysAsSlice))
	for i, k := range myListKeysAsSlice {
		sortedValues[i] = myList[k]
	}

	return strings.Join(sortedValues, ",")
}

func main() {
	data := parseData()
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
