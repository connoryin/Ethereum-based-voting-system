const web3 = require("web3")

async function main() {
    const poll = await ethers.getContractFactory("Poll");

    // Start deployment, returning a promise that resolves to a contract object
    const contract = await poll.deploy(web3.utils.padLeft(web3.utils.asciiToHex("Connor"), 64), web3.utils.padLeft(web3.utils.asciiToHex("Test"), 64), web3.utils.padLeft(web3.utils.asciiToHex("03/10/2022"), 64), web3.utils.padLeft(web3.utils.asciiToHex("04/10/2022"), 64), 5, 3, [1, 2, 3, 4, 5], ["Connor", "Ben", "Alice", "Bob", "Kate"].map((arg) => web3.utils.padLeft(web3.utils.asciiToHex(arg), 64)));
    console.log("Contract deployed to address:", contract.address);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
