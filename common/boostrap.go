package common

func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT authentication
	initKeys()
	// Start MongoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}
