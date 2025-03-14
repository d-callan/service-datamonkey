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

type GardResultResultAllOfImprovementsValue struct {

	// Breakpoint positions
	Breakpoints [][]int32 `json:"breakpoints,omitempty"`

	// Change in AICc score
	DeltaAICc float32 `json:"deltaAICc,omitempty"`
}
