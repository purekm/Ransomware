
import binascii
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

# file path
cipher_file = r"C:/Users/Alpha/Desktop/FLAG.txt.ryk"

# encrypted_data file
with open(cipher_file, "rb") as f:
    enc_data = f.read()

# AES 32 Bytes Key & hex to bytes
hexkey = "133a985d25765d4af3c84fcb1f8296f888d5d8fa028697e186939dbaf283108e"
key = binascii.unhexlify(hexkey)

# NonceSize = 12 & enc_data + tag = 16
nonce = enc_data[:12]          
data = enc_data[12:] 

# decrypt + tag verify
gcm = AESGCM(key)
try:
    plaintext = gcm.decrypt(nonce, data, None)
    print(plaintext)  
except Exception as e:
    print(f"Decryption failed: {e}")
