# MinIO File Upload and Download API

This document describes how to use the MinIO file upload and download endpoints in the Sadewa Portfolio Service.

## Configuration

MinIO credentials are configured in the `.env` file:

```env
MINIO_ENDPOINT=91.108.104.69:9001
MINIO_ACCESS_KEY=nahwan_alsaki
MINIO_SECRET_KEY=sakaya0903#$
MINIO_USE_SSL=false
```

## Supported File Types

The API only accepts the following file types:
- Images: `.png`, `.jpeg`, `.jpg`
- Documents: `.pdf`, `.docx`, `.doc`
- Spreadsheets: `.xls`, `.xlsx`

Any other file types will be rejected with a 400 Bad Request error.

## API Endpoints

### 1. Upload File

**Endpoint:** `POST /api/upload`

**Parameters:**
- `bucket` (required): Bucket name - can be provided as query parameter or form field
- `file` (required): The file to upload (multipart form data)

**Query Parameter Example:**
```
POST /api/upload?bucket=my-bucket
```

**Form Data Example:**
```
POST /api/upload
Content-Type: multipart/form-data

bucket=my-bucket
file=<binary file data>
```

**cURL Example:**
```bash
# Using query parameter
curl -X POST "http://localhost:8089/api/upload?bucket=portfolio-files" \
  -F "file=@/path/to/your/document.pdf"

# Using form field
curl -X POST "http://localhost:8089/api/upload" \
  -F "bucket=portfolio-files" \
  -F "file=@/path/to/your/image.png"
```

**Success Response (200 OK):**
```json
{
  "success": true,
  "message": "File uploaded successfully",
  "file_name": "document.pdf",
  "bucket": "portfolio-files",
  "file_url": "http://91.108.104.69:9001/portfolio-files/document.pdf"
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "File type '.exe' not allowed. Allowed types: .png, .jpeg, .jpg, .pdf, .docx, .doc, .xls, .xlsx"
}
```

### 2. Download File

**Endpoint:** `GET /api/download`

**Parameters:**
- `bucket` (required): Bucket name (query parameter)
- `file` (required): File name (query parameter)

**cURL Example:**
```bash
curl -X GET "http://localhost:8089/api/download?bucket=portfolio-files&file=document.pdf" \
  --output downloaded-file.pdf
```

**Browser Example:**
```
http://localhost:8089/api/download?bucket=portfolio-files&file=image.png
```

**Success Response:**
Returns the file content with appropriate headers:
- `Content-Type`: Based on file extension
- `Content-Disposition`: inline; filename="filename.ext"
- `Content-Length`: File size in bytes

**Error Response (404 Not Found):**
```json
{
  "success": false,
  "error": "File not found"
}
```

**Error Response (400 Bad Request):**
```json
{
  "success": false,
  "error": "File type '.exe' not allowed. Allowed types: .png, .jpeg, .jpg, .pdf, .docx, .doc, .xls, .xlsx"
}
```

## Features

1. **Automatic Bucket Creation**: If the specified bucket doesn't exist, it will be created automatically.

2. **File Type Validation**: Only specified file types are allowed for upload and download.

3. **Content Type Detection**: The API automatically sets the correct MIME type based on file extension.

4. **File URL Generation**: Upload response includes a direct URL to access the file.

## Testing with Postman

### Upload File
1. Create a new POST request to `http://localhost:8089/api/s3/upload?bucket=test-bucket`
2. Go to the "Body" tab
3. Select "form-data"
4. Add a key named "file" and change its type to "File"
5. Choose a file from your computer
6. Send the request

### Download File
1. Create a new GET request to `http://localhost:8089/api/s3/download?bucket=test-bucket&file=yourfile.pdf`
2. Send the request
3. The file will be displayed or downloaded based on its type

## Error Handling
The API handles various error scenarios:
- Missing bucket name
- Missing file in upload request
- Invalid file type
- Bucket creation failures
- File upload failures
- File not found during download
- Network or MinIO connection errors

All errors return appropriate HTTP status codes and JSON error messages.
