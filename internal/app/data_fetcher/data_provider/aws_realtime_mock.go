package provider

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"yunji/common"
	"yunji/configs"
	"yunji/internal/pkg"

	"github.com/google/uuid"

	"yunji/internal/service/store"
)

type AWSRealtimeHackerConfig struct {
	IsAttack        bool
	AttackResource  pkg.AWSResourceType
	AttackOperation pkg.ResourceOperationType
	AttackSource    string
}

var Config = &AWSRealtimeHackerConfig{
	IsAttack:       false,
	AttackResource: pkg.AWSS3,
}

type AWSRealtimeMockProvider struct {
}

func (provider AWSRealtimeMockProvider) FetchData() error {
	store := store.NewStore(configs.Config)
	go MockHackerAttack(store)

	originDetails, err := store.Billing.Select1000ForRealtime(context.Background())
	if err != nil {
		return err
	}

	for {

		// mock interval of operations
		intervalSeed := rand.Intn(1500)
		time.Sleep(time.Duration(intervalSeed) * time.Millisecond)

		// mock distribution of operations
		opSeed := rand.Intn(1000)
		detail := originDetails[opSeed]

		realtimeEvent := common.RealtimeEvent{
			EventID:       uuid.New().String(),
			AccountID:     detail.AccountID,
			ProductCode:   detail.ProductCode,
			ProductName:   detail.ProductName,
			ProductRegion: detail.ProductRegion,
			ResourceID:    fmt.Sprintf("%s-%s", detail.ResourceID, uuid.New().String()[:5]),
			CreatedTime:   time.Now(),
			UsageType:     detail.UsageType,
			Operation:     detail.Operation,
			UsedByTag:     detail.UsedByTag,
		}

		pkg.ProviderLogger.Printf("Accessing resource detail: %v; \n", realtimeEvent)
		store.RealTime.Create(realtimeEvent)
	}
}

func MockHackerAttack(store *store.Store) {
	resource := Config.AttackResource
	originDetails, err := store.Billing.Select1000ByProductCode(context.Background(), string(Config.AttackResource))
	if err != nil {
		fmt.Print(err)
	}

	for {
		if !Config.IsAttack {
			tmp := rand.Intn(2000)
			time.Sleep(time.Duration(tmp) * time.Millisecond)
			continue
		}

		if resource != Config.AttackResource {
			resource = Config.AttackResource
			originDetails, err = store.Billing.Select1000ByProductCode(context.Background(), string(Config.AttackResource))
			if err != nil {
				fmt.Print(err)
			}
		}

		// mock interval of operations
		intervalSeed := rand.Intn(300)
		time.Sleep(time.Duration(intervalSeed) * time.Millisecond)

		// mock distribution of operations
		count := len(originDetails)
		opSeed := rand.Intn(count)
		detail := originDetails[opSeed]

		realtimeEvent := common.RealtimeEvent{
			EventID:       uuid.New().String(),
			AccountID:     detail.AccountID,
			ProductCode:   detail.ProductCode,
			ProductName:   detail.ProductName,
			ProductRegion: detail.ProductRegion,
			ResourceID:    fmt.Sprintf("%s-%s", detail.ResourceID, uuid.New().String()[:5]),
			CreatedTime:   time.Now(),
			UsageType:     detail.UsageType,
			Operation:     detail.Operation,
			UsedByTag:     detail.UsedByTag,
		}

		pkg.ProviderLogger.Printf("!!!Hacker accessing resource detail: %v; \n", realtimeEvent)
		store.RealTime.Create(realtimeEvent)
	}
}
