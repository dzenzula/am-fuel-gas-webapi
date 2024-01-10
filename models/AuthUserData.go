package models

type AuthUserData struct {
	DisplayName string
	Domain      string
	Name        string
	AuthStatus  *int
	Type        string
	Permission  []MyPermission
}

type MyPermission struct {
	Name       string
	Permission int
}
