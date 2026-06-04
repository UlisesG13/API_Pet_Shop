package controllers

import (
    "net/http"
    "proyecto/src/pets/application"
    "strconv"
	"fmt"
    "github.com/gin-gonic/gin"
)

type DeletePetController struct {
    petRemover *application.DeletePet
}

func NewDeletePetController(useCase *application.DeletePet) *DeletePetController {
    return &DeletePetController{petRemover: useCase}
}

func (dp *DeletePetController) Run(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pet ID"})
        return
    }
    err = dp.petRemover.Execute(id)
    if err != nil {
        if err.Error() == fmt.Sprintf("no se encontró ninguna mascota con el ID %d", id) {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Mascota eliminada correctamente"})
}