package filex

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStr(t *testing.T) {
	content := "Hello, World!"
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)
	filePath := dir + "/test.txt"
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	str, err := ToStr(filePath)
	if err != nil {
		t.Errorf("failed to convert file to string: %v", err)
	}
	assert.Equal(t, content, str, content)
}

func TestToBytes(t *testing.T) {
	content := "Hello, World!"
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)
	filePath := dir + "/test.txt"
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	b, err := ToBytes(filePath)
	if err != nil {
		t.Errorf("failed to convert file to bytes: %v", err)
	}
	assert.Equal(t, content, string(b), content)
}

func TestToTrimStr(t *testing.T) {
	content := "   Hello, World!   \n"
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)
	filePath := dir + "/test.txt"
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}
	str, err := ToTrimStr(filePath)
	if err != nil {
		t.Errorf("failed to convert file to trimmed string: %v", err)
	}
	expected := strings.TrimSpace(content)
	assert.Equal(t, expected, str, content)
}
