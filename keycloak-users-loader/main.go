package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var conf config

func main() {
	loadConfiguration()
	token := authenticateToKeycloak()
	content := readCsvFile(conf.Input.File)

	for _,line := range content {
		id := line[0]
		firstName := line[1]
		lastName := line[2]
		email := line[3]
		fmt.Println("Creating user "+id)
		createUser(token, firstName, lastName, email, id)
	}	
}

func loadConfiguration() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetConfigName("config.local")
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
}

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func requiredActions() *[]string {
	return &[]string{"VERIFY_EMAIL",  "UPDATE_PASSWORD"};
}
