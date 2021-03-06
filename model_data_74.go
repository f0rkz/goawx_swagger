/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data74 struct {
	// If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node
	AllParentsMustConverge bool `json:"all_parents_must_converge,omitempty"`
	// 
	DiffMode bool `json:"diff_mode,omitempty"`
	// 
	ExtraData string `json:"extra_data,omitempty"`
	// An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node.
	Identifier string `json:"identifier,omitempty"`
	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int32 `json:"inventory,omitempty"`
	// 
	JobTags string `json:"job_tags,omitempty"`
	// 
	JobType string `json:"job_type,omitempty"`
	// 
	Limit string `json:"limit,omitempty"`
	// 
	ScmBranch string `json:"scm_branch,omitempty"`
	// 
	SkipTags string `json:"skip_tags,omitempty"`
	// 
	UnifiedJobTemplate int32 `json:"unified_job_template,omitempty"`
	// 
	Verbosity string `json:"verbosity,omitempty"`
	// 
	WorkflowJobTemplate string `json:"workflow_job_template"`
}
