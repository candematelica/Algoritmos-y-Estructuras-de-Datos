package tp2

import (
	"strings"
	TDADicc "tdas/diccionario"
	TDALista "tdas/lista"
	Errores "tp2/errores"
)

type publicacion struct {
	user      string
	idUser    int
	textoPost string
	idPost    int
	likes     TDADicc.DiccionarioOrdenado[string, string]
}

func CrearPost(user string, idUser int, textoPost string, idPost int) Post {
	return &publicacion{user, idUser, textoPost, idPost, TDADicc.CrearABB[string, string](strings.Compare)}
}

func (post *publicacion) Creador() (string, int) {
	return post.user, post.idUser
}

func (post *publicacion) IDPost() int {
	return post.idPost
}

func (post *publicacion) Contenido() string {
	return post.textoPost
}

func (post *publicacion) Likear(user string) {
	if !post.likes.Pertenece(user) {
		post.likes.Guardar(user, user)
	}
}

func (post *publicacion) VerLikes() (TDALista.Lista[string], error) {
	listaLikes := TDALista.CrearListaEnlazada[string]()

	if listaLikes.Largo() == 0 {
		return nil, Errores.ErrorMostrarLikes{}
	}

	post.likes.Iterar(func(_, user string) bool {
		listaLikes.InsertarUltimo(user)
		return true
	})

	return listaLikes, nil
}
