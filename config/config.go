package config

var Port = getString("PORT", "3000")
var WebFolderPath = getString("WEB_FOLDER_PATH", "")
var EncryptorSecret = getString("ENCRYPTOR_SECRET", "")
var JWTSecret = getString("JWT_SECRET", "")

var DBConnectionString = getString("DB_CONNECTION_STRING", "")
var MongoDBConnectionString = getString("MONGO_DB_CONNECTION_STRING", "")
var MongoDBName = getString("MONGO_DB_NAME", "")
