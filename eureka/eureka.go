package eureka

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var serviceInstaces map[string]map[string]Instance = make(map[string]map[string]Instance)

func mapToListInstances(mapInstance map[string]Instance) []Instance {
	var instances []Instance
	for _, instance := range mapInstance {
		instances = append(instances, instance)
	}

	return instances
}

func mapToArray() Applications {
	var applications []Application

	for name, rawInstances := range serviceInstaces {
		application := &Application{
			Name:      name,
			Instances: mapToListInstances(rawInstances),
		}

		applications = append(applications, *application)
	}

	return Applications{
		Applications: applications,
	}
}

func addInstance(c *gin.Context) {
	appId := c.Param("appId")

	var json NewInstance
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err)
		return
	}

	instance := json.Instance

	if mapInstance, ok := serviceInstaces[strings.ToUpper(appId)]; ok {
		mapInstance[instance.InstanceId] = instance
		serviceInstaces[strings.ToUpper(appId)] = mapInstance
	} else {
		mapInstance := make(map[string]Instance)
		mapInstance[instance.InstanceId] = instance
		serviceInstaces[strings.ToUpper(appId)] = mapInstance
	}

	c.JSON(204, gin.H{})
}

func removeInstance(c *gin.Context) {
	appId := c.Param("appId")
	instanceId := c.Param("instanceId")

	delete(serviceInstaces[strings.ToUpper(appId)], instanceId)
}

func sendHeartbeat(c *gin.Context) {
	appId := c.Param("appId")
	instanceId := c.Param("instanceId")
	status := c.Query("status")

	instance := serviceInstaces[strings.ToUpper(appId)][instanceId]
	instance.Status = status

	serviceInstaces[strings.ToUpper(appId)][instanceId] = instance
}

func getAllService(c *gin.Context) {
	c.JSON(200, gin.H{
		"applications": mapToArray(),
	})
}

func getAllInstance(c *gin.Context) {
	appId := c.Param("appId")

	if mapInstance, ok := serviceInstaces[strings.ToUpper(appId)]; ok {
		var instances []Instance
		for _, instance := range mapInstance {
			instances = append(instances, instance)
		}

		c.JSON(200, gin.H{
			"name":      appId,
			"instances": mapToListInstances(mapInstance),
		})
	} else {
		c.JSON(204, gin.H{})
	}
}

func getInstace(c *gin.Context) {
	appId := c.Param("appId")
	instanceId := c.Param("instanceId")

	c.JSON(200, gin.H{
		"instance": serviceInstaces[appId][instanceId],
	})
}

func updateMetadata(c *gin.Context) {
	appId := c.Param("appId")
	instanceId := c.Param("instanceId")
	key := c.Query("key")

	fmt.Println(appId + instanceId)
	fmt.Println(key)
}

func getInstanceAllService(c *gin.Context) {
	instanceId := c.Param("instanceId")

	fmt.Println(instanceId)
}

func updateInstanceStatus(c *gin.Context) {
	appId := c.Param("appId")
	instanceId := c.Param("instanceId")
	value := c.Query("value")

	fmt.Println(appId + instanceId)
	fmt.Println(value)
}

func Run(port string) {
	router := gin.Default()

	router.POST("/eureka/apps/:appId", addInstance)

	router.DELETE("/eureka/app/:appId/:instanceId", removeInstance)

	router.PUT("/eureka/apps/:appId/:instanceId", sendHeartbeat)

	router.GET("/eureka/apps", getAllService)

	router.GET("/eureka/apps/:appId", getAllInstance)

	router.GET("/eureka/apps/:appId/:instanceId", getInstace)

	router.PUT("/eureka/apps/:appId/:instanceId/metadata", updateMetadata)

	router.GET("/eureka/instances/:instanceId", getInstanceAllService)

	router.PUT("/eureka/apps/:appId/:instanceId/status", updateInstanceStatus)

	router.Run(":" + port)
}
