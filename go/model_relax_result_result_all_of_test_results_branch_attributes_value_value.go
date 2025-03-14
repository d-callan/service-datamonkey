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

type RelaxResultResultAllOfTestResultsBranchAttributesValueValue struct {

	OriginalName string `json:"original name,omitempty"`

	NucleotideGTR float32 `json:"Nucleotide GTR,omitempty"`

	MG94xREVWithSeparateRatesForBranchSets float32 `json:"MG94xREV with separate rates for branch sets,omitempty"`

	GeneralDescriptive float32 `json:"General descriptive,omitempty"`

	KGeneralDescriptive float32 `json:"k (general descriptive),omitempty"`

	RELAXNull float32 `json:"RELAX null,omitempty"`

	RELAXAlternative float32 `json:"RELAX alternative,omitempty"`

	KRELAXAlternative float32 `json:"k (RELAX alternative),omitempty"`

	Tested bool `json:"tested,omitempty"`
}
