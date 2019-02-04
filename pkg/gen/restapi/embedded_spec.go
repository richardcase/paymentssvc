// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample Payments API",
    "title": "Payments API",
    "contact": {
      "email": "fruitcase@pm.me"
    },
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/live": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "probes"
        ],
        "summary": "Is the service alive",
        "operationId": "getLive",
        "responses": {
          "200": {
            "description": "service is alive"
          },
          "400": {
            "description": "service isn't alive"
          }
        }
      }
    },
    "/payments": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Get a list of payments",
        "operationId": "listPayments",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Payment"
              }
            }
          },
          "400": {
            "description": "Invalid status value",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Creates a new payment",
        "operationId": "addPayment",
        "parameters": [
          {
            "description": "Payment attributes object for the new payment",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PaymentAttributes"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Payment has been created",
            "headers": {
              "Location": {
                "type": "string",
                "description": "The location to retrieve the payment instance"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/payments/{id}": {
      "get": {
        "description": "Returns a single payment",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Find a payment by ID",
        "operationId": "getPaymentById",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "payment was found",
            "schema": {
              "$ref": "#/definitions/Payment"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Update an existing payment",
        "operationId": "updatePayment",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Payment attributes object that needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PaymentAttributes"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Payment was updated successfully"
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Delete an existing payment",
        "operationId": "deletePayment",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to update",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Payment was deleted succesfully"
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "405": {
            "description": "Validation exception",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/ready": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "probes"
        ],
        "summary": "Is the service able to serve traffic",
        "operationId": "getReady",
        "responses": {
          "200": {
            "description": "service is ready"
          },
          "400": {
            "description": "service isn't ready"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Party": {
      "type": "object",
      "properties": {
        "account_name": {
          "type": "string"
        },
        "account_number": {
          "type": "string"
        }
      }
    },
    "Payment": {
      "type": "object",
      "required": [
        "id",
        "version",
        "organisation_id",
        "attributes"
      ],
      "properties": {
        "attributes": {
          "$ref": "#/definitions/PaymentAttributes"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "organisation_id": {
          "type": "string",
          "format": "uuid"
        },
        "version": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PaymentAttributes": {
      "type": "object",
      "required": [
        "amount",
        "currency",
        "payment_type",
        "payment_scheme",
        "beneficiary_party",
        "debtor_party"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "float"
        },
        "beneficiary_party": {
          "$ref": "#/definitions/Party"
        },
        "currency": {
          "type": "string"
        },
        "debtor_party": {
          "$ref": "#/definitions/Party"
        },
        "payment_scheme": {
          "description": "Which payment scheme are we using",
          "type": "string",
          "enum": [
            "FPS",
            "BACS",
            "CHAPS",
            "SEPA"
          ]
        },
        "payment_type": {
          "type": "string",
          "enum": [
            "Credit",
            "Debit"
          ]
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything related to payments",
      "name": "payment"
    },
    {
      "description": "Endpoints related to k8s probes",
      "name": "probes"
    },
    {
      "description": "Operations related to metrics",
      "name": "metrics"
    }
  ],
  "externalDocs": {
    "description": "Find out more about the service",
    "url": "https://github.com/richardcase/paymentsapi/blob/master/docs/design.md"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample Payments API",
    "title": "Payments API",
    "contact": {
      "email": "fruitcase@pm.me"
    },
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/live": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "probes"
        ],
        "summary": "Is the service alive",
        "operationId": "getLive",
        "responses": {
          "200": {
            "description": "service is alive"
          },
          "400": {
            "description": "service isn't alive"
          }
        }
      }
    },
    "/payments": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Get a list of payments",
        "operationId": "listPayments",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Payment"
              }
            }
          },
          "400": {
            "description": "Invalid status value",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Creates a new payment",
        "operationId": "addPayment",
        "parameters": [
          {
            "description": "Payment attributes object for the new payment",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PaymentAttributes"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Payment has been created",
            "headers": {
              "Location": {
                "type": "string",
                "description": "The location to retrieve the payment instance"
              }
            }
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/payments/{id}": {
      "get": {
        "description": "Returns a single payment",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Find a payment by ID",
        "operationId": "getPaymentById",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "payment was found",
            "schema": {
              "$ref": "#/definitions/Payment"
            }
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Update an existing payment",
        "operationId": "updatePayment",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Payment attributes object that needs to be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PaymentAttributes"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Payment was updated successfully"
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "payment"
        ],
        "summary": "Delete an existing payment",
        "operationId": "deletePayment",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of payment to update",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Payment was deleted succesfully"
          },
          "400": {
            "description": "Invalid ID supplied",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Payment not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "405": {
            "description": "Validation exception",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/ready": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "tags": [
          "probes"
        ],
        "summary": "Is the service able to serve traffic",
        "operationId": "getReady",
        "responses": {
          "200": {
            "description": "service is ready"
          },
          "400": {
            "description": "service isn't ready"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "Party": {
      "type": "object",
      "properties": {
        "account_name": {
          "type": "string"
        },
        "account_number": {
          "type": "string"
        }
      }
    },
    "Payment": {
      "type": "object",
      "required": [
        "id",
        "version",
        "organisation_id",
        "attributes"
      ],
      "properties": {
        "attributes": {
          "$ref": "#/definitions/PaymentAttributes"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "organisation_id": {
          "type": "string",
          "format": "uuid"
        },
        "version": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "PaymentAttributes": {
      "type": "object",
      "required": [
        "amount",
        "currency",
        "payment_type",
        "payment_scheme",
        "beneficiary_party",
        "debtor_party"
      ],
      "properties": {
        "amount": {
          "type": "number",
          "format": "float"
        },
        "beneficiary_party": {
          "$ref": "#/definitions/Party"
        },
        "currency": {
          "type": "string"
        },
        "debtor_party": {
          "$ref": "#/definitions/Party"
        },
        "payment_scheme": {
          "description": "Which payment scheme are we using",
          "type": "string",
          "enum": [
            "FPS",
            "BACS",
            "CHAPS",
            "SEPA"
          ]
        },
        "payment_type": {
          "type": "string",
          "enum": [
            "Credit",
            "Debit"
          ]
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything related to payments",
      "name": "payment"
    },
    {
      "description": "Endpoints related to k8s probes",
      "name": "probes"
    },
    {
      "description": "Operations related to metrics",
      "name": "metrics"
    }
  ],
  "externalDocs": {
    "description": "Find out more about the service",
    "url": "https://github.com/richardcase/paymentsapi/blob/master/docs/design.md"
  }
}`))
}