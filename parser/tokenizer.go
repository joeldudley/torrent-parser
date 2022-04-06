package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// TODO - Validation of bytes read, to check they're valid UTF-8/ints.

const (
	dictStartByte = byte('d')
	listStartByte = byte('l')
	endByte       = byte('e')
	intByte       = byte('i')
	zeroByteCode  = 48
	nineByteCode  = 57
)

// tokenize converts the Bencoded file at filePath into a slice of token.
func tokenize(filePath string) []token {
	file, err := os.Open(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("Could not read file %s.", filePath)
		log.Fatal(errMsg)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var tokens []token
	reader := bufio.NewReader(file)

	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			// Reached end of bytestream.
			break
		}

		switch readByte {
		case readByte == dictStartByte:
			tokens = append(tokens, token{dictStartToken, []byte{}})
		case readByte == listStartByte:
			tokens = append(tokens, token{listStartToken, []byte{}})
		case readByte == endByte:
			tokens = append(tokens, token{endToken, []byte{}})
		case readByte == intByte:
			tokens = append(tokens, extractIntToken(reader))
		case readByte >= zeroByteCode && readByte <= nineByteCode:
			tokens = append(tokens, extractStringToken(readByte, reader))
		default:
			errMsg := fmt.Sprintf("Unmatched token while tokenising file %s.", filePath)
			log.Fatal(errMsg)
		}
	}

	return tokens
}

// printTokens pretty-prints a slice of token.
func printTokens(tokens []token) {
	for _, nextToken := range tokens {
		switch nextToken.tokenType {
		case dictStartToken:
			fmt.Println("Found dict start")
		case listStartToken:
			fmt.Println("Found list start")
		case endToken:
			fmt.Println("Found end")
		case stringToken:
			fmt.Print("Found string: ")
			fmt.Println(nextToken.value)
		case intToken:
			fmt.Print("Found integer: ")
			fmt.Println(nextToken.value)
		}
	}
}

func extractIntToken(reader *bufio.Reader) token {
	var intBytes []byte

	for {
		readByte, err := reader.ReadByte()

		if err != nil {
			errMsg := fmt.Sprintf(
				"Reached end of input while tokenising an integer. Bytes so far: %s.",
				string(intBytes))
			log.Fatal(errMsg)
		}

		if readByte == byte('e') {
			break // Reached end of integer.
		} else {
			intBytes = append(intBytes, readByte)
		}
	}

	return token{intToken, intBytes}
}

func extractStringToken(readByte byte, reader *bufio.Reader) token {
	// The bytes representing the string's length.
	stringLenBytes := []byte{readByte}

	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			break
		}
		if readByte == byte(':') {
			break
		} else {
			stringLenBytes = append(stringLenBytes, readByte)
		}
	}

	length := bytesToInt(stringLenBytes)

	var stringBytes []byte
	for j := 0; j < length; j++ {
		x, err := reader.ReadByte()

		if err != nil {
			errMsg := fmt.Sprintf(
				"Reached end of input while tokenising a string. Bytes so far: %s.",
				string(stringBytes))
			log.Fatal(errMsg)
		}

		stringBytes = append(stringBytes, x)
	}

	return token{stringToken, stringBytes}
}

func bytesToInt(bytes []byte) int {
	length := 0
	for _, k := range bytes {
		length *= 10
		length += int(k) - 48
	}
	return length
}
