package tp2

import TDAPost "tp2/post"

type User interface {
	// Nombre: devuelve el nombre del usuario
	Nombre() string
	// IdUser: devuelve el id del usuario
	IDUser() int
	// Publicar: publica un nuevo post
	Publicar(TDAPost.Post)
	// VerSigFeed: devuelve el siguiente post a ver en el feed del usuario loggeado
	VerSigFeed() (TDAPost.Post, error)
}
