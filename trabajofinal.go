package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// ======================================================
// BIBLIOTECA VIRTUAL - GRUPO A
// Archivo: trabajofinal.go
// Lenguaje: Go
// ======================================================

type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Genero string `json:"genero"`
	Stock  int    `json:"stock"`
}

type Usuario struct {
	ID                int      `json:"id"`
	Nombre            string   `json:"nombre"`
	Correo            string   `json:"correo"`
	Password          string   `json:"password"`
	LibrosDescargados []string `json:"libros_descargados"`
}

type RespuestaAPI struct {
	Estado  string      `json:"estado"`
	Mensaje string      `json:"mensaje"`
	Datos   interface{} `json:"datos,omitempty"`
}

var libros []Libro
var usuarios []Usuario

const archivoLibros = "libros.json"
const archivoUsuarios = "usuarios.json"

const usuarioAdmin = "GrupoA"
const claveAdmin = "GrupoA111"

var lector = bufio.NewReader(os.Stdin)

// ======================================================
// LEER DATOS POR CONSOLA
// ======================================================

func leerTexto(mensaje string) string {
	fmt.Print(mensaje)
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(mensaje string) int {
	for {
		texto := leerTexto(mensaje)
		numero, err := strconv.Atoi(texto)

		if err == nil {
			return numero
		}

		fmt.Println("Ingrese un numero valido.")
	}
}

// ======================================================
// DATOS INICIALES Y JSON
// ======================================================

func librosIniciales() []Libro {
	return []Libro{
		{ID: 1, Titulo: "El Quijote", Autor: "Miguel de Cervantes", Genero: "Novela", Stock: 5},
		{ID: 2, Titulo: "Clean Code", Autor: "Robert C. Martin", Genero: "Programacion", Stock: 3},
		{ID: 3, Titulo: "Cien años de soledad", Autor: "Gabriel Garcia Marquez", Genero: "Novela", Stock: 4},
		{ID: 4, Titulo: "Programacion en Go", Autor: "Alan Donovan", Genero: "Programacion", Stock: 2},
		{ID: 5, Titulo: "La ciudad y los perros", Autor: "Mario Vargas Llosa", Genero: "Novela", Stock: 4},
		{ID: 6, Titulo: "Introduccion a la programacion", Autor: "Joyanes Aguilar", Genero: "Educativo", Stock: 6},
		{ID: 7, Titulo: "Sistemas de informacion gerencial", Autor: "Kenneth C. Laudon", Genero: "Gestion empresarial", Stock: 3},
	}
}

func cargarDatos() {
	dataLibros, err := os.ReadFile(archivoLibros)

	if err != nil || len(dataLibros) == 0 {
		libros = librosIniciales()
		guardarLibros()
	} else {
		json.Unmarshal(dataLibros, &libros)

		if len(libros) == 0 {
			libros = librosIniciales()
			guardarLibros()
		}
	}

	dataUsuarios, err := os.ReadFile(archivoUsuarios)

	if err != nil || len(dataUsuarios) == 0 {
		usuarios = []Usuario{}
		guardarUsuarios()
	} else {
		json.Unmarshal(dataUsuarios, &usuarios)
	}
}

func guardarLibros() {
	data, err := json.MarshalIndent(libros, "", "  ")

	if err != nil {
		fmt.Println("Error al guardar libros:", err)
		return
	}

	os.WriteFile(archivoLibros, data, 0644)
}

func guardarUsuarios() {
	data, err := json.MarshalIndent(usuarios, "", "  ")

	if err != nil {
		fmt.Println("Error al guardar usuarios:", err)
		return
	}

	os.WriteFile(archivoUsuarios, data, 0644)
}

// ======================================================
// VALIDACIONES
// ======================================================

func validarLibro(libro Libro) error {
	if strings.TrimSpace(libro.Titulo) == "" {
		return errors.New("el titulo es obligatorio")
	}

	if strings.TrimSpace(libro.Autor) == "" {
		return errors.New("el autor es obligatorio")
	}

	if strings.TrimSpace(libro.Genero) == "" {
		return errors.New("el genero es obligatorio")
	}

	if libro.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}

	return nil
}

func validarUsuario(usuario Usuario) error {
	if strings.TrimSpace(usuario.Nombre) == "" {
		return errors.New("el nombre es obligatorio")
	}

	if strings.TrimSpace(usuario.Correo) == "" {
		return errors.New("el correo es obligatorio")
	}

	if strings.TrimSpace(usuario.Password) == "" {
		return errors.New("la contraseña es obligatoria")
	}

	return nil
}

func existeCorreo(correo string) bool {
	for _, usuario := range usuarios {
		if strings.EqualFold(usuario.Correo, correo) {
			return true
		}
	}

	return false
}

func generarIDLibro() int {
	mayor := 0

	for _, libro := range libros {
		if libro.ID > mayor {
			mayor = libro.ID
		}
	}

	return mayor + 1
}

func generarIDUsuario() int {
	mayor := 0

	for _, usuario := range usuarios {
		if usuario.ID > mayor {
			mayor = usuario.ID
		}
	}

	return mayor + 1
}

// ======================================================
// FUNCIONES PRINCIPALES
// ======================================================

func agregarLibro(libro Libro) error {
	if libro.ID == 0 {
		libro.ID = generarIDLibro()
	}

	if err := validarLibro(libro); err != nil {
		return err
	}

	libros = append(libros, libro)
	guardarLibros()

	return nil
}

func actualizarLibro(id int, libroActualizado Libro) error {
	if err := validarLibro(libroActualizado); err != nil {
		return err
	}

	for i := range libros {
		if libros[i].ID == id {
			libros[i].Titulo = libroActualizado.Titulo
			libros[i].Autor = libroActualizado.Autor
			libros[i].Genero = libroActualizado.Genero
			libros[i].Stock = libroActualizado.Stock

			guardarLibros()
			return nil
		}
	}

	return errors.New("libro no encontrado")
}

func eliminarLibro(id int) error {
	for i, libro := range libros {
		if libro.ID == id {
			libros = append(libros[:i], libros[i+1:]...)
			guardarLibros()
			return nil
		}
	}

	return errors.New("libro no encontrado")
}

func registrarUsuario(usuario Usuario) error {
	if err := validarUsuario(usuario); err != nil {
		return err
	}

	if existeCorreo(usuario.Correo) {
		return errors.New("ya existe un usuario con ese correo")
	}

	usuario.ID = generarIDUsuario()
	usuario.LibrosDescargados = []string{}

	usuarios = append(usuarios, usuario)
	guardarUsuarios()

	return nil
}

func iniciarSesion(correo string, password string) (*Usuario, error) {
	for i := range usuarios {
		if strings.EqualFold(usuarios[i].Correo, correo) && usuarios[i].Password == password {
			return &usuarios[i], nil
		}
	}

	return nil, errors.New("correo o contraseña incorrectos")
}

func buscarLibroPorID(id int) (*Libro, error) {
	for i := range libros {
		if libros[i].ID == id {
			return &libros[i], nil
		}
	}

	return nil, errors.New("libro no encontrado")
}

func buscarUsuarioPorID(id int) (*Usuario, error) {
	for i := range usuarios {
		if usuarios[i].ID == id {
			return &usuarios[i], nil
		}
	}

	return nil, errors.New("usuario no encontrado")
}

func descargarLibro(usuarioID int, libroID int) error {
	usuario, err := buscarUsuarioPorID(usuarioID)

	if err != nil {
		return err
	}

	libro, err := buscarLibroPorID(libroID)

	if err != nil {
		return err
	}

	if libro.Stock <= 0 {
		return errors.New("no hay stock disponible")
	}

	usuario.LibrosDescargados = append(usuario.LibrosDescargados, libro.Titulo)
	libro.Stock--

	guardarUsuarios()
	guardarLibros()

	return nil
}

// ======================================================
// PROGRAMACIÓN FUNCIONAL
// ======================================================

type FiltroLibro func(Libro) bool

func filtrarLibros(filtro FiltroLibro) []Libro {
	var resultado []Libro

	for _, libro := range libros {
		if filtro(libro) {
			resultado = append(resultado, libro)
		}
	}

	return resultado
}

func buscarPorTitulo(titulo string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.Contains(strings.ToLower(libro.Titulo), strings.ToLower(titulo))
	}
}

func buscarPorAutor(autor string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(autor))
	}
}

func buscarPorGenero(genero string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.Contains(strings.ToLower(libro.Genero), strings.ToLower(genero))
	}
}

// ======================================================
// MENÚ DE CONSOLA
// ======================================================

func mostrarCatalogo() {
	fmt.Println("\n===== CATALOGO DE LIBROS =====")

	if len(libros) == 0 {
		fmt.Println("No hay libros registrados.")
		return
	}

	for _, libro := range libros {
		fmt.Printf("ID: %d | Titulo: %s | Autor: %s | Genero: %s | Stock: %d\n",
			libro.ID, libro.Titulo, libro.Autor, libro.Genero, libro.Stock)
	}
}

func mostrarResultados(resultados []Libro) {
	if len(resultados) == 0 {
		fmt.Println("No se encontraron resultados.")
		return
	}

	fmt.Println("\n===== RESULTADOS =====")

	for _, libro := range resultados {
		fmt.Printf("ID: %d | Titulo: %s | Autor: %s | Genero: %s | Stock: %d\n",
			libro.ID, libro.Titulo, libro.Autor, libro.Genero, libro.Stock)
	}
}

func menuRegistrarUsuario() {
	fmt.Println("\n===== REGISTRAR USUARIO =====")

	nombre := leerTexto("Nombre: ")
	correo := leerTexto("Correo: ")
	password := leerTexto("Contraseña: ")

	usuario := Usuario{
		Nombre:   nombre,
		Correo:   correo,
		Password: password,
	}

	err := registrarUsuario(usuario)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuario registrado correctamente.")
}

func menuLoginUsuario() {
	fmt.Println("\n===== INICIAR SESION USUARIO =====")

	correo := leerTexto("Correo: ")
	password := leerTexto("Contraseña: ")

	usuario, err := iniciarSesion(correo, password)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bienvenido:", usuario.Nombre)
	menuUsuario(usuario)
}

func menuUsuario(usuario *Usuario) {
	for {
		fmt.Println("\n===== MENU USUARIO =====")
		fmt.Println("1. Ver catalogo")
		fmt.Println("2. Buscar por titulo")
		fmt.Println("3. Buscar por autor")
		fmt.Println("4. Buscar por genero")
		fmt.Println("5. Descargar libro")
		fmt.Println("6. Ver mi perfil")
		fmt.Println("7. Cerrar sesion")

		opcion := leerEntero("Seleccione una opcion: ")

		switch opcion {
		case 1:
			mostrarCatalogo()

		case 2:
			titulo := leerTexto("Ingrese titulo: ")
			resultados := filtrarLibros(buscarPorTitulo(titulo))
			mostrarResultados(resultados)

		case 3:
			autor := leerTexto("Ingrese autor: ")
			resultados := filtrarLibros(buscarPorAutor(autor))
			mostrarResultados(resultados)

		case 4:
			genero := leerTexto("Ingrese genero: ")
			resultados := filtrarLibros(buscarPorGenero(genero))
			mostrarResultados(resultados)

		case 5:
			mostrarCatalogo()
			idLibro := leerEntero("Ingrese ID del libro a descargar: ")

			err := descargarLibro(usuario.ID, idLibro)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Libro descargado correctamente.")
			}

		case 6:
			fmt.Println("\n===== MI PERFIL =====")
			fmt.Println("ID:", usuario.ID)
			fmt.Println("Nombre:", usuario.Nombre)
			fmt.Println("Correo:", usuario.Correo)
			fmt.Println("Libros descargados:", usuario.LibrosDescargados)

		case 7:
			fmt.Println("Sesion cerrada.")
			return

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func menuLoginAdmin() {
	fmt.Println("\n===== INICIAR SESION ADMINISTRADOR =====")

	usuario := leerTexto("Usuario administrador: ")
	clave := leerTexto("Contraseña: ")

	if usuario == usuarioAdmin && clave == claveAdmin {
		fmt.Println("Bienvenido administrador.")
		menuAdmin()
	} else {
		fmt.Println("Credenciales incorrectas.")
	}
}

func menuAdmin() {
	for {
		fmt.Println("\n===== MENU ADMINISTRADOR =====")
		fmt.Println("1. Ver catalogo")
		fmt.Println("2. Agregar libro")
		fmt.Println("3. Actualizar libro")
		fmt.Println("4. Eliminar libro")
		fmt.Println("5. Ver usuarios registrados")
		fmt.Println("6. Cerrar sesion")

		opcion := leerEntero("Seleccione una opcion: ")

		switch opcion {
		case 1:
			mostrarCatalogo()

		case 2:
			menuAgregarLibro()

		case 3:
			menuActualizarLibro()

		case 4:
			menuEliminarLibro()

		case 5:
			menuVerUsuarios()

		case 6:
			fmt.Println("Sesion de administrador cerrada.")
			return

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func menuAgregarLibro() {
	fmt.Println("\n===== AGREGAR LIBRO =====")

	titulo := leerTexto("Titulo: ")
	autor := leerTexto("Autor: ")
	genero := leerTexto("Genero: ")
	stock := leerEntero("Stock: ")

	libro := Libro{
		Titulo: titulo,
		Autor:  autor,
		Genero: genero,
		Stock:  stock,
	}

	err := agregarLibro(libro)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Libro agregado correctamente.")
}

func menuActualizarLibro() {
	fmt.Println("\n===== ACTUALIZAR LIBRO =====")

	mostrarCatalogo()

	id := leerEntero("Ingrese ID del libro: ")
	titulo := leerTexto("Nuevo titulo: ")
	autor := leerTexto("Nuevo autor: ")
	genero := leerTexto("Nuevo genero: ")
	stock := leerEntero("Nuevo stock: ")

	libro := Libro{
		Titulo: titulo,
		Autor:  autor,
		Genero: genero,
		Stock:  stock,
	}

	err := actualizarLibro(id, libro)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Libro actualizado correctamente.")
}

func menuEliminarLibro() {
	fmt.Println("\n===== ELIMINAR LIBRO =====")

	mostrarCatalogo()

	id := leerEntero("Ingrese ID del libro a eliminar: ")

	err := eliminarLibro(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Libro eliminado correctamente.")
}

func menuVerUsuarios() {
	fmt.Println("\n===== USUARIOS REGISTRADOS =====")

	if len(usuarios) == 0 {
		fmt.Println("No hay usuarios registrados.")
		return
	}

	for _, usuario := range usuarios {
		fmt.Println("-------------------------")
		fmt.Println("ID:", usuario.ID)
		fmt.Println("Nombre:", usuario.Nombre)
		fmt.Println("Correo:", usuario.Correo)
		fmt.Println("Libros descargados:", usuario.LibrosDescargados)
	}
}

// ======================================================
// SERVICIOS WEB REST
// ======================================================

func responderJSON(w http.ResponseWriter, estado int, respuesta RespuestaAPI) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(estado)
	json.NewEncoder(w).Encode(respuesta)
}

// Ruta principal para evitar 404 al abrir localhost:8080
func servicioInicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Biblioteca Virtual - Grupo A")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Rutas disponibles:")
	fmt.Fprintln(w, "GET  /api/estado")
	fmt.Fprintln(w, "GET  /api/libros")
	fmt.Fprintln(w, "GET  /api/usuarios")
	fmt.Fprintln(w, "GET  /api/buscar?genero=Novela")
	fmt.Fprintln(w, "POST /api/usuarios/registrar")
	fmt.Fprintln(w, "POST /api/usuarios/login")
	fmt.Fprintln(w, "POST /api/usuarios/descargar?usuario_id=1&libro_id=2")
}

func servicioEstado(w http.ResponseWriter, r *http.Request) {
	responderJSON(w, http.StatusOK, RespuestaAPI{
		Estado:  "ok",
		Mensaje: "Biblioteca Virtual funcionando correctamente",
		Datos:   "Proyecto Grupo A",
	})
}

func servicioLibros(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responderJSON(w, http.StatusOK, RespuestaAPI{"ok", "catalogo de libros", libros})
		return
	}

	if r.Method == http.MethodPost {
		var libro Libro
		json.NewDecoder(r.Body).Decode(&libro)

		err := agregarLibro(libro)

		if err != nil {
			responderJSON(w, http.StatusBadRequest, RespuestaAPI{"error", err.Error(), nil})
			return
		}

		responderJSON(w, http.StatusCreated, RespuestaAPI{"ok", "libro agregado", libro})
		return
	}

	responderJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{"error", "metodo no permitido", nil})
}

// NUEVA RUTA: GET /api/usuarios
func servicioUsuarios(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{"error", "metodo no permitido", nil})
		return
	}

	responderJSON(w, http.StatusOK, RespuestaAPI{
		Estado:  "ok",
		Mensaje: "usuarios registrados",
		Datos:   usuarios,
	})
}

func servicioBuscar(w http.ResponseWriter, r *http.Request) {
	titulo := r.URL.Query().Get("titulo")
	autor := r.URL.Query().Get("autor")
	genero := r.URL.Query().Get("genero")

	var resultado []Libro

	if titulo != "" {
		resultado = filtrarLibros(buscarPorTitulo(titulo))
	} else if autor != "" {
		resultado = filtrarLibros(buscarPorAutor(autor))
	} else if genero != "" {
		resultado = filtrarLibros(buscarPorGenero(genero))
	} else {
		responderJSON(w, http.StatusBadRequest, RespuestaAPI{"error", "ingrese titulo, autor o genero", nil})
		return
	}

	responderJSON(w, http.StatusOK, RespuestaAPI{"ok", "resultado de busqueda", resultado})
}

func servicioRegistrar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responderJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{"error", "metodo no permitido", nil})
		return
	}

	var usuario Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	err := registrarUsuario(usuario)

	if err != nil {
		responderJSON(w, http.StatusBadRequest, RespuestaAPI{"error", err.Error(), nil})
		return
	}

	responderJSON(w, http.StatusCreated, RespuestaAPI{"ok", "usuario registrado", nil})
}

func servicioLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responderJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{"error", "metodo no permitido", nil})
		return
	}

	var datos Usuario
	json.NewDecoder(r.Body).Decode(&datos)

	usuario, err := iniciarSesion(datos.Correo, datos.Password)

	if err != nil {
		responderJSON(w, http.StatusUnauthorized, RespuestaAPI{"error", err.Error(), nil})
		return
	}

	responderJSON(w, http.StatusOK, RespuestaAPI{"ok", "inicio de sesion correcto", usuario})
}

func servicioDescargar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responderJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{"error", "metodo no permitido", nil})
		return
	}

	usuarioID, err1 := strconv.Atoi(r.URL.Query().Get("usuario_id"))
	libroID, err2 := strconv.Atoi(r.URL.Query().Get("libro_id"))

	if err1 != nil || err2 != nil {
		responderJSON(w, http.StatusBadRequest, RespuestaAPI{"error", "usuario_id y libro_id deben ser numeros", nil})
		return
	}

	err := descargarLibro(usuarioID, libroID)

	if err != nil {
		responderJSON(w, http.StatusBadRequest, RespuestaAPI{"error", err.Error(), nil})
		return
	}

	responderJSON(w, http.StatusOK, RespuestaAPI{"ok", "libro descargado", nil})
}

func iniciarServidor() {
	http.HandleFunc("/", servicioInicio)
	http.HandleFunc("/api/estado", servicioEstado)
	http.HandleFunc("/api/libros", servicioLibros)
	http.HandleFunc("/api/usuarios", servicioUsuarios)
	http.HandleFunc("/api/buscar", servicioBuscar)
	http.HandleFunc("/api/usuarios/registrar", servicioRegistrar)
	http.HandleFunc("/api/usuarios/login", servicioLogin)
	http.HandleFunc("/api/usuarios/descargar", servicioDescargar)

	fmt.Println("\nServidor iniciado en http://localhost:8080")
	fmt.Println("Pagina principal: http://localhost:8080/")
	fmt.Println("Ver estado:       http://localhost:8080/api/estado")
	fmt.Println("Ver libros:       http://localhost:8080/api/libros")
	fmt.Println("Ver usuarios:     http://localhost:8080/api/usuarios")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}

// ======================================================
// FUNCIÓN PRINCIPAL
// ======================================================

func main() {
	cargarDatos()

	for {
		fmt.Println("\n===================================")
		fmt.Println("     BIBLIOTECA VIRTUAL GRUPO A")
		fmt.Println("===================================")
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Iniciar sesion como usuario")
		fmt.Println("3. Iniciar sesion como administrador")
		fmt.Println("4. Ver catalogo")
		fmt.Println("5. Iniciar servidor web REST")
		fmt.Println("6. Salir")

		opcion := leerEntero("Seleccione una opcion: ")

		switch opcion {
		case 1:
			menuRegistrarUsuario()

		case 2:
			menuLoginUsuario()

		case 3:
			menuLoginAdmin()

		case 4:
			mostrarCatalogo()

		case 5:
			iniciarServidor()

		case 6:
			fmt.Println("Programa finalizado.")
			return

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}