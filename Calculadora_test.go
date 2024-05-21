package Calculadora

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := Soma(3, 2)
	esperado := 5.0
	if resultado != esperado {
		t.Errorf("Soma(3, 2) = %f; esperado %f", resultado, esperado)
	}
}

func TestSubtrai(t *testing.T) {
	resultado := Subtrai(3, 2)
	esperado := 1.0
	if resultado != esperado {
		t.Errorf("Subtrai(3, 2) = %f; esperado %f", resultado, esperado)
	}
}

func TestMultiplica(t *testing.T) {
	resultado := Multiplica(3, 2)
	esperado := 6.0
	if resultado != esperado {
		t.Errorf("Multiplica(3, 2) = %f; esperado %f", resultado, esperado)
	}
}

func TestDivide(t *testing.T) {
	t.Run("Divisão por um número diferente de zero", func(t *testing.T) {
		resultado, err := Divide(6, 2)
		esperado := 3.0
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}
		if resultado != esperado {
			t.Errorf("Divide(6, 2) = %f; esperado %f", resultado, esperado)
		}
	})

	t.Run("Divisão por zero", func(t *testing.T) {
		_, err := Divide(6, 0)
		if err == nil {
			t.Error("Esperado erro ao dividir por zero, mas não houve erro")
		}
	})
}
