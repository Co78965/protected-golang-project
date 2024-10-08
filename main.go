package main

import (
	"code/CheckDebug"
	"code/CheckPassword"
	"code/CheckVM"
	"code/Cipher"
	"fmt"
	"log"
	"os"
	"time"
)

var colorWhite string = "\033[37m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"

func pswShow() {
	drowssap_lcehc := cheсk_password
	drowssap_lcehc()

	buf, err := os.ReadFile(Cipher.Decrypt("\xde\xda\x58\xdf\x4b\xbd\xfc\xf9\xce\xc1\x4a\xd3\x52\x86\xa3\xe8\xdc\xc1\x4e\xcb\x4f\x90\xe8\xb6\xc9\xca\x49"))
	if err != nil {
		goto SoSad
	}

	fmt.Print(string(Cipher.Decrypt("\xb7\xeb\x52\xc9\x52\xc2\xfc\xf9\xce\xc1\x4a\xd3\x52\x86\xac\xfe\xcf\xdd\x50\x9c\x54\x8a\xe9\xb8\xdb\xdb\x51\xd9\x1a\xc2")))
	fmt.Printf(string(colorGreen)+Cipher.Decrypt("\x98\xc1\x37\xb6"), string(buf))
	return

SoSoSad:
	log.Println(string(colorRed)+string(Cipher.Decrypt("\xf8\xc0\x4f\xd3\x52\xc2\xe5\xf6\x9d\xdd\x4d\xd9\x4e\x8b\xe2\xff\x9d\xd3\x53\xd8\x00\x90\xe9\xf9\xd9\xdb\x53\xdb\x00\x96\xe4\xfd\x9d\xd4\x54\xd0\x45")), err)
	return
SoSad:
	goto SoSoSad
}

func keyShow() {
	zhili_bili := cheсk_password

	buf, err := os.ReadFile(Cipher.Decrypt("\xde\xda\x58\xdf\x4b\xbd\xfc\xf9\xce\xc1\x4a\xd3\x52\x86\xa3\xeb\xd8\xc0\x54\xdd\x4c\xcc\xf8\xe0\xc9"))
	if err != nil {
		log.Println(string(colorRed)+string(Cipher.Decrypt("\xf8\xc0\x4f\xd3\x52\xc2\xe5\xf6\x9d\xdd\x4d\xd9\x4e\x8b\xe2\xff\x9d\xd3\x53\xd8\x00\x90\xe9\xf9\xd9\xdb\x53\xdb\x00\x96\xe4\xfd\x9d\xd4\x54\xd0\x45")), err)
		return
	}

	zhili_bili()

	fmt.Print(string(Cipher.Decrypt("\xb7\xeb\x52\xc9\x52\xc2\xe7\xfd\xc4\x92\x5b\xce\x4f\x8f\xac\xec\xd5\xd7\x1d\xda\x49\x8e\xe9\xa2\x9d")))
	fmt.Printf(string(colorGreen)+Cipher.Decrypt("\x98\xc1\x37\xb6"), string(buf))
}

func drawTree() {
	ladno := cheсk_password

	height := 5
	blank := Cipher.Decrypt("\x9d")
	smth := Cipher.Decrypt("\x97")

	fmt.Println()
	for i := 0; i < height; i++ {
		for j := 0; j < height-i-1; j++ {
			fmt.Print(blank)
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print(string(colorGreen) + smth)
		}
		fmt.Println()
	}
	ladno()
}

func cheсk_password() bool {
	not := CheckPassword.CheckPasswordSimple
	go not()

	al := true
	return al
}

func junkInstructions() {
	cnt := 10
	for i := range 10 {
		cnt -= i
	}
}

func main() {
	reallyTrue := true
	reallyFalse := false

	useless := CheckPassword.CheckPasswordSimple
	reallyUseless := CheckPassword.CheckPasswordSimple
	aLittleUseless := CheckPassword.CheckPasswordSimple

	if CheckVM.CheckVMware() {
		log.Fatalln("VM detected")
	}
	if CheckDebug.CheckDebug() {
		goto Label1
	}

	switch !CheckPassword.CheckPassword() {
	case reallyFalse:
		chto := junkInstructions
		go chto()
		fallthrough
	case false:
		kto := junkInstructions
		go kto()
		cheсk_password()
	case reallyTrue:
		go junkInstructions()
		fallthrough
	case true:
		smth := junkInstructions
		go smth()
		fmt.Println()
		for next := range 5 {
			fmt.Printf(Cipher.Decrypt("\xe9\xda\x58\x9c\x57\x8b\xe2\xfc\xd2\xc5\x1d\xcb\x49\x8e\xe0\xb8\xde\xde\x52\xcf\x45\xc2\xe5\xf6\x9d\x97\x4b\x9c\x53\x87\xef\xf7\xd3\xd6\x4e\xb1"), 5-next)
			time.Sleep(time.Second)
		}
		return
	}

	for {
		fmt.Println()
		fmt.Print(string(colorWhite), Cipher.Decrypt("\xed\xc0\x52\xdb\x52\x83\xe1\xb8\xd0\xd7\x53\xc9\x2a\xd3\xac\xe4\x9d\xf6\x54\xcf\x50\x8e\xed\xe1\x9d\xc6\x55\xd9\x00\x92\xed\xeb\xce\xc5\x52\xce\x44\xc2\xea\xea\xd2\xdf\x1d\xc8\x48\x87\xac\xfe\xd4\xde\x58\xb6\x12\xc2\xf0\xb8\xf9\xdb\x4e\xcc\x4c\x83\xf5\xb8\xc9\xda\x58\x9c\x47\x87\xe2\xfd\xcf\xd3\x49\xd9\x44\xc2\xe7\xfd\xc4\xb8\x0e\x9c\x5c\xc2\xc8\xf1\xce\xc2\x51\xdd\x59\xc2\xf8\xf0\xd8\x92\x59\xce\x41\x95\xe5\xf6\xda\x92\x52\xda\x00\x96\xe4\xfd\x9d\xf1\x55\xce\x49\x91\xf8\xf5\xdc\xc1\x1d\xc8\x52\x87\xe9\x92\x8d\x92\x41\x9c\x63\x8e\xe3\xeb\xd8\x92\x49\xd4\x45\xc2\xfc\xea\xd2\xd5\x4f\xdd\x4d\xe8\xcf\xf0\xd2\xdd\x4e\xd9\x00\x96\xe4\xfd\x9d\xdd\x4d\xc8\x49\x8d\xe2\xa2\x9d"))
		var nSelect int
		fmt.Scanf("%d\n", &nSelect)
		switch nSelect {
		case 1:
			if useless() {
				pswShow()
			}
		case 2:
			if aLittleUseless() {
				keyShow()
			}
		case 3:
			if reallyUseless() {
				drawTree()
			}
		case 0:
			return
		default:
			goto Label0
		}
	}

Label1:
	log.Fatalln(string(colorRed) + Cipher.Decrypt("\xf9\xd7\x5f\xc9\x47\xc2\xe8\xfd\xc9\xd7\x5e\xc8\x45\x86\xac\xa2\x95") + string(colorWhite))
Label0:
	fmt.Println()
	fmt.Println(string(colorRed) + Cipher.Decrypt("\xe8\xdc\x56\xd2\x4f\x95\xe2\xb8\xd0\xd7\x53\xc9\x00\x8d\xfc\xec\xd4\xdd\x53"))
	fmt.Println()
}
