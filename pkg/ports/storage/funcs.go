package storage

import (
	"log/slog"
	"regexp"
	"strconv"

	"gorm.io/gorm"
)

func GetDatabaseServer(db *gorm.DB) (string, error) {

	var version string
	if err := db.Raw("SELECT version();").Scan(&version).Error; err != nil {
		return "", err
	}
	return version, nil
}

// IsVersionGreaterThanMajor checks if the PostgreSQL version is greater than `major`
func IsVersionGreaterThanMajor(version string, major int) bool {
	re := regexp.MustCompile(`PostgreSQL (\d+)`)
	matches := re.FindStringSubmatch(version)

	if len(matches) > 1 {
		majorVersion, err := strconv.Atoi(matches[1])
		if err != nil {
			slog.Error("Failed to parse PostgreSQL version", slog.Any("err", err))
			return false
		}
		return majorVersion >= major
	}
	slog.Warn("Unable to parse PostgreSQL version string.")
	return false
}
func EnsureDatabaseServerVersion(db *gorm.DB) {

	databaseServerVer, err := GetDatabaseServer(db)
	if err != nil {
		slog.Warn("Get Database server version failed", slog.Any("err", err))
	} else if databaseServerVer != "" {
		slog.Info("Database connection established", "databaseServerVer", databaseServerVer)
		if !IsVersionGreaterThanMajor(databaseServerVer, 15) {
			slog.Warn("Database server upgrade needed", "databaseServerVer", databaseServerVer)

		}
	}
}
