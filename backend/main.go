package main

import (
	"context"
	docs "github.com/astrikov-d/gpt-api/backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/sheeiavellie/go-yandexgpt"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Get response from Yandex GPT and return result or error.
func getGPTResponse(prompt string, gptContext string) (yandexgpt.YandexGPTResponse, bool) {
	client := yandexgpt.NewYandexGPTClientWithAPIKey(YandexGPTAPIKey)
	request := yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI(YandexGPTCatalogId, yandexgpt.YandexGPTModelLite),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 0.7,
			MaxTokens:   2000,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: prompt,
			},
			{
				Role: yandexgpt.YandexGPTMessageRoleSystem,
				Text: gptContext,
			},
		},
	}

	response, err := client.CreateRequest(context.Background(), request)
	if err != nil {
		panic(err)
	}

	return response, true
}

// getGPTResponseHandler godoc
// @Summary Get GPT response and return result.
// @Accept json
// @Param prompt body string true "prompt to pass to the GPT model"
// @Param context body string true "context to pass to the GPT model"
// @Produce json
// @Success 200
// @Router /api/gpt/get_response/ [post]
func getGPTResponseHandler(c *gin.Context) {
	var response yandexgpt.YandexGPTResponse
	var request getGPTResponseRequest
	var result bool

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	
	response, result = getGPTResponse(request.Prompt, request.Context)
	if result {
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.IndentedJSON(http.StatusBadRequest, response)
	}
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/"
	r.POST("/api/gpt/get_response/", getGPTResponseHandler)
	r.GET("/api/docs/", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
