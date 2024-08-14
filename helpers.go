package main

import (
	"strconv"
	"strings"

	"github.com/krzemLech/go-todo-app/config"
)

func formatFilters(pageParam, perPageParam string) (int64, int64, error) {
	var err error = nil
	var page, perPage int
	page, err = strconv.Atoi(pageParam)
	perPage, err = strconv.Atoi(perPageParam)
	if err != nil {
		return 1, 5, err
	}
	return int64(page), int64(perPage), nil
}

func checkProfane(todo string) bool {
	profaneWords := strings.Split(config.Envs.ProfaneWords, ",")
	for _, word := range profaneWords {
		if strings.Contains(todo, word) {
			return true
		}
	}
	return false
}