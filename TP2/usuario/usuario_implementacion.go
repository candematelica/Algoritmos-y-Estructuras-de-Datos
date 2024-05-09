package tp2

import (
	TDAHeap "tdas/cola_prioridad"
	Errores "tp2/errores"
	FuncCmp "tp2/func_comparacion"
	TDAPost "tp2/post"
)

type usuario struct {
	nombre string
	id     int
	feed   TDAHeap.ColaPrioridad[*TDAPost.Post]
}

func CrearUsuario(nombre string, id int) User {
	return &usuario{nombre, id, TDAHeap.CrearHeap[*TDAPost.Post](FuncCmp.CmpAfinidad(id))}
}

func (user *usuario) Nombre() string {
	return user.nombre
}

func (user *usuario) IDUser() int {
	return user.id
}

func (user *usuario) Publicar(post TDAPost.Post) {
	user.feed.Encolar(&post)
}

func (user *usuario) VerSigFeed() (TDAPost.Post, error) {
	if user.feed.EstaVacia() {
		return nil, Errores.ErrorVerSigFeed{}
	}

	return *user.feed.Desencolar(), nil
}
