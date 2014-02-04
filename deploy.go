package main

import (
	"fmt"
)

func tomcatUrl(deploy Deploy) string {
	host := getValue(deploy.Properties, "host", "localhost")
	port := getValue(deploy.Properties, "port", "8080")
	path := getValue(deploy.Properties, "path", "/")
	user := getValue(deploy.Properties, "user", "admin")
	password := getValue(deploy.Properties, "password", "password")
	return fmt.Sprintf("http://%s:%s@%s:%s/manager/deploy?path=%s&update=true", user, password, host, port, path)
}

func getValue(properties map[string]string, key, defaultValue string) string {
	if value, ok := properties[key]; ok {
        return value
    }
    return defaultValue
}
