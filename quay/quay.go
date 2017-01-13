package quay

const QuayURLBase = "https://quay.io/api/v1/repository/"

type QuayPermission struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type QuayPermissions struct {
	Items []QuayPermission
}

type QuayPermissionsResponse struct {
	Items map[string]interface{} `json:"permissions"`
}

type QuayRepository struct {
	Namespace   string `json:"namespace"`
	Visibility  string `json:"visibility"`
	Name        string `json:"name"`
	Description string `json:"description"`
}