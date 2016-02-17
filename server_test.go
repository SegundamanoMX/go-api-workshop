package main

import (
	_ "bytes"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var router *gin.Engine

func initAPI() {
	router = gin.New()
	router.POST("/user", postUser)
	router.GET("/hello", getHello)
	router.GET("/hello/:name", getHello)
}

func TestPostUserReturnsWithStatusOK(t *testing.T) {
	initAPI()
	data := `{"name": "sergio", "pass":"pizza"}`
	reader := strings.NewReader(data)
	request, _ := http.NewRequest("POST", "/user", reader)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, request)
	expected := `{"status":"you are italian"}
`
	assert.Equal(t, resp.Body.String(), expected)
}

func TestPostUserReturnsWithUnauthorized(t *testing.T) {
	data := `{"name": "sergio", "pass":"tacos"}`
	reader := strings.NewReader(data)
	request, _ := http.NewRequest("POST", "/user", reader)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, request)
	expected := `{"status":"unauthorized"}
`
	assert.Equal(t, resp.Body.String(), expected)
}

func TestGetUserReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, request)
	expected := `{"message":"Hello, World!"}
`
	assert.Equal(t, resp.Body.String(), expected)
}

func TestGetUserNameReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/hello/david", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, request)
	expected := `{"message":"Hello, david"}
`
	assert.Equal(t, resp.Body.String(), expected)
}
