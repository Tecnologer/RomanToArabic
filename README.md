# RomanToArabic [![Build Status](https://travis-ci.org/Tecnologer/RomanToArabic.svg?branch=master)](https://travis-ci.org/Tecnologer/RomanToArabic)
Hay dos funciones:
1. **Romano a Arabigo**
2. **Arabigo a Romano**

### Romano a Arabigo

* Recibe como entrada una cadena de caracteres, los caracteres permitidos son: `I, V, X, L, C, D y M`
* Imprime un numero entero en digitos del 1 al 9

#### Ejemplo

* Entrada: `MCMXCV` <br>
* Salida: **1995**

### Arabigo a Romano
* Recibe un numero entero mayor a *0* y menor a *4000*
* Imprime el equivalente en numero romano

#### Ejemplo

* Entrada: `1995` <br>
* Salida: **MCMXCV**

## Ejecutar desde codigo

1. Descargar codigo: `go get github.com/tecnologer/romantoarabic`<br>
    Con esto tambien quedara disponible para usarse en otros proyestos.
2. Abrir la carpeta del proyecto: `cd $GOPATH/github.com/tecnologer/romantoarabic`
3. Ejecutar el codigo: `go run main.go`

## Ejecutar desde binarios
Puede ejecutar el binario correcpondiente al sistema operativo que este utilizando. Si no esta disponible para su sistema, puede crearlo con `go build`.


