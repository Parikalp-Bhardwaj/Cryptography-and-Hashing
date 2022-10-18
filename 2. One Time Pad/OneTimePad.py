from random import randint
import matplotlib.pyplot as plt
# we need the alphabet because we convert letters into numerical values to be able to use
# mathematical operations (note we encrypt the spaces as well)
ALPHABET=" ABCDEFGHIJKLMNOPQRSTUVWXYZ().,"


# In cryptography, a one-time pad is a system in which a randomly generated private key 
# is used only once to encrypt a message that is then decrypted by the receiver using a 
# matching one-time pad and key.



#One time pad (OTP) encryption
def encrypt(text,key):
    text = text.upper()
    #key is randomly generates of list
    #the length of key equal to length of text
    cipher_text = ''
    # consider all the plain_text letters: enumerate returns the item + it's index
    for idx,char in enumerate(text):
        # the value with which we shift the given letter
        # otp = now every key will assign to char 
        otp = key[idx]
        # xi = character that we want to encrypt
        char_idx = ALPHABET.find(char)
        # formula is Ei(xi) = (xi + opt) % 26
        # we have i-th letter of the key for encrypting the i-th letter
        cipher_text += ALPHABET[(char_idx + otp) % len(ALPHABET)]
        
    return cipher_text
        
        
def decrypt(text,key):
    text = text.upper()
    plain_text = ""
    for idx,char in enumerate(text):
        # the value with which we shift the given letter
        # xi = character that we want to encrypt
        char_idx = ALPHABET.find(char)
        # otp = now every key will assign to char 
        opt = key[idx]
        # formula is Ei(xi) = (xi - opt) % 26
        # we have i-th letter of the key for encrypting the i-th letter
        plain_text += ALPHABET[(char_idx - opt) % len(ALPHABET)]
    
    return plain_text

def random_sequence(text):
    random = []
    for _ in range(len(text)):
        random.append(randint(0,len(ALPHABET)-1))
    return random


def frequency_analysis(text):
    #the text are analyse
    text = text.upper()
    letter_frequency = {}
    
    
    for letter in ALPHABET:
        letter_frequency[letter] = 0
        
    #let's consider the text we want to analyse
    for letter in text:
        # we keep incrementing the occurrence of the given text
        if letter in ALPHABET:
            letter_frequency[letter] +=1
    
    return letter_frequency
            

def plot_distribution(letter_frequency):
    plt.bar(letter_frequency.keys(),letter_frequency.values())
    plt.show()


if __name__ =="__main__":
    message = "Rust has different behaviour than other languages. In a language where variables are always references (like Java or Python)"
    key = random_sequence(message)
    print("Key are ",key)
    print("Original text is ",message.upper())
    cipher = encrypt(message,key)
    print("Encrypt Message ",cipher)
    real_text = decrypt(cipher,key)
    print("Decrypt text ",real_text)
    plot_distribution(frequency_analysis(cipher))