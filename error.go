package main
 
import "errors"

var ErrorsNombre = errors.New("les nombres ne sont pas acceptés")
var ErrorsArgument= errors.New("argument manquant")
var ErrorsBadArgument = errors.New("argument invalide")