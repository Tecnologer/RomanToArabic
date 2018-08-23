windows:
	GOOS=windows go build -o dist/RomanConverter-win-x64.exe

linux:
	go build -o dist/RomanConverter-lin-x64

windowsx86:
	GOOS=windows GOARCH=386 go build -o dist/RomanConverter-win-x86.exe

linuxx86:
	GOARCH=386 go build -o dist/RomanConverter-lin-x86

all:
	GOOS=windows go build -o dist/RomanConverter-win-x64.exe
	go build -o dist/RomanConverter-lin-x64
	GOOS=windows GOARCH=386 go build -o dist/RomanConverter-win-x86.exe
	GOARCH=386 go build -o dist/RomanConverter-lin-x86