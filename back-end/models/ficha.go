package models_test

type Jugador struct {
	NombreJugador string
	Color         Color
}

// Constructor tipo "builder"
func NewJugador(nombre string, color Color) *Jugador {
	return &Jugador{
		NombreJugador: nombre,
		Color:         color,
	}
}


