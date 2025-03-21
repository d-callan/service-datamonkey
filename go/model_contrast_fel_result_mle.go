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

type ContrastFelResultMle struct {

	// Headers describing the columns in the content array
	Headers [][]string `json:"headers,omitempty"`

	// Site-specific results with statistical values
	Content map[string][][]float32 `json:"content,omitempty"`
}
