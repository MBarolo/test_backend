package config

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type SQLiteConnection struct {
	DB *sql.DB
}

var DB *sql.DB

// NewSQLiteConnection: Crea y retorna una nueva conexión a SQLite
func NewSQLiteConnection() *SQLiteConnection {
	if err := InitDB(); err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err.Error())
	}
	return &SQLiteConnection{DB: DB}
}

// InitDB: Función para incilaizar la base de datos
func InitDB() error {
	// ruta al archivo .db
	path := os.Getenv("SQLITE_PATH")
	if path == "" {
		path = "./app.db"
	}

	var err error
	DB, err = sql.Open("sqlite", path)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		log.Println("Error al conectarse a la base de datos:", err.Error())
		return err
	}

	if err = createTables(); err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// createTables: Función que crea las tablas si es que no existen.
func createTables() error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        hashed_password TEXT NOT NULL,
        first_name TEXT NOT NULL,
        last_name TEXT NOT NULL,
		deleted INTEGER NOT NULL DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS bikes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        is_available INTEGER NOT NULL DEFAULT 1,
        latitude REAL NOT NULL,
        longitude REAL NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		cost_per_minute INTEGER NOT NULL
    );

    CREATE TABLE IF NOT EXISTS rentals (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        bike_id INTEGER NOT NULL,
        rental_status TEXT NOT NULL DEFAULT 'running',
        start_time DATETIME NOT NULL,
        end_time DATETIME,
        start_latitude REAL NOT NULL,
        start_longitude REAL NOT NULL,
        end_latitude REAL,
        end_longitude REAL,
		duration INTEGER,
		cost INTEGER,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (bike_id) REFERENCES bikes(id) ON DELETE CASCADE,
        CHECK (rental_status IN ('running', 'ended'))
    );

    CREATE INDEX IF NOT EXISTS idx_bikes_available ON bikes(is_available);
    CREATE INDEX IF NOT EXISTS idx_rentals_user ON rentals(user_id);
    CREATE INDEX IF NOT EXISTS idx_rentals_bike ON rentals(bike_id);
    CREATE INDEX IF NOT EXISTS idx_rentals_status ON rentals(rental_status);
    `

	_, err := DB.Exec(schema)
	if err != nil {
		log.Println("Error al inicializar la base de datos: ", err.Error())
		return err
	}

	return nil
}
