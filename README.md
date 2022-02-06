# ed25519-login

Login to websites using a base64 encoded [Ed25519](https://en.wikipedia.org/wiki/EdDSA#Ed25519) key. A simple alternative to [webauthn](https://webauthn.guide/).

# Advantages of this approach

  * Much simpler than webauthn.
  * The website only stores users' public Ed25519 keys.

## Build
```bash
$ make
```

## Usage
```bash
# Generate a keypair
$ ./ed25519-login -gen -privkey /home/user/ed

$ ls -alh /home/user/ed*
-r-------- 1 user user 88 Feb  6 17:13 /home/user/ed
-r-------- 1 user user 44 Feb  6 17:13 /home/user/ed.pub

# Sign a message
$ ./ed25519-login -sign -message hi -privkey /home/user/ed
2022/02/06 17:14:08 signature: cdF4uV7L4ZupvpJfEHXC1QfjmKBGCUc/U72KRiPv3xfU1vneLFgHTpPECUjGITVuAcQwhrIGYNO3XtB+gtz+Cg==

# Verify a signature
$ ./ed25519-login -verify -signature cdF4uV7L4ZupvpJfEHXC1QfjmKBGCUc/U72KRiPv3xfU1vneLFgHTpPECUjGITVuAcQwhrIGYNO3XtB+gtz+Cg== -message hi -pubkey /home/user/ed.pub
2022/02/06 17:15:04 Signature verified: true
```

## By default, ed25519-login signs the current Unix epoch time
```bash
$ ./ed25519-login -sign -privkey /home/user/ed
2022/02/06 17:15:26 signature: m2A8sZxRbSXCJIhwnZCPVFSBy/c/kIytxG0bgcn+PH0H35jgv88Y4Hlof8YD4A7NLWFsa5FHstm5Dc4BthMGDw==

$ ./ed25519-login -sign -privkey /home/user/ed
2022/02/06 17:15:35 signature: 8QGMX8MnE8khfVZ9VWScT0VkvXD9XCK/AesPdMIFxaZAQQTpFjr2PlDbrgcTZjPIUTR32bpnpoXDAf2USnyxDg==

$ ./ed25519-login -sign -privkey /home/user/ed
2022/02/06 17:15:41 signature: NRnVqh5o6dm4XB7KYVqSrEBHdDAMoOjC1+0a6Ht0D2YQk4KEfIJGg0Jmbibtz8Ag+e62i49IuIN2MYa/6ibACw==
```

## Login Process

  1. Register your base64 encoded public Ed25519 key with the website.
  2. Use your private Ed25519 key to sign the current [Unix Epoch Time](https://en.wikipedia.org/wiki/Unix_time)
  3. Paste that signature into the website's login form.

## Notes

  * Requires Go 1.16 or higher.
  * A base64 encoded public Ed25519 key looks like this *uv8AWTxoUzWJp2RDGczJXf/Z+Cq484+wEM602zjTLNM=*
