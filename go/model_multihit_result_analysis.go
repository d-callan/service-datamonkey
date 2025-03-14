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

// MultihitResultAnalysis - Metadata about the analysis
type MultihitResultAnalysis struct {

	Authors string `json:"authors,omitempty"`

	Contact string `json:"contact,omitempty"`

	Info string `json:"info,omitempty"`

	Requirements string `json:"requirements,omitempty"`

	Version string `json:"version,omitempty"`
}
