# pycrypto-aescfb128.py
# command line application to encrypt/decrypt messages
# usage: crypto.py encrypt|decrypt payload

import sys
from Crypto.Cipher import AES
from base64 import b64encode, b64decode

BLOCK_SIZE = 16

# pad adds \x00 to end of input up to a multiple of BLOCk_SIZE
# returns padded string and padding length
def pad(payload):
    padding_len = len(payload) % BLOCK_SIZE
    padded_len = len(payload) + padding_len
    return payload.ljust(padded_len, "\x00"), padding_len

# unpad removes last n bytes from the payload
def unpad(payload, n):
    if n > len(payload):
        return ""
    return payload[:len(payload)-n]

# encrypt encrypts payload using AES-CFB(key, iv) and returns it in base64
def encrypt(payload, key, iv):
    aes = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
    padded, padding_len = pad(payload)
    encrypted = aes.encrypt(padded)
    encrypted = unpad(encrypted, padding_len)
    return b64encode(encrypted)

# decrypt decodes the payload from base64 and decrypts with AES-CFB(key, iv)
def decrypt(payload, key, iv):
    decoded = b64decode(payload)
    padded, padding_len = pad(decoded)
    aes = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
    decrypted = aes.decrypt(padded)
    return unpad(decrypted, padding_len)


def main():
    if len(sys.argv) != 3:
        print "invalid arguments - usage: crypto.py encrypt|decrypt payload"
        sys.exit(2)

    key = "0123456789012345"
    iv = "9876543210987654"

    # read argument 2, it must encrypt or decrypt
    action = sys.argv[1]
    payload = sys.argv[2]

    if action == "encrypt":
        print encrypt(payload, key, iv)
    elif action == "decrypt":
        print decrypt(payload, key, iv)
    else:
        print "first argument can only be encrypt or decrypt"

if __name__ == "__main__":
    main()
