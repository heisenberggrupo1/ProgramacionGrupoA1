package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ==========================================
// 1. CAPA DE MODELOS (Models)
// ==========================================

type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Genero string `json:"genero"`
}

type Usuario struct {
	ID                int      `json:"id"`
	Nombre            string   `json:"nombre"`
	Correo            string   `json:"correo"`
	Password          string   `json:"password"`
	LibrosDescargados []string `json:"libros_descargados"`
}

// ==========================================
// 2. ARCHIVOS JSON (Data / Almacenamiento)
// ==========================================

var libros []Libro
var usuarios []Usuario

const archivoLibros = "libros.json"
const archivoUsuarios = "usuarios.json"

func inicializarDatos() {
	// Inicializar Libros
	if _, err := os.Stat(archivoLibros); os.IsNotExist(err) {
		libros = []Libro{
			{ID: 1, Titulo: "El Quijote", Autor: "Miguel de Cervantes", Genero: "Novela"},
			{ID: 2, Titulo: "Clean Code", Autor: "Robert Martin", Genero: "Programacion"},
		}
		guardarLibros()
	} else {
		file, _ := os.ReadFile(archivoLibros)
		json.Unmarshal(file, &libros)
	}

	// Inicializar Usuarios
	if _, err := os.Stat(archivoUsuarios); os.IsNotExist(err) {
		guardarUsuarios()
	} else {
		file, _ := os.ReadFile(archivoUsuarios)
		json.Unmarshal(file, &usuarios)
	}
}

func guardarLibros() {
	data, _ := json.MarshalIndent(libros, "", "  ")
	os.WriteFile(archivoLibros, data, 0644)
}

func guardarUsuarios() {
	data, _ := json.MarshalIndent(usuarios, "", "  ")
	os.WriteFile(archivoUsuarios, data, 0644)
}

// ==========================================
// 3. CAPA DE SERVICIOS (Programación Funcional)
// ==========================================

type FiltroLibro func(Libro) bool

func FiltrarLibros(filtro FiltroLibro) []Libro {
	var resultado []Libro
	for _, l := range libros {
		if filtro(l) {
			resultado = append(resultado, l)
		}
	}
	return resultado
}

func BuscarPorTitulo(titulo string) FiltroLibro {
	return func(l Libro) bool { return strings.EqualFold(l.Titulo, titulo) }
}

func BuscarPorAutor(autor string) FiltroLibro {
	return func(l Libro) bool { return strings.EqualFold(l.Autor, autor) }
}

func BuscarPorGenero(genero string) FiltroLibro {
	return func(l Libro) bool { return strings.EqualFold(l.Genero, genero) }
}

// ==========================================
// 4. FUNCIONES DE ADMINISTRADOR
// ==========================================

func iniciarSesionAdmin() bool {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Usuario administrador: ")
	scanner.Scan()
	usuario := strings.TrimSpace(scanner.Text())

	fmt.Print("Contraseña: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	if usuario == "GrupoA" && password == "GrupoA111" {
		fmt.Println("Bienvenido Administrador.")
		return true
	}
	fmt.Println("Usuario o contraseña incorrectos.")
	return false
}

func agregarLibro() {
	scanner := bufio.NewScanner(os.Stdin)
	var id int

	fmt.Print("ID del libro: ")
	fmt.Scanln(&id)

	fmt.Print("Título: ")
	scanner.Scan()
	titulo := scanner.Text()

	fmt.Print("Autor: ")
	scanner.Scan()
	autor := scanner.Text()

	fmt.Print("Género: ")
	scanner.Scan()
	genero := scanner.Text()

	libro := Libro{ID: id, Titulo: titulo, Autor: autor, Genero: genero}
	libros = append(libros, libro)
	guardarLibros()
	fmt.Println("Libro agregado correctamente.")
}

func cambiarLibro() {
	scanner := bufio.NewScanner(os.Stdin)
	var id int
	fmt.Print("Ingrese el ID del libro que desea cambiar: ")
	fmt.Scanln(&id)

	for i := range libros {
		if libros[i].ID == id {
			fmt.Print("Nuevo título: ")
			scanner.Scan()
			libros[i].Titulo = scanner.Text()

			fmt.Print("Nuevo autor: ")
			scanner.Scan()
			libros[i].Autor = scanner.Text()

			fmt.Print("Nuevo género: ")
			scanner.Scan()
			libros[i].Genero = scanner.Text()

			guardarLibros()
			fmt.Println("Libro actualizado correctamente.")
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}

func borrarLibro() {
	var id int
	fmt.Print("Ingrese el ID del libro a borrar: ")
	fmt.Scanln(&id)

	for i, libro := range libros {
		if libro.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			guardarLibros()
			fmt.Println("Libro borrado correctamente.")
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}

// ==========================================
// 5. FUNCIONES DE USUARIO
// ==========================================

func registrarUsuario() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Nombre completo: ")
	scanner.Scan()
	nombre := scanner.Text()

	fmt.Print("Correo: ")
	scanner.Scan()
	correo := strings.TrimSpace(scanner.Text())

	fmt.Print("Contraseña: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	usuario := Usuario{
		ID:                len(usuarios) + 1,
		Nombre:            nombre,
		Correo:            correo,
		Password:          password,
		LibrosDescargados: []string{},
	}
	usuarios = append(usuarios, usuario)
	guardarUsuarios()
	fmt.Println("Usuario registrado correctamente.")
}

func iniciarSesion() *Usuario {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Correo: ")
	scanner.Scan()
	correo := strings.TrimSpace(scanner.Text())

	fmt.Print("Contraseña: ")
	scanner.Scan()
	password := strings.TrimSpace(scanner.Text())

	for i := range usuarios {
		if usuarios[i].Correo == correo && usuarios[i].Password == password {
			fmt.Println("Inicio de sesión exitoso.")
			return &usuarios[i]
		}
	}
	fmt.Println("Credenciales incorrectas.")
	return nil
}

func mostrarLibros() {
	if len(libros) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}
	fmt.Println("\n--- CATÁLOGO DE LIBROS ---")
	for _, libro := range libros {
		fmt.Printf("[%d] %s - %s (Género: %s)\n", libro.ID, libro.Titulo, libro.Autor, libro.Genero)
	}
}

func descargarLibro(usuario *Usuario) {
	var id int
	fmt.Print("Ingrese el ID del libro que desea descargar: ")
	fmt.Scanln(&id)

	for _, libro := range libros {
		if libro.ID == id {
			usuario.LibrosDescargados = append(usuario.LibrosDescargados, libro.Titulo)
			guardarUsuarios()
			fmt.Println("Libro descargado correctamente:", libro.Titulo)
			return
		}
	}
	fmt.Println("Libro no encontrado.")
}

func verPerfil(usuario *Usuario) {
	fmt.Println("\n--- PERFIL DEL USUARIO ---")
	fmt.Println("ID:", usuario.ID)
	fmt.Println("Nombre:", usuario.Nombre)
	fmt.Println("Correo:", usuario.Correo)
	fmt.Println("\nLibros descargados:")
	if len(usuario.LibrosDescargados) == 0 {
		fmt.Println("No tiene libros descargados.")
	} else {
		for _, libro := range usuario.LibrosDescargados {
			fmt.Println("-", libro)
		}
	}
}

// ==========================================
// 6. MENÚS DE INTERACCIÓN
// ==========================================

func ejecutarBusqueda(criterio string, filtro FiltroLibro) {
	resultados := FiltrarLibros(filtro)
	if len(resultados) == 0 {
		fmt.Println("No se encontraron libros.")
		return
	}
	fmt.Printf("\n--- RESULTADOS DE BÚSQUEDA (%s) ---\n", strings.ToUpper(criterio))
	for _, libro := range resultados {
		fmt.Printf("[%d] %s - %s (Género: %s)\n", libro.ID, libro.Titulo, libro.Autor, libro.Genero)
	}
}

func menuBusqueda() {
	scanner := bufio.NewScanner(os.Stdin)
	var opcion int

	fmt.Println("\nBUSCAR LIBRO")
	fmt.Println("1. Buscar por título")
	fmt.Println("2. Buscar por autor")
	fmt.Println("3. Buscar por género")
	fmt.Print("Seleccione una opción: ")
	fmt.Scanln(&opcion)

	fmt.Print("Ingrese el término a buscar: ")
	scanner.Scan()
	busqueda := scanner.Text()

	switch opcion {
	case 1:
		ejecutarBusqueda("Título", BuscarPorTitulo(busqueda))
	case 2:
		ejecutarBusqueda("Autor", BuscarPorAutor(busqueda))
	case 3:
		ejecutarBusqueda("Género", BuscarPorGenero(busqueda))
	default:
		fmt.Println("Opción inválida.")
	}
}

func menuUsuario(usuario *Usuario) {
	var opcion int
	for {
		fmt.Println("\nMENÚ USUARIO")
		fmt.Println("1. Ver catálogo")
		fmt.Println("2. Buscar libro")
		fmt.Println("3. Descargar libro")
		fmt.Println("4. Ver perfil")
		fmt.Println("5. Cerrar sesión")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			mostrarLibros()
		case 2:
			menuBusqueda()
		case 3:
			descargarLibro(usuario)
		case 4:
			verPerfil(usuario)
		case 5:
			fmt.Println("Sesión cerrada.")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func menuAdmin() {
	var opcion int
	for {
		fmt.Println("\nMENÚ ADMINISTRADOR")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Cambiar libro")
		fmt.Println("3. Borrar libro")
		fmt.Println("4. Ver catálogo")
		fmt.Println("5. Cerrar sesión admin")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			agregarLibro()
		case 2:
			cambiarLibro()
		case 3:
			borrarLibro()
		case 4:
			mostrarLibros()
		case 5:
			fmt.Println("Sesión de administrador cerrada.")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

// ==========================================
// 7. CLIENTE / PUNTO DE ENTRADA (Main)
// ==========================================

func main() {
	inicializarDatos()
	var opcion int

	for {
		fmt.Println("\n===================================")
		fmt.Println("   BIBLIOTECA VIRTUAL - GRUPO A")
		fmt.Println("===================================")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Iniciar sesión")
		fmt.Println("3. Usuario administrador")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			registrarUsuario()
		case 2:
			usuario := iniciarSesion()
			if usuario != nil {
				menuUsuario(usuario)
			}
		case 3:
			if iniciarSesionAdmin() {
				menuAdmin()
			}
		case 4:
			fmt.Println("Programa finalizado.")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
