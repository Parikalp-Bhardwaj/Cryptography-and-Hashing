from pydoc import plain
from Crypto.Cipher import AES
from Crypto import Random
import hashlib
from base64 import b64encode, b64decode

class AESCipher:
    def __init__(self,key):
        self.block_size = AES.block_size
        self.key = hashlib.sha256(key.encode()).digest()
        
    def add_padding(self,plain_text):
        bytes_to_pad = self.block_size - len(plain_text) % self.block_size
        # PKCS or CMS "Cryptographic Message Syntax"
        ascii_string = chr(bytes_to_pad)
        padding_string = ascii_string * bytes_to_pad
        return plain_text + padding_string
    
    
    def remove_padding(self,txt):
        last_character = txt[len(txt) -1:]
        return txt[:-ord(last_character)]
    
    def encrypt(self,plain_text):
        plain_text = self.add_padding(plain_text)
        # IV (Initialization Vector - Seed)
        # if we use the same key and same plain text
        # the result will be different !!!
        iv = Random.new().read(self.block_size)
        cipher = AES.new(self.key,AES.MODE_CBC,iv)
        encrypted_text = cipher.encrypt(plain_text.encode())
        return b64encode(iv+encrypted_text).decode('utf-8')
    
    
    def decrypt(self,encrpted_text):
        encrpted_text = b64decode(encrpted_text)
        iv = encrpted_text[:self.block_size]
        cipher = AES.new(self.key,AES.MODE_CBC,iv)
        plain_text = cipher.decrypt(encrpted_text[self.block_size:]).decode('utf-8')
        return self.remove_padding(plain_text)
    
if __name__ == "__main__":
    aes = AESCipher("My Secret Message")
    msg = "This is a secret message"
    encrypted = aes.encrypt(msg)
    print(encrypted)
    print(aes.decrypt(encrypted))

    
    
    
    
    
        
        
        