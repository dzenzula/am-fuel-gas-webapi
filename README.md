# am-fuel-gas-webapi

### генерация компонентов swagger:
-   создать модуль `go mod init main`
-   использовать `go mod tidy`
-   установить утилиту swag `go get -u github.com/swaggo/swag/cmd/swag`
-   установить последнюю версию сваггера`go get -u github.com/swaggo/swag`
-   автоматическая генерация компонентов swagger `swag init`
	- docs/docs.go
	- docs/swagger.json
	- docs/swagger.yaml
-   введена переменная **@BasePath**, которая является префиксом к основному пути приложения
	https://krr-tst-padev02.europe.mittalco.com/ `am-fuel-gas-webapi` /swagger/index.html
> в данном проекте генерирует разработчик
