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
        "/api/config/degrade/{ns}/{name}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "zxy",
                        "description": "命名空间",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "gr7b1231",
                        "description": "名称",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "500": {
                        "description": "熔断规则查询失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "命名空间",
                        "name": "ns",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "gr7b1231",
                        "description": "名称",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "成功",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "熔断规则删除失败",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
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