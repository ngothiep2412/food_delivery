package restaurantbiz

import (
	"errors"
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	"golang.org/x/net/context"
	"testing"
)

type mokeCreateStore struct{}

func (mokeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Test" {
		return common.ErrDB(errors.New("something went wrong in DB"))
	}

	data.Id = 200

	return nil
}

func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurant(mokeCreateStore{})

	dataTest := restaurantmodel.RestaurantCreate{Name: ""}

	err := biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil || err.Error() != "Invalid request" {
		t.Errorf("Failed")

	}

	dataTest2 := restaurantmodel.RestaurantCreate{Name: "Test"}
	err = biz.CreateRestaurant(context.Background(), &dataTest2)

	if err != nil {
		t.Errorf("Failed")
	}

	dataTest3 := restaurantmodel.RestaurantCreate{Name: "Test-2"}

	err = biz.CreateRestaurant(context.Background(), &dataTest3)

	if err != nil {
		t.Errorf("Failed")
	}

	//t.Log("TestNewCreateRestaurantBiz Passed")

}
