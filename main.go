package main

import (
	"fmt"
	"strings"
)

// Определяем структуру CacheValue
type CacheValue struct {
	Imsi     string
	LacCell  string
	SectorID int
	Service  string
}

func main() {
	// Мапа cache остается неизменной
	cache := make(map[string]string)
	retryBlock := make(map[string]CacheValue) // Теперь мапа хранит CacheValue

	// Заполняем cache
	cache["77014151777"] = ".."
	cache["77014151778"] = ".."
	cache["77014151779"] = ".."

	// Заполняем retryBlock
	retryBlock["77014151777.VS"] = CacheValue{Imsi: "imsi1", LacCell: "lac1", SectorID: 1, Service: "VS"}
	retryBlock["77014151777.DT"] = CacheValue{Imsi: "imsi2", LacCell: "lac2", SectorID: 2, Service: "DT"}

	// Результирующая мапа
	result := make(map[string]string)

	// Перебор cache и проверка на наличие в retryBlock
	for msisdn, val := range cache {
		if hasRetryBlock(msisdn, retryBlock) {
			continue
		}
		result[msisdn] = val
	}

	fmt.Println("Result:", result)
}

// Проверяем наличие retryBlock с учетом нового типа CacheValue
func hasRetryBlock(msisdn string, retryBlock map[string]CacheValue) bool {
	for key := range retryBlock {
		fmt.Println("Checking key:", key)
		if strings.HasPrefix(key, msisdn+".") {
			fmt.Println(key, msisdn)
			return true
		}
	}
	return false
}
