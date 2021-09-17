// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

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
            "name": "LyricTian",
            "email": "tiannianshou@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/demos": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "查询数据",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "分页索引",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "分页大小",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态(1:启用 2:禁用)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "查询结果",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/schema.ListResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "list": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/schema.Demo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
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
                "tags": [
                    "DemoManage"
                ],
                "summary": "创建数据",
                "parameters": [
                    {
                        "description": "创建数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Demo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.IDResult"
                        }
                    },
                    "400": {
                        "description": "{error:{code:0,message:无效的请求参数}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        },
        "/api/v1/demos/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "查询指定数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "唯一标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Demo"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "404": {
                        "description": "{error:{code:0,message:资源不存在}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "更新数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "唯一标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新数据",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.Demo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{status:OK}",
                        "schema": {
                            "$ref": "#/definitions/schema.StatusResult"
                        }
                    },
                    "400": {
                        "description": "{error:{code:0,message:无效的请求参数}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "删除数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "唯一标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{status:OK}",
                        "schema": {
                            "$ref": "#/definitions/schema.StatusResult"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        },
        "/api/v1/demos/{id}/disable": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "禁用数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "唯一标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{status:OK}",
                        "schema": {
                            "$ref": "#/definitions/schema.StatusResult"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        },
        "/api/v1/demos/{id}/enable": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "DemoManage"
                ],
                "summary": "启用数据",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "唯一标识",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{status:OK}",
                        "schema": {
                            "$ref": "#/definitions/schema.StatusResult"
                        }
                    },
                    "401": {
                        "description": "{error:{code:0,message:未授权}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    },
                    "500": {
                        "description": "{error:{code:0,message:服务器错误}}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.Demo": {
            "type": "object",
            "required": [
                "code",
                "name"
            ],
            "properties": {
                "code": {
                    "description": "编号",
                    "type": "string"
                },
                "created_at": {
                    "description": "创建时间",
                    "type": "string"
                },
                "creator": {
                    "description": "创建者",
                    "type": "integer"
                },
                "id": {
                    "description": "唯一标识",
                    "type": "integer"
                },
                "memo": {
                    "description": "备注",
                    "type": "string"
                },
                "name": {
                    "description": "名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态(1:启用 2:禁用)",
                    "type": "integer"
                },
                "updated_at": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "schema.ErrorItem": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码",
                    "type": "integer"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        },
        "schema.ErrorResult": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "错误项",
                    "$ref": "#/definitions/schema.ErrorItem"
                }
            }
        },
        "schema.IDResult": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "schema.ListResult": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "object"
                },
                "pagination": {
                    "$ref": "#/definitions/schema.PaginationResult"
                }
            }
        },
        "schema.PaginationResult": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "schema.StatusResult": {
            "type": "object",
            "properties": {
                "status": {
                    "description": "状态(OK)",
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
	Version:     "8.0.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "gin-admin-tpl",
	Description: "RBAC scaffolding based on GIN + GORM + CASBIN + WIRE.",
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
