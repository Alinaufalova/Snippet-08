package main
import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	_"github.com/jackc/pgx"

)
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Psql data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	pool, err := pgxpool.Connect(context.Background(), "user = admin password=04317740 host=localhost port=5432 dbname=snippetbox sslmode=disable pool_max_conns=10" )
	if err != nil{
		log.Fatalf("Unable to connection to database: %v\n", err)
	}

	defer pool.Close()

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil


}