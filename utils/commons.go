package utils

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gobwas/glob"
)

var g glob.Glob

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

	g = glob.MustCompile(mustValue)
	state := g.Match(toValidateValue)

	return state
}

func TmpFolderCreation(folderName string) (folderPath string, err error) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	folderToCreate := path + "/" + folderName
	err = os.MkdirAll(folderToCreate, 0755)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return folderToCreate, nil
}
