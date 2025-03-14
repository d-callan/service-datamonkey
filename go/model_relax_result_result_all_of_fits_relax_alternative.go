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

type RelaxResultResultAllOfFitsRelaxAlternative struct {

	LogLikelihood float32 `json:"Log Likelihood,omitempty"`

	EstimatedParameters int32 `json:"estimated parameters,omitempty"`

	AICC float32 `json:"AIC-c,omitempty"`

	RateDistributions RelaxResultResultAllOfFitsRelaxAlternativeRateDistributions `json:"Rate Distributions,omitempty"`

	DisplayOrder int32 `json:"display order,omitempty"`
}
