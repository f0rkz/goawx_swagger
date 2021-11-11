/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data57 struct {
	// 
	Description string `json:"description,omitempty"`
	// 
	DiffMode bool `json:"diff_mode,omitempty"`
	// Enables processing of this schedule.
	Enabled bool `json:"enabled,omitempty"`
	// 
	ExtraData string `json:"extra_data,omitempty"`
	// Inventory applied as a prompt, assuming job template prompts for inventory
	Inventory int32 `json:"inventory,omitempty"`
	// 
	JobTags string `json:"job_tags,omitempty"`
	// 
	JobType string `json:"job_type,omitempty"`
	// 
	Limit string `json:"limit,omitempty"`
	// 
	Name string `json:"name"`
	// A value representing the schedules iCal recurrence rule.
	Rrule string `json:"rrule"`
	// 
	ScmBranch string `json:"scm_branch,omitempty"`
	// 
	SkipTags string `json:"skip_tags,omitempty"`
	// 
	UnifiedJobTemplate int32 `json:"unified_job_template"`
	// 
	Verbosity string `json:"verbosity,omitempty"`
}
