/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data37 struct {
	// 
	Description string `json:"description,omitempty"`
	// Optional custom messages for notification template.
	Messages string `json:"messages,omitempty"`
	// 
	Name string `json:"name"`
	// 
	NotificationConfiguration string `json:"notification_configuration,omitempty"`
	// 
	NotificationType string `json:"notification_type"`
	// 
	Organization int32 `json:"organization"`
}
