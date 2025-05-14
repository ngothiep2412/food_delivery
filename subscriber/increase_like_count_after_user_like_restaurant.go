package subscriber

import (
	"g05-food-delivery/component/appctx"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"g05-food-delivery/pubsub"
	"golang.org/x/net/context"
	"log"
)

type HasRestaurantId interface { // interface là con trỏ
	GetRestaurantId() int
	GetUserId() int
}

//func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
//	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)
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
//			likeData, ok := msg.Data().(HasRestaurantId)
//			if !ok {
//				log.Printf("Error: cannot assert data to HasRestaurantId, got %T", msg.Data())
//				continue
//			}
//
//			log.Printf("Increasing like count for restaurant ID: %d", likeData.GetRestaurantId())
//			if err := store.IncreaseLikeCount(ctx, likeData.GetRestaurantId()); err != nil {
//				log.Printf("Error increasing like count: %v", err)
//			} else {
//				log.Printf("Successfully increased like count for restaurant ID: %d", likeData.GetRestaurantId())
//			}
//		}
//	}()
//}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase count like after user like restaurant ",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
			likeData := message.Data().(HasRestaurantId)

			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Push notification when/user like restaurant ",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)
			log.Println("Push notification when user likes restaurant id:", likeData.GetRestaurantId())

			// Nếu có handler token,.. -> nên xài pubsub khác để lắng nghe

			return nil
		},
	}
}

func EmitIncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Realtime emit after user likes restaurant ",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := message.Data().(HasRestaurantId)

			appCtx.GetRealtimeEngine().EmitToUser(likeData.GetUserId(), string(message.Channel()), likeData)

			return nil
		},
	}
}
