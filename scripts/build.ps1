$RELEASE_BIN_DIR='.\bin\'
$BINARY_NAME='websocket-tester'
$PACKAGE_NAME='websocket-tester/cmd/common'
function removeDir() {
    Remove-Item -Path $RELEASE_BIN_DIR -Recurse
}

function createDir() {
    if (Test-Path -Path $RELEASE_BIN_DIR) {
        echo "path exists"
    } else {
        echo "path not exists"
        New-Item -Path $RELEASE_BIN_DIR -ItemType Directory
    }
}

function goBuild() {
    param(
        [string]$os,
        [string]$arch,
        [string]$program
    )
    $suffix=''
    if ($os -like "windows") {
        $suffix='.exe'
    }
    $versionCode=Get-Date -format "yyyyMMdd"
    $goVersion=go version
    $gitHash=git log --pretty=format:'%h' -n 1
    $buildTime=git log --pretty=format:'%cd' -n 1
    set CGO_ENABLED=0
    go env -w CGO_ENABLED=0
    set GOOS=$os
    go env -w GOOS=$os
    set GOARCH=$arch
    go env -w GOARCH=$arch
    go build -o "$RELEASE_BIN_DIR$BINARY_NAME-${program}-${os}_$arch$suffix" -ldflags "-w -s -X '$PACKAGE_NAME._version=1.0.$versionCode' -X '$PACKAGE_NAME._goVersion=$goVersion' -X '$PACKAGE_NAME._gitHash=$gitHash' -X '$PACKAGE_NAME._buildTime=$buildTime'" ./cmd/$program
}

function goBuildAll()
{
    param(
        [string]$os,
        [string]$arch
    )
    goBuild $os $arch "client"
    goBuild $os $arch "server"
}

function main() {
    removeDir
    go clean
    go mod tidy
    createDir
    goBuildAll linux amd64
    goBuildAll linux arm64
    goBuildAll darwin arm64
    goBuildAll darwin amd64
    goBuildAll windows amd64
    goBuildAll windows arm64
}

main