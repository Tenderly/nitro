import ethers, {utils, providers, Wallet} from "ethers"

const nitroL2FundedPK =
  '0xe887f7d17d07cc7b8004053fb8826f6657084e88904bb61590e498ca04704cf2'

const shouldFail = true

const main = async () => {
    const l2Provider = new providers.JsonRpcProvider('http://127.0.0.1:7545')
    const l2Funder = new Wallet(nitroL2FundedPK).connect(l2Provider)

    const balance = await l2Provider.getBalance(await l2Funder.getAddress())
    console.log("got funds", balance)

    const to = "0x27Cc3B0a2FA10218B116799380410AD7B93D9314"
    const value = utils.parseEther('1')

    const expected = await l2Funder.estimateGas({ to, value })
    console.log("got expected", expected)

    try {
        const l2FundTx = await l2Funder.sendTransaction({ to, value, gasLimit: expected })
        await l2FundTx.wait()
        console.log("gas estimation works")
    } catch(e) {
        console.log("Doesn't work with gas limit", expected.toNumber())
    }
    try {
        const hardcode = 300000
        const l2FundTx = await l2Funder.sendTransaction({ to, value, gasLimit: hardcode })
        await l2FundTx.wait()
        console.log("hard code works with gas limit", hardcode)
    }catch (e) {
        console.log("error here too")
    }
}


main()
.then(() => console.log("Done"))
.catch((err) => console.error(err));
