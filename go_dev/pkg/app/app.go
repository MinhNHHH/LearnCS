package app

type Application struct {
	JWTSecret string
	DSN       string
}

func NewApplication(JWTSecret, DSN string) *Application {
	return &Application{
		JWTSecret: JWTSecret,
		DSN:       DSN,
	}
}
