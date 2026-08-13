package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"log"
	//"net/http"
	"os"
	"path/filepath"
	"runtime"
)


type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepository UserRepository
	templateDir string
	tp *TemplateRenderer
}


func main() {
	
	db, err := connectToDatabase("users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		userRepository: NewSQLUserRepository(db),
		templateDir: templateDir(),
		tp: NewTemplateRenderer(templateDir(), false),
	}

	log.Println("Listening on :8080")
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

func templateDir() string {
	_, thisFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(thisFile), "templates")
}

func connectToDatabase(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", name)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
		