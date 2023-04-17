package main

import (
	"net"
	"log"
	"os"
)

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:; -")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet)
		newIndex := (indeks + chiffer) % len(alphabet)
		kryptertMelding[i] = alphabet[newIndex]
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}

func main() {
	conn, err := net.Dial("tcp", "172.17.0.4:37867")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])
kryptertMelding := Krypter([]rune(os.Args[1]), ALF_SEM03, 4)
log.Println("Kryptert melding: ", string(kryptertMelding))
_, err = conn.Write([]byte(string(kryptertMelding)))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
dekryptertMelding := string(Krypter([]rune(string(buf[:n])), ALF_SEM03, len(ALF_SEM03)-4))

	log.Printf("reply from proxy: %s", dekryptertMelding)
}

