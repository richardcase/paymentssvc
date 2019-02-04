package main

import (
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/richardcase/paymentssvc/pkg/config"
	"github.com/richardcase/paymentssvc/pkg/domain"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi"
	"github.com/richardcase/paymentssvc/pkg/gen/restapi/operations"
	"github.com/richardcase/paymentssvc/pkg/handlers"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	eh "github.com/looplab/eventhorizon"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
	repo "github.com/looplab/eventhorizon/repo/mongodb"
)

var (
	dbURL    string
	dbPrefix string
	colName  string
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	flag.StringVar(&dbURL, "dburl", "localhost:27017", "The URL to the database")
	flag.StringVar(&dbPrefix, "dbprefix", "contoso", "The prefix to use when creating the datastore")
	flag.StringVar(&colName, "colname", "payments", "The database collection name to use")
}

func main() {

	logger, err := configureLogging()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to configure logging")
	}

	logger.Info("starting payments service")
	logger.Infof("Version: %s", version)
	logger.Infof("Build Date: %s", date)
	logger.Infof("Git Commit: %s", commit)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  payments-server [OPTIONS]\n\n")

		title := "Payments API"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "This is a sample Payments API"
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	// Configure event horizin dependencies
	eventStore, err := eventstore.NewEventStore(dbURL, dbPrefix)
	if err != nil {
		logger.Fatalf("could not create event store: %s", err)
	}

	eventBus := eventbus.NewEventBus(nil)
	go func() {
		for e := range eventBus.Errors() {
			log.Printf("eventbus: %s", e.Error())
		}
	}()

	repo, err := repo.NewRepo(dbURL, dbPrefix, colName)
	if err != nil {
		logger.Fatalf("could not create payments repository: %s", err)
	}
	repo.SetEntityFactory(func() eh.Entity { return &domain.Payment{} }) //TODO: this needs to be moved

	config := config.Config{
		Logger:     logger,
		EventStore: eventStore,
		EventBus:   eventBus,
		Repo:       repo,
	}

	api := operations.NewPaymentsAPI(swaggerSpec)

	// Add our handlers
	_, err = handlers.New(config, api)
	if err != nil {
		logger.Fatalf("could not create new handler: %s", err)
	}

	// get server with flag values filled out
	server = restapi.NewServer(api)

	//nolint
	defer server.Shutdown()

	//server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		logger.Fatalln(err)
	}
}

func configureLogging() (*logrus.Entry, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, errors.Wrap(err, "getting hostname")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter((&logrus.JSONFormatter{}))

	fields := logrus.Fields{
		"hostname":  hostname,
		"component": "payments-svc",
	}

	return logrus.StandardLogger().WithFields(fields), nil
}
