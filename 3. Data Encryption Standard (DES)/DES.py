from Crypto.Cipher import DES

def add_padding(text,blocksize=64):
    pad_len = blocksize - (len(text) % blocksize)
    padding = '$' * pad_len
    return text + padding

def remove_padding(text):
    counter = 0

    print(text[::-1])
    for c in text[::-1]:
        if c == "$":
            counter+=1
        else:
            break
    print("t",text)
    return text[:-counter]

def encrypt(plain_text,key):
    des = DES.new(key,DES.MODE_ECB)
    return des.encrypt(plain_text)

def decrypt(cipher_text,key):
    des = DES.new(key,DES.MODE_ECB)
    return des.decrypt(cipher_text).decode('UTF-8')


if __name__ == '__main__':
    key = "secretaa"
    plain_text = 'This is the secret message we want to encrypt'
    plain_text = add_padding(plain_text)
    cipher_text = encrypt(plain_text,key)
    print(cipher_text)
    decrypted_message = decrypt(cipher_text,key)
    print("d",decrypted_message)
    decrypted_message = remove_padding(decrypted_message)
    print(decrypted_message)
    