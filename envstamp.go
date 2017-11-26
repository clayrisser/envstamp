package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must supply file path as argument")
	}
	filePath := os.Args[1]
	template := loadTemplate(filePath)
	content := setEnvsInTemplate(template)
	writeFile(content, filePath)
}

func loadTemplate(filePath string) string {
	template, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(template[:])
}

func writeFile(content string, filePath string) {
	data := []byte(content)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}
	mode := fileInfo.Mode()
	err = ioutil.WriteFile(filePath, data, mode)
	if err != nil {
		log.Fatal(err)
	}
}

func setEnvsInTemplate(template string) string {
	matches := getMatchesFromTemplate(template)
	for _, match := range matches {
		env := getEnvFromMatch(match)
		template = strings.Replace(template, match, env, -1)
	}
	return template
}

func getMatchesFromTemplate(template string) []string {
	r, err := regexp.Compile("\\${((\\\\})|[^}])*}")
	if err != nil {
		log.Fatal(err)
	}
	return r.FindAllString(template, -1)
}

func getEnvFromMatch(match string) string {
	envDefault := ""
	r, err := regexp.Compile("[^\\\\]:")
	if err != nil {
		log.Fatal(err)
	}
	envName := match[2 : len(match)-1]
	if r.Match([]byte(match)) {
		index := r.FindIndex([]byte(match))[0]
		envDefault = envName[index:]
		envName = envName[:index-1]
	}
	env := os.Getenv(envName)
	if len(env) <= 0 {
		return envDefault
	}
	return env
}
