// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/CalculateImbalance": {
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
                        "type": "string",
                        "description": "Id batch расчета",
                        "name": "batch",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Флаг раздельности расчета",
                        "name": "separate",
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
                                "$ref": "#/definitions/models.PostImbalanceCalculation"
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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id batch расчета",
                        "name": "cloneId",
                        "in": "query"
                    }
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
                    },
                    {
                        "type": "string",
                        "description": "Тэг типы: 'day', 'month' or empty",
                        "name": "tag",
                        "in": "query"
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
        "/api/GetScales": {
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
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetScales"
                        }
                    }
                }
            }
        },
        "/api/PrepareImbalanceCalculation": {
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
                        "description": "Дата расчета",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
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
        "/api/RemoveImbalanceCalculation": {
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
                        "description": "Id batch расчета",
                        "name": "batch",
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
        },
        "/api/UpdateScale": {
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
                        "description": "Данные шкалы",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateScale"
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
                "imbalanceCalculation": {
                    "$ref": "#/definitions/models.ImbalanceCalculation"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Node"
                    }
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
                "dateCoefficient": {
                    "type": "string"
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
        "models.GetScales": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.ImbalanceCalculation": {
            "type": "object",
            "properties": {
                "aggregateTotal": {
                    "description": "Сумма поагрегатного потребления",
                    "type": "string"
                },
                "aggregateTotal12": {
                    "type": "string"
                },
                "aggregateTotal3": {
                    "type": "string"
                },
                "autoTotal": {
                    "type": "string"
                },
                "calculationDate": {
                    "type": "string"
                },
                "density": {
                    "type": "string"
                },
                "grp10Manual": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "internalImbalance": {
                    "description": "Внутренний небаланс",
                    "type": "string"
                },
                "internalImbalance12": {
                    "type": "string"
                },
                "internalImbalance3": {
                    "type": "string"
                },
                "manualTotal": {
                    "type": "string"
                },
                "nitka1Auto": {
                    "type": "string"
                },
                "nitka1Manual": {
                    "type": "string"
                },
                "nitka2Auto": {
                    "type": "string"
                },
                "nitka2Manual": {
                    "type": "string"
                },
                "nitka3Auto": {
                    "type": "string"
                },
                "nitka3Manual": {
                    "type": "string"
                },
                "percentageImbalance": {
                    "description": "Процент небаланса",
                    "type": "string"
                },
                "percentageImbalance12": {
                    "type": "string"
                },
                "percentageImbalance3": {
                    "type": "string"
                },
                "pgRedisTotal": {
                    "description": "Сумма внутреннего небаланса",
                    "type": "string"
                },
                "separately": {
                    "type": "string"
                },
                "sumWithImbalance": {
                    "description": "Сумма потребления с небалансом",
                    "type": "string"
                },
                "sumWithImbalance12": {
                    "type": "string"
                },
                "sumWithImbalance3": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
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
                "adjustment": {
                    "type": "string"
                },
                "batchId": {
                    "type": "string"
                },
                "consumption": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "distributed": {
                    "type": "string"
                },
                "flagBalance": {
                    "type": "string"
                },
                "flagRedistribution": {
                    "type": "string"
                },
                "gasRedistribution": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "parentId": {
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
                "flag": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.PostImbalanceCalculation": {
            "type": "object",
            "properties": {
                "imbalanceCalculation": {
                    "$ref": "#/definitions/models.SetImbalanceData"
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SetImbalanceFlag"
                    }
                }
            }
        },
        "models.SetImbalanceData": {
            "type": "object",
            "properties": {
                "grp10Manual": {
                    "type": "string"
                },
                "nitka1Auto": {
                    "type": "string"
                },
                "nitka1Manual": {
                    "type": "string"
                },
                "nitka2Auto": {
                    "type": "string"
                },
                "nitka2Manual": {
                    "type": "string"
                },
                "nitka3Auto": {
                    "type": "string"
                },
                "nitka3Manual": {
                    "type": "string"
                }
            }
        },
        "models.SetImbalanceFlag": {
            "type": "object",
            "properties": {
                "adjustment": {
                    "type": "string"
                },
                "flagBalance": {
                    "type": "string"
                },
                "flagRedistribution": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
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
        "models.UpdateScale": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
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
