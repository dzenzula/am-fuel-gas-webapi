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
        "/api/GetCalculatedImbalanceDetails": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id batch расчета",
                        "name": "batch",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetCalculatedImbalanceDetails"
                        }
                    }
                }
            }
        },
        "/api/GetCalculationsList": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CalculationList"
                        }
                    }
                }
            }
        },
        "/api/GetDensityCoefficientDetails": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
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
                            "$ref": "#/definitions/models.GetDensityCoefficient"
                        }
                    }
                }
            }
        },
        "/api/GetImbalanceHistory": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ImbalanceCalculationHistory"
                            }
                        }
                    }
                }
            }
        },
        "/api/GetNodesList": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.NodeList"
                        }
                    }
                }
            }
        },
        "/api/GetParameterHistory": {
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
                    },
                    {
                        "type": "integer",
                        "description": "Id параметра",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateHistory"
                        }
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
        "/api/RecalculateDensityCoefficient": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
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
                        "description": "OK"
                    }
                }
            }
        },
        "/api/RecalculateImbalance": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculations"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Дата получения параметров",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Данные расчета небаланс",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.SetImbalanceFlag"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
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
                        "name": "data",
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
        "models.CalculationList": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.DensityCalculationHistory": {
            "type": "object",
            "properties": {
                "calculationDate": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "syncMode": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "models.GetCalculatedImbalanceDetails": {
            "type": "object",
            "properties": {
                "aggregateTotal": {
                    "type": "string"
                },
                "autoTotal": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "manualTotal": {
                    "type": "string"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Node"
                    }
                },
                "pgRedisTotal": {
                    "type": "string"
                }
            }
        },
        "models.GetDensityCoefficient": {
            "type": "object",
            "properties": {
                "calculationHistory": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DensityCalculationHistory"
                    }
                },
                "densityCoefficient": {
                    "type": "number"
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
                "timestamp": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "models.ImbalanceCalculationHistory": {
            "type": "object",
            "properties": {
                "calculationDate": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "syncMode": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
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
        "models.Node": {
            "type": "object",
            "properties": {
                "consumption": {
                    "type": "string"
                },
                "distributed": {
                    "type": "string"
                },
                "flag": {
                    "type": "string"
                },
                "gasRedistribution": {
                    "type": "string"
                },
                "measuringId": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.NodeList": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.SetImbalanceFlag": {
            "type": "object",
            "properties": {
                "flag": {
                    "type": "string"
                },
                "measuringId": {
                    "type": "integer"
                }
            }
        },
        "models.SetManualFuelGas": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
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
        "models.UpdateHistory": {
            "type": "object",
            "properties": {
                "timestampInsert": {
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
