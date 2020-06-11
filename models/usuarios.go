package models

import "github.com/rafaelarodrigues90/usuarios/lib"

// Usuarios representa a tabela de usuarios na base de dados
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

// UsuarioModel recebe a tabela do db
var UsuarioModel = lib.Session.Collection("usuarios")
