package conf

import "encore.dev/config"

type DBConfig struct {
	Url  config.String
	Size config.Int
}

func foo() {
	// dbConfig := config.Load[*DBConfig]()
	// print(dbConfig)
}
