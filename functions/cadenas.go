package functions

// BuscarByte -> Busca la posicion de un caracter en un array
func BuscarByte(caracter byte, cadena []byte) int64 {
	for contador := int64(len(cadena) - 1); contador >= 0; contador-- {
		if cadena[contador] == caracter {
			return contador
		}
	}
	return 0
}
