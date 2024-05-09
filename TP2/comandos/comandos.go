package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADicc "tdas/diccionario"
	Errores "tp2/errores"
	TDAPost "tp2/post"
	TDAUser "tp2/usuario"
)

const (
	COMANDO_LOGIN         = "login"
	COMANDO_LOGOUT        = "logout"
	COMANDO_PUBLICAR      = "publicar"
	COMANDO_VER_SIG_FEED  = "ver_siguiente_feed"
	COMANDO_LIKEAR        = "likear_post"
	COMANDO_MOSTRAR_LIKES = "mostrar_likes"
	SALUDO                = "Hola"
	DESPEDIDA             = "Adios"
	MSG_OK_PUBLICAR       = "Post publicado"
	MSG_OK_LIKEAR         = "Post likeado"
)

type usuarioLoggeado struct {
	estaLoggeado bool
	user         TDAUser.User
}

func LeerComandos(usuarios TDADicc.Diccionario[string, *TDAUser.User]) {
	posteados := TDADicc.CrearHash[int, *TDAPost.Post]()
	var usuarioLog usuarioLoggeado
	var contadorPublicaciones int

	lectorComando := bufio.NewScanner(os.Stdin)
	for lectorComando.Scan() {
		input := lectorComando.Text()
		comando := strings.Fields(input)
		if len(comando) < 1 {
			fmt.Println(Errores.ErrorComandoInvalido{})
			continue
		}
		switch comando[0] {
		case COMANDO_LOGIN:
			nombreUsuario := strings.Join(comando[1:], " ")
			login(nombreUsuario, &usuarioLog, usuarios)
		case COMANDO_LOGOUT:
			logout(&usuarioLog)
		case COMANDO_PUBLICAR:
			textoPublicacion := strings.Join(comando[1:], " ")
			publicar(usuarioLog, textoPublicacion, usuarios, posteados, &contadorPublicaciones)
		case COMANDO_VER_SIG_FEED:
			verSigFeed(usuarioLog)
		case COMANDO_LIKEAR:
			idPostALikear, _ := strconv.Atoi(comando[1])
			likear(idPostALikear, usuarioLog, posteados)
		case COMANDO_MOSTRAR_LIKES:
			idPost, _ := strconv.Atoi(comando[1])
			mostrarLikes(idPost, posteados)
		}
	}

	return
}

func login(nombreUsuario string, usuarioLog *usuarioLoggeado, usuarios TDADicc.Diccionario[string, *TDAUser.User]) {
	if usuarioLog.estaLoggeado {
		fmt.Println(Errores.ErrorUsuarioYaLoggeado{})
		return
	}
	if !usuarios.Pertenece(nombreUsuario) {
		fmt.Println(Errores.ErrorUsuarioInexistente{})
		return
	}

	(*usuarioLog).estaLoggeado = true
	(*usuarioLog).user = TDAUser.CrearUsuario(nombreUsuario, (*usuarios.Obtener(nombreUsuario)).IDUser())

	fmt.Println(SALUDO, nombreUsuario)
}

func logout(usuarioLog *usuarioLoggeado) {
	if !usuarioLog.estaLoggeado {
		fmt.Println(Errores.ErrorNoHayUsuarioLoggeado{})
		return
	}

	(*usuarioLog).estaLoggeado = false

	fmt.Println(DESPEDIDA)
}

func publicar(usuarioLog usuarioLoggeado, textoPublicacion string, usuarios TDADicc.Diccionario[string, *TDAUser.User], posteados TDADicc.Diccionario[int, *TDAPost.Post], contadorPublicaciones *int) {
	if !usuarioLog.estaLoggeado {
		fmt.Println(Errores.ErrorNoHayUsuarioLoggeado{})
		return
	}

	post := TDAPost.CrearPost(usuarioLog.user.Nombre(), usuarioLog.user.IDUser(), textoPublicacion, *contadorPublicaciones)

	usuarios.Iterar(func(nom string, user *TDAUser.User) bool {
		if (*user).Nombre() != usuarioLog.user.Nombre() {
			(*user).Publicar(post)
		}
		return true
	})
	posteados.Guardar(*contadorPublicaciones, &post)
	(*contadorPublicaciones)++

	fmt.Println(MSG_OK_PUBLICAR)
}

func verSigFeed(usuarioLog usuarioLoggeado) {
	if !usuarioLog.estaLoggeado {
		fmt.Println(Errores.ErrorVerSigFeed{})
		return
	}

	sigPost, err := usuarioLog.user.VerSigFeed()
	if err != nil {
		fmt.Println(err)
		return
	}
	creador, _ := sigPost.Creador()
	likes, _ := sigPost.VerLikes()

	fmt.Printf("Post ID %d\n", sigPost.IDPost())
	fmt.Printf("%s dijo: %s\n", creador, sigPost.Contenido())
	fmt.Printf("Likes: %d\n", likes.Largo())
}

func likear(idPostALikear int, usuarioLog usuarioLoggeado, posteados TDADicc.Diccionario[int, *TDAPost.Post]) {
	if !posteados.Pertenece(idPostALikear) || !usuarioLog.estaLoggeado {
		fmt.Println(Errores.ErrorLikear{})
		return
	}

	post := (*posteados.Obtener(idPostALikear))
	post.Likear(usuarioLog.user.Nombre())

	fmt.Println(MSG_OK_LIKEAR)
}

func mostrarLikes(idPost int, posteados TDADicc.Diccionario[int, *TDAPost.Post]) {
	if !posteados.Pertenece(idPost) {
		fmt.Println(Errores.ErrorMostrarLikes{})
		return
	}

	likes, err := (*posteados.Obtener(idPost)).VerLikes()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("El post tiene %d likes:\n", likes.Largo())
	likes.Iterar(func(user string) bool {
		fmt.Printf("\t%s\n", user)
		return true
	})
}
