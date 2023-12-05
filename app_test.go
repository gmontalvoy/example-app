package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// Crear un request para pasar al handler.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un ResponseRecorder (httptest.ResponseWriter) para grabar las respuestas.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	// El handler satisface http.Handler, así que podemos llamar a su método ServeHTTP directamente
	// y pasarle nuestro Request y ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Comprobar el código de estado de la respuesta.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler devolvió un código incorrecto: obtuvo %v, esperaba %v", status, http.StatusOK)
	}

	// Comprobar el cuerpo de la respuesta.
	expected := `Hola, este es un servidor web básico en Go!`
	if rr.Body.String() != expected {
		t.Errorf("handler devolvió un cuerpo inesperado: obtuvo %v, esperaba %v", rr.Body.String(), expected)
	}
}
