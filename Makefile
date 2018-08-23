windows:
	GOOS=windows go build -o dist/RomanConverter-win-x64.exe

linux:
	go build -o dist/RomanConverter-lin-x64

darwin:
	GOOS=darwin go build -o dist/RomanConverter-darwin-x64.app

windowsx86:
	GOOS=windows GOARCH=386 go build -o dist/RomanConverter-win-x86.exe

linuxx86:
	GOARCH=386 go build -o dist/RomanConverter-lin-x86

darwin:
	GOOS=darwin GOARCH=386 go build -o dist/RomanConverter-darwin-x86.app

all:
	GOOS=windows go build -o dist/RomanConverter-win-x64.exe
	go build -o dist/RomanConverter-lin-x64
	GOOS=windows GOARCH=386 go build -o dist/RomanConverter-win-x86.exe
	GOARCH=386 go build -o dist/RomanConverter-lin-x86
	GOOS=darwin go build -o dist/RomanConverter-darwin-x64.app
	GOOS=darwin GOARCH=386 go build -o dist/RomanConverter-darwin-x86.app