package rest

type NewPersonRequest struct {
	Email    string `json:"email" binding:"required,email"`
	IsActive bool   `json:"is_active"`
	Password string `json:"password" binding:"required"`
}

type NewPermissionRequest struct {
	Title    string `json:"title" binding:"required"`
	Service  string `json:"service" binding:"required"`
	Function string `json:"function" binding:"required"`
	Verb     string `json:"verb" binding:"required"`
}

type NewRoleRequest struct {
	Title string `json:"title" binding:"required"`
}

type NewTenantRequest struct {
	IsActive  bool   `json:"is_active" binding:"required"`
	Name      string `json:"name" binding:"required"`
	ShortName string `json:"short_name" binding:"required"`
}
