package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func path_to_value(path string, rex string) []string {
	file_path := strings.ReplaceAll(path, "\\", "/")
	file_path = strings.ReplaceAll(file_path, "\"", "")
	file, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
	}
	var file_content string = string(file)

	var file_content_gjsonGet []gjson.Result = gjson.Get(file_content, rex).Array()

	var return_value []string
	for _, value := range file_content_gjsonGet {
		return_value = append(return_value, value.String())
	}
	return return_value
}

func string_array_to_file(string_array []string, path string) {
	var return_value string
	for _, value := range string_array {
		return_value += value + "\n"
	}
	err := os.WriteFile(path, []byte(return_value), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// func string_Array_to_file(byte_array []byte, path string) {
// 	file_path := strings.ReplaceAll(path, "\\", "/")
// 	file_path = strings.ReplaceAll(file_path, "\"", "")
// 	var return_value []byte
// 	for _, value := range byte_array {
// 		valueBytes := []byte{value}
// 		valueBytes = append(valueBytes, []byte("\n")...)
// 		return_value = append(return_value, valueBytes...)
// 	}
// 	err := os.WriteFile(file_path, return_value, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 4)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	flag.Parse()
	var path string = flag.Arg(0)
	var rex string = flag.Arg(1)
	if path == "" || rex == "" {
		fmt.Println("path or rex is empty")
		return
	}
	var return_value_string []string = path_to_value(path, rex)

	var save_file_name string = strings.Split(path, "\\")[len(strings.Split(path, "\\"))-1]
	save_file_name = strings.ReplaceAll(save_file_name, ".json", "")

	var extract_part string = strings.Split(rex, ".")[len(strings.Split(rex, "."))-1]
	var save_path string = "./" + save_file_name + "_" + extract_part + ".txt"

	string_array_to_file(return_value_string, save_path)
}
