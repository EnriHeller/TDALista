package tp2

type lector interface {

	//Procesa de forma completa un archivo de log.
	Procesar(string) (string, error)

}
