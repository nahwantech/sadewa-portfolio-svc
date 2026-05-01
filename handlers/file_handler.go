package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"sadewa-portfolio-svc/config"

	"github.com/minio/minio-go/v7"
)

// Allowed file extensions
var allowedExtensions = map[string]bool{
	".png":  true,
	".jpeg": true,
	".jpg":  true,
	".pdf":  true,
	".docx": true,
	".doc":  true,
	".xls":  true,
	".xlsx": true,
}

// UploadResponse represents the response for upload endpoint
type UploadResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	FileName string `json:"file_name,omitempty"`
	Bucket   string `json:"bucket,omitempty"`
	FileURL  string `json:"file_url,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// UploadFileHandler handles file upload to MinIO
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Only accept POST requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(32 << 20) // 32 MB max memory
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Failed to parse form data",
		})
		return
	}

	// Get bucket name from query parameter or form data
	bucketName := r.URL.Query().Get("bucket")
	if bucketName == "" {
		bucketName = r.FormValue("bucket")
	}

	if bucketName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Bucket name is required (use 'bucket' query parameter or form field)",
		})
		return
	}

	// Get file from request
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Failed to get file from request",
		})
		return
	}
	defer file.Close()

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if !allowedExtensions[ext] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("File type '%s' not allowed. Allowed types: .png, .jpeg, .jpg, .pdf, .docx, .doc, .xls, .xlsx", ext),
		})
		return
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := config.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to check bucket existence: %v", err),
		})
		return
	}

	if !exists {
		err = config.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{
				Success: false,
				Error:   fmt.Sprintf("Failed to create bucket: %v", err),
			})
			return
		}
	}

	// Set content type based on file extension
	contentType := getContentType(ext)

	// Upload file to MinIO
	objectName := header.Filename
	_, err = config.MinioClient.PutObject(ctx, bucketName, objectName, file, header.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to upload file: %v", err),
		})
		return
	}

	// Generate file URL
	fileURL := fmt.Sprintf("http://%s/%s/%s", config.MinioClient.EndpointURL().Host, bucketName, objectName)

	// Success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UploadResponse{
		Success:  true,
		Message:  "File uploaded successfully",
		FileName: objectName,
		Bucket:   bucketName,
		FileURL:  fileURL,
	})
}

// DownloadFileHandler handles file download from MinIO
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept GET requests
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	// Get bucket name and file name from query parameters
	bucketName := r.URL.Query().Get("bucket")
	fileName := r.URL.Query().Get("file")

	if bucketName == "" || fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "Bucket name and file name are required",
		})
		return
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(fileName))
	if !allowedExtensions[ext] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("File type '%s' not allowed. Allowed types: .png, .jpeg, .jpg, .pdf, .docx, .doc, .xls, .xlsx", ext),
		})
		return
	}

	// Get object from MinIO
	ctx := context.Background()
	object, err := config.MinioClient.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to get file: %v", err),
		})
		return
	}
	defer object.Close()

	// Get object info to set headers
	objectInfo, err := object.Stat()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{
			Success: false,
			Error:   "File not found",
		})
		return
	}

	// Set content type and headers
	contentType := getContentType(ext)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", objectInfo.Size))

	// Copy file content to response
	_, err = io.Copy(w, object)
	if err != nil {
		// Can't write error response after headers are sent
		fmt.Printf("Error streaming file: %v\n", err)
	}
}

// getContentType returns the MIME type based on file extension
func getContentType(ext string) string {
	contentTypes := map[string]string{
		".png":  "image/png",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".pdf":  "application/pdf",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".doc":  "application/msword",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}

	if ct, ok := contentTypes[ext]; ok {
		return ct
	}
	return "application/octet-stream"
}
