import hashlib
import qrcode
from PIL import Image
import time

class AdanichainBlock:
    def __init__(self, info, previous_hash, transaction):
        self.info = info
        self.transaction = transaction
        self.previous_hash = previous_hash
        string_to_hash = "".join(transaction) + previous_hash
        self.block_hash = hashlib.sha256(string_to_hash.encode()).hexdigest()

class Oil:
    def __init__(self, name, variety, refinery, size, production_date, expiry_date):
        self.name = name
        self.variety = variety
        self.refinery = refinery
        self.size = size
        self.production_date = production_date
        self.expiry_date = expiry_date
        self.info = name + ";" + variety + ";" + refinery + ";" + size + ";" + production_date + ";" + expiry_date

then_time = time.time()

refinery_id = hashlib.sha256("PUNE REFINERY".encode()).hexdigest()
first_item = Oil("palm_oilseeds", "Grade 1", str(refinery_id), "10000 Kg", "2020-07-13", "2020-08-3")

print("Item info: ")
print(first_item.info + "\n")

next_supply_chain_point1 = hashlib.sha256("distribution house 1".encode()).hexdigest()

genesis_block = AdanichainBlock(first_item.info, first_item.info, [str(next_supply_chain_point1)])
print("Item info: " + genesis_block.info + "\n")
print("First block hash: ")
print(genesis_block.block_hash + "\n\n")

next_supply_chain_point2 = hashlib.sha256("shop 1".encode()).hexdigest()
second_block = AdanichainBlock(first_item.info, genesis_block.block_hash, [str(next_supply_chain_point2)])
print("Item info: " + second_block.info + "\n")
print("Second block hash: ")
print(second_block.block_hash + "\n\n")

qr_code = qrcode.QRCode(
            version = 30,
            error_correction = qrcode.constants.ERROR_CORRECT_H,
            box_size = 4
            )
qr_code.add_data(second_block.info + ";" + second_block.block_hash)
embedded_data = second_block.info + ";" + second_block.block_hash
print("[!] Embedded data in QR code: \n" + embedded_data + "\n")
qr_code.make(fit=True)
img = qr_code.make_image(fill_color="#451bfe", back_color="white").convert('RGB')

img.save("qrcode_sample.png")

#execution time
time_def = time.time() - then_time
print("[!] Execution time: " + str(time_def))

