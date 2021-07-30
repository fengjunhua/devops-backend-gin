package swag

import ginSwagger "github.com/swaggo/gin-swagger"

func InitSwagger(){

	URL := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
}