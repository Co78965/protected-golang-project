package Cipher

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const Flag = false
const ExpectedCRC32 = 0x43b8f634   //0x196d401e
const ExpectedCRC32_2 = 0xa1469479 //0x7a244a08

const key = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" //some random numbers here

func LoadCRCValues() (uint32, uint32, error) {
	// Загружаем переменные из .env файла
	err := godotenv.Load("Cipher/crc.env")
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка загрузки .env файла: %w", err)
	}

	// Получаем значения CRC1 и CRC2
	crc1Str := os.Getenv("CRC1")
	crc2Str := os.Getenv("CRC2")

	// Проверяем, что значения не пустые
	if crc1Str == "" || crc2Str == "" {
		return 0, 0, fmt.Errorf("CRC1 или CRC2 не найдены в .env файле")
	}

	// Преобразуем CRC1 из строки в uint32
	crc1, err := strconv.ParseUint(crc1Str, 16, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования CRC1: %w", err)
	}

	// Преобразуем CRC2 из строки в uint32
	crc2, err := strconv.ParseUint(crc2Str, 16, 32)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования CRC2: %w", err)
	}

	// Возвращаем значения как uint32
	return uint32(crc1), uint32(crc2), nil
}

func Encrypt(input string) (output string) {
	for i := 0; i < len(input); i++ {
		output += fmt.Sprintf("\\x%02x", input[i]^key[i%len(key)])
	}
	return output
}

func Decrypt(input string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}
