package statetransfer

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/nitro/solgen/go/classicgen"
)

var ArbosAddressTable = common.HexToAddress("0x0000000000000000000000000000000000000066")

func openClassicArbAddressTable(client *ethclient.Client) (*classicgen.ArbAddressTableCaller, error) {
	return classicgen.NewArbAddressTableCaller(ArbosAddressTable, client)
}

func scanAndCopyAddressTable(reader AddressReader, writer *JsonListWriter) (uint64, common.Address, error) {
	length := uint64(0)
	address := &common.Address{}
	for reader.More() {
		var err error
		address, err = reader.GetNext()
		if err != nil {
			return length, common.Address{}, err
		}
		err = writer.Write(address)
		if err != nil {
			return length, common.Address{}, err
		}
		AddressSeen(*address)
		length += 1
	}
	return length, *address, nil
}

func verifyAndFillAddressTable(ethClient *ethclient.Client, callopts *bind.CallOpts, prevLength uint64, lastAddress common.Address, writer *JsonListWriter) error {
	classicArbAddressTable, err := openClassicArbAddressTable(ethClient)
	if err != nil {
		return err
	}
	if prevLength > 0 {
		// sanity test for reorgs, etc.. assume all is o.k. if last is o.k.
		lastIndex := big.NewInt(int64(prevLength) - 1)
		foundAddress, err := classicArbAddressTable.LookupIndex(callopts, lastIndex)
		if err != nil {
			return err
		}
		if foundAddress != lastAddress {
			return fmt.Errorf("addresstable index %v expected %s found %s", lastIndex, lastAddress, foundAddress)
		}
	}

	numAddresses, err := classicArbAddressTable.Size(callopts)
	if err != nil {
		return fmt.Errorf("addresstable.Size error: %w", err)
	}
	numAddressesInt := numAddresses.Int64()
	if (!numAddresses.IsInt64()) || numAddressesInt < int64(prevLength) {
		return fmt.Errorf("addresstable size %v expected at least %v", numAddresses, prevLength)
	}
	fmt.Println("current Num of addresses ", numAddresses)

	for cIndex := int64(prevLength); cIndex < numAddressesInt; cIndex++ {
		cAddress, err := classicArbAddressTable.LookupIndex(callopts, big.NewInt(cIndex))
		if err != nil {
			return err
		}
		err = writer.Write(cAddress)
		if err != nil {
			return err
		}
		AddressSeen(cAddress)
	}
	return nil
}
