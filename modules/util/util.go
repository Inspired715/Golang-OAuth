package util

type App struct {
	JwtSecret string
}

var AppSetting = &App{}

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(AppSetting.JwtSecret)
}
