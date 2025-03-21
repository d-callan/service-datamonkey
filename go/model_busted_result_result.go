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

type BustedResultResult struct {

	Analysis AnalysisInfo `json:"analysis,omitempty"`

	Input InputInfo `json:"input,omitempty"`

	Tested map[string]map[string]string `json:"tested,omitempty"`

	Timers map[string]TimersInfoValue `json:"timers,omitempty"`

	Runtime string `json:"runtime,omitempty" validate:"regexp=^[0-9]+\\\\.[0-9]+\\\\.[0-9]+$"`

	DataPartitions map[string]PartitionsInfoValue `json:"data partitions,omitempty"`

	Background float32 `json:"background,omitempty"`

	Fits map[string]BustedResultResultAllOfFitsValue `json:"fits,omitempty"`

	BranchAttributes BustedResultResultAllOfBranchAttributes `json:"branch attributes,omitempty"`

	TestResults BustedResultResultAllOfTestResults `json:"test results,omitempty"`

	Substitutions map[string]map[string]BustedResultResultAllOfSubstitutionsValueValue `json:"substitutions,omitempty"`

	SynonymousSitePosteriors [][]float32 `json:"Synonymous site-posteriors,omitempty"`

	SiteLogLikelihood BustedResultResultAllOfSiteLogLikelihood `json:"Site Log Likelihood,omitempty"`

	EvidenceRatios BustedResultResultAllOfEvidenceRatios `json:"Evidence Ratios,omitempty"`
}
