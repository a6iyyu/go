// Package main is the entry point of the application.
package main

import (
	"flag"
	"fmt"
	"os"

	c "sentra-medika/constants"
	"sentra-medika/seeders"
	s "sentra-medika/services"
	u "sentra-medika/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	u.ConnectToDatabase()

	seed := flag.Bool("seed", false, "Seed initial data into the database")
	flag.Parse()

	if *seed {
		fmt.Println("🌱 Seeding initial data...")

		doctor, patient := seeders.Users(u.DB)

		if doctor.ID != [16]byte{} && patient.ID != [16]byte{} {
			seeders.MedicalRecords(u.DB, doctor.ID, patient.ID)
		}

		fmt.Println("✨ Seeding completed.")
		os.Exit(0)
	}

	r := gin.Default()

	r.POST(c.Login, s.Login)
	r.POST(c.Refresh, s.Refresh)

	protected := r.Group("/")
	protected.Use(s.Middleware())
	{
		protected.POST(c.Logout, s.Logout)

		admin := protected.Group("/admin")
		admin.Use(s.RoleGuard("admin"))
		{
			admin.POST(c.Register, s.Register)
			admin.POST(c.Users, s.CreateUser)
			admin.GET(c.Users, s.GetUsers)
			admin.PUT(c.UserByID, s.UpdateUser)
			admin.DELETE(c.UserByID, s.DeleteUser)
		}

		doctor := protected.Group("/medical")
		doctor.Use(s.RoleGuard("doctor"))
		{
			doctor.POST(c.Records, s.CreateRecord)
			doctor.GET(c.Records, s.GetRecords)
			doctor.PUT(c.RecordByID, s.UpdateRecord)
			doctor.DELETE(c.RecordByID, s.DeleteRecord)
		}

		patient := protected.Group("/medical")
		patient.Use(s.RoleGuard("patient"))
		{
			patient.GET(c.MyRecords, s.GetMyRecords)
		}
	}

	r.Run(":7000")
}