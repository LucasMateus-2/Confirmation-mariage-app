package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucas/confirmation-mariage-app/internal/model"
	"github.com/lucas/confirmation-mariage-app/internal/service"
	"github.com/lucas/confirmation-mariage-app/pkg/response"
)

type GuestHandler struct {
	guestService *service.GuestService
}

func NewGuestHandler(guestService *service.GuestService) *GuestHandler {
	return &GuestHandler{guestService: guestService}
}

// GET /guests — lista todos os convidados (protegido)
func (h *GuestHandler) ListAll(c *gin.Context) {
	guests, err := h.guestService.ListAll()

	if err != nil {
		log.Println("erro ao listar guests:", err)
		response.InternalError(c)
		return
	}
	response.OK(c, guests)
}

// GET /guests/:id — detalhe de um convidado (protegido, uso dos noivos)
func (h *GuestHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "id inválido")
		return
	}

	guest, err := h.guestService.GetByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.OK(c, guest)
}

// GET /confirm/:id — dados públicos para a tela de confirmação (sem login)
func (h *GuestHandler) GetForConfirmation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "id inválido")
		return
	}

	guest, err := h.guestService.GetByID(id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}
	response.OK(c, guest)
}

// PATCH /guests/:id/confirm — confirma presença do convidado e agregados (público)
// PATCH /guests/:id/confirm — confirma presença do convidado e agregados (público)
func (h *GuestHandler) Confirm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "id inválido")
		return
	}

	var input model.ConfirmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "payload inválido")
		return
	}

	if err := h.guestService.Confirm(id, input); err != nil {
		log.Println("erro ao confirmar presença:", err)
		response.InternalError(c)
		return
	}

	response.OK(c, gin.H{"message": "presença confirmada"})
}

// GET /dashboard — resumo geral (protegido)
func (h *GuestHandler) Dashboard(c *gin.Context) {
	summary, err := h.guestService.Dashboard()

	if err != nil {
		log.Println("erro ao listar guests:", err)
		response.InternalError(c)
		return
	}
	response.OK(c, summary)
}

func (h *GuestHandler) SearchByName(c *gin.Context) {
	name := c.Query("q")
	if name == "" {
		response.BadRequest(c, "parâmetro 'q' é obrigatório")
		return
	}

	guests, err := h.guestService.SearchByName(name)
	if err != nil {
		log.Println("erro ao buscar guests por nome:", err)
		response.InternalError(c)
		return
	}
	response.OK(c, guests)
}
