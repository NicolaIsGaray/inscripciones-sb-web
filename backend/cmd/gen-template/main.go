// Command gen-template regenera la plantilla de importación de alumnos con
// listas desplegables en las columnas Instancia (C), Titulo (D) y Escuela (E).
//
// Uso:
//
//	go run ./cmd/gen-template
//
// Escribe in-place en frontend/public/plantilla-alumnos.xlsx, preservando
// encabezados, estilos y la hoja "Instrucciones". Es idempotente: si se corre
// dos veces, reemplaza las validaciones previas en vez de duplicarlas.
//
// El backend valida las listas en modo advertencia: Excel avisa si el valor no
// está en la lista pero permite continuar. El import sigue leyendo las columnas
// por posición (A=DNI, B=nombre, C=instancia, D=título, E=escuela) y no exige
// valores de lista, así que las validaciones no cambian el comportamiento.
package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Opciones de las listas desplegables. Editarlas por ciclo lectivo y volver a
// correr el comando.
var (
	instanceOptions = []string{"Primera", "Segunda", "Tercera", "Cuarta"}
	titleOptions    = []string{
		"Discapacidad",
		"Hermanos/as",
		"Hijos/as del personal",
		"Origen a 300 m",
		"Federados (Ed. Física)",
		"Cercanía del domicilio",
		"Abanderados y escoltas titulares",
	}
	schoolOptions = []string{"secundaria", "deportiva"}
)

const (
	dataSheet = "Hoja1"
	helpSheet = "Instrucciones"
	// La fila 1 es el encabezado; la validación cubre desde la fila 2.
	lastRow = 1000
)

func main() {
	tplPath := templatePath()

	f, err := excelize.OpenFile(tplPath)
	if err != nil {
		log.Fatalf("no se pudo abrir la plantilla %s: %v", tplPath, err)
	}
	defer f.Close()

	addListValidation(f, "C", instanceOptions)
	addListValidation(f, "D", titleOptions)
	addListValidation(f, "E", schoolOptions)

	if err := addHelpNote(f); err != nil {
		log.Fatalf("no se pudo actualizar la hoja %s: %v", helpSheet, err)
	}

	if err := f.Save(); err != nil {
		log.Fatalf("no se pudo guardar la plantilla: %v", err)
	}

	fmt.Printf("Plantilla regenerada con listas desplegables: %s\n", tplPath)
}

// templatePath resuelve la ruta de la plantilla a partir del código fuente, de
// modo que el comando funcione desde cualquier directorio de trabajo.
func templatePath() string {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("no se pudo resolver la ruta de la plantilla")
	}
	// main.go vive en backend/cmd/gen-template: tres niveles arriba está la
	// raíz del repositorio.
	return filepath.Clean(filepath.Join(
		filepath.Dir(thisFile), "..", "..", "..",
		"frontend", "public", "plantilla-alumnos.xlsx",
	))
}

// addListValidation agrega (o reemplaza) la lista desplegable de una columna.
func addListValidation(f *excelize.File, column string, options []string) {
	sqref := fmt.Sprintf("%s2:%s%d", column, column, lastRow)

	// Idempotencia: descartar validaciones previas sobre el mismo rango.
	if err := f.DeleteDataValidation(dataSheet, sqref); err != nil {
		log.Fatalf("no se pudo descartar la validación previa de %s: %v", sqref, err)
	}

	dv := excelize.NewDataValidation(true)
	dv.SetSqref(sqref)
	if err := dv.SetDropList(options); err != nil {
		log.Fatalf("la lista de la columna %s excede el límite de caracteres de Excel: %v", column, err)
	}
	dv.SetInput("Elegí una opción", "Desplegá la lista y elegí un valor. Si escribís otro valor, Excel advierte pero permite continuar.")
	// Advertencia (no bloqueo): Excel muestra el aviso y acepta "Sí" para
	// conservar el valor escrito a mano.
	dv.SetError(excelize.DataValidationErrorStyleWarning, "Valor fuera de la lista", "El valor no está en la lista sugerida. Podés continuar igual, pero verificá que esté bien escrito.")

	if err := f.AddDataValidation(dataSheet, dv); err != nil {
		log.Fatalf("no se pudo agregar la validación de %s: %v", column, err)
	}
}

// addHelpNote escribe la nota sobre las listas desplegables al final de la
// hoja Instrucciones. Es idempotente: quita cualquier nota previa (de corridas
// anteriores) y escribe una única línea al final del contenido.
func addHelpNote(f *excelize.File) error {
	if err := removeHelpNotes(f); err != nil {
		return err
	}

	rows, err := f.GetRows(helpSheet)
	if err != nil {
		return err
	}
	// Última fila con contenido (ignorando filas vacías de relleno).
	lastContent := 0
	for i, row := range rows {
		if !isEmptyCells(row) {
			lastContent = i + 1
		}
	}

	// Fila destino, con una fila de separación debajo del contenido.
	row := lastContent + 2
	cell, err := excelize.CoordinatesToCellName(1, row)
	if err != nil {
		return err
	}
	note := "9. Las columnas Instancia, Titulo y Escuela tienen listas desplegables: " +
		"elegí un valor de la lista. Si escribís otro valor, Excel advierte pero permite continuar."
	if err := f.SetCellValue(helpSheet, cell, note); err != nil {
		return err
	}
	// Estilo cuerpo (índice 1) igual al resto de las líneas de instrucciones.
	return f.SetCellStyle(helpSheet, cell, cell, 1)
}

// removeHelpNotes quita las notas previas (empiezan con "9. " en la columna A),
// de abajo hacia arriba para no desplazar los índices restantes.
func removeHelpNotes(f *excelize.File) error {
	rows, err := f.GetRows(helpSheet)
	if err != nil {
		return err
	}
	var noteRows []int
	for i, row := range rows {
		if len(row) > 0 && strings.HasPrefix(strings.TrimSpace(row[0]), "9. ") {
			noteRows = append(noteRows, i+1)
		}
	}
	for i := len(noteRows) - 1; i >= 0; i-- {
		if err := f.RemoveRow(helpSheet, noteRows[i]); err != nil {
			return err
		}
	}
	return nil
}

// isEmptyCells indica si todas las celdas de la fila están vacías.
func isEmptyCells(cells []string) bool {
	for _, cell := range cells {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}
