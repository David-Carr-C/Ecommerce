package services

import (
	"encoding/json"
	"net/http"

	"gitlab.com/David-Carr-C/Ecommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserToSend struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}

// GetUserByID obtiene un usuario por su ID
func GetUserByID(c *gin.Context, db *gorm.DB, userID uint64) { // obtiene por parametro a gin, gorm y el id del usuario
	var user models.User              // se crea una variable de tipo User
	result := db.First(&user, userID) // se busca en la base de datos el usuario con el id especificado
	if result.Error != nil {          // si hay un error, se retorna el error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user}) // si no hay error, se retorna el usuario
}

// Login verifica las credenciales solo email y password
func Login(c *gin.Context, db *gorm.DB, userEmail string, userPassword string) { // obtiene por parametro a gin, gorm, email y contraseña
	var user models.User                                    // se crea una variable de tipo User
	result := db.Where("email = ?", userEmail).First(&user) // se busca en la base de datos el usuario con el email especificado
	if result.Error != nil {                                // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if !user.CheckPassword(userPassword) { // se verifica la contraseña del usuario
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"}) // si la contraseña no es correcta, se retorna un mensaje de error
		return
	}

	userToSend := UserToSend{ // se crea un nuevo usuario con los datos necesarios
		ID:       uint64(user.ID),
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}

	jsonToSend, err := json.Marshal(userToSend) // se convierte el usuario a json
	if err != nil {                             // si hay un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cookie := http.Cookie{
		Name:     "user",
		Value:    string(jsonToSend),
		MaxAge:   3600,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, &cookie) // se crea una cookie con los datos del usuario

	c.JSON(http.StatusOK, gin.H{"data": user}) // si no hay error, se retorna el usuario
}

// CreateUser crea un nuevo usuario
// Obtiene por parametro a gin y gorm
func CreateUser(c *gin.Context, db *gorm.DB) {
	var newUser models.User                    // se crea una variable de tipo User
	err := c.ShouldBindJSON(&newUser)          // se obtiene el json enviado por el cliente y se lo asigna a la variable newUser
	handleError(http.StatusBadRequest, c, err) // maneja el error

	err = newUser.SetPassword(newUser.Password)         // se cifra la contraseña del usuario
	handleError(http.StatusInternalServerError, c, err) // maneja el error

	result := db.Create(&newUser)                                // se crea el usuario en la base de datos
	handleError(http.StatusInternalServerError, c, result.Error) // maneja el error

	newUser.Password = "" // se elimina la contraseña del usuario para no mostrarla al cliente

	c.JSON(http.StatusOK, gin.H{"data": newUser}) // se retorna el usuario creado
}

// UpdateUser actualiza un usuario existente
// Obtiene por parametro a gin, gorm y el id del usuario
func UpdateUser(c *gin.Context, db *gorm.DB, userID uint64) {
	var user models.User              // se crea una variable de tipo User
	result := db.First(&user, userID) // se busca en la base de datos el usuario con el id especificado

	handleError(http.StatusInternalServerError, c, result.Error) // maneja el error

	err := c.ShouldBindJSON(&user)             // se obtiene el json enviado por el cliente y se lo asigna a la variable user
	handleError(http.StatusBadRequest, c, err) // maneja el error

	result = db.Save(&user)                                      // se actualiza el usuario en la base de datos
	handleError(http.StatusInternalServerError, c, result.Error) // maneja el error

	user.Password = "" // se elimina la contraseña del usuario para no mostrarla al cliente

	c.JSON(http.StatusOK, gin.H{"data": user}) // se retorna el usuario actualizado
}

// DeleteUser elimina un usuario por su ID
// Obtiene por parametro a gin, gorm y el id del usuario
func DeleteUser(c *gin.Context, db *gorm.DB, userID uint64) {
	result := db.Delete(&models.User{}, userID)                  // elimina el usuario de la base de datos
	handleError(http.StatusInternalServerError, c, result.Error) // maneja el error

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"}) // retorna un mensaje de exito
}

// handleError maneja los errores
// Obtiene por parametro el codigo de estado, gin y el error
func handleError(statusCode int, c *gin.Context, err error) {
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		c.Abort()
	}
}
