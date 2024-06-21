package main

import (
	 "fmt"
	 "unsafe" // For pointers
)

/**
 * Using cgo to link the static library; as a result, users shall install the gcc package;
 * for example, "apk add gcc musl-dev" in golang:1.19.3-alpine
 */

/*
#cgo LDFLAGS: -L./Libs -ljsonDecoder
#include "Headers/BuildDLL/BuildDLL.h"
*/
import "C"

/**
 * Main function
 */
func main() {
	// The plainText string -> byte array -> unsigned char*

	var plainText = "Hello world!1" ;
	var plainTextByteFormat = []byte(plainText);

	// For the parameter of the function in the library 
	var plainTextLength = uint32(len(plainTextByteFormat));
	var cPlainTextByteFormat = (*C.uchar)(unsafe.Pointer(&plainTextByteFormat[0]));
	
	// Creating a byte result and defining the pointer for C; the f
	var cipherByteArray = make ([]byte, plainTextLength * 3);
	var cCipherPointer = (*C.uchar)(unsafe.Pointer(&cipherByteArray[0]));
	
	// The parameters are the plaintext pointer, the unsigned integer length of the plaintext and the cipher text pointer to the memory space, respectively.
	var result = C.___encryptSEDecoder(cPlainTextByteFormat, C.uint(plainTextLength), cCipherPointer);


	fmt.Printf( "The length of the encrypted string in the hex format: %d\n", result);
	fmt.Println( "The encrypted string in the hex format:", string(cipherByteArray)[:result]);

	// Decryption phase
	// Creating a pointer to the copiedCipherBytes
	var copiedCipherBytes = make ([]byte, result);
	var copiedCipherPointer = (*C.uchar)(unsafe.Pointer(&copiedCipherBytes[0]));
	copy(copiedCipherBytes, cipherByteArray);
	
	var copiedPlainText = make ([]byte, result);
	var copiedPlainPointer = (*C.uchar)(unsafe.Pointer(&copiedPlainText[0]));

	// The parameters are the cipher text pointer, the unsigned integer length of the cipher text and the plaintext pointer to the memory space, respectively.
	result = C.___decryptSEDecoder(copiedCipherPointer, C.uint(result), copiedPlainPointer);

	fmt.Printf( "The length of the encrypted string in the hex format: %d\n", result);
	fmt.Println( "The encrypted string in the hex format:", string(copiedPlainText)[:result]);	
}
