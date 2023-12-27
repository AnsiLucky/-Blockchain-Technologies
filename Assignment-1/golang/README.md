# Cryptography and Blockchain Implementation Examples on Golang #
## Review on files ##
- `signature.go` - implementation of public and private key
- `merkle-tree-blockchain.go` - blockchain implementation with RSA Encryption and Decryption
- `merkle-tree-blockchain(v2.0.0).go` - blockchain implementation with our own Cryptography Methods
<hr>

## signature.go ##

`getPrivateKey()` - generate the 'Private Key', which seems like 
`num1xnum2`, `num1`-letters encryption,  
`num2` - number encryption. `1 <= num1 <= 25`;`1 <= num2 <= 9`
<br><br>
`getPublicKeyByPrivateKey()` - returns the 'Public Key', via 'Private Key', which seems like
  `num1xnum2`, `num1`-letters decryption, `num2` - number decryption. `1 <= num1 <= 25`;`1 <= num2 <= 9`
  <br><br>
`getSignature(text, privateKey string) (string, error)` - get the hash
of text (by sha256 algorithm), after encrypt with privateKey and return 
signature in type `string`.<br><br>
`verifySignature(text, signature, publicKey string) (bool, error)` - 
decrypt the signature with 'Public Key' => check is `hashFromSignature == textHash`
and return the `boolean` value.

## merkle-tree-blockchain(v2.0.0).go ##
same with 'merkle-tree-blockchain(v2.0.0).go', but use the 
`RSA Cryptography Algorithms for Encryption/Decryption`.

## merkle-tree-blockchain(v2.0.0).go ##
#### Block struct ####
-	`Transactions      `
-	`PreviousBlockHash `
-	`MerkleRoot        `
-	`Timestamp         `
-	`Nonce             `
- `func (b *Block) ComputeMerkleRoot() []byte`
- `func constructMerkleTree(hashes [][]byte) []byte`
#### Transaction struct ####
- `SenderPublicKey    `
- `RecipientPublicKey `
- `Amount             `
<br><br>

`HashBlock()` - create the string by concatenation of 
`fmt.Sprintf("%x%x%f", tx.SenderPublicKey, tx.RecipientPublicKey, tx.Amount)`
=> create the string by concatenation of `fmt.Sprintf("%s%x%d%d", transactionStr, b.PreviousBlockHash, b.Timestamp.UnixNano(), b.Nonce)`
and return the hash of this string (by sha256 algorithm).

`GenerateKeyPair()` - generate the 'pair of key' by `getPrivateKey()` 
and `getPublicKeyByPrivateKey()` functions

`getPrivateKey()` - generate the 'Private Key', which seems like
`num1xnum2`, `num1`-letters encryption/decryption,  
`num2` - number encryption/decryption. <br>`1 <= num1 <= 25`;`1 <= num2 <= 9`

`getPublicKeyByPrivateKey()` - returns the 'Public Key', via 'Private Key', which seems like
`num1xnum2`, `num1`-letters encryption/decryption,  
`num2` - number encryption/decryption. <br>`1 <= num1 <= 25`;`1 <= num2 <= 9`

`SignMessage(privateKey string, message []byte) string` - get the hash
of text and after encrypt with privateKey.

`VerifySignature(text, signature, publicKey string) (bool, error)` -
decrypt the signature with 'Public Key' => check is `hashFromSignature == textHash`
and return the `boolean` value.

`CreateGenesisBlock() Block` - create genesis block with the previous 
hash `"Genesis Block"`.

`AddTransactionToBlock` - add Transaction and Renew the Merkle Tree Root.

`DisplayBlockchain(blockchain []Block)` - traverse all list and output.

## How to Run ##
1. Ensure you have Go installed on your system.
2. Clone the repository:
   `gh repo clone AnsiLucky/Blockchain-Technologies`
3. Navigate to the project directory:
   `cd Blockchain-Technologies`
4. Run the Go program:
   `go run <file-which-you-need>`


