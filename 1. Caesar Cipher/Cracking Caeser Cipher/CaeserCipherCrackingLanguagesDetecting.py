# When Cracking a given cipher it nay be useful to detect whether
# the decrypted language is eglish or not
ALPHABET=" ABCDEFGHIJKLMNOPQRSTUVWXYZ"

ENGLISH_WORDS = []

def get_data():
    # opening english words
    words = open("./english_words.txt",'r')
    # we can use a dictionary and check whether the given words are
    # are present in a dictionary or not
    
    for word in words.read().split("\n"):
        ENGLISH_WORDS.append(word)
    
    words.close()
    print(len(ENGLISH_WORDS))

#count the number of english words in a given text
def count_words(text):
    text = text.upper()
    # transform the text into a list of words (split by spaces)
    words = text.split(' ')
    # matches counts the number of english words in the text
    matches = 1
    # consider all the words in the text and check whether the
    # given word in english or not
    for word in words:
        if word in ENGLISH_WORDS:
            matches +=1
    
    return matches

#decides whether a given text is english or not
def is_text_english(text):
    # number of English words in a given text
    matches = count_words(text)
    # you can define your own classification alogrithm in this case
    # the assumption: if 70% of the words in the text are english words
    # then it is an enlgish text
    if (float(matches) / len(text.split(' ')))* 100 >= 70:
        return True
    
    #not an english text
    return False

# now we will try brute force to get the decrepted text
def crake_caesar(cipher_text):
   # we try all the possible key values so the size of the ALPHABET
    for key in range(len(ALPHABET)):

        # reinitialize this to be an empty string
        plain_text = ''

        # we just have to make a simple caesar decryption
        for c in cipher_text:
            index = ALPHABET.find(c)
            index = (index - key) % len(ALPHABET)
            plain_text = plain_text + ALPHABET[index]

        # print the actual decrypted string with the given key
        if is_text_english(plain_text):
            print("We have managed to crack Caesar cipher, the key is: %s, the message is %s" % (key, plain_text))


if __name__ =="__main__":
    get_data()
    # plain_text = "The most important property of a program is that it is correct. For the Caesar Cipher, Latin letters are rotated, in either direction, by a fixed shift amount. Decoding is the reverse of encoding."
    # encrypted = "KHOOR"
    encrypted = "FU@SWRJUDSK@CDQGCKDVKLQJAB"
    print(crake_caesar(encrypted))
    
    
    