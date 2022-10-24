package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/0xPolygon/polygon-edge/consensus/ibft/signer"
	"github.com/0xPolygon/polygon-edge/validators"

	hex "github.com/0xPolygon/polygon-edge/helper/hex"
)

func main() {
	extraData := os.Args[1]
	decoded := Decode(&extraData)
	printData(&decoded)
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

func printData(extra *signer.IstanbulExtra) {
	fmt.Printf("validators: %d:\taddress:bls\n", extra.Validators.Len())
	for i := 0; i < extra.Validators.Len(); i++ {
		validator := extra.Validators.At(uint64(i))
		fmt.Printf("validator: %v\n", validator)
	}
	fmt.Printf("\nProposerSeal: %v\n", extra.ProposerSeal)
	fmt.Printf("CommittedSeals: %v\n", extra.CommittedSeals.Num())
	fmt.Printf("ParentCommittedSeals: %v\n", extra.ParentCommittedSeals.Num())
}
