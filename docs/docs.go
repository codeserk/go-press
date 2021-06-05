// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/login": {
            "post": {
                "description": "Tries to login using some credentials.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Tries to login using some credentials.",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "User login with email and password",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User response and JWT",
                        "schema": {
                            "$ref": "#/definitions/LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "description": "Registers a new user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Registers a new user with email and password",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "User registration parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User response and JWT",
                        "schema": {
                            "$ref": "#/definitions/RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets all the realms accessible for the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Realm"
                ],
                "summary": "Gets all the realms accessible for the current user",
                "operationId": "get-realms",
                "responses": {
                    "200": {
                        "description": "Realms to which the current user has access",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/realm.Entity"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new realm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Realm"
                ],
                "summary": "Creates a new realm",
                "operationId": "create-realm",
                "parameters": [
                    {
                        "description": "Realm parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/http.createRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/realm.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm/{realmId}/node": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets all the nodes in the given realm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Gets all the nodes in the given realm",
                "operationId": "get-nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/node.Entity"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Creates a new node",
                "operationId": "create-node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "description": "Node parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/CreateNodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/node.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm/{realmId}/node/{nodeId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates a node",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Node"
                ],
                "summary": "Updates a node",
                "operationId": "update-node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Node ID",
                        "name": "nodeId",
                        "in": "path"
                    },
                    {
                        "description": "Node parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/UpdateNodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/node.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm/{realmId}/schema": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets all the schemas in the given realm",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schema"
                ],
                "summary": "Gets all the schemas in the given realm",
                "operationId": "get-schemas",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.Entity"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new schema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schema"
                ],
                "summary": "Creates a new schema",
                "operationId": "create-schema",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "description": "Schema parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/CreateSchemaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm/{realmId}/schema/{schemaId}/field": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets all the fields of the given schema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "Gets all the fields of the given schema",
                "operationId": "get-fields",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Schema ID",
                        "name": "schemaId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/field.Entity"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new field",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "Creates a new field",
                "operationId": "create-field",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Schema ID",
                        "name": "schemaId",
                        "in": "path"
                    },
                    {
                        "description": "Field parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/CreateFieldRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/field.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/realm/{realmId}/schema/{schemaId}/field/{fieldId}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates a field",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Field"
                ],
                "summary": "Updates a field",
                "operationId": "update-field",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Realm ID",
                        "name": "realmId",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Schema ID",
                        "name": "schemaId",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Field ID",
                        "name": "fieldId",
                        "in": "path"
                    },
                    {
                        "description": "Field parameters",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/UpdateFieldRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/field.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/user/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Gets the current user and JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Gets the current user and JWT",
                "operationId": "me",
                "responses": {
                    "200": {
                        "description": "User and JWT",
                        "schema": {
                            "$ref": "#/definitions/MeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateFieldRequest": {
            "type": "object",
            "required": [
                "key",
                "primitive"
            ],
            "properties": {
                "data": {
                    "type": "object"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "primitive": {
                    "type": "string"
                }
            }
        },
        "CreateNodeRequest": {
            "type": "object",
            "required": [
                "name",
                "schemaId",
                "slug"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "schemaId": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "scene",
                        "nested"
                    ]
                }
            }
        },
        "CreateSchemaRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "scene",
                        "nested"
                    ]
                }
            }
        },
        "LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@test.com"
                },
                "password": {
                    "type": "string",
                    "example": "test"
                }
            }
        },
        "LoginResponse": {
            "type": "object",
            "properties": {
                "jwt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.Entity"
                }
            }
        },
        "MeResponse": {
            "type": "object",
            "properties": {
                "jwt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.Entity"
                }
            }
        },
        "RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "RegisterResponse": {
            "type": "object",
            "properties": {
                "jwt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/user.Entity"
                }
            }
        },
        "UpdateFieldRequest": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "primitive": {
                    "type": "string"
                }
            }
        },
        "UpdateNodeRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                }
            }
        },
        "field.Entity": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "additionalProperties": true
                },
                "id": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "primitive": {
                    "type": "string"
                },
                "schemaId": {
                    "type": "string"
                }
            }
        },
        "http.createRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "node.Entity": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "realmId": {
                    "type": "string"
                },
                "schemaId": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "scene",
                        "nested"
                    ]
                }
            }
        },
        "realm.Entity": {
            "type": "object",
            "properties": {
                "authorId": {
                    "description": "Fist iteration, a realm belongs only to one user.",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schema.Entity": {
            "type": "object",
            "properties": {
                "authorId": {
                    "type": "string"
                },
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/field.Entity"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "realmId": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "scene",
                        "nested"
                    ]
                }
            }
        },
        "user.Entity": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "util.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:5050",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "GoPress",
	Description: "This is a sample server Petstore server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
