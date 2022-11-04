package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/0xPolygon/polygon-edge/consensus/ibft/signer"
	"github.com/0xPolygon/polygon-edge/helper/hex"
	"github.com/0xPolygon/polygon-edge/validators"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid arg length!")
		return
	}
	extraData := os.Args[1]
	blockHash := os.Args[2]
	decoded := Decode(&extraData)
	address := recoverAddress(&decoded, &blockHash)
	printData(&decoded, &address)
}

func Decode(extraData *string) signer.IstanbulExtra {
	extra := &signer.IstanbulExtra{
		Validators:           validators.NewBLSValidatorSet(),
		ProposerSeal:         []byte{},
		CommittedSeals:       &signer.AggregatedSeal{},
		ParentCommittedSeals: &signer.AggregatedSeal{},
	}
	//remove 0x and leading 0s and convert to byte array
	dataStr := strings.TrimLeft(strings.TrimPrefix(*extraData, "0x"), "0")
	byteData, err := hex.DecodeHex(dataStr)
	if err != nil {
		return *extra
	}

	//Unmarshal byte array into IstanbulExtra struct
	if err := extra.UnmarshalRLP(byteData); err != nil {
		return *extra
	}
	return *extra
}

func recoverAddress(decoded *signer.IstanbulExtra, blockHash *string) string {
	//string to byte array
	blockHashByteData, err := hex.DecodeHex(*blockHash)
	if err != nil {
		panic(err)
	}
	//get ecdsa pubkey
	sigPublicKeyECDSA, err := crypto.SigToPub(crypto.Keccak256(blockHashByteData), decoded.ProposerSeal)
	if err != nil {
		panic(err)
	}
	//convert to readable string
	return crypto.PubkeyToAddress(*sigPublicKeyECDSA).String()
}

func printData(extra *signer.IstanbulExtra, address *string) {
	fmt.Printf("\nvalidator set: %d\taddress:bls\n", extra.Validators.Len())
	for i := 0; i < extra.Validators.Len(); i++ {
		validator := extra.Validators.At(uint64(i))
		fmt.Printf("%v\n", validator)
	}

	fmt.Printf("\nBlock Proposer: %v", *address)
	fmt.Printf("\nProposerSeal: %v\n", hex.EncodeToHex(extra.ProposerSeal))
	fmt.Printf("CommittedSeals: %v\n", extra.CommittedSeals.Num())
	fmt.Printf("ParentCommittedSeals: %v\n", extra.ParentCommittedSeals.Num())

}
