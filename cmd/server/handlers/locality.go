package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/carry/usecases"
	"github.com/Gopher-Rangers/mercadofresco-gopherrangers/pkg/web"
	"github.com/gin-gonic/gin"
)

type Locality struct {
	service usecases.ServiceLocality
}

func NewLocality(l usecases.ServiceLocality) Locality {
	return Locality{l}
}

func (l Locality) GetCarryLocalityByID(ctx *gin.Context) {

	localityID, err := strconv.Atoi(ctx.Query("id"))

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "O id passado não é um número!"))
		return
	}

	locality, err := l.service.GetCarryLocalityByID(localityID)

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusNotFound, "A localidade não foi encontrada!"))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, locality))

}

func (l Locality) GetAllCarriesLocalityByID(ctx *gin.Context) {

	localities, err := l.service.GetAllCarriesLocalityByID()

	if err != nil {
		ctx.JSON(web.DecodeError(http.StatusBadRequest, "erro ao acessar o banco de dados"))
		return
	}

	ctx.JSON(web.NewResponse(http.StatusOK, localities))

}
