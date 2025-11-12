package models

// Definimos un nuevo tipo llamado Color
type TipoDePosicion int

// Declaramos las constantes con iota
// iodta sirve para poner un contador autimatico
const (
	BASE = iota
	SEGURO
	SALIDA
	NORMAL
)
