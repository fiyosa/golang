package cmd

import "tutorial/database/seeder"

func InitSeed() {
	if !seeder.RolePermissionSeeder() {
		return
	}
	if !seeder.UserSeeder() {
		return
	}
}
