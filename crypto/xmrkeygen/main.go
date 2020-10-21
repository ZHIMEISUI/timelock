package main


import (
	// "encoding/binary"
	"encoding/hex"
	"fmt"
	// "hash/crc32"
	// "math/rand"
	// "strings"
	"time"
	// "unsafe"
	"math/big"

	// "edwards25519/ed25519"
	"github.com/moneroutil"
	// "golang.org/x/crypto/sha3"
	"github.com/xlab-si/emmy/crypto/schnorr"
	"github.com/xlab-si/emmy/crypto/common"
)

func main() {
	privKeyA, pubKeyA := moneroutil.NewKeyPair()
	privKeyB, pubKeyB := moneroutil.NewKeyPair()
	fmt.Printf("  %s %s\n", "privKeyA:     	", hex.EncodeToString(privKeyA[:]))
	fmt.Printf("  %s %s\n", "pubKeyA:     		", hex.EncodeToString(pubKeyA[:]))
	fmt.Printf("  %s %s\n", "privKeyB:     	", hex.EncodeToString(privKeyB[:]))
	fmt.Printf("  %s %s\n", "pubKeyB:     		", hex.EncodeToString(pubKeyB[:]))

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("-----------------------split line-----------------------------")
	fmt.Println("--------------------------------------------------------------")
	t1 := time.Now()
	elapsed := time.Since(t1)
	// numTries := 1
	numMixins := 10
	// for i := 0; i < numTries; i++ {
	// 	hash := moneroutil.Hash(*moneroutil.RandomScalar())
		
	// 	mixins := make([]moneroutil.Key, numMixins)
	// 	for j := 0; j < numMixins; j++ {
	// 		mixins[j] = *moneroutil.RandomPubKey()
	// 	}
	// 	t1 = time.Now()
	// 	keyImage, pubKeys, sig := moneroutil.CreateSignature(&hash, mixins, privKeyA)
	// 	elapsed = time.Since(t1)
	// 	fmt.Println("  original ringsig creation elapsed: ", elapsed)
	// 	fmt.Printf("  %s %s\n", "keyImage:     	", hex.EncodeToString(keyImage[:]))
	// 	for _, pub := range pubKeys{
	// 		fmt.Printf("  %s %s\n", "pubKeys:  		", hex.EncodeToString(pub[:]))
	// 	}

	// 	t1 = time.Now()
	// 	if !moneroutil.VerifySignature(&hash, &keyImage, pubKeys, sig) {
	// 		elapsed = time.Since(t1)
	// 		fmt.Println("  original ringsig verification elapsed: ", elapsed)
	// 		var pubKeyStr string
	// 		for _, pk := range pubKeys {
	// 			pubKeyStr += fmt.Sprintf("%x ", pk)
	// 		}
	// 		fmt.Printf("%d: failed on verify: %x %x %s%x", i, hash, keyImage, pubKeyStr, sig.Serialize())
	// 	} else {
	// 		elapsed = time.Since(t1)
	// 		fmt.Println("  original ringsig verification elapsed: ", elapsed)
	// 		fmt.Println("valid sig!!! XD")
	// 	}

	// }
	

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("---------------------test 2of2sig-----------------------------")
	fmt.Println("--------------------------------------------------------------")
	
	numMixins = 10
	// keyImagePrimes := make([]moneroutil.Key, 2)
	// pubKeys := make([]moneroutil.Key, 2)
	// pubKeys[0] = *pubKeyA
	// pubKeys[1] = *pubKeyB
	mixins := make([]moneroutil.Key, numMixins)
	for j := 0; j < numMixins; j++ {
		mixins[j] = *moneroutil.RandomPubKey()
	}

	hash := moneroutil.Hash(*moneroutil.RandomScalar())
	// jointKey, primeKeys, _ := moneroutil.GenerateJointPubKey(pubKeys)
	// jointPubKey := new(moneroutil.Key)
	// moneroutil.ScAdd(jointPubKey, pubKeyA, pubKeyB)
	// fmt.Printf("  %s %s\n", "jointPubKey:   	", hex.EncodeToString(jointPubKey[:]))

	// fmt.Println("-----------------------user A------------------------------")
	// // compute for user A
	// // var middlenonce_a moneroutil.Key
	// r_a := moneroutil.RandomScalar()
	// t_a := moneroutil.RandomScalar()

	// R_A, T_A, LiA, KmultPKA := moneroutil.PreCompute(r_a, t_a, jointPubKey)
	// fmt.Printf("  %s %s\n", "RA:     		", hex.EncodeToString(R_A[:]))
	// fmt.Printf("  %s %s\n", "TA:     		", hex.EncodeToString(T_A[:]))
	// // fmt.Printf("  %s %s\n", "LiprimeA:   		", hex.EncodeToString(Liprime_A[:]))

	// keyImageA := moneroutil.GenerateKeyImagePrime(*privKeyA, *jointPubKey)
	// keyImagePrimes[0] = keyImageA
	// fmt.Printf("  %s %s\n", "keyImageA:   		", hex.EncodeToString(keyImageA[:]))

	// // jointPrivateKeyA, _ := moneroutil.GenerateJointPrivateKey(pubKeys, *privKeyA, 0)
	// // fmt.Printf("  %s %s\n", "jointPrivateKeyA:   	", hex.EncodeToString(jointPrivateKeyA[:]))


	// fmt.Println("-----------------------user B------------------------------")
	// // compute for user B
	// // var middlenonce_b moneroutil.Key
	// r_b := moneroutil.RandomScalar()
	// t_b := moneroutil.RandomScalar()

	// R_B, T_B, LiB, KmultPKB := moneroutil.PreCompute(r_b, t_b, jointPubKey)
	// fmt.Printf("  %s %s\n", "RB:     		", hex.EncodeToString(R_B[:]))
	// fmt.Printf("  %s %s\n", "TB:     		", hex.EncodeToString(T_B[:]))
	// // fmt.Printf("  %s %s\n", "LiprimeB:   		", hex.EncodeToString(Liprime_B[:]))

	// keyImageB := moneroutil.GenerateKeyImagePrime(*privKeyB, *jointPubKey)
	// keyImagePrimes[1] = keyImageB
	// fmt.Printf("  %s %s\n", "keyImageB:   		", hex.EncodeToString(keyImageB[:]))

	// // jointPrivateKeyB, _ := moneroutil.GenerateJointPrivateKey(pubKeys, *privKeyB, 0)
	// // fmt.Printf("  %s %s\n", "jointPrivateKeyB:   	", hex.EncodeToString(jointPrivateKeyB[:]))
	
	
	// fmt.Println("---------------------common parameters-----------------------")

	// randomscalar := make([]moneroutil.Key, 22)
	// for i:=0; i<22; i++{
	// 	randomscalar[i] = *moneroutil.RandomScalar()
	// }
	// // for i,primek := range primeKeys{
	// // 	fmt.Print(i)
	// // 	fmt.Printf("  %s %s\n", "primek:   		", hex.EncodeToString(primek[:]))
	// // }

	// // fmt.Printf("  %s %s\n", "jointPubKey:   	", hex.EncodeToString(jointKey[:]))

	// var q_a, q_b, Li, KmultPK moneroutil.Key
	// moneroutil.ScAdd(&q_a, r_a, t_a)
	// moneroutil.ScAdd(&q_b, r_b, t_b)
	// moneroutil.ScAdd(&KmultPK, &KmultPKA, &KmultPKB)
	// moneroutil.ScAdd(&Li, &LiA, &LiB)


	// t1 = time.Now()
	// c_a, keyImage, sigA, mixpubkeys, incompleteSigA, privIndex, err := moneroutil.PreCompute2of2RingSig(&hash, mixins, keyImagePrimes, randomscalar, privKeyA, jointPubKey, &q_a, &Li, &KmultPK) 
	// adptorSigA := moneroutil.CreateAdptorSig(*t_a, sigA)
	// elapsed = time.Since(t1)
	// fmt.Println("  adaptor signature creation elapsed: ", elapsed)

	// c_b, keyImage, sigB, mixpubkeys, _, privIndex, err := moneroutil.PreCompute2of2RingSig(&hash, mixins, keyImagePrimes, randomscalar, privKeyB, jointPubKey, &q_b, &Li, &KmultPK) 
	// fmt.Printf("  %s %s\n", "keyImage:   		", hex.EncodeToString(keyImage[:]))
	// fmt.Printf("  %s %s\n", "c_a:   		", hex.EncodeToString(c_a[:]))
	// fmt.Printf("  %s %s\n", "c_b:   		", hex.EncodeToString(c_b[:]))
	// for i,mix := range mixpubkeys{
	// 	fmt.Print(i)
	// 	fmt.Printf("  %s %s\n", "pubkeys:   		", hex.EncodeToString(mix[:]))
	// }

	// // fmt.Println(incompleteSigA[1])
	// // fmt.Println(incompleteSigB[1])
	// if err != nil{
	// 	fmt.Print("err: ")
	// 	fmt.Println(err)
	// }
	
	// fmt.Printf("  %s %s\n", "creating partial signature SigA for 2of2Sig	", hex.EncodeToString(sigA[:]))
	// fmt.Printf("  %s %s\n", "creating partial signature SigB for 2of2Sig	", hex.EncodeToString(sigB[:]))

	// fmt.Printf("  %s %s\n", "Creating adptor signature A: ", hex.EncodeToString(sigA[:]))
	// fmt.Print("  verify adptor signature A: ")
	// t1 = time.Now()
	// if moneroutil.VerifyAdptorSig(adptorSigA, *pubKeyA, c_a, R_A){
	// 	elapsed = time.Since(t1)
	// 	fmt.Println("  adaptor signature verification elapsed: ", elapsed)
	// 	fmt.Println("  valid adptorSigA!!! XD")
	// } else {
	// 	elapsed = time.Since(t1)
	// 	fmt.Println("  adaptor signature verification elapsed: ", elapsed)
	// 	fmt.Println("  invalid adptorSigA.... -_-|||")
	// }

	// if moneroutil.VerifyAdptorSig(sigA, *pubKeyA, c_a, LiA){
	// 	fmt.Println("  valid SigA!!! XD")
	// } else {
	// 	fmt.Println("  invalid SigA.... -_-|||")
	// }

	// adptorSigB := moneroutil.CreateAdptorSig(*t_b, sigB)
	// fmt.Printf("  %s %s\n", "Creating adptor signature B: ", hex.EncodeToString(sigB[:]))
	// fmt.Print("  verify adptor signature B: ")
	// if moneroutil.VerifyAdptorSig(adptorSigB, *pubKeyB, c_b, R_B){
	// 	fmt.Println("  valid adptorSigB!!! XD")
	// } else {
	// 	fmt.Println("  invalid adptorSigB.... -_-|||")
	// }


	// sigPrimes := make([]moneroutil.Key, 2)
	// sigPrimes[0] = sigA
	// sigPrimes[1] = sigB
	// ringSig, sig := moneroutil.Create2of2Signature(sigPrimes, incompleteSigA, privIndex)
	// fmt.Println()

	// t1 = time.Now()
	// if moneroutil.VerifySignature(&hash, &keyImage, mixpubkeys, ringSig) {
	// 	elapsed = time.Since(t1)
	// 	fmt.Println("  2of2 ring signature verification elapsed: ", elapsed)
	// 	fmt.Println("  valid ringSig!!! XD")
	// } else {
	// 	elapsed = time.Since(t1)
	// 	fmt.Println("  2of2 ring signature verification elapsed: ", elapsed)
	// 	fmt.Println("  invalid ringSig.... -_-|||")
	// }

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("-------------------- test 2of2 ringsignature ------------------")
	fmt.Println("---------------------------------------------------------------")
	randomscalar := make([]moneroutil.Key, 22)
	for i:=0; i<22; i++{
		randomscalar[i] = *moneroutil.RandomScalar()
	}
	// privKeyA, pubKeyA
	// privKeyB, pubKeyB
	var jointPublicKey, Li, KmultPK, keyImage moneroutil.Key
	var KeyImageA, KmultPKA, R_A, LiA, k_a moneroutil.Key
	var KeyImageB, KmultPKB, R_B, LiB, k_b moneroutil.Key
	moneroutil.KeyAdd(&jointPublicKey, pubKeyA, pubKeyB)
	
	r_a := moneroutil.RandomScalar()
	t_a := moneroutil.RandomScalar()
	// moneroutil.KeyAdd(&k_a, r_a, t_a)
	moneroutil.ScAdd(&k_a, r_a, t_a)
	r_b := moneroutil.RandomScalar()
	t_b := moneroutil.RandomScalar()
	// moneroutil.KeyAdd(&k_b, r_b, t_b)
	moneroutil.ScAdd(&k_b, r_b, t_b)
	
	R_A, _, LiA, KmultPKA = moneroutil.PreCompute(*r_a, *t_a, jointPublicKey)
	KeyImageA = moneroutil.GenerateKeyImagePrime(*privKeyA, jointPublicKey)

	R_B, _, LiB, KmultPKB = moneroutil.PreCompute(*r_b, *t_b, jointPublicKey)
	KeyImageB = moneroutil.GenerateKeyImagePrime(*privKeyB, jointPublicKey)


	moneroutil.KeyAdd(&Li, &LiA, &LiB)
	moneroutil.KeyAdd(&KmultPK, &KmultPKA, &KmultPKB)
	moneroutil.KeyAdd(&keyImage, &KeyImageA, &KeyImageB)
	// func PreCompute2of2RingSig(prefixHash *Hash, mixins, randomscalar []Key, keyImage, privKeyPrime, jointPubKey, k, Li, KmultPK *Key) (c, sigPrime Key, pubKeys []Key, sig RingSignature, privIndex int, err error) {
	var c, sigA moneroutil.Key 
	var pubKeys []moneroutil.Key
	var ringsig moneroutil.RingSignature
	var privIndex int
	testNums := 1000
	t1 = time.Now()
	for i := 0; i < testNums; i++{
		c, sigA, pubKeys, ringsig, privIndex, _ = moneroutil.PreCompute2of2RingSig(&hash, mixins, randomscalar, &keyImage, privKeyA, &jointPublicKey, &k_a, &Li, &KmultPK) 
	}
	elapsed = time.Since(t1)
	fmt.Println("  incomplete ringsig creation: ", elapsed)

	c, sigB, pubKeys, _, privIndex, err := moneroutil.PreCompute2of2RingSig(&hash, mixins, randomscalar, &keyImage, privKeyB, &jointPublicKey, &k_b, &Li, &KmultPK) 


	adptorSigA := moneroutil.CreateAdptorSig(*t_a, sigA)
	if moneroutil.VerifyAdptorSig(adptorSigA, *pubKeyA, c, R_A){
		fmt.Println("  valid adptorSigA!!! XD")
	} else {
		fmt.Println("  invalid adptorSigA.... -_-|||")
	}
	t1 = time.Now()
	for i := 0; i < testNums; i++{
		moneroutil.VerifyAdptorSig(adptorSigA, *pubKeyA, c, R_A)
	}
	elapsed = time.Since(t1)
	fmt.Println("  adaptor signature verification elapsed: ", elapsed)

	adptorSigB := moneroutil.CreateAdptorSig(*t_b, sigB)
	if moneroutil.VerifyAdptorSig(adptorSigB, *pubKeyB, c, R_B){
		fmt.Println("  valid adptorSigB!!! XD")
	} else {
		fmt.Println("  invalid adptorSigB.... -_-|||")
	}


	// var t, R, jointsig moneroutil.Key
	// moneroutil.KeyAdd(&jointsig, &sigA, &sigB)
	// moneroutil.ScAdd(&jointsig, &sigA, &sigB)
	// moneroutil.ScAdd(&t, t_a, t_b)
	// adptorSig := moneroutil.CreateAdptorSig(t, jointsig)
	// moneroutil.KeyAdd(&R, &R_A, &R_B)
	// if moneroutil.VerifyAdptorSig(adptorSig, jointPublicKey, c_b, R){
	// 	fmt.Println("  valid jointAdaptorsig!!! XD")
	// } else {
	// 	fmt.Println("  invalid jointAdaptorsig.... -_-|||")
	// }

	// if moneroutil.VerifyAdptorSig(jointsig, jointPublicKey, c_b, Li){
	// 	fmt.Println("  valid jointsig!!! XD")
	// } else {
	// 	fmt.Println("  invalid jointsig.... -_-|||")
	// }

	sigPrimes := make([]moneroutil.Key, 2)
	sigPrimes[0]=sigA
	sigPrimes[1]=sigB
	// var js *moneroutil.Key
	ringSig, _ := moneroutil.Create2of2Signature(sigPrimes, ringsig, privIndex)
	if moneroutil.VerifySignature(&hash, &keyImage, pubKeys, ringSig) {
		fmt.Println("  valid 2of2ringSig!!! XD")
	} else {
		fmt.Println("  invalid 2of2ringSig.... -_-|||")
	}
	
	t1 = time.Now()
	for i := 0; i < testNums; i++{
		moneroutil.VerifySignature(&hash, &keyImage, pubKeys, ringSig)
	}
	elapsed = time.Since(t1)
	fmt.Println("  2of2 ring signature verification elapsed: ", elapsed)

	// moneroutil.ScAdd(&jointPrivateKey, privKeyA, privKeyB)
	// moneroutil.ScAdd(&ktest, &k_a, &k_b)
	// ki, pks, sig := moneroutil.CreateSignature(&hash, mixins, randomscalar, &ktest, &jointPrivateKey)
	// if moneroutil.VerifySignature(&hash, &ki, pks, sig) {
	// 	fmt.Println("  valid ringSig!!! XD")
	// } else {
	// 	fmt.Println("  invalid ringSig.... -_-|||")
	// }


	fmt.Println("---------------------------------------------------------------")
	fmt.Println("------------------------ ZKP verification ---------------------")
	fmt.Println("---------------------------------------------------------------")

	group, err := schnorr.NewGroup(256)
	if err != nil {
		fmt.Println("error when creating Schnorr group: %v", err)
	}

	var bases [3]*big.Int
	for i := 0; i < len(bases); i++ {
		r := common.GetRandomInt(group.Q)
		bases[i] = group.Exp(group.G, r)
	}

	var secrets [3]*big.Int
	for i := 0; i < 3; i++ {
		secrets[i] = common.GetRandomInt(group.Q)
	}

	// y = g_1^x_1 * ... * g_k^x_k where g_i are bases and x_i are secrets
	y := big.NewInt(1)
	for i := 0; i < 3; i++ {
		f := group.Exp(bases[i], secrets[i])
		y = group.Mul(y, f)
	}

	prover, err := schnorr.NewProver(group, secrets[:], bases[:], y)
	if err != nil {
		fmt.Println("  error when creating Prover!")
	}
	verifier := schnorr.NewVerifier(group)

	proofRandomData := prover.GetProofRandomData()
	verifier.SetProofRandomData(proofRandomData, bases[:], y)

	challenge := verifier.GetChallenge()
	proofData := prover.GetProofData(challenge)
	t1 = time.Now()
	for i := 0; i < testNums; i++{
		verifier.Verify(proofData)
		// if verifier.Verify(proofData){
		// 	elapsed = time.Since(t1)
		// 	fmt.Println("  ZKP Verification: ", elapsed)
		// 	fmt.Println("  Valid ZKP! XD")
		// }else{
		// 	fmt.Println("  Invalid ZKP...! -_-|||")
		// }
	}
	elapsed = time.Since(t1)
	fmt.Println("  ZKP Verification: ", elapsed)

	// fmt.Println("---------------------------------------------------------------")
	// fmt.Println("--------------------------test ed25519-------------------------")
	// fmt.Println("---------------------------------------------------------------")
	// pk_A, sk_A, err_A := ed25519.GenerateKey(nil)
	// pk_B, sk_B, err_B := ed25519.GenerateKey(nil)
	// if err_A != nil || err_B != nil {
	// 	fmt.Println(err_A)
	// 	fmt.Println(err_B)
	// } else {
	// 	fmt.Printf("  %s %s\n", "pk_A:    		", hex.EncodeToString(pk_A[:]))
	// 	fmt.Printf("  %s %s\n", "sk_A:    		", hex.EncodeToString(sk_A[:]))
	// 	fmt.Printf("  %s %s\n", "pk_B:    		", hex.EncodeToString(pk_B[:]))
	// 	fmt.Printf("  %s %s\n", "sk_B:    		", hex.EncodeToString(sk_B[:]))
	// }

	// mess := []byte("TESTMESSAGE")

	// // only one pk,sk
	// sign_A := ed25519.Sign(sk_A, pk_A, mess)
	// fmt.Printf("  %s %s\n", "sign_A:  		", hex.EncodeToString(sign_A[:]))

	// if ed25519.Verify(pk_A, mess, sign_A) {
	// 	fmt.Println("this is a valid signature!")
	// }else{
	// 	fmt.Println("this is an invalid signature!")
	// }

	// fmt.Println("---------------------------------------------------------------")
	// fmt.Println("---------------------------------------------------------------")
	// fmt.Println("---------------------------------------------------------------")


	// // mess := []byte("TESTMESSAGE")
	// privKeyA, pubKeyA, err := moneroutil.GenerateKey()
	// privKeyB, pubKeyB, err := moneroutil.GenerateKey()

	// // privKeyA, pubKeyA := moneroutil.NewKeyPair()
	// // privKeyB, pubKeyB := moneroutil.NewKeyPair()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println()
	// fmt.Printf("  %s %s\n", "privKeyA:  		", hex.EncodeToString(privKeyA[:]))
	// fmt.Printf("  %s %s\n", "privKeyB:  		", hex.EncodeToString(privKeyB[:]))
	// fmt.Printf("  %s %s\n", "pubKeyA:  		", hex.EncodeToString(pubKeyA[:]))
	// fmt.Printf("  %s %s\n", "pubKeyB:  		", hex.EncodeToString(pubKeyB[:]))

	// fmt.Println("---------------------test sig-----------------------------")
	
	// sigsigle := moneroutil.Sign(privKeyA, pubKeyA, mess)
	// fmt.Printf("  %s %s\n", "sigsigle:  		", hex.EncodeToString(sigsigle[:]))
	// if moneroutil.Verify(pubKeyA, mess, sigsigle) {
	// 	fmt.Println("valid sig!!! XD")
	// } else {
	// 	fmt.Println("invalid sig.... -_-|||")
	// }

	// fmt.Println("---------------------test 2of2sig-------------------------")
	
	// pubKeys := make([]moneroutil.Key, 2)
	// pubKeys[0] = pubKeyA
	// pubKeys[1] = pubKeyB
	// jointPubKey, primeKeys, err := moneroutil.GenerateJointPubKey(pubKeys)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("  %s %s\n", "jointPubKey:  	", hex.EncodeToString(jointPubKey[:]))
	// for _,primeKey := range primeKeys{
	// 	fmt.Printf("  %s %s\n", "primeKeys:  		", hex.EncodeToString(primeKey[:]))
	// }
	// // fmt.Printf("  %s %s\n", "error:  			", err)
	// jointPrivateKeyA, _ := moneroutil.GenerateJointPrivateKey(pubKeys, privKeyA, 0)
	// fmt.Printf("  %s %s\n", "jointPrivateKeyA:  	", hex.EncodeToString(jointPrivateKeyA[:]))
	// jointPrivateKeyB, _ := moneroutil.GenerateJointPrivateKey(pubKeys, privKeyB, 1)
	// fmt.Printf("  %s %s\n", "jointPrivateKeyB:  	", hex.EncodeToString(jointPrivateKeyB[:]))
	

	// noncePoints_A := moneroutil.GenerateNoncePoint(privKeyA, mess)
	// noncePoints_B := moneroutil.GenerateNoncePoint(privKeyB, mess)
	// np := []moneroutil.CurvePoint{noncePoints_A, noncePoints_B}
	// SigA := moneroutil.JointSign(privKeyA, jointPrivateKeyA, jointPubKey, np, mess)
	// fmt.Printf("  %s %s\n", "SigA:  		", hex.EncodeToString(SigA[:]))

	// SigB := moneroutil.JointSign(privKeyB, jointPrivateKeyB, jointPubKey, np, mess)
	// fmt.Printf("  %s %s\n", "SigB:  		", hex.EncodeToString(SigB[:]))

	// Sig_add := moneroutil.AddSignature(SigA, SigB)
	// fmt.Printf("  %s %s\n", "Sig_add:  		", hex.EncodeToString(Sig_add[:]))
	// if moneroutil.Verify(jointPubKey, mess, Sig_add) {
	// 	fmt.Println("valid sig!!! XD")
	// } else {
	// 	fmt.Println("invalid sig.... -_-|||")
	// }

}

// English Mnemonic: https://github.com/monero-project/monero/blob/master/src/mnemonics/english.h
var words = &[...]string{
	"abbey",
	"abducts",
	"ability",
	"ablaze",
	"abnormal",
	"abort",
	"abrasive",
	"absorb",
	"abyss",
	"academy",
	"aces",
	"aching",
	"acidic",
	"acoustic",
	"acquire",
	"across",
	"actress",
	"acumen",
	"adapt",
	"addicted",
	"adept",
	"adhesive",
	"adjust",
	"adopt",
	"adrenalin",
	"adult",
	"adventure",
	"aerial",
	"afar",
	"affair",
	"afield",
	"afloat",
	"afoot",
	"afraid",
	"after",
	"against",
	"agenda",
	"aggravate",
	"agile",
	"aglow",
	"agnostic",
	"agony",
	"agreed",
	"ahead",
	"aided",
	"ailments",
	"aimless",
	"airport",
	"aisle",
	"ajar",
	"akin",
	"alarms",
	"album",
	"alchemy",
	"alerts",
	"algebra",
	"alkaline",
	"alley",
	"almost",
	"aloof",
	"alpine",
	"already",
	"also",
	"altitude",
	"alumni",
	"always",
	"amaze",
	"ambush",
	"amended",
	"amidst",
	"ammo",
	"amnesty",
	"among",
	"amply",
	"amused",
	"anchor",
	"android",
	"anecdote",
	"angled",
	"ankle",
	"annoyed",
	"answers",
	"antics",
	"anvil",
	"anxiety",
	"anybody",
	"apart",
	"apex",
	"aphid",
	"aplomb",
	"apology",
	"apply",
	"apricot",
	"aptitude",
	"aquarium",
	"arbitrary",
	"archer",
	"ardent",
	"arena",
	"argue",
	"arises",
	"army",
	"around",
	"arrow",
	"arsenic",
	"artistic",
	"ascend",
	"ashtray",
	"aside",
	"asked",
	"asleep",
	"aspire",
	"assorted",
	"asylum",
	"athlete",
	"atlas",
	"atom",
	"atrium",
	"attire",
	"auburn",
	"auctions",
	"audio",
	"august",
	"aunt",
	"austere",
	"autumn",
	"avatar",
	"avidly",
	"avoid",
	"awakened",
	"awesome",
	"awful",
	"awkward",
	"awning",
	"awoken",
	"axes",
	"axis",
	"axle",
	"aztec",
	"azure",
	"baby",
	"bacon",
	"badge",
	"baffles",
	"bagpipe",
	"bailed",
	"bakery",
	"balding",
	"bamboo",
	"banjo",
	"baptism",
	"basin",
	"batch",
	"bawled",
	"bays",
	"because",
	"beer",
	"befit",
	"begun",
	"behind",
	"being",
	"below",
	"bemused",
	"benches",
	"berries",
	"bested",
	"betting",
	"bevel",
	"beware",
	"beyond",
	"bias",
	"bicycle",
	"bids",
	"bifocals",
	"biggest",
	"bikini",
	"bimonthly",
	"binocular",
	"biology",
	"biplane",
	"birth",
	"biscuit",
	"bite",
	"biweekly",
	"blender",
	"blip",
	"bluntly",
	"boat",
	"bobsled",
	"bodies",
	"bogeys",
	"boil",
	"boldly",
	"bomb",
	"border",
	"boss",
	"both",
	"bounced",
	"bovine",
	"bowling",
	"boxes",
	"boyfriend",
	"broken",
	"brunt",
	"bubble",
	"buckets",
	"budget",
	"buffet",
	"bugs",
	"building",
	"bulb",
	"bumper",
	"bunch",
	"business",
	"butter",
	"buying",
	"buzzer",
	"bygones",
	"byline",
	"bypass",
	"cabin",
	"cactus",
	"cadets",
	"cafe",
	"cage",
	"cajun",
	"cake",
	"calamity",
	"camp",
	"candy",
	"casket",
	"catch",
	"cause",
	"cavernous",
	"cease",
	"cedar",
	"ceiling",
	"cell",
	"cement",
	"cent",
	"certain",
	"chlorine",
	"chrome",
	"cider",
	"cigar",
	"cinema",
	"circle",
	"cistern",
	"citadel",
	"civilian",
	"claim",
	"click",
	"clue",
	"coal",
	"cobra",
	"cocoa",
	"code",
	"coexist",
	"coffee",
	"cogs",
	"cohesive",
	"coils",
	"colony",
	"comb",
	"cool",
	"copy",
	"corrode",
	"costume",
	"cottage",
	"cousin",
	"cowl",
	"criminal",
	"cube",
	"cucumber",
	"cuddled",
	"cuffs",
	"cuisine",
	"cunning",
	"cupcake",
	"custom",
	"cycling",
	"cylinder",
	"cynical",
	"dabbing",
	"dads",
	"daft",
	"dagger",
	"daily",
	"damp",
	"dangerous",
	"dapper",
	"darted",
	"dash",
	"dating",
	"dauntless",
	"dawn",
	"daytime",
	"dazed",
	"debut",
	"decay",
	"dedicated",
	"deepest",
	"deftly",
	"degrees",
	"dehydrate",
	"deity",
	"dejected",
	"delayed",
	"demonstrate",
	"dented",
	"deodorant",
	"depth",
	"desk",
	"devoid",
	"dewdrop",
	"dexterity",
	"dialect",
	"dice",
	"diet",
	"different",
	"digit",
	"dilute",
	"dime",
	"dinner",
	"diode",
	"diplomat",
	"directed",
	"distance",
	"ditch",
	"divers",
	"dizzy",
	"doctor",
	"dodge",
	"does",
	"dogs",
	"doing",
	"dolphin",
	"domestic",
	"donuts",
	"doorway",
	"dormant",
	"dosage",
	"dotted",
	"double",
	"dove",
	"down",
	"dozen",
	"dreams",
	"drinks",
	"drowning",
	"drunk",
	"drying",
	"dual",
	"dubbed",
	"duckling",
	"dude",
	"duets",
	"duke",
	"dullness",
	"dummy",
	"dunes",
	"duplex",
	"duration",
	"dusted",
	"duties",
	"dwarf",
	"dwelt",
	"dwindling",
	"dying",
	"dynamite",
	"dyslexic",
	"each",
	"eagle",
	"earth",
	"easy",
	"eating",
	"eavesdrop",
	"eccentric",
	"echo",
	"eclipse",
	"economics",
	"ecstatic",
	"eden",
	"edgy",
	"edited",
	"educated",
	"eels",
	"efficient",
	"eggs",
	"egotistic",
	"eight",
	"either",
	"eject",
	"elapse",
	"elbow",
	"eldest",
	"eleven",
	"elite",
	"elope",
	"else",
	"eluded",
	"emails",
	"ember",
	"emerge",
	"emit",
	"emotion",
	"empty",
	"emulate",
	"energy",
	"enforce",
	"enhanced",
	"enigma",
	"enjoy",
	"enlist",
	"enmity",
	"enough",
	"enraged",
	"ensign",
	"entrance",
	"envy",
	"epoxy",
	"equip",
	"erase",
	"erected",
	"erosion",
	"error",
	"eskimos",
	"espionage",
	"essential",
	"estate",
	"etched",
	"eternal",
	"ethics",
	"etiquette",
	"evaluate",
	"evenings",
	"evicted",
	"evolved",
	"examine",
	"excess",
	"exhale",
	"exit",
	"exotic",
	"exquisite",
	"extra",
	"exult",
	"fabrics",
	"factual",
	"fading",
	"fainted",
	"faked",
	"fall",
	"family",
	"fancy",
	"farming",
	"fatal",
	"faulty",
	"fawns",
	"faxed",
	"fazed",
	"feast",
	"february",
	"federal",
	"feel",
	"feline",
	"females",
	"fences",
	"ferry",
	"festival",
	"fetches",
	"fever",
	"fewest",
	"fiat",
	"fibula",
	"fictional",
	"fidget",
	"fierce",
	"fifteen",
	"fight",
	"films",
	"firm",
	"fishing",
	"fitting",
	"five",
	"fixate",
	"fizzle",
	"fleet",
	"flippant",
	"flying",
	"foamy",
	"focus",
	"foes",
	"foggy",
	"foiled",
	"folding",
	"fonts",
	"foolish",
	"fossil",
	"fountain",
	"fowls",
	"foxes",
	"foyer",
	"framed",
	"friendly",
	"frown",
	"fruit",
	"frying",
	"fudge",
	"fuel",
	"fugitive",
	"fully",
	"fuming",
	"fungal",
	"furnished",
	"fuselage",
	"future",
	"fuzzy",
	"gables",
	"gadget",
	"gags",
	"gained",
	"galaxy",
	"gambit",
	"gang",
	"gasp",
	"gather",
	"gauze",
	"gave",
	"gawk",
	"gaze",
	"gearbox",
	"gecko",
	"geek",
	"gels",
	"gemstone",
	"general",
	"geometry",
	"germs",
	"gesture",
	"getting",
	"geyser",
	"ghetto",
	"ghost",
	"giant",
	"giddy",
	"gifts",
	"gigantic",
	"gills",
	"gimmick",
	"ginger",
	"girth",
	"giving",
	"glass",
	"gleeful",
	"glide",
	"gnaw",
	"gnome",
	"goat",
	"goblet",
	"godfather",
	"goes",
	"goggles",
	"going",
	"goldfish",
	"gone",
	"goodbye",
	"gopher",
	"gorilla",
	"gossip",
	"gotten",
	"gourmet",
	"governing",
	"gown",
	"greater",
	"grunt",
	"guarded",
	"guest",
	"guide",
	"gulp",
	"gumball",
	"guru",
	"gusts",
	"gutter",
	"guys",
	"gymnast",
	"gypsy",
	"gyrate",
	"habitat",
	"hacksaw",
	"haggled",
	"hairy",
	"hamburger",
	"happens",
	"hashing",
	"hatchet",
	"haunted",
	"having",
	"hawk",
	"haystack",
	"hazard",
	"hectare",
	"hedgehog",
	"heels",
	"hefty",
	"height",
	"hemlock",
	"hence",
	"heron",
	"hesitate",
	"hexagon",
	"hickory",
	"hiding",
	"highway",
	"hijack",
	"hiker",
	"hills",
	"himself",
	"hinder",
	"hippo",
	"hire",
	"history",
	"hitched",
	"hive",
	"hoax",
	"hobby",
	"hockey",
	"hoisting",
	"hold",
	"honked",
	"hookup",
	"hope",
	"hornet",
	"hospital",
	"hotel",
	"hounded",
	"hover",
	"howls",
	"hubcaps",
	"huddle",
	"huge",
	"hull",
	"humid",
	"hunter",
	"hurried",
	"husband",
	"huts",
	"hybrid",
	"hydrogen",
	"hyper",
	"iceberg",
	"icing",
	"icon",
	"identity",
	"idiom",
	"idled",
	"idols",
	"igloo",
	"ignore",
	"iguana",
	"illness",
	"imagine",
	"imbalance",
	"imitate",
	"impel",
	"inactive",
	"inbound",
	"incur",
	"industrial",
	"inexact",
	"inflamed",
	"ingested",
	"initiate",
	"injury",
	"inkling",
	"inline",
	"inmate",
	"innocent",
	"inorganic",
	"input",
	"inquest",
	"inroads",
	"insult",
	"intended",
	"inundate",
	"invoke",
	"inwardly",
	"ionic",
	"irate",
	"iris",
	"irony",
	"irritate",
	"island",
	"isolated",
	"issued",
	"italics",
	"itches",
	"items",
	"itinerary",
	"itself",
	"ivory",
	"jabbed",
	"jackets",
	"jaded",
	"jagged",
	"jailed",
	"jamming",
	"january",
	"jargon",
	"jaunt",
	"javelin",
	"jaws",
	"jazz",
	"jeans",
	"jeers",
	"jellyfish",
	"jeopardy",
	"jerseys",
	"jester",
	"jetting",
	"jewels",
	"jigsaw",
	"jingle",
	"jittery",
	"jive",
	"jobs",
	"jockey",
	"jogger",
	"joining",
	"joking",
	"jolted",
	"jostle",
	"journal",
	"joyous",
	"jubilee",
	"judge",
	"juggled",
	"juicy",
	"jukebox",
	"july",
	"jump",
	"junk",
	"jury",
	"justice",
	"juvenile",
	"kangaroo",
	"karate",
	"keep",
	"kennel",
	"kept",
	"kernels",
	"kettle",
	"keyboard",
	"kickoff",
	"kidneys",
	"king",
	"kiosk",
	"kisses",
	"kitchens",
	"kiwi",
	"knapsack",
	"knee",
	"knife",
	"knowledge",
	"knuckle",
	"koala",
	"laboratory",
	"ladder",
	"lagoon",
	"lair",
	"lakes",
	"lamb",
	"language",
	"laptop",
	"large",
	"last",
	"later",
	"launching",
	"lava",
	"lawsuit",
	"layout",
	"lazy",
	"lectures",
	"ledge",
	"leech",
	"left",
	"legion",
	"leisure",
	"lemon",
	"lending",
	"leopard",
	"lesson",
	"lettuce",
	"lexicon",
	"liar",
	"library",
	"licks",
	"lids",
	"lied",
	"lifestyle",
	"light",
	"likewise",
	"lilac",
	"limits",
	"linen",
	"lion",
	"lipstick",
	"liquid",
	"listen",
	"lively",
	"loaded",
	"lobster",
	"locker",
	"lodge",
	"lofty",
	"logic",
	"loincloth",
	"long",
	"looking",
	"lopped",
	"lordship",
	"losing",
	"lottery",
	"loudly",
	"love",
	"lower",
	"loyal",
	"lucky",
	"luggage",
	"lukewarm",
	"lullaby",
	"lumber",
	"lunar",
	"lurk",
	"lush",
	"luxury",
	"lymph",
	"lynx",
	"lyrics",
	"macro",
	"madness",
	"magically",
	"mailed",
	"major",
	"makeup",
	"malady",
	"mammal",
	"maps",
	"masterful",
	"match",
	"maul",
	"maverick",
	"maximum",
	"mayor",
	"maze",
	"meant",
	"mechanic",
	"medicate",
	"meeting",
	"megabyte",
	"melting",
	"memoir",
	"menu",
	"merger",
	"mesh",
	"metro",
	"mews",
	"mice",
	"midst",
	"mighty",
	"mime",
	"mirror",
	"misery",
	"mittens",
	"mixture",
	"moat",
	"mobile",
	"mocked",
	"mohawk",
	"moisture",
	"molten",
	"moment",
	"money",
	"moon",
	"mops",
	"morsel",
	"mostly",
	"motherly",
	"mouth",
	"movement",
	"mowing",
	"much",
	"muddy",
	"muffin",
	"mugged",
	"mullet",
	"mumble",
	"mundane",
	"muppet",
	"mural",
	"musical",
	"muzzle",
	"myriad",
	"mystery",
	"myth",
	"nabbing",
	"nagged",
	"nail",
	"names",
	"nanny",
	"napkin",
	"narrate",
	"nasty",
	"natural",
	"nautical",
	"navy",
	"nearby",
	"necklace",
	"needed",
	"negative",
	"neither",
	"neon",
	"nephew",
	"nerves",
	"nestle",
	"network",
	"neutral",
	"never",
	"newt",
	"nexus",
	"nibs",
	"niche",
	"niece",
	"nifty",
	"nightly",
	"nimbly",
	"nineteen",
	"nirvana",
	"nitrogen",
	"nobody",
	"nocturnal",
	"nodes",
	"noises",
	"nomad",
	"noodles",
	"northern",
	"nostril",
	"noted",
	"nouns",
	"novelty",
	"nowhere",
	"nozzle",
	"nuance",
	"nucleus",
	"nudged",
	"nugget",
	"nuisance",
	"null",
	"number",
	"nuns",
	"nurse",
	"nutshell",
	"nylon",
	"oaks",
	"oars",
	"oasis",
	"oatmeal",
	"obedient",
	"object",
	"obliged",
	"obnoxious",
	"observant",
	"obtains",
	"obvious",
	"occur",
	"ocean",
	"october",
	"odds",
	"odometer",
	"offend",
	"often",
	"oilfield",
	"ointment",
	"okay",
	"older",
	"olive",
	"olympics",
	"omega",
	"omission",
	"omnibus",
	"onboard",
	"oncoming",
	"oneself",
	"ongoing",
	"onion",
	"online",
	"onslaught",
	"onto",
	"onward",
	"oozed",
	"opacity",
	"opened",
	"opposite",
	"optical",
	"opus",
	"orange",
	"orbit",
	"orchid",
	"orders",
	"organs",
	"origin",
	"ornament",
	"orphans",
	"oscar",
	"ostrich",
	"otherwise",
	"otter",
	"ouch",
	"ought",
	"ounce",
	"ourselves",
	"oust",
	"outbreak",
	"oval",
	"oven",
	"owed",
	"owls",
	"owner",
	"oxidant",
	"oxygen",
	"oyster",
	"ozone",
	"pact",
	"paddles",
	"pager",
	"pairing",
	"palace",
	"pamphlet",
	"pancakes",
	"paper",
	"paradise",
	"pastry",
	"patio",
	"pause",
	"pavements",
	"pawnshop",
	"payment",
	"peaches",
	"pebbles",
	"peculiar",
	"pedantic",
	"peeled",
	"pegs",
	"pelican",
	"pencil",
	"people",
	"pepper",
	"perfect",
	"pests",
	"petals",
	"phase",
	"pheasants",
	"phone",
	"phrases",
	"physics",
	"piano",
	"picked",
	"pierce",
	"pigment",
	"piloted",
	"pimple",
	"pinched",
	"pioneer",
	"pipeline",
	"pirate",
	"pistons",
	"pitched",
	"pivot",
	"pixels",
	"pizza",
	"playful",
	"pledge",
	"pliers",
	"plotting",
	"plus",
	"plywood",
	"poaching",
	"pockets",
	"podcast",
	"poetry",
	"point",
	"poker",
	"polar",
	"ponies",
	"pool",
	"popular",
	"portents",
	"possible",
	"potato",
	"pouch",
	"poverty",
	"powder",
	"pram",
	"present",
	"pride",
	"problems",
	"pruned",
	"prying",
	"psychic",
	"public",
	"puck",
	"puddle",
	"puffin",
	"pulp",
	"pumpkins",
	"punch",
	"puppy",
	"purged",
	"push",
	"putty",
	"puzzled",
	"pylons",
	"pyramid",
	"python",
	"queen",
	"quick",
	"quote",
	"rabbits",
	"racetrack",
	"radar",
	"rafts",
	"rage",
	"railway",
	"raking",
	"rally",
	"ramped",
	"randomly",
	"rapid",
	"rarest",
	"rash",
	"rated",
	"ravine",
	"rays",
	"razor",
	"react",
	"rebel",
	"recipe",
	"reduce",
	"reef",
	"refer",
	"regular",
	"reheat",
	"reinvest",
	"rejoices",
	"rekindle",
	"relic",
	"remedy",
	"renting",
	"reorder",
	"repent",
	"request",
	"reruns",
	"rest",
	"return",
	"reunion",
	"revamp",
	"rewind",
	"rhino",
	"rhythm",
	"ribbon",
	"richly",
	"ridges",
	"rift",
	"rigid",
	"rims",
	"ringing",
	"riots",
	"ripped",
	"rising",
	"ritual",
	"river",
	"roared",
	"robot",
	"rockets",
	"rodent",
	"rogue",
	"roles",
	"romance",
	"roomy",
	"roped",
	"roster",
	"rotate",
	"rounded",
	"rover",
	"rowboat",
	"royal",
	"ruby",
	"rudely",
	"ruffled",
	"rugged",
	"ruined",
	"ruling",
	"rumble",
	"runway",
	"rural",
	"rustled",
	"ruthless",
	"sabotage",
	"sack",
	"sadness",
	"safety",
	"saga",
	"sailor",
	"sake",
	"salads",
	"sample",
	"sanity",
	"sapling",
	"sarcasm",
	"sash",
	"satin",
	"saucepan",
	"saved",
	"sawmill",
	"saxophone",
	"sayings",
	"scamper",
	"scenic",
	"school",
	"science",
	"scoop",
	"scrub",
	"scuba",
	"seasons",
	"second",
	"sedan",
	"seeded",
	"segments",
	"seismic",
	"selfish",
	"semifinal",
	"sensible",
	"september",
	"sequence",
	"serving",
	"session",
	"setup",
	"seventh",
	"sewage",
	"shackles",
	"shelter",
	"shipped",
	"shocking",
	"shrugged",
	"shuffled",
	"shyness",
	"siblings",
	"sickness",
	"sidekick",
	"sieve",
	"sifting",
	"sighting",
	"silk",
	"simplest",
	"sincerely",
	"sipped",
	"siren",
	"situated",
	"sixteen",
	"sizes",
	"skater",
	"skew",
	"skirting",
	"skulls",
	"skydive",
	"slackens",
	"sleepless",
	"slid",
	"slower",
	"slug",
	"smash",
	"smelting",
	"smidgen",
	"smog",
	"smuggled",
	"snake",
	"sneeze",
	"sniff",
	"snout",
	"snug",
	"soapy",
	"sober",
	"soccer",
	"soda",
	"software",
	"soggy",
	"soil",
	"solved",
	"somewhere",
	"sonic",
	"soothe",
	"soprano",
	"sorry",
	"southern",
	"sovereign",
	"sowed",
	"soya",
	"space",
	"speedy",
	"sphere",
	"spiders",
	"splendid",
	"spout",
	"sprig",
	"spud",
	"spying",
	"square",
	"stacking",
	"stellar",
	"stick",
	"stockpile",
	"strained",
	"stunning",
	"stylishly",
	"subtly",
	"succeed",
	"suddenly",
	"suede",
	"suffice",
	"sugar",
	"suitcase",
	"sulking",
	"summon",
	"sunken",
	"superior",
	"surfer",
	"sushi",
	"suture",
	"swagger",
	"swept",
	"swiftly",
	"sword",
	"swung",
	"syllabus",
	"symptoms",
	"syndrome",
	"syringe",
	"system",
	"taboo",
	"tacit",
	"tadpoles",
	"tagged",
	"tail",
	"taken",
	"talent",
	"tamper",
	"tanks",
	"tapestry",
	"tarnished",
	"tasked",
	"tattoo",
	"taunts",
	"tavern",
	"tawny",
	"taxi",
	"teardrop",
	"technical",
	"tedious",
	"teeming",
	"tell",
	"template",
	"tender",
	"tepid",
	"tequila",
	"terminal",
	"testing",
	"tether",
	"textbook",
	"thaw",
	"theatrics",
	"thirsty",
	"thorn",
	"threaten",
	"thumbs",
	"thwart",
	"ticket",
	"tidy",
	"tiers",
	"tiger",
	"tilt",
	"timber",
	"tinted",
	"tipsy",
	"tirade",
	"tissue",
	"titans",
	"toaster",
	"tobacco",
	"today",
	"toenail",
	"toffee",
	"together",
	"toilet",
	"token",
	"tolerant",
	"tomorrow",
	"tonic",
	"toolbox",
	"topic",
	"torch",
	"tossed",
	"total",
	"touchy",
	"towel",
	"toxic",
	"toyed",
	"trash",
	"trendy",
	"tribal",
	"trolling",
	"truth",
	"trying",
	"tsunami",
	"tubes",
	"tucks",
	"tudor",
	"tuesday",
	"tufts",
	"tugs",
	"tuition",
	"tulips",
	"tumbling",
	"tunnel",
	"turnip",
	"tusks",
	"tutor",
	"tuxedo",
	"twang",
	"tweezers",
	"twice",
	"twofold",
	"tycoon",
	"typist",
	"tyrant",
	"ugly",
	"ulcers",
	"ultimate",
	"umbrella",
	"umpire",
	"unafraid",
	"unbending",
	"uncle",
	"under",
	"uneven",
	"unfit",
	"ungainly",
	"unhappy",
	"union",
	"unjustly",
	"unknown",
	"unlikely",
	"unmask",
	"unnoticed",
	"unopened",
	"unplugs",
	"unquoted",
	"unrest",
	"unsafe",
	"until",
	"unusual",
	"unveil",
	"unwind",
	"unzip",
	"upbeat",
	"upcoming",
	"update",
	"upgrade",
	"uphill",
	"upkeep",
	"upload",
	"upon",
	"upper",
	"upright",
	"upstairs",
	"uptight",
	"upwards",
	"urban",
	"urchins",
	"urgent",
	"usage",
	"useful",
	"usher",
	"using",
	"usual",
	"utensils",
	"utility",
	"utmost",
	"utopia",
	"uttered",
	"vacation",
	"vague",
	"vain",
	"value",
	"vampire",
	"vane",
	"vapidly",
	"vary",
	"vastness",
	"vats",
	"vaults",
	"vector",
	"veered",
	"vegan",
	"vehicle",
	"vein",
	"velvet",
	"venomous",
	"verification",
	"vessel",
	"veteran",
	"vexed",
	"vials",
	"vibrate",
	"victim",
	"video",
	"viewpoint",
	"vigilant",
	"viking",
	"village",
	"vinegar",
	"violin",
	"vipers",
	"virtual",
	"visited",
	"vitals",
	"vivid",
	"vixen",
	"vocal",
	"vogue",
	"voice",
	"volcano",
	"vortex",
	"voted",
	"voucher",
	"vowels",
	"voyage",
	"vulture",
	"wade",
	"waffle",
	"wagtail",
	"waist",
	"waking",
	"wallets",
	"wanted",
	"warped",
	"washing",
	"water",
	"waveform",
	"waxing",
	"wayside",
	"weavers",
	"website",
	"wedge",
	"weekday",
	"weird",
	"welders",
	"went",
	"wept",
	"were",
	"western",
	"wetsuit",
	"whale",
	"when",
	"whipped",
	"whole",
	"wickets",
	"width",
	"wield",
	"wife",
	"wiggle",
	"wildly",
	"winter",
	"wipeout",
	"wiring",
	"wise",
	"withdrawn",
	"wives",
	"wizard",
	"wobbly",
	"woes",
	"woken",
	"wolf",
	"womanly",
	"wonders",
	"woozy",
	"worry",
	"wounded",
	"woven",
	"wrap",
	"wrist",
	"wrong",
	"yacht",
	"yahoo",
	"yanks",
	"yard",
	"yawning",
	"yearbook",
	"yellow",
	"yesterday",
	"yeti",
	"yields",
	"yodel",
	"yoga",
	"younger",
	"yoyo",
	"zapped",
	"zeal",
	"zebra",
	"zero",
	"zesty",
	"zigzags",
	"zinger",
	"zippers",
	"zodiac",
	"zombie",
	"zones",
	"zoom",
}