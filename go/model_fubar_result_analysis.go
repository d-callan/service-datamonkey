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

type FubarResultAnalysis struct {

	// Information about the FUBAR analysis
	Info string `json:"info,omitempty"`

	// Version of the FUBAR method
	Version string `json:"version,omitempty"`

	// Citation for the FUBAR method
	Citation string `json:"citation,omitempty"`

	// Authors of the FUBAR method
	Authors string `json:"authors,omitempty"`

	// Contact information for the authors
	Contact string `json:"contact,omitempty"`

	// Requirements for running FUBAR
	Requirements string `json:"requirements,omitempty"`
}
