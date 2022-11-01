# !/bin/bash
function newadv() {
    #usage newapp 2015/01
    app=$1
    echo $app
    wd=$(pwd)
    
    appdir=$wd/$app
    mkdir -p $appdir

    cd $appdir
    go mod init $app
    rm -f main.go && printf 'package main\n\nimport "fmt"\n\nfunc main(){\n\tfmt.Println("Hello")\n}\n' > main.go
    rm -f main_test.go && printf 'package main\n\nimport("testing" \n "github.com/stretchr/testify/require" \n)  \nfunc Test(t *testing.T){\n\tfor _, tc := range []struct { \n\t desc string \n } {\n{\n desc:\"\", \n},\n } { \n\tt.Run(tc.desc, func(t *testing.T) { require.Equal(t, 1, 1) \n})\n}\n}\n' > main_test.go
    go fmt
    go get -u "github.com/stretchr/testify/require"/
    go build
    go run main.go
    go test ./...
    #code .
}
