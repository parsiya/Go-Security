# pycrypto1
from Crypto.Cipher import AES
from base64 import b64decode

key = "0123456789012345"
iv = "9876543210987654"
ciphertext = "jaJW8QJbKqHEg5zyFURe2o565/wDVca9"

# decode from base64
decoded = b64decode(ciphertext)

# encrypt with AES-CFB
aes = AES.new(key, AES.MODE_CFB, iv)
decrypted = aes.decrypt(decoded)

print decrypted