package application

func BuildApp() App {
	//
	return App{}
}

type App struct {
}

func (app App) Services() AppService {
	return AppService{
		app: app,
	}
}
