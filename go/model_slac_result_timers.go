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

// SlacResultTimers - Timing information for the analysis
type SlacResultTimers struct {

	TotalTime AbsrelResultResultAllOfTimersValue `json:"Total_time,omitempty"`

	ModelFitting AbsrelResultResultAllOfTimersValue `json:"Model_fitting,omitempty"`

	PrimarySLACAnalysis AbsrelResultResultAllOfTimersValue `json:"Primary_SLAC_analysis,omitempty"`

	AncestorSamplingAnalysis AbsrelResultResultAllOfTimersValue `json:"Ancestor_sampling_analysis,omitempty"`
}
