package main

import (
	"fmt"
	"encoding/hex"

	"edwards25519/ed25519"
	// "ed25519"
)


func main() {
	pk_A, sk_A, err_A := ed25519.GenerateKey(nil)
	pk_B, sk_B, err_B := ed25519.GenerateKey(nil)
	if err_A != nil || err_B != nil {
		fmt.Println(err_A)
		fmt.Println(err_B)
	} else {
		fmt.Printf("  %s %s\n", "pk_A:    ", hex.EncodeToString(pk_A[:]))
		fmt.Printf("  %s %s\n", "sk_A:    ", hex.EncodeToString(sk_A[:]))
		fmt.Printf("  %s %s\n", "pk_B:    ", hex.EncodeToString(pk_B[:]))
		fmt.Printf("  %s %s\n", "sk_B:    ", hex.EncodeToString(sk_B[:]))
	}

	mess := []byte("a message to be signed")

	// only one pk,sk
	sign_A := ed25519.Sign(sk_A, pk_A, mess)
	fmt.Printf("  %s %s\n", "sign_A:  ", hex.EncodeToString(sign_A[:]))

	if ed25519.Verify(pk_A, mess, sign_A) {
		fmt.Println("this is a valid signature!")
	}else{
		fmt.Println("this is an invalid signature!")
	}

	//joint pk,sk
	pk_arr := []ed25519.PublicKey{pk_A, pk_B}
	jointKey, primeKey, err := ed25519.GenerateJointKey(pk_arr)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("  %s %s\n", "jointKey:    ", hex.EncodeToString(jointKey[:]))
		fmt.Print("primeKey:")
		fmt.Println(primeKey)
	}
	jointPrivateKey_A, err := ed25519.GenerateJointPrivateKey(pk_arr, sk_A,0)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("  %s %s\n", "jointPrivateKey_A:    ", hex.EncodeToString(jointPrivateKey_A[:]))
		// fmt.Println(jointPrivateKey_A)
	}
	jointPrivateKey_B, err := ed25519.GenerateJointPrivateKey(pk_arr, sk_B,1)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Printf("  %s %s\n", "jointPrivateKey_B:    ", hex.EncodeToString(jointPrivateKey_B[:]))
		// fmt.Println(jointPrivateKey_B)
	}

	// H(R1 + R2 + ... + Rn || J(P1, P2, ..., Pn) || m) = e
	// si = ri + e * x'i
	// JointSign(privateKey, jointPrivateKey PrivateKey, noncePoints []CurvePoint, message []byte) []byte {}
	noncePoints_A := ed25519.GenerateNoncePoint(sk_A, mess)
	noncePoints_B := ed25519.GenerateNoncePoint(sk_B, mess)
	np := []ed25519.CurvePoint{noncePoints_A, noncePoints_B}
	jointSign_A := ed25519.JointSign(sk_A, jointPrivateKey_A, np, mess)
	fmt.Printf("  %s %s\n", "jointSign_A:    ", hex.EncodeToString(jointSign_A[:]))
	// fmt.Print("jointSign_A: ")
	// fmt.Println(hex.EncodeToString(jointSign_A[:]))

	jointSign_B := ed25519.JointSign(sk_B, jointPrivateKey_B, np, mess)
	// fmt.Print("jointSign_B: ")
	// fmt.Println(jointSign_B)
	fmt.Printf("  %s %s\n", "jointSign_B:    ", hex.EncodeToString(jointSign_B[:]))
	addsig := ed25519.AddSignature(jointSign_A, jointSign_B)
	// fmt.Print("addsig: ")
	// fmt.Println(addsig)
	fmt.Printf("  %s %s\n", "addsig:    ", hex.EncodeToString(addsig[:]))
	
	if ed25519.Verify(jointKey, mess, addsig) {
		fmt.Println("this is a valid signature!")
	}else{
		fmt.Println("this is an invalid signature!")
	}

	//AdaptorSig
	
}
