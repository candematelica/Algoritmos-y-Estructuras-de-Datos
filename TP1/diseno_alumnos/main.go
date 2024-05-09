package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDALista "tdas/lista"
	Errores "tp1/errores"
	Votos "tp1/votos"
)

const (
	ARCHIVO_LISTAS    = 1
	ARCHIVO_PADRON    = 2
	CANT_PARAMETROS   = 3
	RANGO_DNI         = 10000000
	MSG_OK            = "OK"
	COMANDO_INGRESAR  = "ingresar"
	COMANDO_VOTAR     = "votar"
	COMANDO_DESHACER  = "deshacer"
	COMANDO_FIN_VOTAR = "fin-votar"
	VOTO_PRESIDENTE   = "Presidente"
	VOTO_GOBERNADOR   = "Gobernador"
	VOTO_INTENDENTE   = "Intendente"
)

func convertirDni(input string) (int, error) {
	dni, err := strconv.Atoi(input)
	if err != nil || dni <= 0 {
		return 0, Errores.DNIError{}
	}

	return dni, nil
}

func digito(elem, indice_digito int) int {
	elem /= indice_digito
	return int(elem % 10)
}

func criterioCountingSort(padron []Votos.Votante, rango, indice_digito int) []Votos.Votante {
	frecuencias := make([]int, rango+1)
	sumas_acumuladas := make([]int, rango+1)
	padron_ordenado := make([]Votos.Votante, len(padron))

	for _, votante := range padron {
		frecuencias[digito(votante.LeerDNI(), indice_digito)]++
	}

	for i := 1; i <= rango; i++ {
		sumas_acumuladas[i] = sumas_acumuladas[i-1] + frecuencias[i-1]
	}

	for _, votante := range padron {
		padron_ordenado[sumas_acumuladas[digito(votante.LeerDNI(), indice_digito)]] = votante
		sumas_acumuladas[digito(votante.LeerDNI(), indice_digito)]++
	}

	return padron_ordenado
}

func radixSort(padron []Votos.Votante) []Votos.Votante {
	for exp := 1; exp <= RANGO_DNI; exp *= 10 {
		padron = criterioCountingSort(padron, 9, exp)
	}

	return padron
}

func busquedaBinVotante(padron []Votos.Votante, inicio, fin, valor_buscado int) (int, error) {
	if fin < inicio {
		return -1, Errores.DNIFueraPadron{}
	}

	centro := (inicio + fin) / 2
	if padron[centro].LeerDNI() > valor_buscado {
		return busquedaBinVotante(padron, inicio, centro-1, valor_buscado)
	} else if padron[centro].LeerDNI() < valor_buscado {
		return busquedaBinVotante(padron, centro+1, fin, valor_buscado)
	}

	return centro, nil
}

func buscarVotante(padron []Votos.Votante, dni int) (int, error) {
	return busquedaBinVotante(padron, 0, len(padron), dni)
}

// Cargamos los partidos en el slice de tal manera que el primero (si es que lo hay) queda en la primer posicion (la 0),
// por lo que su numero de lista resultaria de su posicion + 1
func cargarPartidos(archivoListas *os.File) []Votos.Partido {
	var listas []Votos.Partido
	lectorListas := bufio.NewScanner(archivoListas)

	for lectorListas.Scan() {
		infoPartido := strings.Split(lectorListas.Text(), ",")
		candidatos := [Votos.CANT_VOTACION]string{infoPartido[1], infoPartido[2], infoPartido[3]}

		nuevoPartido := Votos.CrearPartido(infoPartido[0], candidatos)

		listas = append(listas, nuevoPartido)
	}

	return listas
}

func cargarPadron(archivoPadron *os.File) []Votos.Votante {
	padron := make([]Votos.Votante, 0)
	lectorPadron := bufio.NewScanner(archivoPadron)

	for lectorPadron.Scan() {
		dniPadron, _ := convertirDni(lectorPadron.Text())
		votantePadron := Votos.CrearVotante(dniPadron)
		padron = append(padron, votantePadron)
	}

	return padron
}

func convertirTipoVoto(tipo string) Votos.TipoVoto {
	switch tipo {
	case VOTO_PRESIDENTE:
		return Votos.PRESIDENTE
	case VOTO_GOBERNADOR:
		return Votos.GOBERNADOR
	case VOTO_INTENDENTE:
		return Votos.INTENDENTE
	}

	err := Errores.ErrorTipoVoto{}
	fmt.Println(err.Error())
	return -1
}

func convertirAlternativa(alternativa string) (int, error) {
	numAlternativa, err := strconv.Atoi(alternativa)
	if err != nil {
		return -1, Errores.ErrorAlternativaInvalida{}
	}
	return numAlternativa, nil
}

func esAlternativaValida(alternativa int, listas []Votos.Partido) bool {
	if alternativa >= 0 && alternativa <= len(listas) {
		return true
	}
	err := Errores.ErrorAlternativaInvalida{}
	fmt.Println(err.Error())

	return false
}

func cargarResultados(listas []Votos.Partido, votos TDALista.Lista[Votos.Voto]) {
	votosEnBlanco := Votos.CrearVotosEnBlanco()
	var votosImpugnados int
	votos.Iterar(func(voto Votos.Voto) bool {
		if voto.Impugnado == true {
			votosImpugnados++
		} else {
			if voto.VotoPorTipo[Votos.PRESIDENTE] == 0 {
				votosEnBlanco.VotadoPara(Votos.PRESIDENTE)
			} else {
				listas[(voto.VotoPorTipo[Votos.PRESIDENTE])-1].VotadoPara(Votos.PRESIDENTE)
			}
			if voto.VotoPorTipo[Votos.GOBERNADOR] == 0 {
				votosEnBlanco.VotadoPara(Votos.GOBERNADOR)
			} else {
				listas[(voto.VotoPorTipo[Votos.GOBERNADOR])-1].VotadoPara(Votos.GOBERNADOR)
			}
			if voto.VotoPorTipo[Votos.INTENDENTE] == 0 {
				votosEnBlanco.VotadoPara(Votos.INTENDENTE)
			} else {
				listas[(voto.VotoPorTipo[Votos.INTENDENTE])-1].VotadoPara(Votos.INTENDENTE)
			}
		}

		return true
	})

	fmt.Print("Presidente:\n")
	fmt.Print(votosEnBlanco.ObtenerResultado(Votos.PRESIDENTE))
	for _, partido := range listas {
		fmt.Print(partido.ObtenerResultado(Votos.PRESIDENTE))
	}
	fmt.Print("\n")

	fmt.Print("Gobernador:\n")
	fmt.Print(votosEnBlanco.ObtenerResultado(Votos.GOBERNADOR))
	for _, partido := range listas {
		fmt.Print(partido.ObtenerResultado(Votos.GOBERNADOR))
	}
	fmt.Print("\n")

	fmt.Print("Intendente:\n")
	fmt.Print(votosEnBlanco.ObtenerResultado(Votos.INTENDENTE))
	for _, partido := range listas {
		fmt.Print(partido.ObtenerResultado(Votos.INTENDENTE))
	}
	fmt.Print("\n")

	if votosImpugnados == 1 {
		fmt.Printf("Votos Impugnados: %d voto\n", votosImpugnados)
	} else {
		fmt.Printf("Votos Impugnados: %d votos\n", votosImpugnados)
	}
}

// Abro, chequeo la apertura y delvuelvo los archivos abiertos
func leerArchivos() (*os.File, *os.File, error) {
	if len(os.Args) != CANT_PARAMETROS {
		return nil, nil, Errores.ErrorParametros{}
	}

	ruta_archivoListas := os.Args[ARCHIVO_LISTAS]
	ruta_archivoPadron := os.Args[ARCHIVO_PADRON]

	archivoListas, errorLista := os.Open(ruta_archivoListas)
	archivoPadron, errorPadron := os.Open(ruta_archivoPadron)
	if errorLista != nil || errorPadron != nil {
		return nil, nil, Errores.ErrorLeerArchivo{}
	}

	return archivoListas, archivoPadron, nil
}

func finVotar(fila_votantes *TDACola.Cola[Votos.Votante], votosFinales *TDALista.Lista[Votos.Voto], padron []Votos.Votante) {
	if (*fila_votantes).EstaVacia() {
		fmt.Println(Errores.FilaVacia{})
		return
	} else {
		votante := (*fila_votantes).VerPrimero()
		voto, err := votante.FinVoto()
		if err != nil {
			fmt.Println(err)
			(*fila_votantes).Desencolar()
			return
		}
		(*votosFinales).InsertarUltimo(voto)
		posVotantePadron, _ := buscarVotante(padron, votante.LeerDNI())
		padron[posVotantePadron] = votante
		(*fila_votantes).Desencolar()
	}
	fmt.Println(MSG_OK)

	return
}

func deshacer(fila_votantes *TDACola.Cola[Votos.Votante]) {
	if (*fila_votantes).EstaVacia() {
		fmt.Println(Errores.FilaVacia{})
		return
	} else {
		votante := (*fila_votantes).VerPrimero()
		err := votante.Deshacer()
		if err != nil {
			fmt.Println(err)
			if _, ok := err.(Errores.ErrorVotanteFraudulento); ok {
				(*fila_votantes).Desencolar()
			}
			return
		}
		fmt.Println(MSG_OK)
	}
}

func ingresar(dni string, padron []Votos.Votante, fila_votantes *TDACola.Cola[Votos.Votante]) {
	dniVotante, err := convertirDni(dni)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	posVotanteNuevoPadron, err := buscarVotante(padron, dniVotante)
	if err != nil {
		fmt.Println(err)
		return
	}
	(*fila_votantes).Encolar(padron[posVotanteNuevoPadron])
	fmt.Println(MSG_OK)

	return
}

func votar(tipoVoto Votos.TipoVoto, alternativa int, listas []Votos.Partido, fila_votantes *TDACola.Cola[Votos.Votante]) {
	if (*fila_votantes).EstaVacia() {
		fmt.Println(Errores.FilaVacia{})
	} else if tipoVoto != -1 && esAlternativaValida(alternativa, listas) {
		votante := (*fila_votantes).VerPrimero()
		err := votante.Votar(tipoVoto, alternativa)
		if err != nil {
			fmt.Println(err)
			(*fila_votantes).Desencolar()
			return
		} else {
			fmt.Println(MSG_OK)
		}
	}
}

func leerComandos(listas []Votos.Partido, padron []Votos.Votante) TDALista.Lista[Votos.Voto] {
	fila_votantes := TDACola.CrearColaEnlazada[Votos.Votante]()
	votosFinales := TDALista.CrearListaEnlazada[Votos.Voto]()

	lectorComando := bufio.NewScanner(os.Stdin)
	for lectorComando.Scan() {
		input := lectorComando.Text()
		comando := strings.Fields(input)
		if len(comando) == 1 {
			if COMANDO_DESHACER == comando[0] {
				deshacer(&fila_votantes)
			} else if COMANDO_FIN_VOTAR == comando[0] {
				finVotar(&fila_votantes, &votosFinales, padron)
			}
		} else if len(comando) == 2 && comando[0] == COMANDO_INGRESAR {
			ingresar(comando[1], padron, &fila_votantes)
		} else if len(comando) == 3 && comando[0] == COMANDO_VOTAR {
			tipoVoto := convertirTipoVoto(comando[1])
			alternativa, err := convertirAlternativa(comando[2])
			if err != nil {
				fmt.Println(err)
			} else {
				votar(tipoVoto, alternativa, listas, &fila_votantes)
			}
		} else {
			fmt.Println(Errores.ErrorComandoInvalido{})
		}
	}

	if !fila_votantes.EstaVacia() {
		fmt.Println(Errores.ErrorCiudadanosSinVotar{})
	}

	return votosFinales
}

func main() {
	archivoListas, archivoPadron, err := leerArchivos()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivoPadron.Close()
	defer archivoListas.Close()

	listas := cargarPartidos(archivoListas)
	padron := cargarPadron(archivoPadron)
	padron = radixSort(padron)

	votos := leerComandos(listas, padron)
	cargarResultados(listas, votos)

	return
}
