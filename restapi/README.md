Open Service Broker Actions (Demo) Server
====

# Overview

This is an example of how to implement custom/generic actions as part of an Open Service Broker server.

## Pre-requisites
* [goswagger](https://goswagger.io/) swagger code generation tool has been installed. This will enable automatic generation of files from swagger.yml file.
* Within your workspace you have [setup](https://golang.org/doc/code.html) the standard folders expected for a GO workspace, and the GOPATH variable pointing to it.

## How to build
* cd to your go workspace's src folder.
* `git clone <crb git repo path> osb_api`.

**Note**: You must name the destination folder osb_api as go uses folder names as package names.
* cd to the project folder.
* Generate required classes from swagger spec file for building.
* `swagger generate server -f combined-swagger.yaml -A osbapi`.

**Note**: You must name the application 'osbapi' as goswagger uses it to name the packages and files, some of which cannot be modified.
* Download and install dependencies using one of the following methods:
  * Using godeps tool:
    * Run "godep save ./..." to compile the "Godeps/Godeps.json" and generate the lists of required packages in "vendor/".
  * Using native go get command:
    * Run `go get -u -f ./...`
* cd cmd/osbapi-server/
* Run `go build` or `go install`.

## How to run
* Start the server as local standalone app `./osbapi-server --host=localhost --port=8080`


## Example rest calls


```
$ curl -X POST "http://localhost:8080/v2/service_instances/test_instance/backups" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"backup_type\": \"incremental\", \"expires\": \"12/31/2020\", \"backup_location\": \"http://s3.amazonaws.com/bucket/backup1.tar\", \"backup_location_protocol\": \"s3\", \"backup_location_credentials\": \"admin/password\", \"note\": \"Changing expiration date of this backup.\", \"compression\": \"tar\", \"encryption\": \"AES128\"}"
```

```
$ curl -X POST "http://localhost:8080/v2/service_instances/test_instance/restores" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"backup_location\": \"http://s3.amazonaws.com/bucket/backup1.tar\", \"note\": \"Had to restore from an error at 12pm EST.\"}"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test_instance" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"service_id\": \"string\", \"plan_id\": \"string\", \"context\": {}, \"organization_guid\": \"string\", \"space_guid\": \"string\", \"parameters\": {}}"
```

```
$ curl -X GET "http://localhost:8080/v2/service_instances/test_instance/backups" -H "accept: application/json"
```

```
$ curl -X GET "http://localhost:8080/v2/service_instances/test_instance/backups/test_backup" -H "accept: application/json"
```

```
$ curl -X GET "http://localhost:8080/v2/service_instances/test_instance/restores" -H "accept: application/json"
```

```
$ curl -X GET "http://localhost:8080/v2/actions/swagger" -H "accept: application/json"
```

```
$ curl -X GET "http://localhost:8080/v2/actions/swagger2" -H "accept: application/json"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test_instance/backups/backup_instance" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"expires\": \"12/31/2020\", \"backup_location\": \"http://s3.amazonaws.com/bucket/backup1.tar\", \"backup_location_protocol\": \"s3\", \"backup_location_credentials\": \"admin/password\", \"note\": \"Changing expiration date of this backup.\", \"compression\": \"tar\", \"encryption\": \"AES128\"}"
```

```
$ curl -X DELETE "http://localhost:8080/v2/service_instances/test_instance/backups/test_backup" -H "accept: application/json"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test/pause" -H "accept: application/json"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test/restart" -H "accept: application/json"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test/stop" -H "accept: application/json"
```

```
$ curl -X PUT "http://localhost:8080/v2/service_instances/test/start" -H "accept: application/json"
```