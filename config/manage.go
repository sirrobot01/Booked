package config

import (
	"booked/common"
	"booked/models"
	"bufio"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strings"
)

type Manager struct {
	Debug bool
	Host  string
	Port  string
	ENV   string
	DB    *gorm.DB
}

func (m *Manager) Runserver(migrate bool) {
	if migrate {
		m.Migrate()
	} else {
		m.InitDB()
	}
	if m.Debug {
		println("Debug mode is on")
	}
	println("Running server on " + m.Host + ":" + m.Port)
	r := Router(m.DB)
	err := r.Run(m.Host + ":" + m.Port)
	if err != nil {
		return
	}
}

func (m *Manager) InitDB() {
	db, err := gorm.Open(sqlite.Open("booked.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	m.DB = db
}

func (m *Manager) Migrate() {
	m.InitDB()
	println("Migrating database")
	err := m.DB.AutoMigrate(models.AllModels...)
	if err != nil {
		return
	}
	println("Migrations done")
}

func (m *Manager) CreateAdmin() {
	m.InitDB()

	println("Creating admin")
	reader := bufio.NewReader(os.Stdin)
	print("Enter username: ")
	username, _ := reader.ReadString('\n')
	print("Enter password: ")
	password, _ := reader.ReadString('\n')

	print("Enter email: ")
	email, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	email = strings.TrimSpace(email)

	var u models.User
	m.DB.Where("username = ?", username).First(&u)
	if u.Username != "" {
		println("User already exists")
		return
	}
	u = *common.CreateUser(username, "", "", email, password, true, true)
	m.DB.Create(&u)
	println("Admin created")
}

func (m *Manager) Test() {
	m.InitDB()
	println("Running tests")
}

func (m *Manager) Run() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	switch os.Args[1] {
	default:
		println("Command not found")
	case "runserver":
		m.Runserver(false) // TODO: Add migrate flag
	case "migrate":
		m.Migrate()
	case "createsuperuser":
		m.CreateAdmin()
	case "test":
		m.Test()
	}
}
