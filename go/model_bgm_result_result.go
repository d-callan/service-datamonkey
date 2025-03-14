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

type BgmResultResult struct {

	Analysis AnalysisInfo `json:"analysis,omitempty"`

	Input InputInfo `json:"input,omitempty"`

	Tested map[string]map[string]string `json:"tested,omitempty"`

	Timers map[string]TimersInfoValue `json:"timers,omitempty"`

	Runtime string `json:"runtime,omitempty" validate:"regexp=^[0-9]+\\\\.[0-9]+\\\\.[0-9]+$"`

	DataPartitions map[string]PartitionsInfoValue `json:"data partitions,omitempty"`

	// Model fit information
	Fits map[string]BgmResultResultAllOfFitsValue `json:"fits,omitempty"`

	MLE BgmResultResultAllOfMle `json:"MLE,omitempty"`

	// Branch-specific attributes
	BranchAttributes map[string]map[string]map[string]float32 `json:"branch attributes,omitempty"`
}
