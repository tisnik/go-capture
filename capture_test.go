/*
Copyright © 2020 Pavel Tisnovsky

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package capture_test

import (
	"fmt"
	"github.com/tisnik/go-capture"
	"os"
	"testing"
)

// TestNoOutput checks if empty standard output is captured properly
func TestNoOutput(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Standard should be empty")
	}
}

// TestEmptyOutput checks if empty standard output is captured properly
func TestEmptyOutput(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Print("")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Standard should be empty")
	}
}

// TestOutputWithoutNewlines checks if standard output created by fmt.Print is captured properly
func TestOutputWithoutNewlines(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Print("Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello!" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestOutputWithNewlines checks if standard output created by fmt.Println is captured properly
func TestOutputWithNewlines(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Println("Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello!\n" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestOutputToStdErr checks whether output to stderr is captured or not
func TestOutputToStdErr(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Fprint(os.Stderr, "Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestOutputToStdOutAndStdErr checks whether output to stderr is captured or not
func TestOutputToStdOutAndStdErr(t *testing.T) {
	captured, err := capture.StandardOutput(func() {
		fmt.Fprint(os.Stdout, "Hello to stdout!")
		fmt.Fprint(os.Stderr, "Hello to stderr!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello to stdout!" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}

// TestNoOutput checks if empty error output is captured properly
func TestNoErrorOutput(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
	})
	if err != nil {
		t.Fatal("Unable to capture error output", err)
	}
	if captured != "" {
		t.Fatal("Error should be empty")
	}
}

// TestEmptyErrorOutput checks if empty error output is captured properly
func TestEmptyErrorOutput(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
		fmt.Print("")
	})
	if err != nil {
		t.Fatal("Unable to capture error output", err)
	}
	if captured != "" {
		t.Fatal("Error should be empty")
	}
}

// TestErrorOutputWithoutNewlines checks if error output created by fmt.Fprint is captured properly
func TestErrorOutputWithoutNewlines(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
		fmt.Fprint(os.Stderr, "Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture error output", err)
	}
	if captured != "Hello!" {
		t.Fatal("Incorrect error output has been captured:", captured)
	}
}

// TestErrorOutputWithNewlines checks if error output created by fmt.Fprintln is captured properly
func TestErrorOutputWithNewlines(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
		fmt.Fprintln(os.Stderr, "Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture error output", err)
	}
	if captured != "Hello!\n" {
		t.Fatal("Incorrect error output has been captured:", captured)
	}
}

// TestOutputToStdout checks whether output to stdout is captured or not
func TestOutputToStdou(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
		fmt.Fprint(os.Stdout, "Hello!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "" {
		t.Fatal("Incorrect error output has been captured:", captured)
	}
}

// TestOutputToStdOutAndStdErr2 checks whether output to stderr is captured or not
func TestOutputToStdOutAndStdErr2(t *testing.T) {
	captured, err := capture.ErrorOutput(func() {
		fmt.Fprint(os.Stdout, "Hello to stdout!")
		fmt.Fprint(os.Stderr, "Hello to stderr!")
	})
	if err != nil {
		t.Fatal("Unable to capture standard output", err)
	}
	if captured != "Hello to stderr!" {
		t.Fatal("Incorrect output has been captured:", captured)
	}
}
