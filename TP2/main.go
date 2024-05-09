package tp2

import (
	"fmt"
	Comandos "tp2/comandos"
	FuncArchivos "tp2/func_archivos"
)

func main() {
	archivoUsuarios, err := FuncArchivos.LeerArchivos()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivoUsuarios.Close()

	usuarios := FuncArchivos.CargarUsuarios(archivoUsuarios)
	Comandos.LeerComandos(usuarios)

	return
}
