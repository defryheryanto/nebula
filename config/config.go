package config

var Port = getString("PORT", "3000")
var WebFolderPath = getString("WEB_FOLDER_PATH", "")

var DBConnectionString = getString("DB_CONNECTION_STRING", "")
