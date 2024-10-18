// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplateapi = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/discovery": {
            "post": {
                "description": "Discovers new Printers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "discovery"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DiscoveryResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns the health and status of the various services that make up the API.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.HealthResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/status": {
            "post": {
                "description": "Refreshes the status of a machine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "discovery"
                ],
                "parameters": [
                    {
                        "description": "Status Refresh Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StatusRefreshRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.DiscoveryData": {
            "type": "object",
            "properties": {
                "BrandName": {
                    "description": "Brand Name",
                    "type": "string"
                },
                "FirmwareVersion": {
                    "description": "Firmware Version",
                    "type": "string"
                },
                "MachineModel": {
                    "description": "Machine Model",
                    "type": "string"
                },
                "MachineName": {
                    "description": "Machine Name",
                    "type": "string"
                },
                "MachineID": {
                    "description": "Motherboard ID (16-bit)",
                    "type": "string"
                },
                "MachineIP": {
                    "description": "Motherboard IP Address",
                    "type": "string"
                },
                "ProtocolVersion": {
                    "description": "Protocol Version",
                    "type": "string"
                }
            }
        },
        "models.DiscoveryResponse": {
            "type": "object",
            "properties": {
                "discovered": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.DiscoveryData"
                    }
                }
            }
        },
        "models.HealthResponse": {
            "type": "object"
        },
        "models.StatusRefreshRequest": {
            "type": "object",
            "properties": {
                "machine_id": {
                    "type": "string"
                },
                "machine_ip": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfoapi holds exported Swagger Info so clients can modify it
var SwaggerInfoapi = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{"https"},
	Title:            "Flux API V1",
	Description:      "API for Flux, V1",
	InfoInstanceName: "api",
	SwaggerTemplate:  docTemplateapi,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoapi.InstanceName(), SwaggerInfoapi)
}
