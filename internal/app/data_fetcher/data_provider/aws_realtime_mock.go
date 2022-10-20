package provider

import (
	"math/rand"
	"time"
	"yunji/internal/pkg"

	"github.com/google/uuid"
)

type AWSRealtimeHackerConfig struct {
	IsAttack        bool
	AttackResource  pkg.AWSResourceType
	AttackOperation pkg.ResourceOperationType
	AttackSource    string
}

var Config = &AWSRealtimeHackerConfig{
	IsAttack:        true,
	AttackResource:  pkg.AWSEC2,
	AttackOperation: pkg.Create,
	AttackSource:    "hacker_ip",
}

type AWSRealtimeMockProvider struct {
}

func (provider AWSRealtimeMockProvider) FetchData() error {
	if Config.IsAttack {
		go MockHackerAttack()
	}

	for {
		// mock interval of operations
		intervalSeed := rand.Intn(1500)
		time.Sleep(time.Duration(intervalSeed) * time.Millisecond)

		// mock distribution of operations
		opSeed := rand.Intn(10)
		operation := pkg.Read
		if opSeed == 9 {
			operation = pkg.Delete
		} else if opSeed >= 7 {
			operation = pkg.Create
		} else if opSeed >= 5 {
			operation = pkg.Update
		}

		resourceSeed := rand.Intn(5)
		resourceType := pkg.AWSResourceTypes[resourceSeed]

		resourceID := uuid.New()

		pkg.ProviderLogger.Printf("Someone %s resource %s of ID: %s at %s; \n", operation, resourceType, resourceID, time.Now().String())
	}
}

func MockHackerAttack() {
	for {
		// mock interval of operations
		intervalSeed := rand.Intn(300)
		time.Sleep(time.Duration(intervalSeed) * time.Millisecond)

		// mock distribution of operations
		operation := Config.AttackOperation

		resourceType := Config.AttackResource

		resourceID := uuid.New()

		pkg.ProviderLogger.Printf("%s %s resource %s of ID: %s at %s; \n", Config.AttackSource, operation, resourceType, resourceID, time.Now().String())
	}
}
