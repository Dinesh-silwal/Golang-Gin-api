# golang-restapi

# Initialize a go module : make sure to run it over Go-env folder
    $ go mod init
# Removed unused dependencies
    $ go mod tidy
# Installing dependencies : updating go.mod file
    $ go get <package_name>
    $ go get github.com/gin-gonic/gin github.com/jinzhu/gorm
# download required dependencies
    $ go mod download
#Crea Configure mux-router
    $ go get -u github.com/gorilla/mux
