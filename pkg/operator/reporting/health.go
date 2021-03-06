package reporting

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"

	"github.com/bentitopolis/operator-metering/pkg/db"
	"github.com/bentitopolis/operator-metering/pkg/hive"
	"github.com/bentitopolis/operator-metering/pkg/presto"
)

type PrestoHealthChecker struct {
	logger       logrus.FieldLogger
	queryer      db.Queryer
	tableManager HiveTableManager

	databaseName, tableName string
	// ensures only at most a single testRead query is running against Presto
	// at one time
	healthCheckSingleFlight singleflight.Group
}

func NewPrestoHealthChecker(logger logrus.FieldLogger, queryer db.Queryer, tableManager HiveTableManager, databaseName, tableName string) *PrestoHealthChecker {
	return &PrestoHealthChecker{
		logger:       logger,
		queryer:      queryer,
		tableManager: tableManager,
		databaseName: databaseName,
		tableName:    tableName,
	}
}

func (checker *PrestoHealthChecker) TestWriteToPrestoSingleFlight() bool {
	const key = "presto-write"
	v, _, _ := checker.healthCheckSingleFlight.Do(key, func() (interface{}, error) {
		defer checker.healthCheckSingleFlight.Forget(key)
		healthy := checker.TestWriteToPresto()
		return healthy, nil
	})
	healthy := v.(bool)
	return healthy
}

func (checker *PrestoHealthChecker) TestReadFromPrestoSingleFlight() bool {
	const key = "presto-read"
	v, _, _ := checker.healthCheckSingleFlight.Do(key, func() (interface{}, error) {
		defer checker.healthCheckSingleFlight.Forget(key)
		healthy := checker.TestReadFromPresto()
		return healthy, nil
	})
	healthy := v.(bool)
	return healthy
}

func (checker *PrestoHealthChecker) TestReadFromPresto() bool {
	_, err := presto.ExecuteSelect(checker.queryer, "SELECT * FROM system.runtime.nodes")
	if err != nil {
		checker.logger.WithError(err).Debugf("cannot query Presto system.runtime.nodes table")
		return false
	}
	return true
}

func (checker *PrestoHealthChecker) TestWriteToPresto() bool {
	logger := checker.logger.WithField("component", "testWriteToPresto")
	columns := []hive.Column{{Name: "check_time", Type: "TIMESTAMP"}}

	params := hive.TableParameters{
		Database: checker.databaseName,
		Name:     checker.tableName,
		Columns:  columns,
	}
	err := checker.tableManager.CreateTable(params, true)
	if err != nil {
		logger.WithError(err).Errorf("cannot create Presto table %s", checker.tableName)
		return false
	}

	// Hive does not support timezones, and now() returns a
	// TIMESTAMP WITH TIMEZONE so we cast the return of now() to a TIMESTAMP.
	err = presto.InsertInto(checker.queryer, checker.tableName, "VALUES (cast(now() AS TIMESTAMP))")
	if err != nil {
		logger.WithError(err).Errorf("cannot insert into Presto table %s", checker.tableName)
		return false
	}
	return true
}
