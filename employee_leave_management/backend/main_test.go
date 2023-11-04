package main

import (
	"mime/multipart"
	"testing"
)

func TestIsValidDate(t *testing.T) {
	// Test valid date
	valid := isValidDate("2023-09-01", "2023-09-05")
	if !valid {
		t.Error("Expected valid date range, but got invalid")
	}

	// Test invalid date
	invalid := isValidDate("2023-09-05", "2023-09-01")
	if invalid {
		t.Error("Expected invalid date range, but got valid")
	}
}

func TestGenerateUniqueFilename(t *testing.T) {
	// Test generating unique filenames
	filename1 := generateUniqueFilename("file.pdf")
	filename2 := generateUniqueFilename("file.png")

	if filename1 == filename2 {
		t.Error("Expected unique filenames, but got the same filename")
	}
}

func TestValidateMedicalCertificate(t *testing.T) {
	// Test valid PDF file
	pdfFile := &multipart.FileHeader{
		Filename: "certificate.pdf",
		Size:     10 * 1024 * 1024, // 10MB (within the limit)
	}
	err := validateMedicalCertificate(pdfFile)
	if err != nil {
		t.Error("Expected valid PDF file, but got an error:", err)
	}

	// Test valid PNG file
	pngFile := &multipart.FileHeader{
		Filename: "certificate.png",
		Size:     5 * 1024 * 1024, // 5MB (within the limit)
	}
	err = validateMedicalCertificate(pngFile)
	if err != nil {
		t.Error("Expected valid PNG file, but got an error:", err)
	}

	// Test oversized file
	oversizedFile := &multipart.FileHeader{
		Filename: "oversized.pdf",
		Size:     16 * 1024 * 1024, // 16MB (exceeds the limit)
	}
	err = validateMedicalCertificate(oversizedFile)
	if err == nil {
		t.Error("Expected error for oversized file, but got none")
	}

	// Test invalid file extension
	invalidFile := &multipart.FileHeader{
		Filename: "invalid.txt",
		Size:     1 * 1024 * 1024,
	}
	err = validateMedicalCertificate(invalidFile)
	if err == nil {
		t.Error("Expected error for invalid file extension, but got none")
	}
}

func TestInitDB(t *testing.T) {
	// Test initializing the database
	db := initDB()
	if db == nil {
		t.Error("Expected a valid database connection, but got nil")
	}
}

//If all tests pass, you will see output similar to the following:
//PASS
//ok your/package   0.123s

//If any test fails, you will see output similar to the following:
// --- FAIL: TestFunctionName (0.00s)
//     main_test.go:20: Expected valid date range, but got invalid
// FAIL
// exit status 1
// FAIL    your/package   0.123s

//If there are no tests to run or no test files found, you may see output like this:
//?       your/package   [no test files]
