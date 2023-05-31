package cmd

import "flag"

func InitCmd() bool {
	option := flag.String("option", "", "Put your option")
	flag.Parse()

	switch *option {
	case "DropAllTable":
		InitDrop()
		return true

	case "MigrateAllTable":
		InitMigrate(true)
		return true

	case "MigrateFreshAllTable":
		InitDrop()
		InitMigrate(false)
		InitSeed()
		return true

	case "SeedData":
		InitSeed()
		return true

	default:
		return false
	}
}
