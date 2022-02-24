package main

import (
	"crypto/x509"
	"encoding/base64"
	"flag"
	"fmt"
	"golang.org/x/crypto/ed25519"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var privkeyfile = flag.String("privkey", "", "the path to the private keyfile.")
	var pubkeyfile = flag.String("pubkey", "", "the path to the public keyfile.")
	var message = flag.String("message", "", "the message to sign.")
	var signature = flag.String("signature", "", "the base64 encoded signature.")
	var gen = flag.Bool("gen", false, "generate a new keypair. Requires -privkey.")
	var sign = flag.Bool("sign", false, "sign the message. Requires -privkey")
	var verify = flag.Bool("verify", false, "verify a signature. Requires -pubkey, -message and -signature.")
	var verbose = flag.Bool("verbose", false, "verbose logging.")
	var help = flag.Bool("help", false, "show help.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	if *gen && *privkeyfile != "" {
		// Generate new random keypair
		pubkey, privkey, err := ed25519.GenerateKey(nil)
		if err != nil {
			log.Fatalln(err)
		}

		// base64 encode the keys
		b64pubkey := base64.StdEncoding.EncodeToString(pubkey)
		b64privkey := base64.StdEncoding.EncodeToString(privkey)

		// create keyfiles (os.Writefile requires Go 1.16 or later)
		err = os.WriteFile(*privkeyfile, []byte(b64privkey), 0400)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(*privkeyfile+".pub", []byte(b64pubkey), 0400)
		if err != nil {
			log.Fatal(err)
		}

		if *verbose {
			// PEM Blocks
			// https://pkg.go.dev/crypto/x509#MarshalPKCS8PrivateKey
			privbytes, err := x509.MarshalPKCS8PrivateKey(privkey)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("------BEGIN PRIVATE KEY------")
			fmt.Println(base64.StdEncoding.EncodeToString(privbytes))
			fmt.Println("------END PRIVATE KEY------")

			pubbytes, err := x509.MarshalPKIXPublicKey(pubkey)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("------BEGIN PUBLIC KEY------")
			fmt.Println(base64.StdEncoding.EncodeToString(pubbytes))
			fmt.Println("------END PUBLIC KEY------")

			fmt.Printf("base64 public key: %s\n", b64pubkey)
			fmt.Printf("base64 private key: %s\n", b64privkey)
		}
	}

	if *sign && *privkeyfile != "" {
		b64privkey, err := os.ReadFile(*privkeyfile)
		if err != nil {
			log.Fatal(err)
		}

		dprivkey, err := base64.StdEncoding.DecodeString(string(b64privkey))
		if err != nil {
			log.Fatalln(err)
		}

		b64pubkey, err := os.ReadFile(*privkeyfile + ".pub")
		if err != nil {
			log.Fatal(err)
		}

		// Sign the message
		if *message == "" {
			now := time.Now()
			seconds := now.Unix()
			strSeconds := strconv.FormatInt(seconds, 10)
			sig := ed25519.Sign(dprivkey, []byte(strSeconds))
			fmt.Printf("Public key: %s\n", string(b64pubkey))
			fmt.Printf("Signature: %s\n", base64.StdEncoding.EncodeToString(sig))
		} else {
			sig := ed25519.Sign(dprivkey, []byte(*message))
			fmt.Printf("Public key: %s\n", string(b64pubkey))
			fmt.Printf("Signature: %s\n", base64.StdEncoding.EncodeToString(sig))
		}
	}

	if *verify && *pubkeyfile != "" && *signature != "" {
		b64pubkey, err := os.ReadFile(*pubkeyfile)
		if err != nil {
			log.Fatal(err)
		}

		dpubkey, err := base64.StdEncoding.DecodeString(string(b64pubkey))
		if err != nil {
			log.Fatalln(err)
		}

		dsignature, err := base64.StdEncoding.DecodeString(*signature)
		if err != nil {
			log.Fatalln(err)
		}

		// verify the signature
		verified := ed25519.Verify(dpubkey, []byte(*message), dsignature)

		fmt.Printf("signature verified: %t\n", verified)
	}
}
