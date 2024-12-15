package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse_input(path string) ([]int32, []int32) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer file.Close()

	var left_list, right_list []int32
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) != 2 {
			log.Fatal("Lines must contain exactly two fields")
		}

		left_field_i, err := strconv.Atoi(words[0])
		left_field := int32(left_field_i)
		if err != nil {
			log.Fatal("Line contains non-numerical data")
		}

		right_field_i, err := strconv.Atoi(words[1])
		right_field := int32(right_field_i)
		if err != nil {
			log.Fatal("Line contains non-numerical data")
		}

		left_list = append(left_list, left_field)
		right_list = append(right_list, right_field)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left_list, right_list
}

func abs_diff(l, r int32) int32 {
	if l >= r {
		return l - r
	} else {
		return r - l
	}
}

func part_1(left_list, right_list []int32) {
	sort.Slice(left_list, func(i, j int) bool {
		return left_list[i] < left_list[j]
	})
	sort.Slice(right_list, func(i, j int) bool {
		return right_list[i] < right_list[j]
	})

	if len(left_list) != len(right_list) {
		log.Fatal("Left and right lists differ in lenght")
	}

	distance_sum := int32(0)
	for i := range left_list {
		distance_sum += abs_diff(left_list[i], right_list[i])
	}

	fmt.Println("Distance sum: ", distance_sum)
}

func part_2(left_list, right_list []int32) {
	counts := make(map[int32]int32)
	for _, value := range right_list {
		counts[value] += 1
	}

	similarity_score := int32(0)
	for _, value := range left_list {
		similarity_score += value * counts[value]
	}

	fmt.Println("Similarity score: ", similarity_score)
}

func main() {
	left_list, right_list := parse_input("./input.txt")
	part_1(left_list, right_list)
	part_2(left_list, right_list)
}
