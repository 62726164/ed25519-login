# ed25519-login

Login to websites using a base64 encoded [Ed25519](https://en.wikipedia.org/wiki/EdDSA#Ed25519) signature. A simple alternative to [webauthn](https://webauthn.guide/).

# Advantages of this approach

  * Much simpler than webauthn.
  * Does not require passwords or password hashing.
    * No passwords to brute-force or stuff.
    * No password complexity rules.
    * No password dumping and cracking.
  * Ed25519 keys are controlled by the end users.
    * Private keys never leave end users' devices.
	* Websites store the users' public keys.

## Disadvantages

  1. If multiple websites required Unix epoch time signatures for login and a user used the same Ed25519 keypair for those sites, then a signature would be valid on all the websites.
  2. This approach is phishable.

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
signature: cdF4uV7L4ZupvpJfEHXC1QfjmKBGCUc/U72KRiPv3xfU1vneLFgHTpPECUjGITVuAcQwhrIGYNO3XtB+gtz+Cg==

# Verify a signature
$ ./ed25519-login -verify -signature cdF4uV7L4ZupvpJfEHXC1QfjmKBGCUc/U72KRiPv3xfU1vneLFgHTpPECUjGITVuAcQwhrIGYNO3XtB+gtz+Cg== -message hi -pubkey /home/user/ed.pub
signature verified: true
```

## By default, ed25519-login signs the current Unix epoch time
```bash
$ ./ed25519-login -sign -privkey /home/user/ed
signature: m2A8sZxRbSXCJIhwnZCPVFSBy/c/kIytxG0bgcn+PH0H35jgv88Y4Hlof8YD4A7NLWFsa5FHstm5Dc4BthMGDw==

$ ./ed25519-login -sign -privkey /home/user/ed
signature: 8QGMX8MnE8khfVZ9VWScT0VkvXD9XCK/AesPdMIFxaZAQQTpFjr2PlDbrgcTZjPIUTR32bpnpoXDAf2USnyxDg==

$ ./ed25519-login -sign -privkey /home/user/ed
signature: NRnVqh5o6dm4XB7KYVqSrEBHdDAMoOjC1+0a6Ht0D2YQk4KEfIJGg0Jmbibtz8Ag+e62i49IuIN2MYa/6ibACw==
```

## Login Process

  1. [Register](https://gen.go350.com/register) your base64 encoded public Ed25519 key with the website.
  2. Use your private Ed25519 key to sign the current [Unix Epoch Time](https://en.wikipedia.org/wiki/Unix_time)
  3. Paste the base64 encoded signature into the website's [login](https://gen.go350.com/login) form.

## Notes

  * Requires Go 1.16 or higher.
  * A base64 encoded public Ed25519 key looks like this: *uv8AWTxoUzWJp2RDGczJXf/Z+Cq484+wEM602zjTLNM=*
  * A base64 encoded Ed25519 signature looks like this: *E3FCTm0qNSu6gl/6oKcf3VABO4u/WEpeKnDOaX+VJFeYmrAA1rF3I9VEN2sD1ogIiTN9F7xtf9Fhwz+jJMm1Cg==*
  * Client systems should use [NTP](https://www.ntp.org/) to ensure accurate time.
