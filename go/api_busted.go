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
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BUSTEDAPI struct {
	HyPhyBaseAPI
}

// NewBUSTEDAPI creates a new BUSTEDAPI instance
func NewBUSTEDAPI(basePath, hyPhyPath string, scheduler SchedulerInterface, datasetTracker DatasetTracker) *BUSTEDAPI {
	return &BUSTEDAPI{
		HyPhyBaseAPI: NewHyPhyBaseAPI(basePath, hyPhyPath, scheduler, datasetTracker),
	}
}

// GetBUSTEDJob retrieves the status and results of a BUSTED job
func (api *BUSTEDAPI) GetBUSTEDJob(c *gin.Context) {
	var request BustedRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse job configuration"})
		return
	}

	adapted, err := AdaptRequest(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to adapt request: %v", err)})
		return
	}

	result, err := api.HandleGetJob(c, adapted, MethodBUSTED)
	if err != nil {
		if err.Error() == "job is not complete" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Parse the raw JSON results into BustedResult
	resultMap := result.(map[string]interface{})

	// Log the raw results for debugging
	rawResults := resultMap["results"].(json.RawMessage)
	log.Printf("Raw results: %s", string(rawResults))

	// Check if the raw results are valid JSON
	var testMap map[string]interface{}
	if err := json.Unmarshal(rawResults, &testMap); err != nil {
		log.Printf("Raw results are not valid JSON: %v", err)
	} else {
		log.Printf("Raw results are valid JSON with %d top-level keys", len(testMap))
		for k := range testMap {
			log.Printf("Found top-level key: %s", k)
		}
	}

	// Create a wrapper structure to match the expected format
	wrappedJSON := fmt.Sprintf(`{"job_id":"test_job","result":%s}`, string(rawResults))
	log.Printf("Wrapped JSON: %s", wrappedJSON)

	var bustedResult BustedResult
	if err := json.Unmarshal([]byte(wrappedJSON), &bustedResult); err != nil {
		log.Printf("Error unmarshaling wrapped results: %v", err)
		// Try to unmarshal as a generic map to see what's in there
		var resultAsMap map[string]interface{}
		if mapErr := json.Unmarshal(rawResults, &resultAsMap); mapErr != nil {
			log.Printf("Error unmarshaling as map: %v", mapErr)
		} else {
			log.Printf("Results as map: %+v", resultAsMap)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to parse results: %v", err)})
		return
	}

	// Log the parsed result structure
	log.Printf("Parsed BustedResult: %+v", bustedResult)

	// Update the results in the resultMap
	resultMap["results"] = bustedResult.Result

	c.JSON(http.StatusOK, resultMap)
}

// StartBUSTEDJob starts a new BUSTED analysis job
func (api *BUSTEDAPI) StartBUSTEDJob(c *gin.Context) {
	var request BustedRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse job configuration"})
		return
	}

	adapted, err := AdaptRequest(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to adapt request: %v", err)})
		return
	}

	result, err := api.HandleStartJob(c, adapted, MethodBUSTED)
	if err != nil {
		if err.Error() == "authentication token required" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, result)
}
