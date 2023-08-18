package resty

type requestType string
type serviceProvider string

const (
	GET requestType = "GET"
	PUT    requestType = "PUT"
	POST   requestType = "POST"
	PATCH  requestType = "PATCH"
	DELETE requestType = "DELETE"
)

const (
	AUTH0           serviceProvider = "Auth0"
	GLOBAL_REGISTRY serviceProvider = "Global Registry"
	MANAGED_BACKUP  serviceProvider = "Managed Backup Service"
	SUB_MGMT        serviceProvider = "Subscription Management Service"
	FUSION_AUTH     serviceProvider = "Fusion Auth"
	REBIT           serviceProvider = "Rebit"
	HUBSPOT         serviceProvider = "Hubspot"
	USR_MGMT        serviceProvider = "User Management Service"
)