## Build
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
$env:URL = "FUNCTION URL"
go build -tags lambda.norpc -o bootstrap .\main.go

## Compress
Compress-Archive -Path bootstrap -DestinationPath deploy.zip -Force

## Deploy
aws lambda update-function-code --function-name food-votes --zip-file fileb://deploy.zip --query 'LastModified' --output text

## Wait 10 seconds
Start-Sleep -s 10

## Test with GET request and ensure 200 status code
$test = Invoke-WebRequest -Method Get -Uri $env:URL/?query=20148
if ($test.StatusCode -eq 200) {
    Write-Host "TESTS PASSED"
} else {
    Write-Host "TEST FAILED"
}

# Clear env variables
$env:GOOS = ""
$env:GOARCH = ""
$env:CGO_ENABLED = ""