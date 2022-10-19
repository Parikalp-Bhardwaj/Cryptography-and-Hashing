# Cryptography in Blockchain

Cryptography is the method of securing important data from unauthorized access. In the blockchain, cryptographic techniques are a part of security protocols. As we know blockchain technology is based on three main pillars; Distributed ledger, Peer-to-peer network, and Cryptographic security. Blockchain uses two types of security approaches i.e. `Cryptography` and `Hashing`. The basic difference between these two is that `cryptography is used to encrypt messages in a P2P (Point-to-Point) network.` Whereas, `hashing is used to secure block information and link blocks in a blockchain.` 

We can break down the word cryptography into two parts; Crypto meaning “hidden” and Graphy meaning “writing”. Therefore, cryptography is a method of converting plaintext into unreadable coded text. 

Two main concepts behind cryptography are `Encryption` and `Decryption`. `Encryption` is coding information in such a way that you and I cannot understand what it means just by looking at it. `Decryption` is the reverse of encryption, i.e. decoding of the coded information.

Key encryption is a cryptographic method that ensures safe transmission of information from point A to point B. This is like an external layer of protection. The internal layer is hashing. Hashing is a process of irreversible encryption of data in a block. All the data present in the block is encrypted using the `SHA256 hashing algorithm` which is irreversible. Thus, applying cryptography at two levels in a blockchain network makes it absolutely secure. 


# Types of Cryptography

 Let’s look at the three primary types of cryptography.
 - Symmetric Cryptography
 - Asymmetric Cryptography
 - Cryptographic Hash
 
 
# Symmetric Cryptography
 
`Symmetric cryptography` – or symmetric key cryptography – was the first type of encryption used online. Symmetric cryptography translates information into a cipher—or encrypted code. To decrypt the cipher, you need a key. 

In `symmetric cryptography`, both the sender and the receiver use the same key to encrypt and decrypt the data. Because it’s so straightforward, symmetric cryptography can process large amounts of data very quickly. 

As you can imagine, though, sharing keys became a problem. Think about when you’re trying to share a password with someone. If you text or email that password, hackers can easily see it. You almost need a password for your password! 

Similarly, sharing keys from the sender to the receiver created a vulnerability that hackers could pretty quickly exploit. 

Suppose `Node A` wants to send some confidential information to `Node B`. To facilitate this transition using the symmetric key method, `Node A` will encrypt the information into an unreadable ciphertext using a key k1 and send it to `Node B`. `Node B` will receive the ciphertext and decrypt it using the same key i.e. k1. This means both `Node A` and `Node B` need to have the same key k1. In the same way, if `Node A` wants to communicate with `Node C`, they both will need a new key k2 between them. Or `Node B` and `Node C` will need yet another new key k3 to carry out a transaction. 


![Symmetric Encryption](https://www.ssl2buy.com/wiki/wp-content/uploads/2015/12/Symmetric-Encryption.png)


This is where the following type of cryptography comes in: `asymmetric cryptography`. 



# Asymmetric Cryptography

With `asymmetric cryptography`, the sender and receiver have different keys. One key is used to encrypt the information, and a separate key is used to decrypt that information at the other end. 

But, if the two people essentially have different passwords, how do you ensure that when the information is sent, only the correct receiver can open it? In short, without sharing keys, how do you tell the code to open for the right person? 

To solve this problem, `asymmetric cryptography` uses a system of two keys per user: a `public key` and a `private key`. Your public key is unique to you, but everyone else can see it, too. No one knows your private key except for you. It’s like a PIN for your bank account.

The public and private keys work together. So, in a transaction, the person sending information can send it to your public key. Then, in order to decrypt data sent to your public key, you have to have the private key to unlock it. 

To send someone a message, you would encrypt it with their public key. Then, only they can unlock it using their private key. Alternatively, if someone adds a digital signature to a set of data with their private key, anyone online can use their public key to decrypt the signature and verify it really is them.

![Symmetric Encryption](https://www.simplilearn.com/ice9/free_resources_article_thumb/alice.PNG)


# Cryptographic Hash

The final type of cryptography is `hashing`. A cryptographic hash is a set of text. Any plaintext information can be put through a hashing algorithm and turned into a unique string of text. The text doesn’t mean anything. For example, the word “Hello” can be turned into the sha1 hash: “f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0.”

Once the original data goes through the cryptographic hash function, you can’t reverse the process. That’s how cryptographic hashing differs from `symmetric` or `asymmetric encryption`, which you can decrypt with a key. There’s no way to start from a hash you’ve never seen before and deduce what the original data was. 

As long as the same hash algorithm is used, the same data will always become the same hash. So, if data along the way was changed, users can tell by comparing it to the final hash. However, hackers have found ways to grab a lot of hashes and then compare them to hashes for common words and phrases. If they find a match, then they know what the hash represents. This is how hackers steal passwords in a data breach. 

Another essential feature of hashing is that you can reduce a whole lot of data to a small string of text in a hash. Hashes are always the same length, regardless of how long or extensive the data is. So, hashing is a way to compress information. We’ll explain more about why this is so important, but first, let’s bring in blockchain.







