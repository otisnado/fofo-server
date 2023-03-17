package utils

import (
	"log"
	"strconv"
	"strings"

	"github.com/gobwas/glob"
)

func ConvertStringToStruct(input string) []string {
	convertedString := strings.Split(input, ",")
	return convertedString
}

func ConvertStringToUintStruct(input string) []int {
	convertedStringArray := strings.Split(input, ",")
	var convertedint = []int{}

	for _, i := range convertedStringArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		convertedint = append(convertedint, j)
	}
	return convertedint
}

func MatchValidator(mustValue string, toValidateValue string) bool {
	var g glob.Glob
	log.Println("Compile value: ", mustValue)
	log.Println("Match value", toValidateValue)

	g = glob.MustCompile(mustValue)
	state := g.Match(toValidateValue)
	log.Println("Match result: ", state)

	return state
}
