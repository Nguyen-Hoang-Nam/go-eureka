package discovery

import (
	eureka "github.com/xuanbo/eureka-client"
)

var client *eureka.Client

func Init() {
	client = eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://localhost:9000/eureka/",
		App:                   "Gin",
		Port:                  8002,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
		Metadata: map[string]interface{}{
			"VERSION":              "1.0.0",
			"NODE_GROUP_ID":        0,
			"PRODUCT_CODE":         "DEFAULT",
			"PRODUCT_VERSION_CODE": "DEFAULT",
			"PRODUCT_ENV_CODE":     "DEFAULT",
			"SERVICE_VERSION_CODE": "DEFAULT",
		},
	})

	client.Start()
}

func GetServices() *eureka.Client {
	return client
}
