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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"createTodo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deleteTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTodos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noOfTodos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"updateTodo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"enumTodo.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"internalType\":\"structTodo.TodoItem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611c53806100536000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b1461014d578063bf3667741461016b578063c33926cb14610189578063dd68afb6146101a757610088565b80633380b99f1461008d5780633a1b3d31146100bd5780636e3c6738146100ed57806381e941991461011d575b600080fd5b6100a760048036038101906100a2919061123d565b6101d7565b6040516100b491906112ce565b60405180910390f35b6100d760048036038101906100d2919061133a565b61036e565b6040516100e4919061153d565b60405180910390f35b6101076004803603810190610102919061155f565b61063c565b604051610114919061153d565b60405180910390f35b6101376004803603810190610132919061158c565b610909565b604051610144919061153d565b60405180910390f35b610155610be2565b6040516101629190611638565b60405180910390f35b610173610c06565b60405161018091906112ce565b60405180910390f35b610191610c0c565b60405161019e9190611792565b60405180910390f35b6101c160048036038101906101bc919061155f565b610ea0565b6040516101ce919061153d565b60405180910390f35b600060018060008282546101eb91906117e3565b9250508190555060006040518060a0016040528060015481526020016000600281111561021b5761021a611389565b5b81526020018581526020018481526020013373ffffffffffffffffffffffffffffffffffffffff168152509050806002600060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548160ff0219169083600281111561029357610292611389565b5b021790555060408201518160020190816102ad9190611a23565b5060608201518160030190816102c39190611a23565b5060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050503360036000600154815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060015491505092915050565b61037661108c565b826003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040f90611b52565b60405180910390fd5b826002600086815260200190815260200160002060010160006101000a81548160ff0219169083600281111561045157610450611389565b5b0217905550600260008581526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff1660028111156104a3576104a2611389565b5b60028111156104b5576104b4611389565b5b81526020016002820180546104c990611846565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611846565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050815260200160038201805461055b90611846565b80601f016020809104026020016040519081016040528092919081815260200182805461058790611846565b80156105d45780601f106105a9576101008083540402835291602001916105d4565b820191906000526020600020905b8154815290600101906020018083116105b757829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505091505092915050565b61064461108c565b816003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106dd90611b52565b60405180910390fd5b600280600085815260200190815260200160002060010160006101000a81548160ff0219169083600281111561071f5761071e611389565b5b0217905550600260008481526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff16600281111561077157610770611389565b5b600281111561078357610782611389565b5b815260200160028201805461079790611846565b80601f01602080910402602001604051908101604052809291908181526020018280546107c390611846565b80156108105780601f106107e557610100808354040283529160200191610810565b820191906000526020600020905b8154815290600101906020018083116107f357829003601f168201915b5050505050815260200160038201805461082990611846565b80601f016020809104026020016040519081016040528092919081815260200182805461085590611846565b80156108a25780601f10610877576101008083540402835291602001916108a2565b820191906000526020600020905b81548152906001019060200180831161088557829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050915050919050565b61091161108c565b836003600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109aa90611b52565b60405180910390fd5b836002600087815260200190815260200160002060020190816109d69190611a23565b50826002600087815260200190815260200160002060030190816109fa9190611a23565b50600260008681526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610a4857610a47611389565b5b6002811115610a5a57610a59611389565b5b8152602001600282018054610a6e90611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610a9a90611846565b8015610ae75780601f10610abc57610100808354040283529160200191610ae7565b820191906000526020600020905b815481529060010190602001808311610aca57829003601f168201915b50505050508152602001600382018054610b0090611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2c90611846565b8015610b795780601f10610b4e57610100808354040283529160200191610b79565b820191906000526020600020905b815481529060010190602001808311610b5c57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509150509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6060600060015467ffffffffffffffff811115610c2c57610c2b611112565b5b604051908082528060200260200182016040528015610c6557816020015b610c5261108c565b815260200190600190039081610c4a5790505b5090506000600190505b6001548111610e98576000600260008381526020019081526020016000209050806040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610cca57610cc9611389565b5b6002811115610cdc57610cdb611389565b5b8152602001600282018054610cf090611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1c90611846565b8015610d695780601f10610d3e57610100808354040283529160200191610d69565b820191906000526020600020905b815481529060010190602001808311610d4c57829003601f168201915b50505050508152602001600382018054610d8290611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610dae90611846565b8015610dfb5780601f10610dd057610100808354040283529160200191610dfb565b820191906000526020600020905b815481529060010190602001808311610dde57829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505083600184610e689190611b72565b81518110610e7957610e78611ba6565b5b6020026020010181905250508080610e9090611bd5565b915050610c6f565b508091505090565b610ea861108c565b600260008381526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900460ff166002811115610ef557610ef4611389565b5b6002811115610f0757610f06611389565b5b8152602001600282018054610f1b90611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610f4790611846565b8015610f945780601f10610f6957610100808354040283529160200191610f94565b820191906000526020600020905b815481529060010190602001808311610f7757829003601f168201915b50505050508152602001600382018054610fad90611846565b80601f0160208091040260200160405190810160405280929190818152602001828054610fd990611846565b80156110265780601f10610ffb57610100808354040283529160200191611026565b820191906000526020600020905b81548152906001019060200180831161100957829003601f168201915b505050505081526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6040518060a0016040528060008152602001600060028111156110b2576110b1611389565b5b81526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61114a82611101565b810181811067ffffffffffffffff8211171561116957611168611112565b5b80604052505050565b600061117c6110e3565b90506111888282611141565b919050565b600067ffffffffffffffff8211156111a8576111a7611112565b5b6111b182611101565b9050602081019050919050565b82818337600083830152505050565b60006111e06111db8461118d565b611172565b9050828152602081018484840111156111fc576111fb6110fc565b5b6112078482856111be565b509392505050565b600082601f830112611224576112236110f7565b5b81356112348482602086016111cd565b91505092915050565b60008060408385031215611254576112536110ed565b5b600083013567ffffffffffffffff811115611272576112716110f2565b5b61127e8582860161120f565b925050602083013567ffffffffffffffff81111561129f5761129e6110f2565b5b6112ab8582860161120f565b9150509250929050565b6000819050919050565b6112c8816112b5565b82525050565b60006020820190506112e360008301846112bf565b92915050565b6112f2816112b5565b81146112fd57600080fd5b50565b60008135905061130f816112e9565b92915050565b6003811061132257600080fd5b50565b60008135905061133481611315565b92915050565b60008060408385031215611351576113506110ed565b5b600061135f85828601611300565b925050602061137085828601611325565b9150509250929050565b611383816112b5565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600381106113c9576113c8611389565b5b50565b60008190506113da826113b8565b919050565b60006113ea826113cc565b9050919050565b6113fa816113df565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561143a57808201518184015260208101905061141f565b60008484015250505050565b600061145182611400565b61145b818561140b565b935061146b81856020860161141c565b61147481611101565b840191505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114aa8261147f565b9050919050565b6114ba8161149f565b82525050565b600060a0830160008301516114d8600086018261137a565b5060208301516114eb60208601826113f1565b50604083015184820360408601526115038282611446565b9150506060830151848203606086015261151d8282611446565b915050608083015161153260808601826114b1565b508091505092915050565b6000602082019050818103600083015261155781846114c0565b905092915050565b600060208284031215611575576115746110ed565b5b600061158384828501611300565b91505092915050565b6000806000606084860312156115a5576115a46110ed565b5b60006115b386828701611300565b935050602084013567ffffffffffffffff8111156115d4576115d36110f2565b5b6115e08682870161120f565b925050604084013567ffffffffffffffff811115611601576116006110f2565b5b61160d8682870161120f565b9150509250925092565b60006116228261147f565b9050919050565b61163281611617565b82525050565b600060208201905061164d6000830184611629565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600060a083016000830151611697600086018261137a565b5060208301516116aa60208601826113f1565b50604083015184820360408601526116c28282611446565b915050606083015184820360608601526116dc8282611446565b91505060808301516116f160808601826114b1565b508091505092915050565b6000611708838361167f565b905092915050565b6000602082019050919050565b600061172882611653565b611732818561165e565b9350836020820285016117448561166f565b8060005b85811015611780578484038952815161176185826116fc565b945061176c83611710565b925060208a01995050600181019050611748565b50829750879550505050505092915050565b600060208201905081810360008301526117ac818461171d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117ee826112b5565b91506117f9836112b5565b9250828201905080821115611811576118106117b4565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061185e57607f821691505b60208210810361187157611870611817565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118d97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261189c565b6118e3868361189c565b95508019841693508086168417925050509392505050565b6000819050919050565b600061192061191b611916846112b5565b6118fb565b6112b5565b9050919050565b6000819050919050565b61193a83611905565b61194e61194682611927565b8484546118a9565b825550505050565b600090565b611963611956565b61196e818484611931565b505050565b5b818110156119925761198760008261195b565b600181019050611974565b5050565b601f8211156119d7576119a881611877565b6119b18461188c565b810160208510156119c0578190505b6119d46119cc8561188c565b830182611973565b50505b505050565b600082821c905092915050565b60006119fa600019846008026119dc565b1980831691505092915050565b6000611a1383836119e9565b9150826002028217905092915050565b611a2c82611400565b67ffffffffffffffff811115611a4557611a44611112565b5b611a4f8254611846565b611a5a828285611996565b600060209050601f831160018114611a8d5760008415611a7b578287015190505b611a858582611a07565b865550611aed565b601f198416611a9b86611877565b60005b82811015611ac357848901518255600182019150602085019450602081019050611a9e565b86831015611ae05784890151611adc601f8916826119e9565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f596f7520617265206e6f7420616c6c6f77656420746f2064656c657465000000600082015250565b6000611b3c601d83611af5565b9150611b4782611b06565b602082019050919050565b60006020820190508181036000830152611b6b81611b2f565b9050919050565b6000611b7d826112b5565b9150611b88836112b5565b9250828203905081811115611ba057611b9f6117b4565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611be0826112b5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c1257611c116117b4565b5b60018201905091905056fea2646970667358221220dba241c3c8203508dff97eced85624fb37d92efa9dec4a13ee9cc1ac3816984764736f6c63430008100033",
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
