package database

import "database/sql"

type Champion struct {
	Name string
	// Image string
	Role     string
	StrongVS string
	WeakVS   string
}

type Website struct{
	*sql.DB
}
