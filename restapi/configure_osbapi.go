// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"
	graceful "github.com/tylerb/graceful"

	"osb_api/models"
	"osb_api/restapi/operations"
	"osb_api/restapi/operations/delete"
	"osb_api/restapi/operations/get"
	"osb_api/restapi/operations/post"
	"osb_api/restapi/operations/put"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name osbapi --spec ../combined-swagger.yaml

var serviceInstances = make(map[string]*models.ServiceResponse)

var backupInstances = make(map[string]map[string]*models.Backup)

var restoreInstances = make(map[string]map[string]*models.Restore)

func getID() string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(time.Now().Format("2006-01-02 15:04:05"), "")
}

func newBackup(instanceID string, backupInstance *models.NewBackup) *models.Backup {

	backupID := getID()

	payload := &models.Backup{
		BackupDate:             time.Now().Format("2006-01-02 15:04:05"),
		BackupID:               backupID,
		BackupLocation:         backupInstance.BackupLocation,
		BackupLocationProtocol: backupInstance.BackupLocationProtocol,
		BackupType:             backupInstance.BackupType,
		Compression:            backupInstance.Compression,
		Encryption:             backupInstance.Encryption,
		Expires:                backupInstance.Expires,
		Note:                   backupInstance.Note,
		Size:                   rand.Int63(),
		User:                   "Sanjeev",
	}

	if backupInstances[instanceID] == nil {
		backupInstances[instanceID] = make(map[string]*models.Backup)
	}

	backupInstances[instanceID][backupID] = payload

	return payload
}

func updateBackup(instanceID string, backupID string, backupInstance *models.UpdateBackup) *models.Backup {

	backupInstances[instanceID][backupID].Expires = backupInstance.Expires
	backupInstances[instanceID][backupID].BackupLocation = backupInstance.BackupLocation
	backupInstances[instanceID][backupID].Note = backupInstance.Note
	backupInstances[instanceID][backupID].Compression = backupInstance.Compression
	backupInstances[instanceID][backupID].BackupLocationProtocol = backupInstance.BackupLocationProtocol
	backupInstances[instanceID][backupID].Encryption = backupInstance.Encryption

	return backupInstances[instanceID][backupID]
}

func newRestore(instanceID string, restoreInstance *models.NewRestore) *models.Restore {

	RestoreID := getID()

	payload := &models.Restore{
		RestoreDate:    time.Now().Format("2006-01-02 15:04:05"),
		RestoreID:      RestoreID,
		BackupLocation: "todo",
		Note:           restoreInstance.Note,
		User:           "Mike",
	}

	if restoreInstances[instanceID] == nil {
		restoreInstances[instanceID] = make(map[string]*models.Restore)
	}

	restoreInstances[instanceID][RestoreID] = payload

	return payload
}

func configureFlags(api *operations.OsbapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OsbapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PostCreateBackupHandler = post.CreateBackupHandlerFunc(func(params post.CreateBackupParams) middleware.Responder {
		return get.NewGetBackupOK().WithPayload(newBackup(params.InstanceID, params.BackupBody))
	})
	api.PostCreateRestoreHandler = post.CreateRestoreHandlerFunc(func(params post.CreateRestoreParams) middleware.Responder {
		return post.NewCreateRestoreCreated().WithPayload(newRestore(params.InstanceID, params.RestoreBody))
	})
	api.PutCreateServiceInstanceHandler = put.CreateServiceInstanceHandlerFunc(func(params put.CreateServiceInstanceParams) middleware.Responder {

		var discURL string

		switch pick := rand.Intn(2); pick {
		case 0:
			discURL = "/v2/actions/swagger"
		default:
			discURL = "/v2/actions/swagger2"
		}

		payload := &models.ServiceResponse{
			DashboardURL: "some dashboard url",
			DiscoveryURL: discURL,
			Operation:    "created",
		}

		serviceInstances[params.InstanceID] = payload

		return put.NewCreateServiceInstanceCreated().WithPayload(payload)
	})
	api.DeleteDeleteBackupHandler = delete.DeleteBackupHandlerFunc(func(params delete.DeleteBackupParams) middleware.Responder {
		backupInstances[params.InstanceID][params.BackupID] = nil

		return delete.NewDeleteBackupOK()
	})
	api.GetGetBackupHandler = get.GetBackupHandlerFunc(func(params get.GetBackupParams) middleware.Responder {
		return get.NewGetBackupOK().WithPayload(backupInstances[params.InstanceID][params.BackupID])
	})
	api.GetGetBackupsHandler = get.GetBackupsHandlerFunc(func(params get.GetBackupsParams) middleware.Responder {

		var backups = []*models.Backup{}

		for k := range backupInstances[params.InstanceID] {
			backups = append(backups, backupInstances[params.InstanceID][k])
		}

		return get.NewGetBackupsOK().WithPayload(backups)
	})
	api.GetGetRestoresHandler = get.GetRestoresHandlerFunc(func(params get.GetRestoresParams) middleware.Responder {

		var restores = []*models.Restore{}

		for k := range restoreInstances[params.InstanceID] {
			restores = append(restores, restoreInstances[params.InstanceID][k])
		}

		return get.NewGetRestoresOK().WithPayload(restores)
	})
	api.GetGetSwaggerHandler = get.GetSwaggerHandlerFunc(func(params get.GetSwaggerParams) middleware.Responder {
		return get.NewGetSwaggerOK().WithPayload(bandrSwagger)
	})
	api.PutUpdateBackupHandler = put.UpdateBackupHandlerFunc(func(params put.UpdateBackupParams) middleware.Responder {
		return put.NewUpdateBackupOK().WithPayload(updateBackup(params.InstanceID, params.BackupID, params.UpdateBackupBody))
	})
	api.GetGetSwagger2Handler = get.GetSwagger2HandlerFunc(func(params get.GetSwagger2Params) middleware.Responder {
		return get.NewGetSwagger2OK().WithPayload(ssprSwagger)
	})
	api.PutPauseServiceHandler = put.PauseServiceHandlerFunc(func(params put.PauseServiceParams) middleware.Responder {
		return put.NewPauseServiceOK()
	})
	api.PutRestartServiceHandler = put.RestartServiceHandlerFunc(func(params put.RestartServiceParams) middleware.Responder {
		return put.NewRestartServiceOK()
	})
	api.PutStartServiceHandler = put.StartServiceHandlerFunc(func(params put.StartServiceParams) middleware.Responder {
		return put.NewStartServiceOK()
	})
	api.PutStopServiceHandler = put.StopServiceHandlerFunc(func(params put.StopServiceParams) middleware.Responder {
		return put.NewStopServiceOK()
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
        handleCORS := cors.New(cors.Options{
                AllowedMethods:[]string{"OPTIONS", "GET", "DELETE", "POST", "PUT"},
                AllowedHeaders:[]string{"content-type", "Authorization"},
    AllowCredentials: true,
                })
        return handleCORS.Handler(handler)
}
