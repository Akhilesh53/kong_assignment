package config

import "sync"

type envVars struct {
	postgresHost     string
	postgresPort     string
	postgresUser     string
	postgresPassword string
	postgresDB       string
	processName      string
	logLevel         string
	environment      string
}

var envVarsOnce sync.Once
var envVarsInstance *envVars

func GetEnvVars() *envVars {
	envVarsOnce.Do(func() {
		envVarsInstance = &envVars{
			postgresHost:     "127.0.0.1",
			postgresPort:     "5432",
			postgresUser:     "postgres",
			postgresPassword: "AkhileshPostgres",
			postgresDB:       "postgres",
			processName:      "kong_assignment",
			logLevel:         "info",
			environment:      "prod",
		}
	})
	return envVarsInstance
}

// get methods for env vars
func (e *envVars) GetPostgresHost() string {
	return e.postgresHost
}

func (e *envVars) GetPostgresPort() string {
	return e.postgresPort
}

func (e *envVars) GetPostgresUser() string {
	return e.postgresUser
}

func (e *envVars) GetPostgresPassword() string {
	return e.postgresPassword
}

func (e *envVars) GetPostgresDB() string {
	return e.postgresDB
}

func (e *envVars) GetProcessName() string {
	return e.processName
}

func (e *envVars) GetLogLevel() string {
	return e.logLevel
}

func (e *envVars) GetEnvironment() string {
	return e.environment
}
