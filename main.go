package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type CipherRequest struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type CipherResponse struct {
	Encrypted string `json:"encrypted"`
}

func vigenereEncrypt(plainText, key string) string {
	plainText = strings.ToUpper(plainText)
	key = strings.ToUpper(key)
	var result strings.Builder

	keyIndex := 0
	for i := 0; i < len(plainText); i++ {
		char := plainText[i]
		if char >= 'A' && char <= 'Z' {
			keyChar := key[keyIndex%len(key)]
			encryptedChar := ((char - 'A') + (keyChar - 'A')) % 26
			result.WriteByte('A' + encryptedChar)
			keyIndex++
		} else {
			result.WriteByte(char)
		}
	}
	return result.String()
}

func vigenereDecrypt(cipherText, key string) string {
	cipherText = strings.ToUpper(cipherText)
	key = strings.ToUpper(key)
	var result strings.Builder

	keyIndex := 0
	for i := 0; i < len(cipherText); i++ {
		char := cipherText[i]
		if char >= 'A' && char <= 'Z' {
			keyChar := key[keyIndex%len(key)]
			decryptedChar := ((char - keyChar + 26) % 26)
			result.WriteByte('A' + decryptedChar)
			keyIndex++
		} else {
			result.WriteByte(char)
		}
	}
	return result.String()
}

func saveToFile(input, output, operation string) {
	file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	entry := fmt.Sprintf(
		"WORD: %s\n%s : %s\n\n",
		input,
		operation,
		output,
	)

	// Add header if it's the first time (file is empty)
	info, _ := file.Stat()
	if info.Size() == 0 {
		header := "**********************VIGÈNERE CODE AND DECODE**********************\n\n"
		file.WriteString(header)
	}

	_, err = file.WriteString(entry)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func cipherHandler(w http.ResponseWriter, r *http.Request) {
	var req CipherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	encrypted := vigenereEncrypt(req.Text, req.Key)

	// 🔽 Call saveToFile here
	saveToFile(req.Text, encrypted, "ENCODE")

	res := CipherResponse{
		Encrypted: encrypted,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func decipherHandler(w http.ResponseWriter, r *http.Request) {
	var req CipherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	decrypted := vigenereDecrypt(req.Text, req.Key)

	// 🔽 Call saveToFile here
	saveToFile(req.Text, decrypted, "DECODE")

	res := CipherResponse{
		Encrypted: decrypted,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/vigenere", cipherHandler)
	http.HandleFunc("/vigenere/decode", decipherHandler)
	fmt.Println("API running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
