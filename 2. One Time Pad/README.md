## What is One Time Pad?

One Time Pad algorithm is the improvement of the Vernam Cipher, proposed by An Army Signal Corp officer, Joseph Mauborgne. It is the only available algorithm that is unbreakable(completely secure). It is a method of encrypting alphabetic plain text. It is one of the Substitution techniques which converts plain text into ciphertext. In this mechanism, we assign a number to each character of the Plain-Text.


### The two requirements for the One-Time pad are
- The key should be randomly generated as long as the size of the message.
- The key is to be used to encrypt and decrypt a single message, and then it is discarded.

### In theory it's impossible to break a one time pad
- Generating perfectly random number (as keys) is extremely hard almost impossible to get truly random number with computer (random sequence with small period: Vigenere-cipher)

- The key has the same length as the plaintext: If we are able to exchange this key securly then why not to exchange the plaintext itself?

This type of cryptography uses just a single Key.
So the same key is use both for encryption and decryption as well that is why it is called `symmetric encryption`

![Symmetric Encryption](https://www.ssl2buy.com/wiki/wp-content/uploads/2015/12/Symmetric-Encryption.png)

