package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mlm.gin/discovery"
	"mlm.gin/models"
)

func Hi(c *gin.Context) {
	var json models.Request

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can not parse json",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Gin hi, " + json.Name,
		})
	}
}

func Test(_ *gin.Context) {
	client := discovery.GetServices()
	apps := client.Applications.Applications
	fmt.Printf("%+v", apps)
}

func Hello(c *gin.Context) {
	framework := c.Param("framework")

	client := discovery.GetServices()
	apps := client.Applications.Applications
	frameworkCapital := strings.ToUpper(framework)
	foundService := false

	var hostname string
	var port int
	for _, app := range apps {
		if app.Name == frameworkCapital {
			foundService = true
			hostname = app.Instances[0].HostName
			port = app.Instances[0].Port.Port
		}
	}

	if !foundService {
		c.String(http.StatusInternalServerError, "Can not found service "+framework)
		return
	}

	postBody, err := json.Marshal(map[string]string{
		"name": "Gin",
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Can not send request to "+framework)
		return
	}

	responseBody := bytes.NewBuffer(postBody)
	res, err := http.Post("http://"+hostname+":"+strconv.Itoa(port), "application/json", responseBody)
	if err != nil {
		c.String(http.StatusInternalServerError, "Can not send request to "+framework)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Can not send request to "+framework)
		return
	}

	c.String(http.StatusOK, string(body))
}
