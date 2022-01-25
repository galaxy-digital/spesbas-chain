sfc = web3.dbas.contract(abi).at("0xdbac000000000000000000000000000000000000")
sfc.lastValidatorID()
sfc.getValidatorID("0x5bc37f132a7De496f202cF6A9E897Eb75C2ba6Ec")
personal.unlockAccount("0x5bc37f132a7De496f202cF6A9E897Eb75C2ba6Ec", "BX$%$k%jSHeyA#$DJS#hHd$%ahjS", 60)
# Register your validator
tx = sfc.createValidator("0x5bc37f132a7De496f202cF6A9E897Eb75C2ba6Ec", {from:"0x5bc37f132a7De496f202cF6A9E897Eb75C2ba6Ec", value: web3.toWei("200000000.0", "dbas")})
# Check your registration transaction
dbas.getTransactionReceipt(tx)
# Get your validator id
sfc.getValidatorID("0x5bc37f132a7De496f202cF6A9E897Eb75C2ba6Ec")
