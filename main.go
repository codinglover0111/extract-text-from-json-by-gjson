package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

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
	if return_value == "" {
		panic("return_value is empty \n check your rex")
	}
	if fileExists(path) {
		fmt.Print("file exists, do you want to overwrite it? (y/n): ")
		var input string
		fmt.Scanln(&input)
		if input == "y" {
			fmt.Println("overwriting file")
		} else {
			fmt.Println("not overwriting file")
			return
		}
	}
	err := os.WriteFile(path, []byte(return_value), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	flag.Parse()
	var path string = flag.Arg(0)
	var rex string = flag.Arg(1)
	var set_cpu string = flag.Arg(2)
	if set_cpu == "" {
		runtime.GOMAXPROCS(1)
	} else {
		want_cpu_num, err := strconv.Atoi(set_cpu)
		if err != nil {
			fmt.Println(err)
		}
		if want_cpu_num > runtime.NumCPU() {
			fmt.Println("set cpu as you can")
			runtime.GOMAXPROCS(runtime.NumCPU())
		} else {
			fmt.Println("set cpu as " + set_cpu)
			runtime.GOMAXPROCS(want_cpu_num)
		}
	}

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
