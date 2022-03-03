package domain

type User struct {
	Id          int32  `json:"id"`
	Verified    bool   `json:"verified"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
}
