package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Helper function to capture stdout for testing
func captureStdout(f func()) string {
	oldStdout := os.Stdout // Keep original stdout
	r, w, _ := os.Pipe()   // Create a pipe
	os.Stdout = w          // Redirect stdout to the write end of the pipe

	outC := make(chan string)
	// Goroutine to read from the read end of the pipe
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r) // Copy all output from pipe to buffer
		r.Close()        // Close the read end of the pipe
		outC <- buf.String()
	}()

	f() // Execute the function that prints to stdout

	w.Close()             // Close the write end of the pipe
	os.Stdout = oldStdout // Restore original stdout
	return <-outC         // Get the captured output
}

// Helper function to simulate stdin for testing
func simulateStdin(input string, f func()) {
	oldStdin := os.Stdin // Keep original stdin
	r, w, _ := os.Pipe() // Create a pipe
	os.Stdin = r         // Redirect stdin to the read end of the pipe

	// Goroutine to write input to the write end of the pipe
	go func() {
		defer w.Close() // Close the write end of the pipe when done
		io.WriteString(w, input)
	}()

	f() // Execute the function that reads from stdin

	os.Stdin = oldStdin // Restore original stdin
}

// TestCalcSum testa a função calc_sum
func TestCalcSum(t *testing.T) {
	tests := []struct {
		name string
		n1   float64
		n2   float64
		want float64
	}{
		{"Soma Positiva", 10, 5, 15},
		{"Soma com Zero", 7, 0, 7},
		{"Soma com Negativo", -3, 8, 5},
		{"Soma com Decimais", 2.5, 3.5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				calc_sum(tt.n1, tt.n2)
			})

			got, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
			if err != nil {
				t.Errorf("calc_sum() failed to parse output: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("calc_sum(%f, %f) = %f; want %f", tt.n1, tt.n2, got, tt.want)
			}
		})
	}
}

// TestCalcMinus testa a função calc_minus
func TestCalcMinus(t *testing.T) {
	tests := []struct {
		name string
		n1   float64
		n2   float64
		want float64
	}{
		{"Subtracao Positiva", 10, 5, 5},
		{"Subtracao com Zero", 7, 0, 7},
		{"Subtracao Resultado Negativo", 3, 8, -5},
		{"Subtracao com Decimais", 5.5, 2.5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				calc_minus(tt.n1, tt.n2)
			})

			got, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
			if err != nil {
				t.Errorf("calc_minus() failed to parse output: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("calc_minus(%f, %f) = %f; want %f", tt.n1, tt.n2, got, tt.want)
			}
		})
	}
}

// TestCalcDivision testa a função calc_division
func TestCalcDivision(t *testing.T) {
	tests := []struct {
		name string
		n1   float64
		n2   float64
		want float64
	}{
		{"Divisao Normal", 10, 2, 5},
		{"Divisao com Decimais", 7, 2, 3.5},
		{"Divisao por 1", 9, 1, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				calc_division(tt.n1, tt.n2)
			})

			got, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
			if err != nil {
				t.Errorf("calc_division() failed to parse output: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("calc_division(%f, %f) = %f; want %f", tt.n1, tt.n2, got, tt.want)
			}
		})
	}

	// Teste específico para divisão por zero
	t.Run("Divisao por Zero", func(t *testing.T) {
		output := captureStdout(func() {
			calc_division(10, 0) // Divisão por zero
		})

		// Em Go, 10/0 resulta em +Inf. Precisamos verificar isso.
		if !strings.Contains(strings.TrimSpace(output), "+Inf") {
			t.Errorf("calc_division(10, 0) did not result in +Inf. Got: %q", strings.TrimSpace(output))
		}
	})
}

// TestCalcMultiple testa a função calc_multiple
func TestCalcMultiple(t *testing.T) {
	tests := []struct {
		name string
		n1   float64
		n2   float64
		want float64
	}{
		{"Multiplicacao Positiva", 5, 4, 20},
		{"Multiplicacao por Zero", 8, 0, 0},
		{"Multiplicacao com Negativo", -3, 2, -6},
		{"Multiplicacao com Decimais", 2.5, 2, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				calc_multiple(tt.n1, tt.n2)
			})

			got, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
			if err != nil {
				t.Errorf("calc_multiple() failed to parse output: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("calc_multiple(%f, %f) = %f; want %f", tt.n1, tt.n2, got, tt.want)
			}
		})
	}
}

// TestMenu testa a função menu
func TestMenu(t *testing.T) {
	tests := []struct {
		name               string
		input              string
		want               int
		wantErr            bool
		expectedOutputPart string // Part of the expected output (like "CALCULADORA" or error message)
	}{
		{"Opcao Valida 1", "1\n", 1, false, "CALCULADORA"},
		{"Opcao Valida 2", "2\n", 2, false, "CALCULADORA"},
		{"Opcao Valida 5", "5\n", 5, false, "CALCULADORA"},
		{"Entrada Nao Numerica", "abc\n", 0, true, "expected integer"}, // This is the error from Scanln printed by fmt.Println(err)
		{"Entrada Vazia", "\n", 0, true, "expected integer"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedOutput string
			var actualResult int

			simulateStdin(tt.input, func() {
				capturedOutput = captureStdout(func() {
					actualResult = menu()
				})
			})

			if tt.wantErr {
				// Check if the expected error message is present in the captured output
				if !strings.Contains(capturedOutput, tt.expectedOutputPart) {
					t.Errorf("menu() with input '%s' did not output expected error message.\nGot: %q\nWant part: %q",
						tt.input, capturedOutput, tt.expectedOutputPart)
				}
				// The menu function returns 0 on error, so we check that.
				if actualResult != tt.want {
					t.Errorf("menu() with input '%s' returned %d, want %d when error", tt.input, actualResult, tt.want)
				}
			} else {
				// For valid inputs, check the returned value
				if actualResult != tt.want {
					t.Errorf("menu() with input '%s' = %d; want %d", tt.input, actualResult, tt.want)
				}
				// Also, ensure the menu text was printed
				if !strings.Contains(capturedOutput, tt.expectedOutputPart) {
					t.Errorf("menu() with input '%s' did not print menu text. Got: %q", tt.input, capturedOutput)
				}
			}
		})
	}
}
