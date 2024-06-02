package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"github.com/garnet-geo/garnet-geoapi/internal/env"
)

var Conn *pgxpool.Pool

type logrusTracer struct {
	logger *log.Logger
}

func (l *logrusTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	l.logger.Debugln("Executing command", "SQL", data.SQL, "Args", data.Args)
	return ctx
}

func (l *logrusTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// Skip...
}

func InitConnection() {
	config, err := pgxpool.ParseConfig(env.GetDatabaseUrl())
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	config.ConnConfig.Tracer = &logrusTracer{
		logger: log.StandardLogger(),
	}

	Conn, err = pgxpool.NewWithConfig(Context(), config)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	log.Infoln("Created connection to the database")
}

func CloseConnection() {
	Conn.Close()

	log.Infoln("Closed connection to the database")
}

func CheckConnection() error {
	log.Debugln("Checking connection")
	return Conn.Ping(Context())
}

func Context() context.Context {
	return context.Background()
}
