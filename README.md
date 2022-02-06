# ed25519-login

Login to websites using an Ed25519 key. A simple alternative to [https://webauthn.guide/](webauthn).

## Build
```bash
$ make
```

## Usage
```bash
$ ./ed25519-login -gen -privkey /home/user/ed

$ ls -alh /home/user/ed*
-r-------- 1 user user 88 Feb  6 17:13 /home/user/ed
-r-------- 1 user user 44 Feb  6 17:13 /home/user/ed.pub

$ ./ed25519-login -sign -message hi -privkey /home/user/ed
2022/02/06 17:14:08 signature: cdF4uV7L4ZupvpJfEHXC1QfjmKBGCUc/U72KRiPv3xfU1vneLFgHTpPECUjGITVuAcQwhrIGYNO3XtB+gtz+Cg==

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

## Notes

  * Requires Go 1.16 or higher.
  * Sign the current Unix epoch time and paste that signature into the website's login form.
