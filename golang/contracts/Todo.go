// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TodoTodoItem is an auto generated low-level Go binding around an user-defined struct.
type TodoTodoItem struct {
	Id          *big.Int
	Status      uint8
	Title       string
	Description string
	CreatedBy   common.Address
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_createdBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"}],\"name\":\"NewTodo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_updatedBy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"indexed\":true,\"internalType\":\"structTodo.TodoItem\",\"name\":\"_todo\",\"type\":\"tuple\"}],\"name\":\"UpdateTodo\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"createTodo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deleteTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTodos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noOfTodos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"updateTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611fb8806100536000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b1461014d578063bf3667741461016b578063c33926cb14610189578063dd68afb6146101a757610088565b80633380b99f1461008d5780633a1b3d31146100bd5780636e3c6738146100ed57806381e941991461011d575b600080fd5b6100a760048036038101906100a29190611304565b6101d7565b6040516100b49190611395565b60405180910390f35b6100d760048036038101906100d29190611401565b6103ca565b6040516100e49190611604565b60405180910390f35b61010760048036038101906101029190611626565b610698565b6040516101149190611604565b60405180910390f35b61013760048036038101906101329190611653565b6109d0565b6040516101449190611604565b60405180910390f35b610155610ca9565b60405161016291906116ff565b60405180910390f35b610173610ccd565b6040516101809190611395565b60405180910390f35b610191610cd3565b60405161019e9190611859565b60405180910390f35b6101c160048036038101906101bc9190611626565b610f67565b6040516101ce9190611604565b60405180910390f35b600060018060008282546101eb91906118aa565b9250508190555060006040518060a0016040528060015481526020016000600281111561021b5761021a611450565b5b81526020018581526020018481526020013373ffffffffffffffffffffffffffffffffffffffff168152509050806002600060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083600281111561029357610292611450565b5b021790555060408201518160020190816102ad9190611aea565b5060608201518160030190816102c39190611aea565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503360036000600154815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836040516103709190611bf8565b60405180910390206001543373ffffffffffffffffffffffffffffffffffffffff167f9807046998e401797e01776e2f1e343cd6ce275a52ba784c133529ec1ca9a0ac60405160405180910390a460015491505092915050565b6103d2611153565b826003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046b90611c6c565b60405180910390fd5b826002600086815260200190815260200160002060010160006101000a81548160ff021916908360028111156104ad576104ac611450565b5b0217905550600260008581526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff1660028111156104ff576104fe611450565b5b600281111561051157610510611450565b5b81526020016002820180546105259061190d565b80601f01602080910402602001604051908101604052809291908181526020018280546105519061190d565b801561059e5780601f106105735761010080835404028352916020019161059e565b820191906000526020600020905b81548152906001019060200180831161058157829003601f168201915b505050505081526020016003820180546105b79061190d565b80601f01602080910402602001604051908101604052809291908181526020018280546105e39061190d565b80156106305780601f1061060557610100808354040283529160200191610630565b820191906000526020600020905b81548152906001019060200180831161061357829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505091505092915050565b6106a0611153565b816003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610742576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073990611c6c565b60405180910390fd5b600280600085815260200190815260200160002060010160006101000a81548160ff0219169083600281111561077b5761077a611450565b5b0217905550600260008481526020019081526020016000206040516107a09190611ec0565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f0f477eb5273d53875723317bd601052ff9f7935aed8bb40f73ed508aa26b012160405160405180910390a3600260008481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff16600281111561083857610837611450565b5b600281111561084a57610849611450565b5b815260200160028201805461085e9061190d565b80601f016020809104026020016040519081016040528092919081815260200182805461088a9061190d565b80156108d75780601f106108ac576101008083540402835291602001916108d7565b820191906000526020600020905b8154815290600101906020018083116108ba57829003601f168201915b505050505081526020016003820180546108f09061190d565b80601f016020809104026020016040519081016040528092919081815260200182805461091c9061190d565b80156109695780601f1061093e57610100808354040283529160200191610969565b820191906000526020600020905b81548152906001019060200180831161094c57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050915050919050565b6109d8611153565b836003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7190611c6c565b60405180910390fd5b83600260008781526020019081526020016000206002019081610a9d9190611aea565b5082600260008781526020019081526020016000206003019081610ac19190611aea565b50600260008681526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610b0f57610b0e611450565b5b6002811115610b2157610b20611450565b5b8152602001600282018054610b359061190d565b80601f0160208091040260200160405190810160405280929190818152602001828054610b619061190d565b8015610bae5780601f10610b8357610100808354040283529160200191610bae565b820191906000526020600020905b815481529060010190602001808311610b9157829003601f168201915b50505050508152602001600382018054610bc79061190d565b80601f0160208091040260200160405190810160405280929190818152602001828054610bf39061190d565b8015610c405780601f10610c1557610100808354040283529160200191610c40565b820191906000526020600020905b815481529060010190602001808311610c2357829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509150509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6060600060015467ffffffffffffffff811115610cf357610cf26111d9565b5b604051908082528060200260200182016040528015610d2c57816020015b610d19611153565b815260200190600190039081610d115790505b5090506000600190505b6001548111610f5f576000600260008381526020019081526020016000209050806040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610d9157610d90611450565b5b6002811115610da357610da2611450565b5b8152602001600282018054610db79061190d565b80601f0160208091040260200160405190810160405280929190818152602001828054610de39061190d565b8015610e305780601f10610e0557610100808354040283529160200191610e30565b820191906000526020600020905b815481529060010190602001808311610e1357829003601f168201915b50505050508152602001600382018054610e499061190d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e759061190d565b8015610ec25780601f10610e9757610100808354040283529160200191610ec2565b820191906000526020600020905b815481529060010190602001808311610ea557829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505083600184610f2f9190611ed7565b81518110610f4057610f3f611f0b565b5b6020026020010181905250508080610f5790611f3a565b915050610d36565b508091505090565b610f6f611153565b600260008381526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610fbc57610fbb611450565b5b6002811115610fce57610fcd611450565b5b8152602001600282018054610fe29061190d565b80601f016020809104026020016040519081016040528092919081815260200182805461100e9061190d565b801561105b5780601f106110305761010080835404028352916020019161105b565b820191906000526020600020905b81548152906001019060200180831161103e57829003601f168201915b505050505081526020016003820180546110749061190d565b80601f01602080910402602001604051908101604052809291908181526020018280546110a09061190d565b80156110ed5780601f106110c2576101008083540402835291602001916110ed565b820191906000526020600020905b8154815290600101906020018083116110d057829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6040518060a00160405280600081526020016000600281111561117957611178611450565b5b81526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611211826111c8565b810181811067ffffffffffffffff821117156112305761122f6111d9565b5b80604052505050565b60006112436111aa565b905061124f8282611208565b919050565b600067ffffffffffffffff82111561126f5761126e6111d9565b5b611278826111c8565b9050602081019050919050565b82818337600083830152505050565b60006112a76112a284611254565b611239565b9050828152602081018484840111156112c3576112c26111c3565b5b6112ce848285611285565b509392505050565b600082601f8301126112eb576112ea6111be565b5b81356112fb848260208601611294565b91505092915050565b6000806040838503121561131b5761131a6111b4565b5b600083013567ffffffffffffffff811115611339576113386111b9565b5b611345858286016112d6565b925050602083013567ffffffffffffffff811115611366576113656111b9565b5b611372858286016112d6565b9150509250929050565b6000819050919050565b61138f8161137c565b82525050565b60006020820190506113aa6000830184611386565b92915050565b6113b98161137c565b81146113c457600080fd5b50565b6000813590506113d6816113b0565b92915050565b600381106113e957600080fd5b50565b6000813590506113fb816113dc565b92915050565b60008060408385031215611418576114176111b4565b5b6000611426858286016113c7565b9250506020611437858286016113ec565b9150509250929050565b61144a8161137c565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106114905761148f611450565b5b50565b60008190506114a18261147f565b919050565b60006114b182611493565b9050919050565b6114c1816114a6565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156115015780820151818401526020810190506114e6565b60008484015250505050565b6000611518826114c7565b61152281856114d2565b93506115328185602086016114e3565b61153b816111c8565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061157182611546565b9050919050565b61158181611566565b82525050565b600060a08301600083015161159f6000860182611441565b5060208301516115b260208601826114b8565b50604083015184820360408601526115ca828261150d565b915050606083015184820360608601526115e4828261150d565b91505060808301516115f96080860182611578565b508091505092915050565b6000602082019050818103600083015261161e8184611587565b905092915050565b60006020828403121561163c5761163b6111b4565b5b600061164a848285016113c7565b91505092915050565b60008060006060848603121561166c5761166b6111b4565b5b600061167a868287016113c7565b935050602084013567ffffffffffffffff81111561169b5761169a6111b9565b5b6116a7868287016112d6565b925050604084013567ffffffffffffffff8111156116c8576116c76111b9565b5b6116d4868287016112d6565b9150509250925092565b60006116e982611546565b9050919050565b6116f9816116de565b82525050565b600060208201905061171460008301846116f0565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600060a08301600083015161175e6000860182611441565b50602083015161177160208601826114b8565b5060408301518482036040860152611789828261150d565b915050606083015184820360608601526117a3828261150d565b91505060808301516117b86080860182611578565b508091505092915050565b60006117cf8383611746565b905092915050565b6000602082019050919050565b60006117ef8261171a565b6117f98185611725565b93508360208202850161180b85611736565b8060005b85811015611847578484038952815161182885826117c3565b9450611833836117d7565b925060208a0199505060018101905061180f565b50829750879550505050505092915050565b6000602082019050818103600083015261187381846117e4565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118b58261137c565b91506118c08361137c565b92508282019050808211156118d8576118d761187b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061192557607f821691505b602082108103611938576119376118de565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026119a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611963565b6119aa8683611963565b95508019841693508086168417925050509392505050565b6000819050919050565b60006119e76119e26119dd8461137c565b6119c2565b61137c565b9050919050565b6000819050919050565b611a01836119cc565b611a15611a0d826119ee565b848454611970565b825550505050565b600090565b611a2a611a1d565b611a358184846119f8565b505050565b5b81811015611a5957611a4e600082611a22565b600181019050611a3b565b5050565b601f821115611a9e57611a6f8161193e565b611a7884611953565b81016020851015611a87578190505b611a9b611a9385611953565b830182611a3a565b50505b505050565b600082821c905092915050565b6000611ac160001984600802611aa3565b1980831691505092915050565b6000611ada8383611ab0565b9150826002028217905092915050565b611af3826114c7565b67ffffffffffffffff811115611b0c57611b0b6111d9565b5b611b16825461190d565b611b21828285611a5d565b600060209050601f831160018114611b545760008415611b42578287015190505b611b4c8582611ace565b865550611bb4565b601f198416611b628661193e565b60005b82811015611b8a57848901518255600182019150602085019450602081019050611b65565b86831015611ba75784890151611ba3601f891682611ab0565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b6000611bd2826114c7565b611bdc8185611bbc565b9350611bec8185602086016114e3565b80840191505092915050565b6000611c048284611bc7565b915081905092915050565b600082825260208201905092915050565b7f596f7520617265206e6f7420616c6c6f77656420746f2064656c657465000000600082015250565b6000611c56601d83611c0f565b9150611c6182611c20565b602082019050919050565b60006020820190508181036000830152611c8581611c49565b9050919050565b60008160001c9050919050565b6000819050919050565b6000611cb6611cb183611c8c565b611c99565b9050919050565b611cc68161137c565b82525050565b6000611cd88383611cbd565b60208301905092915050565b600060ff82169050919050565b6000611d04611cff83611c8c565b611ce4565b9050919050565b611d14816114a6565b82525050565b6000611d268383611d0b565b60208301905092915050565b600081905092915050565b60008154611d4a8161190d565b611d548186611d32565b94506001821660008114611d6f5760018114611d8557611db8565b60ff198316865281151560200286019350611db8565b611d8e8561193e565b60005b83811015611db057815481890152600182019150602081019050611d91565b808801955050505b50505092915050565b6000611dcd8383611d3d565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611e08611e0383611c8c565b611dd5565b9050919050565b611e1881611566565b82525050565b6000611e2a8383611e0f565b60208301905092915050565b60008083016000808401549050611e4c81611ca3565b611e568682611ccc565b95505060018401549050611e6981611cf1565b611e738682611d1a565b95505060028401611e848682611dc1565b95505060038401611e958682611dc1565b95505060048401549050611ea881611df5565b611eb28682611e1e565b955050849250505092915050565b6000611ecc8284611e36565b915081905092915050565b6000611ee28261137c565b9150611eed8361137c565b9250828203905081811115611f0557611f0461187b565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611f458261137c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611f7757611f7661187b565b5b60018201905091905056fea26469706673582212208bf4cf579ee9b81836b311725b6f1730c10b4b2805e2ca0ebbabe5e5cacb442564736f6c63430008100033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TodoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// GetTodo is a free data retrieval call binding the contract method 0xdd68afb6.
//
// Solidity: function getTodo(uint256 id) view returns((uint256,uint8,string,string,address))
func (_Todo *TodoCaller) GetTodo(opts *bind.CallOpts, id *big.Int) (TodoTodoItem, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTodo", id)

	if err != nil {
		return *new(TodoTodoItem), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTodoItem)).(*TodoTodoItem)

	return out0, err

}

// GetTodo is a free data retrieval call binding the contract method 0xdd68afb6.
//
// Solidity: function getTodo(uint256 id) view returns((uint256,uint8,string,string,address))
func (_Todo *TodoSession) GetTodo(id *big.Int) (TodoTodoItem, error) {
	return _Todo.Contract.GetTodo(&_Todo.CallOpts, id)
}

// GetTodo is a free data retrieval call binding the contract method 0xdd68afb6.
//
// Solidity: function getTodo(uint256 id) view returns((uint256,uint8,string,string,address))
func (_Todo *TodoCallerSession) GetTodo(id *big.Int) (TodoTodoItem, error) {
	return _Todo.Contract.GetTodo(&_Todo.CallOpts, id)
}

// GetTodos is a free data retrieval call binding the contract method 0xc33926cb.
//
// Solidity: function getTodos() view returns((uint256,uint8,string,string,address)[])
func (_Todo *TodoCaller) GetTodos(opts *bind.CallOpts) ([]TodoTodoItem, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTodos")

	if err != nil {
		return *new([]TodoTodoItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTodoItem)).(*[]TodoTodoItem)

	return out0, err

}

// GetTodos is a free data retrieval call binding the contract method 0xc33926cb.
//
// Solidity: function getTodos() view returns((uint256,uint8,string,string,address)[])
func (_Todo *TodoSession) GetTodos() ([]TodoTodoItem, error) {
	return _Todo.Contract.GetTodos(&_Todo.CallOpts)
}

// GetTodos is a free data retrieval call binding the contract method 0xc33926cb.
//
// Solidity: function getTodos() view returns((uint256,uint8,string,string,address)[])
func (_Todo *TodoCallerSession) GetTodos() ([]TodoTodoItem, error) {
	return _Todo.Contract.GetTodos(&_Todo.CallOpts)
}

// NoOfTodos is a free data retrieval call binding the contract method 0xbf366774.
//
// Solidity: function noOfTodos() view returns(uint256)
func (_Todo *TodoCaller) NoOfTodos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "noOfTodos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoOfTodos is a free data retrieval call binding the contract method 0xbf366774.
//
// Solidity: function noOfTodos() view returns(uint256)
func (_Todo *TodoSession) NoOfTodos() (*big.Int, error) {
	return _Todo.Contract.NoOfTodos(&_Todo.CallOpts)
}

// NoOfTodos is a free data retrieval call binding the contract method 0xbf366774.
//
// Solidity: function noOfTodos() view returns(uint256)
func (_Todo *TodoCallerSession) NoOfTodos() (*big.Int, error) {
	return _Todo.Contract.NoOfTodos(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// CreateTodo is a paid mutator transaction binding the contract method 0x3380b99f.
//
// Solidity: function createTodo(string title, string description) returns(uint256)
func (_Todo *TodoTransactor) CreateTodo(opts *bind.TransactOpts, title string, description string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "createTodo", title, description)
}

// CreateTodo is a paid mutator transaction binding the contract method 0x3380b99f.
//
// Solidity: function createTodo(string title, string description) returns(uint256)
func (_Todo *TodoSession) CreateTodo(title string, description string) (*types.Transaction, error) {
	return _Todo.Contract.CreateTodo(&_Todo.TransactOpts, title, description)
}

// CreateTodo is a paid mutator transaction binding the contract method 0x3380b99f.
//
// Solidity: function createTodo(string title, string description) returns(uint256)
func (_Todo *TodoTransactorSession) CreateTodo(title string, description string) (*types.Transaction, error) {
	return _Todo.Contract.CreateTodo(&_Todo.TransactOpts, title, description)
}

// DeleteTodo is a paid mutator transaction binding the contract method 0x6e3c6738.
//
// Solidity: function deleteTodo(uint256 id) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactor) DeleteTodo(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "deleteTodo", id)
}

// DeleteTodo is a paid mutator transaction binding the contract method 0x6e3c6738.
//
// Solidity: function deleteTodo(uint256 id) returns((uint256,uint8,string,string,address))
func (_Todo *TodoSession) DeleteTodo(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTodo(&_Todo.TransactOpts, id)
}

// DeleteTodo is a paid mutator transaction binding the contract method 0x6e3c6738.
//
// Solidity: function deleteTodo(uint256 id) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactorSession) DeleteTodo(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.DeleteTodo(&_Todo.TransactOpts, id)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 id, uint8 status) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactor) UpdateStatus(opts *bind.TransactOpts, id *big.Int, status uint8) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "updateStatus", id, status)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 id, uint8 status) returns((uint256,uint8,string,string,address))
func (_Todo *TodoSession) UpdateStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _Todo.Contract.UpdateStatus(&_Todo.TransactOpts, id, status)
}

// UpdateStatus is a paid mutator transaction binding the contract method 0x3a1b3d31.
//
// Solidity: function updateStatus(uint256 id, uint8 status) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactorSession) UpdateStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _Todo.Contract.UpdateStatus(&_Todo.TransactOpts, id, status)
}

// UpdateTodo is a paid mutator transaction binding the contract method 0x81e94199.
//
// Solidity: function updateTodo(uint256 id, string title, string description) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactor) UpdateTodo(opts *bind.TransactOpts, id *big.Int, title string, description string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "updateTodo", id, title, description)
}

// UpdateTodo is a paid mutator transaction binding the contract method 0x81e94199.
//
// Solidity: function updateTodo(uint256 id, string title, string description) returns((uint256,uint8,string,string,address))
func (_Todo *TodoSession) UpdateTodo(id *big.Int, title string, description string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTodo(&_Todo.TransactOpts, id, title, description)
}

// UpdateTodo is a paid mutator transaction binding the contract method 0x81e94199.
//
// Solidity: function updateTodo(uint256 id, string title, string description) returns((uint256,uint8,string,string,address))
func (_Todo *TodoTransactorSession) UpdateTodo(id *big.Int, title string, description string) (*types.Transaction, error) {
	return _Todo.Contract.UpdateTodo(&_Todo.TransactOpts, id, title, description)
}

// TodoNewTodoIterator is returned from FilterNewTodo and is used to iterate over the raw logs and unpacked data for NewTodo events raised by the Todo contract.
type TodoNewTodoIterator struct {
	Event *TodoNewTodo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TodoNewTodoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoNewTodo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TodoNewTodo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TodoNewTodoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoNewTodoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoNewTodo represents a NewTodo event raised by the Todo contract.
type TodoNewTodo struct {
	CreatedBy common.Address
	Id        *big.Int
	Title     common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTodo is a free log retrieval operation binding the contract event 0x9807046998e401797e01776e2f1e343cd6ce275a52ba784c133529ec1ca9a0ac.
//
// Solidity: event NewTodo(address indexed _createdBy, uint256 indexed _id, string indexed _title)
func (_Todo *TodoFilterer) FilterNewTodo(opts *bind.FilterOpts, _createdBy []common.Address, _id []*big.Int, _title []string) (*TodoNewTodoIterator, error) {

	var _createdByRule []interface{}
	for _, _createdByItem := range _createdBy {
		_createdByRule = append(_createdByRule, _createdByItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _titleRule []interface{}
	for _, _titleItem := range _title {
		_titleRule = append(_titleRule, _titleItem)
	}

	logs, sub, err := _Todo.contract.FilterLogs(opts, "NewTodo", _createdByRule, _idRule, _titleRule)
	if err != nil {
		return nil, err
	}
	return &TodoNewTodoIterator{contract: _Todo.contract, event: "NewTodo", logs: logs, sub: sub}, nil
}

// WatchNewTodo is a free log subscription operation binding the contract event 0x9807046998e401797e01776e2f1e343cd6ce275a52ba784c133529ec1ca9a0ac.
//
// Solidity: event NewTodo(address indexed _createdBy, uint256 indexed _id, string indexed _title)
func (_Todo *TodoFilterer) WatchNewTodo(opts *bind.WatchOpts, sink chan<- *TodoNewTodo, _createdBy []common.Address, _id []*big.Int, _title []string) (event.Subscription, error) {

	var _createdByRule []interface{}
	for _, _createdByItem := range _createdBy {
		_createdByRule = append(_createdByRule, _createdByItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _titleRule []interface{}
	for _, _titleItem := range _title {
		_titleRule = append(_titleRule, _titleItem)
	}

	logs, sub, err := _Todo.contract.WatchLogs(opts, "NewTodo", _createdByRule, _idRule, _titleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoNewTodo)
				if err := _Todo.contract.UnpackLog(event, "NewTodo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTodo is a log parse operation binding the contract event 0x9807046998e401797e01776e2f1e343cd6ce275a52ba784c133529ec1ca9a0ac.
//
// Solidity: event NewTodo(address indexed _createdBy, uint256 indexed _id, string indexed _title)
func (_Todo *TodoFilterer) ParseNewTodo(log types.Log) (*TodoNewTodo, error) {
	event := new(TodoNewTodo)
	if err := _Todo.contract.UnpackLog(event, "NewTodo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TodoUpdateTodoIterator is returned from FilterUpdateTodo and is used to iterate over the raw logs and unpacked data for UpdateTodo events raised by the Todo contract.
type TodoUpdateTodoIterator struct {
	Event *TodoUpdateTodo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TodoUpdateTodoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoUpdateTodo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TodoUpdateTodo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TodoUpdateTodoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoUpdateTodoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoUpdateTodo represents a UpdateTodo event raised by the Todo contract.
type TodoUpdateTodo struct {
	UpdatedBy common.Address
	Todo      TodoTodoItem
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateTodo is a free log retrieval operation binding the contract event 0x0f477eb5273d53875723317bd601052ff9f7935aed8bb40f73ed508aa26b0121.
//
// Solidity: event UpdateTodo(address indexed _updatedBy, (uint256,uint8,string,string,address) indexed _todo)
func (_Todo *TodoFilterer) FilterUpdateTodo(opts *bind.FilterOpts, _updatedBy []common.Address, _todo []TodoTodoItem) (*TodoUpdateTodoIterator, error) {

	var _updatedByRule []interface{}
	for _, _updatedByItem := range _updatedBy {
		_updatedByRule = append(_updatedByRule, _updatedByItem)
	}
	var _todoRule []interface{}
	for _, _todoItem := range _todo {
		_todoRule = append(_todoRule, _todoItem)
	}

	logs, sub, err := _Todo.contract.FilterLogs(opts, "UpdateTodo", _updatedByRule, _todoRule)
	if err != nil {
		return nil, err
	}
	return &TodoUpdateTodoIterator{contract: _Todo.contract, event: "UpdateTodo", logs: logs, sub: sub}, nil
}

// WatchUpdateTodo is a free log subscription operation binding the contract event 0x0f477eb5273d53875723317bd601052ff9f7935aed8bb40f73ed508aa26b0121.
//
// Solidity: event UpdateTodo(address indexed _updatedBy, (uint256,uint8,string,string,address) indexed _todo)
func (_Todo *TodoFilterer) WatchUpdateTodo(opts *bind.WatchOpts, sink chan<- *TodoUpdateTodo, _updatedBy []common.Address, _todo []TodoTodoItem) (event.Subscription, error) {

	var _updatedByRule []interface{}
	for _, _updatedByItem := range _updatedBy {
		_updatedByRule = append(_updatedByRule, _updatedByItem)
	}
	var _todoRule []interface{}
	for _, _todoItem := range _todo {
		_todoRule = append(_todoRule, _todoItem)
	}

	logs, sub, err := _Todo.contract.WatchLogs(opts, "UpdateTodo", _updatedByRule, _todoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoUpdateTodo)
				if err := _Todo.contract.UnpackLog(event, "UpdateTodo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTodo is a log parse operation binding the contract event 0x0f477eb5273d53875723317bd601052ff9f7935aed8bb40f73ed508aa26b0121.
//
// Solidity: event UpdateTodo(address indexed _updatedBy, (uint256,uint8,string,string,address) indexed _todo)
func (_Todo *TodoFilterer) ParseUpdateTodo(log types.Log) (*TodoUpdateTodo, error) {
	event := new(TodoUpdateTodo)
	if err := _Todo.contract.UnpackLog(event, "UpdateTodo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
