// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/Authorization/GetCurrentUserInfo": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthUserData"
                        }
                    }
                }
            }
        },
        "/api/Authorization/LogInAuthorization": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "userdata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AuthUserData"
                        }
                    }
                }
            }
        },
        "/api/Authorization/LogOutAuthorization": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/GetParameters": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Дата получения параметров",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetManualFuelGas"
                        }
                    }
                }
            }
        },
        "/api/SetParameters": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Parameters"
                ],
                "parameters": [
                    {
                        "description": "Данные газ",
                        "name": "userdata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SetManualFuelGas"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AuthUserData": {
            "type": "object",
            "properties": {
                "authStatus": {
                    "type": "integer"
                },
                "displayName": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MyPermission"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.GetManualFuelGas": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastUpdateDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "models.MyPermission": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "permission": {
                    "type": "integer"
                }
            }
        },
        "models.SetManualFuelGas": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "models.UserData": {
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/am-fuel-gas-webapi",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
