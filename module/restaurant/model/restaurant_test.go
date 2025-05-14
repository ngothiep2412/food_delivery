package restaurantmodel

import (
	"errors"
	"testing"
)

type testData struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTable := []testData{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameIsEmpty},
		{Input: RestaurantCreate{Name: "Test"}, Expect: nil},
	}

	for _, item := range dataTable {

		err := item.Input.Validate()

		if !errors.Is(err, item.Expect) {
			t.Errorf("Validate restaurant. Input: %v. Expected: %v, Output: %v", item.Input, item.Expect, err)
		}
	}

}
