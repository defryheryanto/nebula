package main

import (
	"context"
	"database/sql"
	"net"

	"github.com/defryheryanto/nebula/config"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUnixDialer struct {
	Addr string
}

func (d *MongoUnixDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	var dial net.Dialer
	dial.LocalAddr = nil
	raddr := net.UnixAddr{Name: d.Addr, Net: "unix"}
	conn, err := dial.DialContext(ctx, "unix", raddr.String())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func setupDatabaseConnection(ctx context.Context) *sql.DB {
	conn, err := sql.Open("postgres", config.DBConnectionString)
	if err != nil {
		panic(err)
	}

	return conn
}

func setupMongoClient(ctx context.Context) *mongo.Client {
	opt := options.Client()
	if config.MongoDBConnectionType != "unix" {
		opt = opt.ApplyURI(config.MongoDBConnectionString)
	} else {
		opt = opt.SetAuth(options.Credential{
			Username: config.MongoDBUsername,
			Password: config.MongoDBPassword,
		})
		opt.SetDialer(&MongoUnixDialer{
			Addr: config.MongoDBConnectionString,
		})
	}
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}

	return client
}
