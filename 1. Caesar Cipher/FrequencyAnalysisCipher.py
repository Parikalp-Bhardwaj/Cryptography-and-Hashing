import matplotlib.pyplot as plt

# In cryptography, frequency analysis is the study of the frequency of letters 
# or groups of letters in a ciphertext. The method is used as an aid to breaking 
# substitution ciphers

LETTERS ="ABCDEFGHIJKLMNOPQRSTUVWXYZ"

def frequency_analysis(text):
    text = text.upper()
    # we use dictionary to store the letter-freq pair
    letter_frequency = {}
    
    # initialize the dicitonary with 0 frequencies
    for letter in LETTERS:
        letter_frequency[letter] = 0
        
    # let's calculate the text
    for letter in text:
        if letter in LETTERS:
            letter_frequency[letter]+=1
            
    return letter_frequency

def caesar_crack(text):
    freq = frequency_analysis(text)
    freq = sorted(freq.items(),key=lambda x:x[1],reverse=True)
    # now how we can find out the key 
    #formula is 
    # key = value of cipher text most frequent letter - value of E
    # Most freq letter is LETTERS.find(freq[0][0])
    # value of E is LETTERS.find('E')
    print(freq)
    print("The possible key value: ",(LETTERS.find(freq[0][0])) - LETTERS.find('E')) 

if __name__ =="__main__":
    # text = "The most important property of a program is that it is correct. For the Caesar Cipher, Latin letters are rotated, in either direction, by a fixed shift amount. Decoding is the reverse of encoding."
    # text = "FU@SWRJUDSK@CDQGCKDVKLQJAB"
    text = "KHOOR"
    freq = caesar_crack(text)
    # plt.bar(freq.keys(),freq.values())
    # plt.show()
    # caesar_crack(freq)
    