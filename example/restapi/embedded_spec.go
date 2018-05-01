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
    "description": "This is a simplifed version of the sample server Petstore server. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). For this sample, you can use the api key ` + "`" + `special-key` + "`" + ` to test the authorization filters.\n",
    "title": "Swagger Petstore",
    "version": "1.0.0"
  },
  "host": "petstore.org",
  "basePath": "/api",
  "paths": {
    "/pet": {
      "get": {
        "tags": [
          "pet"
        ],
        "summary": "List pets",
        "operationId": "PetList",
        "security": [
          {
            "roles": [
              "admin",
              "member"
            ]
          }
        ],
        "parameters": [
          {
            "type": "array",
            "items": {
              "enum": [
                "available",
                "pending",
                "sold"
              ],
              "type": "string",
              "default": "available"
            },
            "collectionFormat": "multi",
            "description": "Status values that need to be considered for filter",
            "name": "status",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      },
      "put": {
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "PetUpdate",
        "security": [
          {
            "roles": [
              "admin"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Updated successfully",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "PetCreate",
        "security": [
          {
            "roles": [
              "admin"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "tags": [
          "pet"
        ],
        "summary": "Get pet by it's ID",
        "operationId": "PetGet",
        "security": [
          {
            "roles": [
              "admin",
              "member"
            ]
          }
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "delete": {
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "PetDelete",
        "security": [
          {
            "roles": [
              "admin"
            ]
          }
        ],
        "parameters": [
          {
            "type": "string",
            "name": "api_key",
            "in": "header"
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted successfully"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    },
    "/store/inventory": {
      "get": {
        "tags": [
          "store"
        ],
        "summary": "Returns pet inventories by status",
        "operationId": "InventoryGet",
        "security": [
          {
            "api_key": []
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "type": "integer",
                "format": "int32"
              }
            }
          }
        }
      }
    },
    "/store/order": {
      "post": {
        "tags": [
          "store"
        ],
        "summary": "Place an order for a pet",
        "operationId": "OrderCreate",
        "parameters": [
          {
            "description": "order placed for purchasing the pet",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid Order"
          }
        }
      }
    },
    "/store/order/{orderId}": {
      "get": {
        "description": "For valid response try integer IDs with value \u003e= 1 and \u003c= 10. Other values will generated exceptions",
        "tags": [
          "store"
        ],
        "summary": "Find purchase order by ID",
        "operationId": "OrderGet",
        "parameters": [
          {
            "maximum": 10,
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be fetched",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      },
      "delete": {
        "description": "For valid response try integer IDs with positive integer value. Negative or non-integer values will generate API errors",
        "tags": [
          "store"
        ],
        "summary": "Delete purchase order by ID",
        "operationId": "OrderDelete",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "ID of the order that needs to be deleted",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted successfully"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean",
          "default": false
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "petId": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int32"
        },
        "shipDate": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "description": "Order Status",
          "type": "string",
          "enum": [
            "placed",
            "approved",
            "delivered"
          ]
        }
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "x-go-custom-tag": "gorm:\"primary_key\" query:\"filter,sort\""
        },
        "name": {
          "type": "string",
          "x-go-custom-tag": "query:\"filter,sort\"",
          "example": "doggie"
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ],
          "x-go-custom-tag": "query:\"filter,sort\""
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "key": {
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "key": []
    }
  ],
  "tags": [
    {
      "description": "Everything about your Pets",
      "name": "pet"
    },
    {
      "description": "Access to Petstore orders",
      "name": "store"
    }
  ]
}`))
}
