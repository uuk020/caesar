package forms

type AccountMp struct {
	MainPassword string `param:"main_password" query:"main_password" form:"main_password" json:"main_password" validate:"required,complexpassword"`
}
type AccountId struct {
	ID string `param:"id" query:"id" form:"id" json:"id" validate:"required,number"`
}

type AccountInfo struct {
	Name     string `form:"name" json:"name" validate:"required,min=1,max=20"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}

type AccountRead struct {
	AccountId
	AccountMp
}

type AccountUpdate struct {
	AccountId
	AccountMp
	AccountInfo
}

type AccountCreate struct {
	AccountInfo
	AccountMp
	Platform string `form:"platform" json:"platform" validate:"required,min=1,max=10"`
	Url      string `form:"url" json:"url" validate:"required,url"`
}

type AccountList struct {
	Page      int    `query:"page" json:"page" validate:"required,number,min=1"`
	PageSize  int    `query:"page_size" json:"page_size" validate:"required,number,min=10"`
	Platform  string `query:"platform" json:"platform" validate:"omitempty,min=1,max=10"`
	DateStart string `query:"date_start" json:"date_start" validate:"omitempty,datetime=2006-01-02"`
	DateEnd   string `query:"date_end" json:"date_end" validate:"omitempty,datetime=2006-01-02"`
}
