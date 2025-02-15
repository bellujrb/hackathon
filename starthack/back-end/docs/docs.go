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
        "/all-fields": {
            "get": {
                "description": "Busca todos os campos do banco de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Busca todos os campos",
                "responses": {
                    "200": {
                        "description": "Campos encontrados com sucesso",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/inter.FieldInput"
                            }
                        }
                    }
                }
            }
        },
        "/classify": {
            "get": {
                "description": "Classifica um campo com base na área fornecida",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Classifica um campo com base na área",
                "parameters": [
                    {
                        "description": "Área a ser classificada",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.ClassifyField"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Classificação do campo",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/external-companies": {
            "post": {
                "description": "Cria uma nova empresa no banco de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Companies"
                ],
                "summary": "Cria uma nova empresa",
                "parameters": [
                    {
                        "description": "Dados da empresa a ser criada",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/external.Company"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Empresa criada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/external-fields": {
            "post": {
                "description": "Cria um novo campo no banco de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Cria um novo campo",
                "parameters": [
                    {
                        "description": "Dados do campo a ser criado",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/external.FieldModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Campo criado com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/external-locations": {
            "post": {
                "description": "Cria uma nova localização no banco de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Locations"
                ],
                "summary": "Cria uma nova localização",
                "parameters": [
                    {
                        "description": "Dados da localização a ser criada",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/external.Location"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Localização criada com sucesso",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fields": {
            "get": {
                "description": "Busca um campo no banco de dados pelo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Busca um campo pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do campo a ser buscado",
                        "name": "id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Campo encontrado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/inter.FieldInput"
                        }
                    }
                }
            },
            "post": {
                "description": "Cria um novo campo no banco de dados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Cria um novo campo",
                "parameters": [
                    {
                        "description": "Dados do campo a ser criado",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/inter.FieldInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Campo criado com sucesso",
                        "schema": {
                            "$ref": "#/definitions/inter.FieldInput"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleta um campo do banco de dados pelo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fields"
                ],
                "summary": "Deleta um campo pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do campo a ser editado",
                        "name": "id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Campo deletado com sucesso"
                    }
                }
            }
        }
    },
    "definitions": {
        "external.Company": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "json_extended_attributes": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "place_id": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "unit_system": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "external.FieldModel": {
            "type": "object",
            "properties": {
                "declared_area": {
                    "type": "number"
                },
                "event_date": {
                    "type": "string"
                },
                "geometry": {
                    "$ref": "#/definitions/external.GeoJSON"
                },
                "json_extended_attributes": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "parent_region_id": {
                    "type": "string"
                },
                "property_id": {
                    "type": "string"
                }
            }
        },
        "external.GeoJSON": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "array",
                            "items": {
                                "type": "number"
                            }
                        }
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "external.Location": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "json_extended_attributes": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "org_id": {
                    "type": "string"
                },
                "place_id": {
                    "type": "string"
                },
                "reference_point": {
                    "$ref": "#/definitions/external.ReferencePoint"
                },
                "state": {
                    "type": "string"
                },
                "time_zone": {
                    "type": "string"
                },
                "zip_code": {
                    "type": "string"
                }
            }
        },
        "external.ReferencePoint": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "inter.ClassifyField": {
            "type": "object",
            "properties": {
                "area": {
                    "type": "integer"
                }
            }
        },
        "inter.FieldInput": {
            "type": "object",
            "properties": {
                "area": {
                    "type": "integer"
                },
                "loc": {
                    "type": "string"
                },
                "name": {
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
	BasePath:         "",
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
