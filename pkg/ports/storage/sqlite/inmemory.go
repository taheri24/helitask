package sqlite

import (
	"bufio"
	"database/sql"
	"embed"
	_ "embed"
	"strings"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/taheri24/helitask/pkg/domain"
	"gorm.io/gorm"
)

//go:embed *.sql
var sqlFiles embed.FS

func NewDb(t *testing.T, scriptName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&domain.TodoItem{})
	if dbConn, err := db.DB(); err != nil {
		panic(err)
	} else if scriptName != "" {
		runScript(scriptName, dbConn)
	}
	return db
}

func runScript(scriptName string, dbConn *sql.DB) {
	fileHand, err := sqlFiles.Open(scriptName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fileHand)
	sb := strings.Builder{}
	for scanner.Scan() {
		queryLn := strings.TrimSpace(scanner.Text())

		sb.WriteString(queryLn + "\n")

		if strings.HasSuffix(queryLn, ";") {
			dbConn.Exec(sb.String())
			sb = strings.Builder{}
		}
	}
}
