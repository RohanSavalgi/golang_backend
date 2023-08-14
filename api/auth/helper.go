package auth

type CustomClaims struct {
	TenantId               string   `json:"https://rebit-sentinel.betsol.com/tenantId"`
	GlobalUserId           string   `json:"https://rebit-sentinel.betsol.com/globalUserId"`
	SubscriptionCustomerId string   `json:"https://rebit-sentinel.betsol.com/subscriptionCustomerId"`
	GlobalSubscriptionId   string   `json:"https://rebit-sentinel.betsol.com/globalSubscriptionId"`
	CrmUserId              string   `json:"https://rebit-sentinel.betsol.com/crmUserId"`
	Email                  string   `json:"https://rebit-sentinel.betsol.com/email"`
	Product                string   `json:"https://rebit-sentinel.betsol.com/product"`
	AuthUserId             string   `json:"sub"`
	Permissions            []string `json:"permissions"`
	OrgId                  string   `json:"org_id"`
	PartnerId              string   `json:"https://rebit-sentinel.betsol.com/partnerId"`
	PartnerTenantId        string   `json:"https://rebit-sentinel.betsol.com/partnerTenantId"`
	OrganizationName       string   `json:"https://rebit-sentinel.betsol.com/organizationName"`
	OrganizationType       string   `json:"https://rebit-sentinel.betsol.com/organizationType"`	
}

var Permissions = struct {
	Owner map[string]string
}{
	Owner: map[string]string {
		"read": "read:user",
	},
}

func (cust CustomClaims) IsAuthorized(permissions []string, checkAllPermissions bool) bool {
	isPresent := func(permission string, arr []string) bool {
		for _, item := range arr {
			if item == permission {
				return true
			}
		}
		return false
	}

	if checkAllPermissions {
		for _, item := range permissions {
			if !isPresent(item, cust.Permissions) {
				return false
			}
		}
		return true
	} else {
		for _, item := range permissions {
			if isPresent(item, cust.Permissions) {
				return true
			}
		}
		return false
	}
}