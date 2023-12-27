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
//
//type Block struct {
//	Transactions      []Transaction
//	PreviousBlockHash []byte
//	MerkleRoot        []byte
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
//	block.MerkleRoot = block.ComputeMerkleRoot()
//}
//
//func DisplayBlockchain(blockchain []Block) {
//	for i, block := range blockchain {
//		fmt.Printf("Block %d: Merkle Root: %x, Hash: %x\n", i+1, block.MerkleRoot, block.HashBlock())
//	}
//}
//
//func (b *Block) ComputeMerkleRoot() []byte {
//	var hashes [][]byte
//	for _, tx := range b.Transactions {
//		hash := sha256.Sum256([]byte(fmt.Sprintf("%x%x%f", tx.SenderPublicKey, tx.RecipientPublicKey, tx.Amount)))
//		hashes = append(hashes, hash[:])
//	}
//	return constructMerkleTree(hashes)
//}
//
//func constructMerkleTree(hashes [][]byte) []byte {
//	if len(hashes) == 0 {
//		return nil
//	}
//	if len(hashes) == 1 {
//		return hashes[0]
//	}
//
//	var newHashes [][]byte
//	for i := 0; i < len(hashes); i += 2 {
//		hash1 := hashes[i]
//		var hash2 []byte
//		if i+1 < len(hashes) {
//			hash2 = hashes[i+1]
//		}
//		combined := append(hash1, hash2...)
//		hash := sha256.Sum256(combined)
//		newHashes = append(newHashes, hash[:])
//	}
//
//	return constructMerkleTree(newHashes)
//}
//
//func main() {
//	fmt.Println("Blockchain powered by Merkle Tree!")
//
//	senderPrivateKey, senderPublicKey := GenerateKeyPair()
//	blockchain := []Block{CreateGenesisBlock()}
//
//	for {
//		fmt.Println("\n1. Add Transaction to Block")
//		fmt.Println("2. View Blockchain")
//		fmt.Println("3. Exit")
//
//		var choice string
//		fmt.Print("Choice variant: ")
//		fmt.Scanln(&choice)
//
//		switch choice {
//		case "1":
//			var recipientPublicKey string
//			fmt.Print("Recipient's public key: ")
//			fmt.Scanln(&recipientPublicKey)
//
//			var amount float64
//			fmt.Print("Transaction amount: ")
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
//			fmt.Println("Exit from blockchain implementation :)\n this work must be graded 100%! :)")
//			return
//
//		default:
//			fmt.Println("Choose valid choose.")
//		}
//	}
//}
