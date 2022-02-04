// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"backend/ent"
	"backend/server/graph/resolver"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/alecthomas/kong"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"

	_ "backend/ent/runtime"

	_ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	postgresHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		panic("POSTGRES_HOST env variable required!")
	}
	postgresPort, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		panic("POSTGRES_PORT env variable required!")
	}
	postgresSSL, ok := os.LookupEnv("POSTGRES_SSL")
	if !ok {
		panic("POSTGRES_SSL env variable required!")
	}
	postgresUsername, ok := os.LookupEnv("POSTGRES_USERNAME")
	if !ok {
		panic("POSTGRES_USERNAME env variable required!")
	}
	postgresPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		panic("POSTGRES_PASSWORD env variable required!")
	}
	minioAddress, ok := os.LookupEnv("MINIO_URL")
	if !ok {
		panic("MINIO_URL env variable required!")
	}
	minioAccessKey, ok := os.LookupEnv("MINIO_ACCESS_KEY")
	if !ok {
		panic("MINIO_ACCESS_KEY env variable required!")
	}

	minioSecretKey, ok := os.LookupEnv("MINIO_SECRET_KEY")
	if !ok {
		panic("MINIO_SECRET_KEY env variable required!")
	}

	_, ok = os.LookupEnv("APP_SECRET")
	if !ok {
		panic("APP_SECRET env variable not found!")
	}

	port, ok := os.LookupEnv("BACKEND_PORT")
	if !ok {
		port = "8085"
	}

	minioURL, err := url.Parse(minioAddress)
	if err != nil {
		panic(err)
	}

	minioClient, err := minio.New(minioURL.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}

	var cli struct {
		Debug bool `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	log, _ := zap.NewDevelopment()
	// client, err := ent.Open(
	// 	"sqlite3",
	// 	"file:ent.sql?cache=shared&_fk=1",
	// )
	sslMode := "enable"
	postgresAddr := fmt.Sprintf("https://%s:%s", postgresHost, postgresPort)

	if postgresSSL == "false" {
		sslMode = "disable"
		postgresAddr = fmt.Sprintf("http://%s:%s", postgresHost, postgresPort)
	}

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s", postgresHost, postgresPort, postgresUsername, postgresPassword, sslMode)

	client, err := ent.Open(dialect.Postgres, dbinfo)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
	}

	for {
		db, err := sql.Open("postgres", dbinfo)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			} else {
				log.Info(err.Error())
			}
		} else {
			log.Info(err.Error())
		}
		log.Info("waiting for postgreSQL on ", zap.String("address", postgresAddr))
		time.Sleep(time.Second)
	}

	client = client.Debug()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("running schema migration", zap.Error(err))
	}

	srv := handler.NewDefaultServer(resolver.NewSchema(client, minioClient))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if cli.Debug {
		srv.Use(&debug.Tracer{})
	}

	router := chi.NewRouter()
	// Inject middleware
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Handle("/graphql", srv)

	log.Info("listening on", zap.String("address", port), zap.String("graphql", ":"+port+"/graphql"))
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Error("http server terminated", zap.Error(err))
	}
}
