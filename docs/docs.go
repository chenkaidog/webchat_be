// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "hertz-contrib",
            "url": "https://github.com/hertz-contrib"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1//ping": {
            "get": {
                "description": "测试 Description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "测试 Summary",
                "responses": {}
            }
        },
        "/api/v1/account/forget_password": {
            "post": {
                "description": "用户忘记密码接口，请求获取验证码进行重置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户忘记密码接口",
                "parameters": [
                    {
                        "description": "password forget request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ForgetPasswordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ForgetPasswordResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/info": {
            "get": {
                "description": "用户信息查询接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户信息查询接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AccountInfoQueryResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "login request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.LoginResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "X-CSRF-TOKEN": {
                                "type": "string",
                                "description": "csrf token"
                            },
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/logout": {
            "post": {
                "description": "用户登出接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登出接口",
                "parameters": [
                    {
                        "description": "logout request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LogoutReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "csrf token",
                        "name": "X-CSRF-TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.LogoutResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/models": {
            "get": {
                "description": "获取模型列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取模型列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "csrf token",
                        "name": "X-CSRF-TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ModelQueryResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/register": {
            "post": {
                "description": "用户注册接口，请求后获取验证码，然后才能创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册接口",
                "parameters": [
                    {
                        "description": "register request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RegisterResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/register_verify": {
            "post": {
                "description": "用户注册验证接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册验证接口",
                "parameters": [
                    {
                        "description": "register request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterVerifyReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RegisterVerifyResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/reset_password": {
            "post": {
                "description": "用户修改密码接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户修改密码接口",
                "parameters": [
                    {
                        "description": "password reset request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ResetPasswordReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.ResetPasswordResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        },
        "/api/v1/account/update_password": {
            "post": {
                "description": "用户修改密码接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户修改密码接口",
                "parameters": [
                    {
                        "description": "password update request body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PasswordUpdateReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "csrf token",
                        "name": "X-CSRF-TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.CommonResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PasswordUpdateResp"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "set-cookie": {
                                "type": "string",
                                "description": "cookie"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.CommonResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AccountInfoQueryResp": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CommonResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "dto.ForgetPasswordReq": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "maxLength": 64
                }
            }
        },
        "dto.ForgetPasswordResp": {
            "type": "object"
        },
        "dto.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 128
                },
                "username": {
                    "type": "string",
                    "maxLength": 64
                }
            }
        },
        "dto.LoginResp": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LogoutReq": {
            "type": "object"
        },
        "dto.LogoutResp": {
            "type": "object"
        },
        "dto.Model": {
            "type": "object",
            "properties": {
                "model_id": {
                    "type": "string"
                },
                "model_name": {
                    "type": "string"
                }
            }
        },
        "dto.ModelQueryResp": {
            "type": "object",
            "properties": {
                "models": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Model"
                    }
                }
            }
        },
        "dto.PasswordUpdateReq": {
            "type": "object",
            "required": [
                "password",
                "password_new"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 128
                },
                "password_new": {
                    "type": "string",
                    "maxLength": 128,
                    "minLength": 8
                }
            }
        },
        "dto.PasswordUpdateResp": {
            "type": "object"
        },
        "dto.RegisterReq": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "password": {
                    "type": "string",
                    "maxLength": 128
                },
                "username": {
                    "type": "string",
                    "maxLength": 64
                }
            }
        },
        "dto.RegisterResp": {
            "type": "object"
        },
        "dto.RegisterVerifyReq": {
            "type": "object",
            "required": [
                "email",
                "verify_code"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "verify_code": {
                    "type": "string",
                    "maxLength": 10
                }
            }
        },
        "dto.RegisterVerifyResp": {
            "type": "object"
        },
        "dto.ResetPasswordReq": {
            "type": "object",
            "required": [
                "password",
                "verify_code"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 128
                },
                "verify_code": {
                    "type": "string",
                    "maxLength": 10
                }
            }
        },
        "dto.ResetPasswordResp": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "HertzTest",
	Description:      "This is a demo using Hertz.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
