package main

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LeaveRecord struct {
	ID                    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FullName              string `json:"fullName" gorm:"not null"`
	LeaveType             string `json:"leaveType" gorm:"not null;type:leave_type_enum"`
	FromDate              string `json:"fromDate" gorm:"not null;type:date"`
	ToDate                string `json:"toDate" gorm:"not null;type:date"`
	Team                  string `json:"team" gorm:"not null"`
	MedicalCertificateUrl string `json:"medicalCertificateUrl"`
	Reporter              string `json:"reporter" gorm:"not null"`
}

// Define a struct to represent the result of KPI_3_Top_5_Employees_Leave_2023 query
type Top5EmployeesLeave2023 struct {
	FullName       string `json:"fullName"`
	TotalLeaveDays int    `json:"totalLeaveDays"`
}

// Define a struct to represent the result of KPI_4_Employees_Leave_Under_Manager_Q1_2023 query
type EmployeesLeaveManagerQ1 struct {
	ManagerName      string `json:"managerName"`
	EmployeesOnLeave int    `json:"employeesOnLeave"`
}

// Define a struct to represent the result of KPI_6_Top_2_Teams_Leave_Type_Distribution_2022 query
type TeamLeaveTypeDistribution struct {
	Team       string `json:"team"`
	LeaveType  string `json:"leaveType"`
	LeaveCount int    `json:"leaveCount"`
}

var (
	db           *gorm.DB
	leaveRecords []LeaveRecord
)

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.Default())

	// Initialize database and migrations
	db = initDB()
	migrateDB(db)

	// Register API routes
	registerRoutes(r)

	r.Run(":8080")
}

func initDB() *gorm.DB {
	dsn := "user=postgres password=123456 dbname=fullstacktaskdb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	return db
}

// func initDB() *gorm.DB {
// 	dbHost := os.Getenv("DB_HOST")
// 	dsn := fmt.Sprintf("host=%s user=postgres password=123456 dbname=fullstacktaskdb sslmode=disable", dbHost)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Error opening database:", err)
// 	}
// 	return db
// }

func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&LeaveRecord{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}
}

func applyLeave(c *gin.Context) {
	var leave LeaveRecord
	if err := c.ShouldBind(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if !isValidDate(leave.FromDate, leave.ToDate) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date range"})
		return
	}

	// Find the maximum ID from the database
	var maxID uint
	if err := db.Model(&LeaveRecord{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the maximum ID"})
		return
	}

	// Increment the maximum ID to generate the next available ID
	leave.ID = maxID + 1

	if err := db.Create(&leave).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply leave"})
		return
	}
	c.JSON(http.StatusCreated, leave)
}

func mergeAndAppendData(c *gin.Context) {
	var data []LeaveRecord

	// Bind the JSON data to a slice of LeaveRecord
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Append the received data to the leaveRecords slice
	leaveRecords = append(leaveRecords, data...)

	// Insert the data into the database
	if err := db.Create(data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		return
	}

	// You can add additional validation or database operations here

	c.JSON(http.StatusOK, gin.H{"message": "Data received and merged successfully"})
}

func getLeaveRecords(c *gin.Context) {
	var leaveRecords []LeaveRecord
	if err := db.Find(&leaveRecords).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leave records"})
		return
	}
	c.JSON(http.StatusOK, leaveRecords)
}

func uploadMedicalCertificate(c *gin.Context) {
	file, err := c.FormFile("medicalCertificate")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload"})
		return
	}

	// Perform validation on the uploaded file
	if err := validateMedicalCertificate(file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate unique filename and save the uploaded file
	uniqueFilename := generateUniqueFilename(file.Filename)
	uploadPath := "medical_certificates/" + uniqueFilename

	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Return the URL of the uploaded file
	c.JSON(http.StatusOK, gin.H{"medicalCertificateUrl": uploadPath})
}

func validateMedicalCertificate(file *multipart.FileHeader) error {
	// Check file size
	if file.Size > 15*1024*1024 { // 15MB
		return errors.New("File size should not exceed 15MB")
	}

	// Check file extension
	allowedExtensions := []string{".pdf", ".png"}
	ext := filepath.Ext(file.Filename)
	isValidExtension := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isValidExtension = true
			break
		}
	}
	if !isValidExtension {
		return errors.New("Invalid file extension. Allowed: pdf, png")
	}

	return nil
}

func generateUniqueFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)
	filename := strings.TrimSuffix(originalFilename, ext)
	uniqueFilename := filename + "_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	return uniqueFilename
}

func isValidDate(fromDate, toDate string) bool {
	from, errFrom := time.Parse("2006-01-02", fromDate)
	to, errTo := time.Parse("2006-01-02", toDate)
	if errFrom != nil || errTo != nil {
		return false
	}
	return !to.Before(from)
}

func getTop5EmployeesLeave2023(c *gin.Context) {
	var result []Top5EmployeesLeave2023
	err := db.Raw(`
		SELECT full_name, total_leave_days FROM kpi3_view;
	`).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Define a handler for KPI_4_Employees_Leave_Under_Manager_Q1_2023
func getEmployeesLeaveManagerQ1(c *gin.Context) {
	var result []EmployeesLeaveManagerQ1
	err := db.Raw(`
		SELECT manager_name, employees_on_leave FROM kpi4_view;
	`).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Define a handler for KPI_6_Top_2_Teams_Leave_Type_Distribution_2022
func getTeamLeaveTypeDistribution(c *gin.Context) {
	var result []TeamLeaveTypeDistribution
	err := db.Raw(`
		SELECT team, leave_type, leave_count FROM kpi6_view;
	`).Scan(&result).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func mainAPIs(r *gin.Engine) {
	r.POST("/api/apply-leave", applyLeave)
	r.GET("/api/leave-records", getLeaveRecords)
	r.POST("/api/upload-medical-certificate", uploadMedicalCertificate)
	r.POST("/api/merge-and-append-data", mergeAndAppendData)
	r.GET("/api/KPI_3_Top_5_Employees_Leave_2023", getTop5EmployeesLeave2023)
	r.GET("/api/KPI_4_Employees_Leave_Under_Manager_Q1_2023", getEmployeesLeaveManagerQ1)
	r.GET("/api/KPI_6_Top_2_Teams_Leave_Type_Distribution_2022", getTeamLeaveTypeDistribution)
}

func registerRoutes(r *gin.Engine) {
	// Main APIs
	mainAPIs(r)
}
