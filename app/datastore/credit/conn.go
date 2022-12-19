package credit

import (
	"context"
	"fmt"
	"sync"

	"carbX/app/system/env"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
    dbPool *pgxpool.Pool
    dbOnce sync.Once
)

func RwInstancePool() *pgxpool.Pool {
    dbOnce.Do(func() {
        conString := createConnectionString()
        db, err := pgxpool.Connect(context.Background(), conString)
        if err != nil {
            panic(err)
        }

        if err := db.Ping(context.Background()); err != nil {
            panic(err)
        }

        dbPool = db
    })
    return dbPool
}

func createConnectionString() string {
    return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.GetConfig().CreditDBUser, env.GetConfig().CreditDBPassword, env.GetConfig().CreditDBHost, env.GetConfig().CreditDBPort, "postgres")
}