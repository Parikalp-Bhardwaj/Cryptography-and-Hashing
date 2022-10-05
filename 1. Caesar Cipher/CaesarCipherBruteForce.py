ALPHABET=" ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){<>?}"

def crake_caesar(cipher_text):
    #we try all possible key value 
    for key in range(len(ALPHABET)):
        plain_text = ''
        for char in cipher_text:
            idx = ALPHABET.find(char)
            # formula how we can decrypt the text 
            # Dn(x) = (x - n) % len(ALPHABET)
            fml = (idx - key) % len(ALPHABET)
            plain_text += ALPHABET[fml]
            
            
        print(f"With key ${key} the result is ${plain_text} ")
        
if __name__ =="__main__":
    decrypt=crake_caesar("FU@SWRJUDSK@CDQGCKDVKLQJAB")
    # print()