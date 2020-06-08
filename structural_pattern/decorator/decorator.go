package decorator

import "fmt"

type DataFile interface {
	WriteFile()
	ReadFile(fileName string)
}

type SimpleDataFile struct {
}

func(s SimpleDataFile) WriteFile() {
	fmt.Println("Writing to a file")
}

func(s SimpleDataFile) ReadFile(fileName string) {
	fmt.Println("Read from a file", fileName)
}

type EncryptedDataFileWrapper struct {
	wrappedDatafile DataFile
}

func(e EncryptedDataFileWrapper) WriteFile() {
	fmt.Println("Encrypting the file")
	e.wrappedDatafile.WriteFile()
}

func(e EncryptedDataFileWrapper) ReadFile(fileName string) {
	fmt.Println("Decrypting the file", fileName)
	e.wrappedDatafile.ReadFile(fileName)
}

type CompressedDataFileWrapper struct {
	wrappedDatafile DataFile
}

func(c CompressedDataFileWrapper) WriteFile() {
	fmt.Println("Compressing the file")
	c.wrappedDatafile.WriteFile()
}

func(c CompressedDataFileWrapper) ReadFile(fileName string) {
	fmt.Println("Decompressing the file", fileName)
	c.wrappedDatafile.ReadFile(fileName)
}

func PlayingAroundWithDecoratorPattern() {
	data := SimpleDataFile{}
	encryptedData := EncryptedDataFileWrapper{wrappedDatafile: data}
	compressedData := CompressedDataFileWrapper{wrappedDatafile: encryptedData}
	compressedData.WriteFile()
	compressedData.ReadFile("coffee.txt")
}
