package subscriber

import (
	"g05-food-delivery/component/appctx"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"g05-food-delivery/pubsub"
	"golang.org/x/net/context"
)

//func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserUnlikeRestaurant)
//
//	store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
//
//	go func() {
//		defer common.AppRecover()
//
//		for {
//			msg := <-c
//			log.Printf("Received message with data: %+v", msg.Data())
//
//			unlikeData, ok := msg.Data().(HasRestaurantId)
//			if !ok {
//				log.Printf("Error: cannot assert data to HasRestaurantId, got %T", msg.Data())
//				continue
//			}
//
//			log.Printf("Decreasing like count for restaurant ID: %d", unlikeData.GetRestaurantId())
//			if err := store.DecreaseLikeCount(ctx, unlikeData.GetRestaurantId()); err != nil {
//				log.Printf("Error decreasing like count: %v", err)
//			} else {
//				log.Printf("Successfully decreased like count for restaurant ID: %d", unlikeData.GetRestaurantId())
//			}
//		}
//	}()
//}

func DecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like after user like restaurant ",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			unlikeData := message.Data().(HasRestaurantId)

			return store.DecreaseLikeCount(ctx, unlikeData.GetRestaurantId())
		},
	}
}
