package cola_prioridad

const (
	CAPACIDAD_INICIAL int = 10
	CANT_INICIAL      int = 0
	COEF_REDIMENSION  int = 4
	VALOR_REDIMENSION int = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	nuevo := make([]T, CAPACIDAD_INICIAL)
	return &colaConPrioridad[T]{datos: nuevo, cant: CANT_INICIAL, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {

	heap := new(colaConPrioridad[T])
	heap.datos = arreglo
	if len(arreglo) == 0 {
		nuevo := make([]T, CAPACIDAD_INICIAL)
		heap.datos = nuevo
	}
	heap.cant = len(arreglo)
	heap.cmp = funcion_cmp
	heapify(heap.datos, heap.cmp)

	return heap
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == CANT_INICIAL
}

func (heap *colaConPrioridad[T]) VerMax() T {

	if heap.cant == CANT_INICIAL {
		panic("La cola está vacía")
	}

	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Encolar(dato T) {

	cap := len(heap.datos)
	if heap.cant == cap {
		heap.redimensionar(heap.cant * VALOR_REDIMENSION)
	}

	heap.datos[heap.cant] = dato
	heap.cant++
	upHeap(heap.cant-1, heap.datos, heap.cmp)

}

func (heap *colaConPrioridad[T]) Desencolar() T {

	if heap.EstaVacia() {
		panic("La cola está vacía")
	}

	dato := heap.datos[0]
	heap.cant--
	if heap.cant > 0 {
		heap.datos[0] = heap.datos[heap.cant]
		downHeap(0, heap.datos[:heap.cant], heap.cmp)
	}

	cap := len(heap.datos)
	if heap.cant*COEF_REDIMENSION <= cap && cap > CAPACIDAD_INICIAL {
		heap.redimensionar(cap / VALOR_REDIMENSION)
	}

	return dato
}

func swap[T any](dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}

func downHeap[T any](i int, arr []T, func_cmp func(T, T) int) {

	if i >= len(arr) || i < 0 {
		return
	}

	hijoIzq := 2*i + 1
	hijoDer := 2*i + 2
	mayor := i

	if hijoIzq < len(arr) && func_cmp(arr[hijoIzq], arr[mayor]) > 0 {
		mayor = hijoIzq
	}

	if hijoDer < len(arr) && func_cmp(arr[hijoDer], arr[mayor]) > 0 {
		mayor = hijoDer
	}

	if mayor != i {
		swap(&arr[i], &arr[mayor])
		downHeap(mayor, arr, func_cmp)
	}

}

func upHeap[T any](i int, arr []T, func_cmp func(T, T) int) {

	padre, iPadre, tienePadre := obtenerPadre(i, arr)

	if !tienePadre || func_cmp(*padre, arr[i]) > 0 {
		return
	}

	swap(&arr[i], &arr[iPadre])
	upHeap(iPadre, arr, func_cmp)
}

func heapify[T any](arr []T, func_cmp func(T, T) int) {

	for i := len(arr) - 1; i >= 0; i-- {
		downHeap(i, arr, func_cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	if len(elementos) == 0 || len(elementos) == 1 {
		return
	}

	heapify(elementos, funcion_cmp)
	swap(&elementos[0], &elementos[len(elementos)-1])

	downHeap(0, elementos[:len(elementos)-1], funcion_cmp)
	HeapSort(elementos[:len(elementos)-1], funcion_cmp)
}

func obtenerPadre[T any](i int, arr []T) (*T, int, bool) {

	iPadre := (i - 1) / 2

	if i == 0 {
		return &arr[0], 0, false
	}

	return &arr[iPadre], iPadre, true
}

func (heap *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}