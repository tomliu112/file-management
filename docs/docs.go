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
        "/delete/{path}": {
            "delete": {
                "description": "删除文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file_path",
                        "name": "path",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/download": {
            "get": {
                "description": "下载文件",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "下载文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file_path",
                        "name": "file_path",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/getList": {
            "get": {
                "description": "获取文件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "getFileList",
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
        "/getPv/{period}": {
            "get": {
                "description": "根据时间按天，月，年，获取pv",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "获取pv",
                "parameters": [
                    {
                        "type": "string",
                        "description": "day/month/year",
                        "name": "period",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/upload": {
            "post": {
                "description": "上传文件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "uploadFile",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file to upload",
                        "name": "file",
                        "in": "formData",
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
