const sigUtil = require('eth-sig-util');
const ethUtil = require('ethereumjs-util');
const Web3 = require('web3');

// Replace these values with your own data
const privateKey = '0x0123456789012345678901234567890123456789012345678901234567890123';
const providerUrl = 'https://mainnet.infura.io/v3/YOUR-PROJECT-ID';
const brokerAddress = '0x0123456789012345678901234567890123456789';
const agreementId = 1;
const amendment = {
    ipfsHash: '0x0123456789012345678901234567890123456789012345678901234567890123',
    pricePerMinute: 100,
};

const web3 = new Web3(new Web3.providers.HttpProvider(providerUrl));

async function main() {
    const chainId = await web3.eth.getChainId();
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const domain = [
        { name: 'name', type: 'string', value: 'Broker' },
        { name: 'version', type: 'string', value: '1' },
        { name: 'chainId', type: 'uint256', value: chainId },
        { name: 'verifyingContract', type: 'address', value: brokerAddress },
    ];

    const amendmentTypes = [
        { name: 'agreementId', type: 'uint256' },
        { name: 'ipfsHash', type: 'bytes32' },
        { name: 'pricePerMinute', type: 'uint256' },
    ];

    const amendmentData = {
        agreementId: agreementId,
        ipfsHash: amendment.ipfsHash,
        pricePerMinute: amendment.pricePerMinute,
    };

    const typedData = {
        types: {
            EIP712Domain: domain,
            Amendment: amendmentTypes,
        },
        domain: domain,
        primaryType: 'Amendment',
        message: amendmentData,
    };

    const signature = sigUtil.signTypedData_v4(Buffer.from(account.privateKey.slice(2), 'hex'), { data: typedData });
    console.log(`Signature: ${signature}`);
}

main()
    .then(() => console.log('Done'))
    .catch((error) => console.error(error));
