package facades

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global variable for the database connection
var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using environment variables instead.")
	} else {
		log.Println(".env file loaded successfully")
	}

	// Validate required environment variables
	// requiredEnvVars := []string{"MYSQL_HOST", "MYSQL_PORT", "MYSQL_DB", "MYSQL_USER", "MYSQL_PASS"}
	// for _, v := range requiredEnvVars {
	//     if os.Getenv(v) == "" {
	//         log.Fatalf("Environment variable %s is not set", v)
	//     }
	// }
	requiredEnvVars := []string{"MYSQL_HOST", "MYSQL_PORT", "MYSQL_DB", "MYSQL_USER"}
	for _, v := range requiredEnvVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Environment variable %s is not set", v)
		}
	}

	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")

	// Build DSN (Data Source Name) for GORM
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName +
		"?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s&readTimeout=10s&writeTimeout=10s"

	// Open the database connection using GORM with a custom logger
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 500 * time.Millisecond, // Lower slow SQL threshold
				LogLevel:      logger.Warn,            // Log level
				Colorful:      true,                   // Enable colorful logs
			},
		),
		PrepareStmt: true, // Enable prepared statement caching
	})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	// Get the database SQL object for connection pooling configuration
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get database object from GORM:", err)
	}

	// Connection pooling configuration
	sqlDB.SetMaxIdleConns(15)                  // Increase max idle connections
	sqlDB.SetMaxOpenConns(200)                 // Increase max open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Reduce connection max lifetime
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Set connection max idle time

	log.Println("Database connection successfully established")
}
