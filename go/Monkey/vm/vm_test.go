package vm

import (
	"fmt"
	"testing"

	"github.com/akif999/compiler_in_go/go/Monkey/ast"
	"github.com/akif999/compiler_in_go/go/Monkey/compiler"
	"github.com/akif999/compiler_in_go/go/Monkey/lexer"
	"github.com/akif999/compiler_in_go/go/Monkey/object"
	"github.com/akif999/compiler_in_go/go/Monkey/parser"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 2}, // FIX ME
	}

	runVmTests(t, tests)
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatal("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.StackTop()

		testExpectedObject(t, tt.expected, stackElem)
	}
}

func testExpectedObject(
	t *testing.T,
	expected interface{},
	actual object.Object,
) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntergerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntergerObject failed: %s", err)
		}
	}
}

func testIntergerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)",
			actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
	}

	return nil
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
