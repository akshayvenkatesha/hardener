// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Hardener appliacnation",
    "title": "Hardener Application",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "post": {
        "tags": [
          "hardener"
        ],
        "operationId": "create",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/createjob"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/createjob"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "get": {
        "tags": [
          "hardener"
        ],
        "responses": {
          "200": {
            "description": "job progress",
            "schema": {
              "$ref": "#/definitions/getjob"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "createjob": {
      "type": "object",
      "required": [
        "hostname",
        "username",
        "password"
      ],
      "properties": {
        "hostname": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "string",
          "minLength": 1
        },
        "password": {
          "type": "string",
          "minLength": 1
        },
        "username": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "getjob": {
      "type": "object",
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "hostname": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "progress": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
}
