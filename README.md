# swagger-tool
> provides convenient tools for swagger

## merge
Merge multiple Swagger JSON files to a JSON file.

* it merges `paths` fields of input files
* it merges `definitions` fields of input files

### Usage
```
swagger-tool merge -i input1.json -i input2.json -o output.json
```

### Why?
[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)'s
[protoc-gen-swagger](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-swagger)
generates seperated Swagger JSON files for each namespace. You may want to merge them to a Swagger JSON file.

### Example
#### input1.json
```
{
  "swagger": "2.0",
  "info": {
    "title": "pb/taeho/auth/auth.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/auth/refreshToken": {
      "post": {
        "operationId": "Refresh",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/authRefreshResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/authRefreshRequest"
            }
          }
        ],
        "tags": [
          "AuthService"
        ]
      }
    }
  },
  "definitions": {
    "authAuthResponse": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "expires_in": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "authParseResponse": {
      "type": "object",
      "properties": {
        "account_id": {
          "type": "string"
        }
      }
    },
    "authRefreshRequest": {
      "type": "object",
      "properties": {
        "refresh_token": {
          "type": "string"
        }
      }
    },
    "authRefreshResponse": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "expires_in": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "authValidateResponse": {
      "type": "object",
      "properties": {
        "is_valid": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    }
  }
}

```

#### input2.json
```
{
  "swagger": "2.0",
  "info": {
    "title": "pb/taeho/account/account.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/account/logIn": {
      "post": {
        "operationId": "LogIn",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/accountLogInResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountLogInRequest"
            }
          }
        ],
        "tags": [
          "AccountService"
        ]
      }
    },
    "/v1/account/register": {
      "post": {
        "operationId": "Register",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/accountRegisterResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountRegisterRequest"
            }
          }
        ],
        "tags": [
          "AccountService"
        ]
      }
    }
  },
  "definitions": {
    "accountAuthType": {
      "type": "string",
      "enum": [
        "NONE",
        "EMAIL"
      ],
      "default": "NONE"
    },
    "accountLogInRequest": {
      "type": "object",
      "properties": {
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "accountLogInResponse": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        },
        "expires_in": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "accountRegisterRequest": {
      "type": "object",
      "properties": {
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "full_name": {
          "type": "string"
        }
      }
    },
    "accountRegisterResponse": {
      "type": "object",
      "properties": {
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        },
        "account_id": {
          "type": "string"
        }
      }
    }
  }
}
```

#### output.json
```
{
  "swagger": "2.0",
  "info": {
    "title": "API",
    "version": "v1"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/account/logIn": {
      "post": {
        "operationId": "LogIn",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountLogInRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/accountLogInResponse"
            }
          }
        },
        "tags": [
          "AccountService"
        ]
      }
    },
    "/v1/account/register": {
      "post": {
        "operationId": "Register",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/accountRegisterRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/accountRegisterResponse"
            }
          }
        },
        "tags": [
          "AccountService"
        ]
      }
    },
    "/v1/auth/refreshToken": {
      "post": {
        "operationId": "Refresh",
        "parameters": [
          {
            "in": "body",
            "name": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/authRefreshRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/authRefreshResponse"
            }
          }
        },
        "tags": [
          "AuthService"
        ]
      }
    }
  },
  "definitions": {
    "accountAuthType": {
      "default": "NONE",
      "enum": [
        "NONE",
        "EMAIL"
      ],
      "type": "string"
    },
    "accountLogInRequest": {
      "properties": {
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        },
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "accountLogInResponse": {
      "properties": {
        "access_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "expires_in": {
          "format": "int64",
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "accountRegisterRequest": {
      "properties": {
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        },
        "email": {
          "type": "string"
        },
        "full_name": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "accountRegisterResponse": {
      "properties": {
        "account_id": {
          "type": "string"
        },
        "auth_type": {
          "$ref": "#/definitions/accountAuthType"
        }
      },
      "type": "object"
    },
    "authAuthResponse": {
      "properties": {
        "access_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "expires_in": {
          "format": "int64",
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "authParseResponse": {
      "properties": {
        "account_id": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "authRefreshRequest": {
      "properties": {
        "refresh_token": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "authRefreshResponse": {
      "properties": {
        "access_token": {
          "type": "string"
        },
        "account_id": {
          "type": "string"
        },
        "expires_in": {
          "format": "int64",
          "type": "string"
        },
        "refresh_token": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "authValidateResponse": {
      "properties": {
        "is_valid": {
          "format": "boolean",
          "type": "boolean"
        }
      },
      "type": "object"
    }
  }
}
```
