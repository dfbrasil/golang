package rotas

import (
	"api/src/controllers"
	"net/http"
)


var rotasUsuarios = []Rota{
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI: "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},

	{
		URI: "/usuarios/{usuarioid}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI: "/usuarios/{usuarioid}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI: "/usuarios/{usuarioid}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},


}