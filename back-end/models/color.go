package models

// Definimos un nuevo tipo llamado Color
type Color int

// Declaramos las constantes con iota
// iodta sirve para poner un contador autimatico
const (
	Rojo Color = iota
	Azul
	Verde
	Amarillo
)
