package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type ErrorComandoInvalido struct{}

func (e ErrorComandoInvalido) Error() string {
	return "ERROR: Comando inválido"
}

type ErrorUsuarioYaLoggeado struct{}

func (e ErrorUsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorNoHayUsuarioLoggeado struct{}

func (e ErrorNoHayUsuarioLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorVerSigFeed struct{}

func (e ErrorVerSigFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorLikear struct{}

func (e ErrorLikear) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
