/*
 * Datamonkey API
 *
 * Datamonkey is a free public server for comparative analysis of sequence alignments using state-of-the-art statistical models. <br> This is the OpenAPI definition for the Datamonkey API. 
 *
 * API version: 1.0.0
 * Contact: spond@temple.edu
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datamonkey

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name		string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method		string
	// Pattern is the pattern of the URI.
	Pattern	 	string
	// HandlerFunc is the handler function of this route.
	HandlerFunc	gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouter add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the ABSRELAPI part of the API
	ABSRELAPI ABSRELAPI
	// Routes for the BGMAPI part of the API
	BGMAPI BGMAPI
	// Routes for the BUSTEDAPI part of the API
	BUSTEDAPI BUSTEDAPI
	// Routes for the CONTRASTFELAPI part of the API
	CONTRASTFELAPI CONTRASTFELAPI
	// Routes for the FELAPI part of the API
	FELAPI FELAPI
	// Routes for the FUBARAPI part of the API
	FUBARAPI FUBARAPI
	// Routes for the FileUploadAndQCAPI part of the API
	FileUploadAndQCAPI FileUploadAndQCAPI
	// Routes for the GARDAPI part of the API
	GARDAPI GARDAPI
	// Routes for the HealthAPI part of the API
	HealthAPI HealthAPI
	// Routes for the MEMEAPI part of the API
	MEMEAPI MEMEAPI
	// Routes for the MULTIHITAPI part of the API
	MULTIHITAPI MULTIHITAPI
	// Routes for the RELAXAPI part of the API
	RELAXAPI RELAXAPI
	// Routes for the SLACAPI part of the API
	SLACAPI SLACAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{ 
		{
			"GetABSRELJob",
			http.MethodPost,
			"/api/v1/methods/absrel-result",
			handleFunctions.ABSRELAPI.GetABSRELJob,
		},
		{
			"StartABSRELJob",
			http.MethodPost,
			"/api/v1/methods/absrel-start",
			handleFunctions.ABSRELAPI.StartABSRELJob,
		},
		{
			"GetBGMJob",
			http.MethodPost,
			"/api/v1/methods/bgm-result",
			handleFunctions.BGMAPI.GetBGMJob,
		},
		{
			"StartBGMJob",
			http.MethodPost,
			"/api/v1/methods/bgm-start",
			handleFunctions.BGMAPI.StartBGMJob,
		},
		{
			"GetBUSTEDJob",
			http.MethodPost,
			"/api/v1/methods/busted-result",
			handleFunctions.BUSTEDAPI.GetBUSTEDJob,
		},
		{
			"StartBUSTEDJob",
			http.MethodPost,
			"/api/v1/methods/busted-start",
			handleFunctions.BUSTEDAPI.StartBUSTEDJob,
		},
		{
			"GetCONTRASTFELJob",
			http.MethodPost,
			"/api/v1/methods/contrast-fel-result",
			handleFunctions.CONTRASTFELAPI.GetCONTRASTFELJob,
		},
		{
			"StartCONTRASTFELJob",
			http.MethodPost,
			"/api/v1/methods/contrast-fel-start",
			handleFunctions.CONTRASTFELAPI.StartCONTRASTFELJob,
		},
		{
			"GetFELJob",
			http.MethodPost,
			"/api/v1/methods/fel-result",
			handleFunctions.FELAPI.GetFELJob,
		},
		{
			"StartFELJob",
			http.MethodPost,
			"/api/v1/methods/fel-start",
			handleFunctions.FELAPI.StartFELJob,
		},
		{
			"GetFUBARJob",
			http.MethodPost,
			"/api/v1/methods/fubar-result",
			handleFunctions.FUBARAPI.GetFUBARJob,
		},
		{
			"StartFUBARJob",
			http.MethodPost,
			"/api/v1/methods/fubar-start",
			handleFunctions.FUBARAPI.StartFUBARJob,
		},
		{
			"GetDatasetsList",
			http.MethodGet,
			"/api/v1/datasets",
			handleFunctions.FileUploadAndQCAPI.GetDatasetsList,
		},
		{
			"PostDataset",
			http.MethodPost,
			"/api/v1/datasets",
			handleFunctions.FileUploadAndQCAPI.PostDataset,
		},
		{
			"GetGARDJob",
			http.MethodPost,
			"/api/v1/methods/gard-result",
			handleFunctions.GARDAPI.GetGARDJob,
		},
		{
			"StartGARDJob",
			http.MethodPost,
			"/api/v1/methods/gard-start",
			handleFunctions.GARDAPI.StartGARDJob,
		},
		{
			"GetHealth",
			http.MethodGet,
			"/api/v1/health",
			handleFunctions.HealthAPI.GetHealth,
		},
		{
			"GetMEMEJob",
			http.MethodPost,
			"/api/v1/methods/meme-result",
			handleFunctions.MEMEAPI.GetMEMEJob,
		},
		{
			"StartMEMEJob",
			http.MethodPost,
			"/api/v1/methods/meme-start",
			handleFunctions.MEMEAPI.StartMEMEJob,
		},
		{
			"GetMULTIHITJob",
			http.MethodPost,
			"/api/v1/methods/multihit-result",
			handleFunctions.MULTIHITAPI.GetMULTIHITJob,
		},
		{
			"StartMULTIHITJob",
			http.MethodPost,
			"/api/v1/methods/multihit-start",
			handleFunctions.MULTIHITAPI.StartMULTIHITJob,
		},
		{
			"GetRELAXJob",
			http.MethodPost,
			"/api/v1/methods/relax-result",
			handleFunctions.RELAXAPI.GetRELAXJob,
		},
		{
			"StartRELAXJob",
			http.MethodPost,
			"/api/v1/methods/relax-start",
			handleFunctions.RELAXAPI.StartRELAXJob,
		},
		{
			"GetSLACJob",
			http.MethodPost,
			"/api/v1/methods/slac-result",
			handleFunctions.SLACAPI.GetSLACJob,
		},
		{
			"StartSLACJob",
			http.MethodPost,
			"/api/v1/methods/slac-start",
			handleFunctions.SLACAPI.StartSLACJob,
		},
	}
}
