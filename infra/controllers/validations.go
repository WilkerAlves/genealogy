package controllers

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateId(id string) (int, error) {
	if len(strings.Trim(id, " ")) < 1 {
		return 0, fmt.Errorf("id is empty")
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("not possible convert value: %s in int. %w", id, err)
	}

	if intId < 1 {
		return 0, fmt.Errorf("id invalid")
	}

	return intId, nil
}
