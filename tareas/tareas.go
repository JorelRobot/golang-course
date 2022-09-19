package main

import "fmt"

type listaTareas struct {
	tareas []*tarea
}

func (lt *listaTareas) agregarTarea(t *tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *listaTareas) eliminarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...) // No olvidar usar los puntos suspensivos al usar slice[index:]
}

func (lt *listaTareas) imprimirTareas() {
	for _, tarea := range lt.tareas {
		fmt.Println()
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
		fmt.Println("Completada:", tarea.completado)
	}
}

func (lt *listaTareas) imprimirTareasCompletadas() {
	for _, tarea := range lt.tareas {

		if tarea.completado {
			fmt.Println("\nTAREA COMPLETADA")
			fmt.Println("Nombre:", tarea.nombre)
			fmt.Println("Descripcion:", tarea.descripcion)
			fmt.Println("Completada:", tarea.completado)
		}
	}
}

type tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *tarea) marcarCompletado() {
	t.completado = true
}

func (t *tarea) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *tarea) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &tarea{
		nombre:      "Completar mi curso en go",
		descripcion: "Hacer esta tarea en esta misma semana",
	}

	t2 := &tarea{
		nombre:      "Completar mi curso en go",
		descripcion: "Hacer esta tarea en esta misma semana",
	}

	t3 := &tarea{
		nombre:      "Completar mi curso en go",
		descripcion: "Hacer esta tarea en esta misma semana",
	}

	fmt.Printf("%+v\n", t) // %+v sirve para imprimir una estructura con un formato de llave valor

	lt := &listaTareas{
		tareas: []*tarea{
			t, t2,
		},
	}

	lt.agregarTarea(t3)

	fmt.Println(len(lt.tareas))

	lt.eliminarTarea(2)
	fmt.Println(len(lt.tareas))

	lt.tareas[1].marcarCompletado()

	lt.imprimirTareas()
	lt.imprimirTareasCompletadas()
}
