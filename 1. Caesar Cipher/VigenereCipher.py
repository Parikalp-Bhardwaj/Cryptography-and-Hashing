ALPHABET=" ABCDEFGHIJKLMNOPQRSTUVWXYZ"

def vigenere_encrypt(plain_text,key):
    # plain_text is a text we want to encrypt + case-insensitive
    plain_text = plain_text.upper()
    key = key.upper()
    cipher_text = ''
    # it will represent the index of the letters of the key
    key_index = 0
    # we have to consider all the characters in the plain_text
    for char in plain_text:
        # xi = character that we want to encrypt
        # ki = key for encrypting the
        # formula is Ei(xi) = (xi + ki) % 26
        # we have i-th letter of the key for encrypting the i-th letter
        idx = (ALPHABET.find(char) + ALPHABET.find(key[key_index])) % len(ALPHABET)
        cipher_text += ALPHABET[idx]
        #increment the key index
        key_index += 1
        
        if key_index == len(key):
            key_index = 0
    
    return cipher_text
        
def vigenere_decrypt(cipher_text,key):
    cipher_text = cipher_text.upper()
    key = key.upper()
    key_index = 0
    plain_text = ''
    for char in cipher_text:
        # xi = character that we want to decrypt
        # ki = key for encrypting the
        # formula is Di(xi) = (xi - ki) % 26
        # we have i-th letter of the key for encrypting the i-th letter
        idx = (ALPHABET.find(char) - ALPHABET.find(key[key_index])) % len(ALPHABET)
        plain_text += ALPHABET[idx]
        
        key_index = key_index + 1
        if key_index == len(key):
            key_index = 0        
    return plain_text


if __name__ == "__main__":
    text = "CRYPTOGRAHY AND HASHING"
    encrypt = vigenere_encrypt(text,"PYTHON")
    print(encrypt)
    print(vigenere_decrypt(encrypt,"PYTHON"))