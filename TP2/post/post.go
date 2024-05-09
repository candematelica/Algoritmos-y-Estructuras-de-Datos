package tp2

import TDALista "tdas/lista"

type Post interface {
	// Creador: devuelve el nombre de usuario de quien creo el post, junto con su id
	Creador() (string, int)
	// IDPost: devuelve el entero identificador del post
	IDPost() int
	// Contenido: devuelve el post
	Contenido() string
	// Likear: se agrega al usuario pasado por parametro a la lista de usuarios likeados
	Likear(string)
	// VerLikes: devuelve una lista que contiene aquellos que le dieron like al post
	VerLikes() (TDALista.Lista[string], error)
}
