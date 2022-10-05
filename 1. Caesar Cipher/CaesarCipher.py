ALPHABET=" ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}"
PRIVATE_KEY = 3

# En(x) = (x+n) % len(ALPHABET) 
# x = character of plain text
# n = PRIVATE_KEY 

def caesar_cipher(plain_text):
    # encryption message
    cipher_text = ''
    plain_text = plain_text.upper()
    
    for chr in plain_text:
        # find the index of every character
        idx = ALPHABET.find(chr)
        # now the formula of caesar encryption is 
        # to the key use modular arithmetic to transform the values within 
        idx = (idx + PRIVATE_KEY) % len(ALPHABET)
        
        cipher_text += ALPHABET[idx]
    return cipher_text
        

# Caesar decryption
# formula  Dn(x) = (x - n) % len(ALPHABET)

def caesar_decypt(cipher_text):
    plain_text = ''
    for c in cipher_text:
        idx = ALPHABET.find(c)
        # Dn(x) = (x - n) % len(ALPHABET)
        idx = (idx - PRIVATE_KEY) % len(ALPHABET)
        plain_text += ALPHABET[idx]
    return plain_text

if __name__ == "__main__":
    # text = "cryptography and hashing?}"
    text = "hello"
    encryption = caesar_cipher(text)
    print(encryption)
    print(caesar_decypt(encryption).lower())