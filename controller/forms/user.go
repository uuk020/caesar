package forms

type Login struct {
	UserName  string `form:"user_name" json:"user_name" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required,min=6"`
	Captcha   string `form:"captcha" json:"captcha" validate:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" validate:"required"`
}

type CommonS struct {
	RealName string `form:"real_name" json:"real_name" validate:"required,min=3"`
	Email    string `form:"email" json:"email" validate:"required,email"`
	Phone    string `form:"phone" json:"phone" validate:"required,chinaphone"`
}

type Register struct {
	UserName string `form:"user_name" json:"user_name" validate:"required,min=5"`
	CommonS
	Password     string `form:"password" json:"password" validate:"required,min=6,max=12,complexpassword"`
	RePassword   string `form:"re_password" json:"re_password" validate:"required,eqfield=Password"`
	MainPassword string `form:"main_password" json:"main_password" validate:"required,nefield=Password,complexpassword"`
	Captcha      string `form:"captcha" json:"captcha" validate:"required,min=5,max=5"`
	CaptchaId    string `form:"captcha_id" json:"captcha_id" validate:"required"`
}

type EmailS struct {
	Email string `form:"email" json:"email" validate:"required,email"`
}

type SendMailS struct {
	EmailS
	Type string `form:"type" json:"type" validate:"required"`
}

type Activation struct {
	EmailS
	Code string `form:"code" json:"code" validate:"required,min=6,max=6"`
}

type ResetPassword struct {
	EmailS
	Code       string `form:"code" json:"code" validate:"required,min=6,max=6"`
	Password   string `form:"password" json:"password" validate:"required,min=6,max=12"`
	RePassword string `form:"re_password" json:"re_password" validate:"required,eqfield=Password"`
}
