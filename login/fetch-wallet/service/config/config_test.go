package config

import (
	"fmt"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	tableName := "blitzbudget"
	// Set environment variable
	os.Setenv("TABLE_NAME", tableName)
	// returns value of TABLE_NAME
	fmt.Println("TABLE_NAME:", os.Getenv("TABLE_NAME"))
	TableName = os.Getenv("TABLE_NAME")

	if TableName != tableName {
		t.Errorf("QueryParameter: TableName do not match, got = %v, want = %v \n", TableName, tableName)
		return
	}

}
