//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Session = newSessionTable("", "session", "")

type sessionTable struct {
	sqlite.Table

	// Columns
	ID        sqlite.ColumnString
	UserUUID  sqlite.ColumnString
	Name      sqlite.ColumnString
	ExpiresAt sqlite.ColumnTimestamp

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type SessionTable struct {
	sessionTable

	EXCLUDED sessionTable
}

// AS creates new SessionTable with assigned alias
func (a SessionTable) AS(alias string) *SessionTable {
	return newSessionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SessionTable with assigned schema name
func (a SessionTable) FromSchema(schemaName string) *SessionTable {
	return newSessionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SessionTable with assigned table prefix
func (a SessionTable) WithPrefix(prefix string) *SessionTable {
	return newSessionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SessionTable with assigned table suffix
func (a SessionTable) WithSuffix(suffix string) *SessionTable {
	return newSessionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSessionTable(schemaName, tableName, alias string) *SessionTable {
	return &SessionTable{
		sessionTable: newSessionTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newSessionTableImpl("", "excluded", ""),
	}
}

func newSessionTableImpl(schemaName, tableName, alias string) sessionTable {
	var (
		IDColumn        = sqlite.StringColumn("id")
		UserUUIDColumn  = sqlite.StringColumn("user_uuid")
		NameColumn      = sqlite.StringColumn("name")
		ExpiresAtColumn = sqlite.TimestampColumn("expires_at")
		allColumns      = sqlite.ColumnList{IDColumn, UserUUIDColumn, NameColumn, ExpiresAtColumn}
		mutableColumns  = sqlite.ColumnList{NameColumn, ExpiresAtColumn}
	)

	return sessionTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserUUID:  UserUUIDColumn,
		Name:      NameColumn,
		ExpiresAt: ExpiresAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
