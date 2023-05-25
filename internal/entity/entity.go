package entity

type CreateSignUpParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserPremium struct {
	PackageId int    `json:"package_id"`
	Username  string `json:"username"`
}

type Package struct {
	PackageId int    `db:"package_id"`
	Name      string `db:"name"`
	Price     int    `db:"price"`
}

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}
