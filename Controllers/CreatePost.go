package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ddd-structure/Application/Service"
	"io"
	"io/ioutil"
	"net/http"
)

func CreatePost(c *gin.Context) {

	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	if err != nil {
		ReturnErr(c.Writer, err)
	}
	if err := c.Request.Body.Close(); err != nil {
		ReturnErr(c.Writer, err)
	}

	var postRequest Service.PostResquest

	err = json.Unmarshal(body, &postRequest)
	if err != nil {
		fmt.Println("There was an error:", err)
	}

	var container DependencyContainer
	_ = GetContainer(&container)

	Service.CreatePost(postRequest, container.GenerateId)

	if err != nil {
		ReturnErr(c.Writer, err)
	}
	response := Response{true, "Ok"}
	MakeResponse(c.Writer, response, http.StatusOK)
}
