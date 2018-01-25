package common

// StartUp bootstrapps the application
func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Initialize Logger objects with Log Level
	setLogLevel(Level(AppConfig.LogLevel))
	// Start a MongoDB session
	createDBSession()
	// Add indexes into MongoDB
	addIndexes()
}
