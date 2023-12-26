package main

//
//import (
//	"crypto"
//	"crypto/rand"
//	"crypto/rsa"
//	"crypto/sha256"
//	"fmt"
//	"time"
//)
//
//type Block struct {
//	Transactions      []Transaction
//	PreviousBlockHash []byte
//	Timestamp         time.Time
//	Nonce             int
//}
//
//type Transaction struct {
//	SenderPublicKey    []byte
//	RecipientPublicKey []byte
//	Amount             float64
//}
//
//func (b *Block) HashBlock() []byte {
//	var transactionStr string
//	for _, tx := range b.Transactions {
//		transactionStr += fmt.Sprintf("%x%x%f", tx.SenderPublicKey, tx.RecipientPublicKey, tx.Amount)
//	}
//	blockStr := fmt.Sprintf("%s%x%d%d", transactionStr, b.PreviousBlockHash, b.Timestamp.UnixNano(), b.Nonce)
//	hash := sha256.Sum256([]byte(blockStr))
//	return hash[:]
//}
//
//func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		panic(err)
//	}
//	return privateKey, &privateKey.PublicKey
//}
//
//func SignMessage(privateKey *rsa.PrivateKey, message []byte) []byte {
//	hashed := sha256.Sum256(message)
//	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], nil)
//	if err != nil {
//		panic(err)
//	}
//	return signature
//}
//
//func VerifySignature(publicKey *rsa.PublicKey, message, signature []byte) bool {
//	hashed := sha256.Sum256(message)
//	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signature, nil)
//	if err != nil {
//		fmt.Println("Signature verification failed:", err)
//		return false
//	}
//	return true
//}
//
//func CreateGenesisBlock() Block {
//	return Block{
//		Transactions:      []Transaction{{}},
//		PreviousBlockHash: []byte("Genesis Block"),
//		Timestamp:         time.Now(),
//		Nonce:             0,
//	}
//}
//
//func AddTransactionToBlock(block *Block, senderPrivateKey *rsa.PrivateKey, recipientPublicKey *rsa.PublicKey, amount float64) {
//	transaction := Transaction{
//		SenderPublicKey:    senderPrivateKey.PublicKey.N.Bytes(),
//		RecipientPublicKey: recipientPublicKey.N.Bytes(),
//		Amount:             amount,
//	}
//	signature := SignMessage(senderPrivateKey, []byte(fmt.Sprintf("%x%x%f", transaction.SenderPublicKey, transaction.RecipientPublicKey, transaction.Amount)))
//
//	if !VerifySignature(recipientPublicKey, []byte(fmt.Sprintf("%x%x%f", transaction.SenderPublicKey, transaction.RecipientPublicKey, transaction.Amount)), signature) {
//		fmt.Println("Transaction signature verification failed.")
//		return
//	}
//
//	block.Transactions = append(block.Transactions, transaction)
//}
//
//func DisplayBlockchain(blockchain []Block) {
//	for i, block := range blockchain {
//		fmt.Printf("Block %d: %x\n", i+1, block.HashBlock())
//	}
//}
//
//func main() {
//	fmt.Println("Welcome to the Blockchain CLI!")
//
//	senderPrivateKey, senderPublicKey := GenerateKeyPair()
//	blockchain := []Block{CreateGenesisBlock()}
//
//	for {
//		fmt.Println("\n1. Add Transaction")
//		fmt.Println("2. View Blockchain")
//		fmt.Println("3. Exit")
//
//		var choice string
//		fmt.Print("Enter your choice: ")
//		fmt.Scanln(&choice)
//
//		switch choice {
//		case "1":
//			var recipientPublicKey string
//			fmt.Print("Enter recipient's public key: ")
//			fmt.Scanln(&recipientPublicKey)
//
//			var amount float64
//			fmt.Print("Enter transaction amount: ")
//			fmt.Scanln(&amount)
//
//			newBlock := Block{
//				Transactions:      []Transaction{},
//				PreviousBlockHash: blockchain[len(blockchain)-1].HashBlock(),
//				Timestamp:         time.Now(),
//				Nonce:             0,
//			}
//			AddTransactionToBlock(&newBlock, senderPrivateKey, senderPublicKey, amount)
//			blockchain = append(blockchain, newBlock)
//			fmt.Println("Transaction added successfully!")
//
//		case "2":
//			DisplayBlockchain(blockchain)
//
//		case "3":
//			fmt.Println("Exiting the Blockchain CLI. Goodbye!")
//			return
//
//		default:
//			fmt.Println("Invalid choice. Please enter a valid option.")
//		}
//	}
//}
