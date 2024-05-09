package tp2

import (
	"bufio"
	"os"
	TDADicc "tdas/diccionario"
	Errores "tp2/errores"
	TDAUser "tp2/usuario"
)

const (
	CANT_PARAMETROS  = 2
	ARCHIVO_USUARIOS = 1
)

func LeerArchivos() (*os.File, error) {
	if len(os.Args) != CANT_PARAMETROS {
		return nil, Errores.ErrorParametros{}
	}

	ruta_archivoUsuarios := os.Args[ARCHIVO_USUARIOS]

	archivoUsuarios, err := os.Open(ruta_archivoUsuarios)
	if err != nil {
		return nil, Errores.ErrorLeerArchivo{}
	}

	return archivoUsuarios, nil
}

func CargarUsuarios(archivoUsuarios *os.File) TDADicc.Diccionario[string, *TDAUser.User] {
	usuarios := TDADicc.CrearHash[string, *TDAUser.User]()
	lector := bufio.NewScanner(archivoUsuarios)

	var i int
	for lector.Scan() {
		nuevoUsuario := TDAUser.CrearUsuario(lector.Text(), i)
		usuarios.Guardar(lector.Text(), &nuevoUsuario)
		i++
	}

	return usuarios
}
