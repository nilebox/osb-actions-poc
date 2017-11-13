package restapi

import (
	"encoding/json"
)

var bandrSwagger = json.RawMessage([]byte(`{
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
	  "description": "This swagger specification defines the recommended approach to enabling backup and restore actions via an Open Service Broker API implementation.",
	  "title": "Open Service Broker - Backup \u0026 Restore Actions",
	  "version": "0.1"
	},
	"host": "localhost:8080",
	"basePath": "/",
	"paths": {
	  "/v2/service_instances/{instance_id}/backups": {
		"get": {
		  "tags": [
			"get"
		  ],
		  "summary": "List all of the backups available on this service instance.",
		  "operationId": "getBackups",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation",
			  "schema": {
				"$ref": "#/definitions/getBackupsOKBody"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"post": {
		  "tags": [
			"post"
		  ],
		  "summary": "Create a new backup for this service instance.",
		  "operationId": "createBackup",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			},
			{
			  "description": "The details of the backup to be taken.",
			  "name": "backup_body",
			  "in": "body",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/new_backup"
			  }
			}
		  ],
		  "responses": {
			"201": {
			  "description": "created",
			  "schema": {
				"$ref": "#/definitions/backup"
			  }
			},
			"202": {
			  "description": "accepted",
			  "schema": {
				"$ref": "#/definitions/status"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/v2/service_instances/{instance_id}/backups/{backup_id}": {
		"get": {
		  "tags": [
			"get"
		  ],
		  "summary": "Get the details of a backup on a service instance.",
		  "operationId": "getBackup",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			},
			{
			  "type": "string",
			  "description": "The id of the backup for a service instance.",
			  "name": "backup_id",
			  "in": "path",
			  "required": true
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation",
			  "schema": {
				"$ref": "#/definitions/backup"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"put": {
		  "tags": [
			"put"
		  ],
		  "summary": "Update the details of a backup on a service instance.",
		  "operationId": "updateBackup",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			},
			{
			  "type": "string",
			  "description": "The id of the backup for a service instance.",
			  "name": "backup_id",
			  "in": "path",
			  "required": true
			},
			{
			  "description": "The details of the backup to be updated.",
			  "name": "update_backup_body",
			  "in": "body",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/update_backup"
			  }
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation",
			  "schema": {
				"$ref": "#/definitions/backup"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"delete": {
		  "tags": [
			"delete"
		  ],
		  "summary": "Delete a backup for this service instance.",
		  "operationId": "deleteBackup",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			},
			{
			  "type": "string",
			  "description": "The id of the backup for a service instance.",
			  "name": "backup_id",
			  "in": "path",
			  "required": true
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation"
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/v2/service_instances/{instance_id}/restores": {
		"get": {
		  "tags": [
			"get"
		  ],
		  "summary": "List all of the previously restored backups on this service instance",
		  "operationId": "getRestores",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation",
			  "schema": {
				"$ref": "#/definitions/getRestoresOKBody"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		},
		"post": {
		  "tags": [
			"post"
		  ],
		  "summary": "Restore a service instance.",
		  "operationId": "createRestore",
		  "parameters": [
			{
			  "type": "string",
			  "description": "The id of the service instance.",
			  "name": "instance_id",
			  "in": "path",
			  "required": true
			},
			{
			  "description": "The details of the restore.",
			  "name": "restore_body",
			  "in": "body",
			  "required": true,
			  "schema": {
				"$ref": "#/definitions/new_restore"
			  }
			}
		  ],
		  "responses": {
			"201": {
			  "description": "created",
			  "schema": {
				"$ref": "#/definitions/restore"
			  }
			},
			"202": {
			  "description": "accepted",
			  "schema": {
				"$ref": "#/definitions/status"
			  }
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  }
	},
	"definitions": {
	  "backup": {
		"type": "object",
		"properties": {
		  "backup_date": {
			"description": "The date/time when the backup was initiated.",
			"type": "string",
			"example": "2017-09-11T21:52:13+00:00"
		  },
		  "backup_id": {
			"description": "A unique id for backups on this service instance.",
			"type": "string",
			"example": "1234"
		  },
		  "backup_location": {
			"description": "The path or url to the backup file. Use 'none' for the default loction.",
			"type": "string",
			"example": "http://s3.amazonaws.com/bucket/backup1.tar"
		  },
		  "backup_location_credentials": {
			"description": "The credentials of the backup location.",
			"type": "string",
			"example": "admin/password"
		  },
		  "backup_location_protocol": {
			"description": "The protocol of the backup location.",
			"type": "string",
			"enum": [
			  "nfs",
			  "s3"
			],
			"example": "s3"
		  },
		  "backup_type": {
			"description": "A unique id for backups on this service instance.",
			"type": "string",
			"enum": [
			  "full",
			  "incremental",
			  "logical",
			  "physical"
			],
			"example": "incremental"
		  },
		  "compression": {
			"description": "The type of compression to be used for the backup file. Use 'none' for no compression.",
			"type": "string",
			"enum": [
			  "none",
			  "tar",
			  "zip"
			],
			"example": "tar"
		  },
		  "encryption": {
			"description": "The type of encryption to be used for the backup file. Use 'none' for no encryption.",
			"type": "string",
			"default": "none",
			"enum": [
			  "none",
			  "AES128",
			  "DES128",
			  "RSA"
			],
			"example": "AES128"
		  },
		  "expires": {
			"description": "The date/time when the backup file will be purged from the system.",
			"type": "string",
			"example": "12/31/2020"
		  },
		  "note": {
			"description": "A custom message about the specific backup action.",
			"type": "string",
			"example": "Changing expiration date of this backup."
		  },
		  "size": {
			"description": "The size (in MB) of the backup file.",
			"type": "integer",
			"example": "203"
		  },
		  "user": {
			"description": "The username of the user that initiated the backup.",
			"type": "string",
			"example": "John Adams"
		  }
		}
	  },
	  "error": {
		"required": [
		  "code",
		  "message"
		],
		"properties": {
		  "code": {
			"type": "integer",
			"format": "int32"
		  },
		  "message": {
			"type": "string"
		  }
		}
	  },
	  "getBackupsOKBody": {
		"type": "array",
		"items": {
		  "$ref": "#/definitions/backup"
		},
		"x-go-gen-location": "operations"
	  },
	  "getRestoresOKBody": {
		"type": "array",
		"items": {
		  "$ref": "#/definitions/restore"
		},
		"x-go-gen-location": "operations"
	  },
	  "new_backup": {
		"type": "object",
		"properties": {
		  "backup_location": {
			"description": "The path or url to the backup file. Use 'none' for the default loction.",
			"type": "string",
			"example": "http://s3.amazonaws.com/bucket/backup1.tar"
		  },
		  "backup_location_credentials": {
			"description": "The credentials of the backup location.",
			"type": "string",
			"example": "admin/password"
		  },
		  "backup_location_protocol": {
			"description": "The protocol of the backup location.",
			"type": "string",
			"enum": [
			  "nfs",
			  "s3"
			],
			"example": "s3"
		  },
		  "backup_type": {
			"description": "A unique id for backups on this service instance.",
			"type": "string",
			"enum": [
			  "full",
			  "incremental",
			  "logical",
			  "physical"
			],
			"example": "incremental"
		  },
		  "compression": {
			"description": "The type of compression to be used for the backup file. Use 'none' for no compression.",
			"type": "string",
			"enum": [
			  "none",
			  "tar",
			  "zip"
			],
			"example": "tar"
		  },
		  "encryption": {
			"description": "The type of encryption to be used for the backup file. Use 'none' for no encryption.",
			"type": "string",
			"default": "none",
			"enum": [
			  "none",
			  "AES128",
			  "DES128",
			  "RSA"
			],
			"example": "AES128"
		  },
		  "expires": {
			"description": "The date/time when the backup file will be purged from the system.",
			"type": "string",
			"example": "12/31/2020"
		  },
		  "note": {
			"description": "A custom message about the specific backup action.",
			"type": "string",
			"example": "Changing expiration date of this backup."
		  }
		}
	  },
	  "new_restore": {
		"type": "object",
		"properties": {
		  "backup_location": {
			"description": "Location of the backup file to restore.",
			"type": "string",
			"example": "http://s3.amazonaws.com/bucket/backup1.tar"
		  },
		  "note": {
			"description": "A custom message about the specific restore action.",
			"type": "string",
			"example": "Had to restore from an error at 12pm EST."
		  }
		}
	  },
	  "restore": {
		"type": "object",
		"properties": {
		  "backup_location": {
			"description": "Location of the backup file to restore.",
			"type": "string",
			"example": "http://s3.amazonaws.com/bucket/backup1.tar"
		  },
		  "note": {
			"description": "A custom message about the specific restore action.",
			"type": "string",
			"example": "Had to restore from an error at 12pm EST."
		  },
		  "restore_date": {
			"description": "The date/time when the restore was initiated.",
			"type": "string",
			"example": "2017-09-11T21:52:13+00:00"
		  },
		  "restore_id": {
			"description": "A unique id for restore on this service instance.",
			"type": "string",
			"example": "5678"
		  },
		  "user": {
			"description": "The username of the user that initiated the restore.",
			"type": "string",
			"example": "John Adams"
		  }
		}
	  },
	  "status": {
		"type": "string"
	  },
	  "update_backup": {
		"type": "object",
		"properties": {
		  "backup_location": {
			"description": "The path or url to the backup file. Use 'none' for the default loction.",
			"type": "string",
			"example": "http://s3.amazonaws.com/bucket/backup1.tar"
		  },
		  "backup_location_credentials": {
			"description": "The credentials of the backup location.",
			"type": "string",
			"example": "admin/password"
		  },
		  "backup_location_protocol": {
			"description": "The protocol of the backup location.",
			"type": "string",
			"enum": [
			  "nfs",
			  "s3"
			],
			"example": "s3"
		  },
		  "compression": {
			"description": "The type of compression to be used for the backup file. Use 'none' for no compression.",
			"type": "string",
			"enum": [
			  "none",
			  "tar",
			  "zip"
			],
			"example": "tar"
		  },
		  "encryption": {
			"description": "The type of encryption to be used for the backup file. Use 'none' for no encryption.",
			"type": "string",
			"default": "none",
			"enum": [
			  "none",
			  "AES128",
			  "DES128",
			  "RSA"
			],
			"example": "AES128"
		  },
		  "expires": {
			"description": "The date/time when the backup file will be purged from the system.",
			"type": "string",
			"example": "12/31/2020"
		  },
		  "note": {
			"description": "A custom message about the specific backup action.",
			"type": "string",
			"example": "Changing expiration date of this backup."
		  }
		}
	  }
	},
	"parameters": {
	  "backup_body": {
		"description": "The details of the backup to be taken.",
		"name": "backup_body",
		"in": "body",
		"required": true,
		"schema": {
		  "$ref": "#/definitions/new_backup"
		}
	  },
	  "backup_id": {
		"type": "string",
		"description": "The id of the backup for a service instance.",
		"name": "backup_id",
		"in": "path",
		"required": true
	  },
	  "instance_id": {
		"type": "string",
		"description": "The id of the service instance.",
		"name": "instance_id",
		"in": "path",
		"required": true
	  },
	  "restore_body": {
		"description": "The details of the restore.",
		"name": "restore_body",
		"in": "body",
		"required": true,
		"schema": {
		  "$ref": "#/definitions/new_restore"
		}
	  },
	  "update_backup_body": {
		"description": "The details of the backup to be updated.",
		"name": "update_backup_body",
		"in": "body",
		"required": true,
		"schema": {
		  "$ref": "#/definitions/update_backup"
		}
	  }
	}
  }`))

var ssprSwagger = json.RawMessage([]byte(`{
	"swagger": "2.0",
	"info": {
	  "version": "0.1",
	  "title": "Open Service Broker - Start, Stop, Pause, Restart Actions",
	  "description": "This swagger specification defines the recommended approach to enabling start, stop, puase and restart actions via an Open Service Broker API implementation."
	},
	"host": "localhost:8080",
	"basePath": "/",
	"schemes": [
	  "http"
	],
	"produces": [
	  "application/json"
	],
	"consumes": [
	  "application/json"
	],
	"parameters": {
	  "instance_id": {
		"name": "instance_id",
		"in": "path",
		"description": "The id of the service instance.",
		"type": "string",
		"required": true
	  }
	},
	"paths": {
	  "/v2/service_instances/{instance_id}/start": {
		"put": {
		  "tags": [
			"put"
		  ],
		  "summary": "Start a service instance that is in a STOPPED state.",
		  "description": "",
		  "operationId": "startService",
		  "parameters": [
			{
			  "$ref": "#/parameters/instance_id"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation"
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/v2/service_instances/{instance_id}/stop": {
		"put": {
		  "tags": [
			"put"
		  ],
		  "summary": "Stop a service instance that is in a RUNNING state.",
		  "description": "",
		  "operationId": "stopService",
		  "parameters": [
			{
			  "$ref": "#/parameters/instance_id"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation"
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/v2/service_instances/{instance_id}/restart": {
		"put": {
		  "tags": [
			"put"
		  ],
		  "summary": "Restart a service instance that is in a RUNNING state.",
		  "description": "",
		  "operationId": "restartService",
		  "parameters": [
			{
			  "$ref": "#/parameters/instance_id"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation"
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  },
	  "/v2/service_instances/{instance_id}/pause": {
		"put": {
		  "tags": [
			"put"
		  ],
		  "summary": "Pause a service instance that is in a RUNNING state.",
		  "description": "",
		  "operationId": "pauseService",
		  "parameters": [
			{
			  "$ref": "#/parameters/instance_id"
			}
		  ],
		  "responses": {
			"200": {
			  "description": "successful operation"
			},
			"default": {
			  "description": "unexpected error",
			  "schema": {
				"$ref": "#/definitions/error"
			  }
			}
		  }
		}
	  }
	},
	"definitions": {
	  "error": {
		"required": [
		  "code",
		  "message"
		],
		"properties": {
		  "code": {
			"type": "integer",
			"format": "int32"
		  },
		  "message": {
			"type": "string"
		  }
		}
	  }
	}
  }`))
