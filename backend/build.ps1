## Build
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
$env:URL = "FUNCTION URL"
go build -o .\deploy\main .\main.go

## Compress
Compress-Archive -Path .\deploy\main -DestinationPath .\deploy\main.zip -Force

## Deploy
aws lambda update-function-code --function-name food-votes --zip-file fileb://deploy/main.zip --query 'LastModified' --output text

## Wait 10 seconds
Start-Sleep -s 10

## Test with GET request and ensure 200 status code
$test = Invoke-WebRequest -Method Get -Uri $env:URL/?zipcode=20148
if ($test.StatusCode -eq 200) {
    Write-Host "Get Local - TEST PASSED"
} else {
    Write-Host "Get Local - TEST FAILED"
}

# Clear env variables
$env:GOOS = ""
$env:GOARCH = ""
$env:CGO_ENABLED = ""