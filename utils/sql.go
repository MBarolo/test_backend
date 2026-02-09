package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GenericScanAll[T any](db *sql.DB, query string, args ...interface{}) ([]*T, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %w", err)
	}
	defer rows.Close()

	// Obtener los nombres de las columnas
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo las columnas: %w", err)
	}

	// Lista para almacenar los resultados
	results := make([]*T, 0)

	for rows.Next() {
		// Crear una instancia de la estructura genérica
		obj := new(T)
		structValue := reflect.ValueOf(obj).Elem()

		// Crear punteros dinámicos para las columnas
		columnValues := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		// Escanear los valores de la fila
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %w", err)
		}

		// Mapear los valores a los campos de la estructura
		for i, column := range columns {
			column = strings.ReplaceAll(column, "_", " ")
			column = strings.Title(column)
			column = strings.ReplaceAll(column, " ", "")
			field := structValue.FieldByName(column)
			//quitar Guiones y espacios en blanco
			if field.IsValid() && field.CanSet() {
				val := columnValues[i]
				if val != nil {
					convertValue(val, field)
				}
			}
		}

		// Agregar la estructura mapeada a los resultados
		results = append(results, obj)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar filas: %w", err)
	}

	return results, nil
}

func convertValue(value interface{}, field reflect.Value) {
	if value == nil {
		// No se asigna nada si el valor es nulo
		return
	}

	// Si el campo es un puntero, manejarlo especialmente
	if field.Kind() == reflect.Ptr {
		// Crear un nuevo valor del tipo apuntado
		elemType := field.Type().Elem()
		newElem := reflect.New(elemType).Elem()

		// Convertir el valor al elemento
		convertValue(value, newElem)

		// Si la conversión fue exitosa, asignar el puntero
		if newElem.IsValid() {
			ptr := reflect.New(elemType)
			ptr.Elem().Set(newElem)
			field.Set(ptr)
		}
		return
	}

	// Si el valor es []uint8, intenta convertirlo al tipo correcto
	if data, ok := value.([]uint8); ok {
		switch field.Kind() {
		case reflect.String:
			field.SetString(string(data)) // Convertir []uint8 a string
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intValue, err := strconv.ParseInt(string(data), 10, 64)
			if err == nil {
				field.SetInt(intValue)
			} else {
				log.Printf("Error convirtiendo a int: %v\n", err)
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintValue, err := strconv.ParseUint(string(data), 10, 64)
			if err == nil {
				field.SetUint(uintValue)
			} else {
				log.Printf("Error convirtiendo a uint: %v\n", err)
			}
		case reflect.Float32, reflect.Float64:
			floatValue, err := strconv.ParseFloat(string(data), 64)
			if err == nil {
				field.SetFloat(floatValue)
			} else {
				log.Printf("Error convirtiendo a float: %v\n", err)
			}
		case reflect.Bool:
			boolValue, err := strconv.ParseBool(string(data))
			if err == nil {
				field.SetBool(boolValue)
			} else {
				log.Printf("Error convirtiendo a bool: %v\n", err)
			}
		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				str := strings.Split(string(data), " ")

				if len(str) == 1 { // str tiene formato 2006-01-02 -> date
					parsedTime, err := time.Parse("2006-01-02", string(data))
					if err == nil {
						field.Set(reflect.ValueOf(parsedTime))
					} else {
						log.Printf("Error convirtiendo a time.Time: %v\n", err)
					}
				} else if len(str) == 2 { // str tiene formato 2006-01-02 15:04:05 -> datetime
					parsedTime, err := time.Parse("2006-01-02 15:04:05", string(data))
					if err == nil {
						field.Set(reflect.ValueOf(parsedTime))
					} else {
						log.Printf("Error convirtiendo a time.Time: %v\n", err)
					}
				} else {
					log.Printf("Error convirtiendo a time.Time: %v\n", data)
				}
			}
		}
		return
	}
	// Si no es []uint8, intenta asignar directamente
	switch field.Kind() {
	case reflect.String:
		if str, ok := value.(string); ok {
			field.SetString(str)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if i, ok := value.(int64); ok {
			field.SetInt(i)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		switch v := value.(type) {
		case uint64:
			field.SetUint(v)
		case int64:
			if v >= 0 {
				field.SetUint(uint64(v))
			} else {
				log.Printf("No se puede asignar int64 negativo a uint: %v\n", v)
			}
		}
	case reflect.Float32, reflect.Float64:
		if f, ok := value.(float64); ok {
			field.SetFloat(f)
		}
	case reflect.Bool:
		switch v := value.(type) {
		case int64:
			field.SetBool(v != 0)
		}
	case reflect.Struct:
		if field.Type() == reflect.TypeOf(time.Time{}) {
			if t, ok := value.(time.Time); ok {
				field.Set(reflect.ValueOf(t))
			}
		}
	}
}
