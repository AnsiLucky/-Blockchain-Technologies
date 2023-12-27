package main

//
//import (
//	"crypto/sha256"
//	"encoding/hex"
//	"errors"
//	"fmt"
//	"html/template"
//	"log"
//	"math/rand"
//	"net/http"
//	"strconv"
//	"strings"
//	"time"
//)
//
//var tpl *template.Template
//
//func init() {
//	tpl = template.Must(template.ParseGlob("./templates/*.html"))
//}
//
//func keys(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		tpl.ExecuteTemplate(w, "keys.html", nil)
//	}
//}
//
//func main() {
//	private, err := getPrivateKey()
//	if err != nil {
//		log.Fatal("GG")
//	}
//	public, err := getPublicKeyByPrivateKey(private)
//	if err != nil {
//		log.Fatal("GG")
//	}
//
//	fmt.Println(private, public)
//
//	signature, err := getSignature("Ansi Sh", private)
//	if err != nil {
//		log.Fatal("GG")
//	}
//
//	fmt.Println(signature)
//	fmt.Println(verifySignature("Ansi Sh", signature, public))
//	// TODO: Blockchain Application Development
//
//	mux := http.NewServeMux()
//	mux.HandleFunc("/keys", keys)
//
//	http.ListenAndServe(":8080", mux)
//
//}
//
//func generateKeyPair() (string, string, error) {
//	privateKey, err := getPrivateKey()
//	if err != nil {
//		return "", "", errors.New("error to get private key")
//	}
//
//	publicKey, err := getPublicKeyByPrivateKey(privateKey)
//	if err != nil {
//		return "", "", errors.New("error to get public key")
//	}
//
//	return privateKey, publicKey, nil
//}
//
//func getPrivateKey() (string, error) {
//	rand.Seed(time.Now().UnixNano())
//	var letter, num int
//LOOP1:
//	for {
//		temp := rand.Intn(25) + 1
//		if temp != 13 {
//			letter = temp
//			break LOOP1
//		}
//	}
//LOOP2:
//	for {
//		temp := rand.Intn(9) + 1
//		if temp != 5 {
//			num = temp
//			break LOOP2
//		}
//	}
//
//	return fmt.Sprintf("%dx%d", letter, num), nil
//}
//
//func getPublicKeyByPrivateKey(privateKey string) (string, error) {
//	temp := strings.Split(privateKey, "x")
//	letter, err := strconv.Atoi(temp[0])
//	if err != nil {
//		return "", errors.New("Cant convert to int")
//	}
//	num, err := strconv.Atoi(temp[1])
//	if err != nil {
//		return "", errors.New("Cant convert to int")
//	}
//	return fmt.Sprintf("%vx%v", 26-letter, 10-num), nil
//}
//
//func getSignature(text, privateKey string) (string, error) {
//	hash := sha256.New()
//	hash.Write([]byte(text))
//	hashBytes := hash.Sum(nil)
//	hashString := hex.EncodeToString(hashBytes)
//	var signature []byte
//	fmt.Println(hashString)
//	ppk := strings.Split(privateKey, "x")
//	letter, _ := strconv.Atoi(ppk[0])
//	num, _ := strconv.Atoi(ppk[1])
//
//	for i := 0; i < len(hashString); i++ {
//		if hashString[i] >= 'a' && hashString[i] <= 'z' {
//			temp := int(hashString[i]) + letter
//			if temp > 'z' {
//				temp -= 26
//			}
//			signature = append(signature, byte(temp))
//		} else {
//			temp := int(hashString[i]) + num
//			if temp > '9' {
//				temp -= 10
//			}
//			signature = append(signature, byte(temp))
//		}
//	}
//
//	return string(signature), nil
//}
//
//func verifySignature(text, signature, publicKey string) (bool, error) {
//	hash := sha256.New()
//	hash.Write([]byte(text))
//	hashBytes := hash.Sum(nil)
//	hashString := hex.EncodeToString(hashBytes)
//
//	var hashSignature []byte
//
//	ppk := strings.Split(publicKey, "x")
//	letter, _ := strconv.Atoi(ppk[0])
//	num, _ := strconv.Atoi(ppk[1])
//
//	for i := 0; i < len(signature); i++ {
//		if signature[i] >= 'a' && signature[i] <= 'z' {
//			temp := int(signature[i]) + letter
//			if temp > 'z' {
//				temp -= 26
//			}
//			hashSignature = append(hashSignature, byte(temp))
//		} else {
//			temp := int(signature[i]) + num
//			if temp > '9' {
//				temp -= 10
//			}
//			hashSignature = append(hashSignature, byte(temp))
//		}
//	}
//	fmt.Println(string(hashSignature))
//
//	if hashString == string(hashSignature) {
//		return true, nil
//	}
//	return false, nil
//}
