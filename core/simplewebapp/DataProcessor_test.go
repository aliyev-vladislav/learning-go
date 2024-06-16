package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser_IsCorrectForValidInput(t *testing.T) {
	testdata := []struct {
		name     string
		data     []byte
		expected Input
	}{
		{"addition", []byte("CALC_1\n+\n2\n2"), Input{"CALC_1", "+", 2, 2}},
		{"subtraction", []byte("CALC_2\n-\n2\n2"), Input{"CALC_2", "-", 2, 2}},
		{"multiplication", []byte("CALC_3\n*\n2\n2"), Input{"CALC_3", "*", 2, 2}},
		{"division", []byte("CALC_4\n/\n2\n2"), Input{"CALC_4", "/", 2, 2}},
	}

	for _, d := range testdata {
		t.Run(d.name, func(t *testing.T) {
			result, err := parser(d.data)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(result, d.expected); diff != "" {
				t.Errorf(diff)
			}
		})
	}

}

func TestParse_ErrorsOnInvalidInput(t *testing.T) {
	testdata := []struct {
		name     string
		data     []byte
		expected Input
	}{
		{"invalid_left_operand", []byte("CALC_5\n+\nchar\n2"), Input{}},
		{"invalid_right_operand", []byte("CALC_6\n+\n2\nchar"), Input{}},
	}

	for _, d := range testdata {
		t.Run(d.name, func(t *testing.T) {
			_, err := parser(d.data)
			if err == nil {
				t.Fatal(err)
			}
		})
	}

}

func TestDataProcessor(t *testing.T) {
	testdata := []struct {
		name        string
		inputValues [][]byte
		output      []Result
	}{
		{"addition", [][]byte{[]byte("CALC_1\n+\n2\n2")}, []Result{{"CALC_1", 4}}},
		{"subtraction", [][]byte{[]byte("CALC_2\n-\n2\n2")}, []Result{{"CALC_2", 0}}},
		{"multiplication", [][]byte{[]byte("CALC_3\n*\n2\n2")}, []Result{{"CALC_3", 4}}},
		{"division", [][]byte{[]byte("CALC_4\n/\n2\n2")}, []Result{{"CALC_4", 1}}},
		{"invalid_op", [][]byte{[]byte("CALC_5\n?\n2\n2"), []byte("CALC_6\n+\n2\n2")}, []Result{{"CALC_6", 4}}},
		{"invalid_operand", [][]byte{[]byte("CALC_6\n+\nchar\n2"), []byte("CALC_7\n*\n3\n3"), []byte("CALC_8\n/\n10\n2")}, []Result{{"CALC_7", 9}, {"CALC_8", 5}}},
	}

	for _, d := range testdata {
		t.Run(d.name, func(t *testing.T) {
			inChan := make(chan []byte)
			outChan := make(chan Result, len(d.output))
			go DataProcessor(inChan, outChan)
			for _, v := range d.inputValues {
				inChan <- v
			}

			close(inChan)
			result := make([]Result, 0)
			for v := range outChan {
				result = append(result, v)
			}
			if diff := cmp.Diff(result, d.output); diff != "" {
				t.Errorf(diff)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	testdata := []struct {
		name         string
		intputValues []Result
		writtenData  []byte
	}{
		{"case1", []Result{{"CALC_1", 10}, {"CALC_2", 4}}, []byte("CALC_1:10\nCALC_2:4\n")},
	}
	for _, d := range testdata {
		t.Run(d.name, func(t *testing.T) {
			var wg sync.WaitGroup
			wg.Add(1)
			in := make(chan Result, len(d.intputValues))
			for _, v := range d.intputValues {
				in <- v
			}
			close(in)
			f, err := os.Create("result_test.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				f.Close()
				os.Remove(f.Name())
			}()
			go func() {
				WriteData(in, f)
				wg.Done()
			}()
			f.Sync()
			wg.Wait()
			fContent, err2 := os.ReadFile(f.Name())
			if err2 != nil {
				t.Fatal(err2)
			}
			if diff := cmp.Diff(fContent, d.writtenData); diff != "" {
				t.Errorf(diff)
			}

		})
	}
}

type errReader int

func (errReader) Read(data []byte) (int, error) {
	return 0, errors.New("bad input")
}

func TestNewController(t *testing.T) {
	testdata := []struct {
		name         string
		intputValues []byte
		expectedCode int
		expectedBody string
	}{
		{"Valid input", []byte("CALC_1\n*\n100\n1000"), http.StatusAccepted, "OK: 1"},
		{"Error reading body", []byte("CALC_2..."), http.StatusBadRequest, "Bad Input"},
	}
	for _, d := range testdata {
		t.Run(d.name, func(t *testing.T) {
			out := make(chan []byte, 100)
			var req *http.Request
			if d.expectedCode == http.StatusBadRequest {
				req = httptest.NewRequest(http.MethodPost, "/", errReader(0))
			} else {
				req = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(d.intputValues))
			}
			res := httptest.NewRecorder()
			handler := NewController(out)
			handler.ServeHTTP(res, req)
			if res.Code != d.expectedCode {
				t.Errorf("Expected %d, but got %d", http.StatusAccepted, res.Code)
			}
			if res.Body.String() != d.expectedBody {
				t.Errorf("Expected body of %s, but got %s", d.expectedBody, res.Body.String())
			}
		})
	}
	t.Run("Channel is backed up", func(t *testing.T) {
		expected := "Too Busy: 1"
		out := make(chan []byte, 1)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("...")))
		res := httptest.NewRecorder()
		req2 := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte("...")))
		res2 := httptest.NewRecorder()
		handler := NewController(out)
		handler.ServeHTTP(res, req)
		handler.ServeHTTP(res2, req2)

		if res2.Code != http.StatusServiceUnavailable {
			t.Errorf("Expected %d, but got %d", http.StatusServiceUnavailable, res.Code)
		}
		if res2.Body.String() != expected {
			t.Errorf("Expected body of %s, but got %s", expected, res.Body.String())
		}

	})

}
