package votos

import (
	TDAPila "tdas/pila"
	Errores "tp1/errores"
)

const (
	DNI_INVALIDO = 0
	DNI_VALIDO   = 1
)

type votanteImplementacion struct {
	numeroDNI int
	votos     TDAPila.Pila[Voto]
	yaVoto    bool
	voto      Voto
}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni, TDAPila.CrearPilaDinamica[Voto](), false, Voto{}}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.numeroDNI
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.yaVoto {
		err := Errores.ErrorVotanteFraudulento{}
		err.Dni = votante.numeroDNI
		return err
	}
	if alternativa == 0 {
		votante.voto.Impugnado = true
	} else if votante.votos.EstaVacia() {
		var votoNuevo Voto
		votoNuevo.VotoPorTipo = [CANT_VOTACION]int{}
		votoNuevo.VotoPorTipo[tipo] = alternativa
		votante.votos.Apilar(votoNuevo)
	} else {
		ultimoVoto := votante.votos.VerTope()
		ultimoVoto.VotoPorTipo[tipo] = alternativa
		votante.votos.Apilar(ultimoVoto)
	}

	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.yaVoto {
		err := Errores.ErrorVotanteFraudulento{}
		err.Dni = votante.numeroDNI
		return err
	}
	if votante.voto.Impugnado {
		votante.voto.Impugnado = false
	} else if votante.votos.EstaVacia() {
		return Errores.ErrorNoHayVotosAnteriores{}
	} else {
		votante.votos.Desapilar()
	}

	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.yaVoto {
		err := Errores.ErrorVotanteFraudulento{}
		err.Dni = votante.numeroDNI
		return Voto{}, err
	}
	if !votante.voto.Impugnado && !votante.votos.EstaVacia() {
		votoDefinitivo := votante.votos.VerTope()
		votante.voto = votoDefinitivo
	}
	votante.yaVoto = true

	return votante.voto, nil
}
