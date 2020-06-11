package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rafaelarodrigues90/usuarios/models"
)

// Home : pagina inicial da aplicação
func Home(c echo.Context) error {
	return c.JSON(http.StatusOK, "Página inicial")
}

// Create ...
func Create(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios
	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, "Registro não adicionado")
		}
		return c.JSON(http.StatusCreated, "Registro armazenado com sucesso")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "campos vazios",
	})
}

// Read ...
func Read(c echo.Context) error {
	var usuarios []models.Usuarios

	err := models.UsuarioModel.Find().All(&usuarios)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Erro ao buscar dados")
	}

	return c.JSON(http.StatusFound, usuarios)
}

// Update ...
func Update(c echo.Context) error {

	usuarioID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario = models.Usuarios{
		ID:    usuarioID,
		Nome:  nome,
		Email: email,
	}

	resultado := models.UsuarioModel.Find("id=?", usuarioID)
	count, _ := resultado.Count()
	if count < 1 {
		return c.JSON(http.StatusNotFound, "Usuário não encontrado")
	}

	err := resultado.Update(usuario)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Erro ao atualizar registro")
	}

	return c.JSON(http.StatusAccepted, "Usuário atualizado")
}

// Delete ...
func Delete(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := models.UsuarioModel.Find(usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, "Usuário não encontrado")
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, "Não foi possível deletar usuário")
	}

	return c.JSON(http.StatusAccepted, "Usuário deletado")
}
