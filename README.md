# CryptoPals Challenges - Go Edition
I've been meaning to complete [CryptoPals](https://cryptopals.com/) for years - I've recently found solace in [GoLang](https://golang.org/), so I'm teaching it to myself through the powers of cryptography.

As such, the syntax, use of functions and a million other things will not be up to what might come out of an experienced Go programmer, but this is just a learning experience for me :D

## Usage
Choose a set number `n`:
```sh
$ cd set{n}
$ go test -v
```
This will run through each of the challenges for that set, verifying the expected output to what my software generates.

## Progress

### [Set 1: Basics](http://cryptopals.com/sets/1)

  - [x] 01. Convert hex to base64
  - [x] 02. Fixed XOR
  - [x] 03. Single-byte XOR cipher
  - [x] 04. Detect single-character XOR
  - [x] 05. Implement repeating-key XOR
  - [x] 06. Break repeating-key XOR
  - [x] 07. AES in ECB mode
  - [x] 08. Detect AES in ECB mode

### [Set 2: Block Crypto](http://cryptopals.com/sets/2)

  - [x] 09. Implement PKCS#7 padding
  - [x] 10. Implement CBC mode
  - [x] 11. An ECB/CBC detection oracle
  - [x] 12. Byte-at-a-time ECB decryption (Simple)
  - [x] 13. ECB cut-and-paste
  - [x] 14. Byte-at-a-time ECB decryption (Harder)
  - [x] 15. PKCS#7 padding validation
  - [x] 16. CBC bitflipping attacks

### [Set 3: Block & Stream Crypto](http://cryptopals.com/sets/3)

  - [x] 17. The CBC padding oracle
  - [ ] 18. Implement CTR, the stream cipher mode
  - [ ] 19. Break fixed-nonce CTR mode using substitions
  - [ ] 20. Break fixed-nonce CTR statistically
  - [ ] 21. Implement the MT19937 Mersenne Twister RNG
  - [ ] 22. Crack an MT19937 seed
  - [ ] 23. Clone an MT19937 RNG from its output
  - [ ] 24. Create the MT19937 stream cipher and break it

### [Set 4: Stream Crypto & Randomness](http://cryptopals.com/sets/4)

  - [ ] 25. Break "random access read/write" AES CTR
  - [ ] 26. CTR bitflipping
  - [ ] 27. Recover the key from CBC with IV=Key
  - [ ] 28. Implement a SHA-1 keyed MAC
  - [ ] 29. Break a SHA-1 keyed MAC using length extension
  - [ ] 30. Break an MD4 keyed MAC using length extension
  - [ ] 31. Implement and break HMAC-SHA1 with an artificial timing leak
  - [ ] 32. Break HMAC-SHA1 with a slightly less artificial timing leak

### [Set 5: Diffie-Hellman and Friends](http://cryptopals.com/sets/5)

  - [ ] 33. Implement Diffie-Hellman
  - [ ] 34. Implement a MITM key-fixing attack on Diffie-Hellman with parameter injection
  - [ ] 35. Implement DH with negotiated groups, and break with malicious "g" parameters
  - [ ] 36. Implement Secure Remote Password (SRP)
  - [ ] 37. Break SRP with a zero key
  - [ ] 38. Offline dictionary attack on simplified SRP
  - [ ] 39. Implement RSA
  - [ ] 40. Implement an E=3 RSA Broadcast attack

### [Set 6: RSA and DSA](http://cryptopals.com/sets/6)

  - [ ] 41. Implement unpadded message recovery oracle
  - [ ] 42. Bleichenbacher's e=3 RSA Attack
  - [ ] 43. DSA key recovery from nonce
  - [ ] 44. DSA nonce recovery from repeated nonce
  - [ ] 45. DSA parameter tampering
  - [ ] 46. RSA parity oracle
  - [ ] 47. Bleichenbacher's PKCS 1.5 Padding Oracle (Simple Case)
  - [ ] 48. Bleichenbacher's PKCS 1.5 Padding Oracle (Complete Case)