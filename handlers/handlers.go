package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mygoapp/models"
	"net/http"
)

var db *gorm.DB

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func ShowPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("D:/Go/mygoapp/mygoapp.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Проверяем, существует ли таблица пользователей, прежде чем выполнять миграцию
	if !db.Migrator().HasTable(&models.User{}) {
		// Создаем таблицу пользователей, если её нет
		err = db.AutoMigrate(&models.User{})
		if err != nil {
			panic("Failed to auto-migrate users table: " + err.Error())
		}
	}
}

func CreateUser(c *gin.Context) {
	// Генерация уникального имени пользователя и электронной почты (псевдокод)
	uniqueUsername := generateUniqueUsername()
	uniqueEmail := generateUniqueEmail()

	// Создание нового пользователя
	newUser := models.User{
		Username: uniqueUsername,
		Email:    uniqueEmail,
	}

	// Выполняем запрос на вставку пользователя и проверяем ошибки
	result := db.Create(&newUser)
	if result.Error != nil {
		panic("Failed to insert user: " + result.Error.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully!",
	})
}

func GetUsers(c *gin.Context) {
	// Получаем список пользователей из базы данных
	users := GetAllUsers()

	// Отправляем список пользователей на страницу
	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}

func GetAllUsers() []models.User {
	var users []models.User
	db.Find(&users)
	return users
}
