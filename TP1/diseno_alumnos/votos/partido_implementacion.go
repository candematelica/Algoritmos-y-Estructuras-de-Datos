package votos

import (
	"fmt"
)

type partidoImplementacion struct {
	nombrePartido string
	candidatos    [CANT_VOTACION]string
	cantVotos     [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	cantVotosEnBlanco [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido { //candidatos es de largo 3 (CANT_VOTACION = INTENDENTE + 1 = 2 + 1 = 3)
	partido := new(partidoImplementacion)
	partido.nombrePartido = nombre
	partido.candidatos[PRESIDENTE] = candidatos[PRESIDENTE]
	partido.candidatos[GOBERNADOR] = candidatos[GOBERNADOR]
	partido.candidatos[INTENDENTE] = candidatos[INTENDENTE]
	return partido
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	if len(partido.cantVotos) == 0 {
		cantVotos := [CANT_VOTACION]int{}
		partido.cantVotos = cantVotos
	}

	switch tipo {
	case PRESIDENTE:
		partido.cantVotos[PRESIDENTE]++
	case GOBERNADOR:
		partido.cantVotos[GOBERNADOR]++
	case INTENDENTE:
		partido.cantVotos[INTENDENTE]++
	}
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	switch tipo {
	case PRESIDENTE:
		if partido.cantVotos[PRESIDENTE] == 1 {
			return fmt.Sprintf("%s - %s: %d voto\n", partido.nombrePartido, partido.candidatos[PRESIDENTE], partido.cantVotos[PRESIDENTE])
		} else {
			return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatos[PRESIDENTE], partido.cantVotos[PRESIDENTE])
		}
	case GOBERNADOR:
		if partido.cantVotos[GOBERNADOR] == 1 {
			return fmt.Sprintf("%s - %s: %d voto\n", partido.nombrePartido, partido.candidatos[GOBERNADOR], partido.cantVotos[GOBERNADOR])
		} else {
			return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatos[GOBERNADOR], partido.cantVotos[GOBERNADOR])
		}
	case INTENDENTE:
		if partido.cantVotos[INTENDENTE] == 1 {
			return fmt.Sprintf("%s - %s: %d voto\n", partido.nombrePartido, partido.candidatos[INTENDENTE], partido.cantVotos[INTENDENTE])
		} else {
			return fmt.Sprintf("%s - %s: %d votos\n", partido.nombrePartido, partido.candidatos[INTENDENTE], partido.cantVotos[INTENDENTE])
		}
	}

	return ""
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	switch tipo {
	case PRESIDENTE:
		blanco.cantVotosEnBlanco[PRESIDENTE]++
	case GOBERNADOR:
		blanco.cantVotosEnBlanco[GOBERNADOR]++
	case INTENDENTE:
		blanco.cantVotosEnBlanco[INTENDENTE]++
	}
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if blanco.cantVotosEnBlanco[tipo] == 1 {
		return fmt.Sprintf("Votos en Blanco: %d voto\n", blanco.cantVotosEnBlanco[tipo])
	} else {
		return fmt.Sprintf("Votos en Blanco: %d votos\n", blanco.cantVotosEnBlanco[tipo])
	}
}
