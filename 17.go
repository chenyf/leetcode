package main

import "fmt"

func letterCombinations(digits string) []string {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return nil
	}
	// 2345
	current := m[string(digits[0])]

	for i := range digits {
		if i == 0 {
			continue
		}
		l := m[string(digits[i])]

		result := []string{}
		for j := range current {
			for k := range l {
				result = append(result, current[j]+l[k])
			}
		}
		current = result
	}
	return current
}

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("234"))
	fmt.Println(letterCombinations("222"))
}
