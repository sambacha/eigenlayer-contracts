// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StrategyBaseTVLLimits

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
	_ = abi.ConvertType
)

// StrategyBaseTVLLimitsMetaData contains all meta data concerning the StrategyBaseTVLLimits contract.
var StrategyBaseTVLLimitsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"newShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTVLLimits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_maxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_underlyingToken\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxPerDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTotalDeposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTVLLimits\",\"inputs\":[{\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newMaxTotalDeposits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxPerDepositUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxTotalDepositsUpdated\",\"inputs\":[{\"name\":\"previousValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001ec738038062001ec7833981016040819052620000349162000116565b6001600160a01b038116608052806200004c62000054565b505062000148565b600054610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000114576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6000602082840312156200012957600080fd5b81516001600160a01b03811681146200014157600080fd5b9392505050565b608051611d4e6200017960003960008181610216015281816107a901528181610b610152610c2c0152611d4e6000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80635c975abb116100de578063ab5921e111610097578063df6fadc111610071578063df6fadc114610366578063e3dae51c14610381578063f3e7387514610394578063fabc1cbc146103a757600080fd5b8063ab5921e11461032b578063ce7c2ac214610340578063d9caed121461035357600080fd5b80635c975abb146102c857806361b01b5d146102d05780637a8b2637146102d9578063886f1195146102ec5780638c871019146103055780638f6a62401461031857600080fd5b80633a98ef391161014b578063485cc95511610125578063485cc9551461026b578063553ca5f81461027e578063595c6a67146102915780635ac86ab71461029957600080fd5b80633a98ef391461023857806343fe08b01461024f57806347e7ef241461025857600080fd5b8063019e27291461019357806310d67a2f146101a857806311c70c9d146101bb578063136439dd146101ce5780632495a599146101e157806339b70e3814610211575b600080fd5b6101a66101a13660046118fd565b6103ba565b005b6101a66101b6366004611947565b61049d565b6101a66101c9366004611964565b610550565b6101a66101dc366004611986565b610605565b6032546101f4906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101f47f000000000000000000000000000000000000000000000000000000000000000081565b61024160335481565b604051908152602001610208565b61024160645481565b61024161026636600461199f565b610749565b6101a66102793660046119cb565b610907565b61024161028c366004611947565b6109d5565b6101a66109e9565b6102b86102a7366004611a13565b6001805460ff9092161b9081161490565b6040519015158152602001610208565b600154610241565b61024160655481565b6102416102e7366004611986565b610ab5565b6000546101f4906201000090046001600160a01b031681565b610241610313366004611986565b610b00565b610241610326366004611947565b610b0b565b610333610b19565b6040516102089190611a60565b61024161034e366004611947565b610b39565b6101a6610361366004611a93565b610bce565b60645460655460408051928352602083019190915201610208565b61024161038f366004611986565b610db4565b6102416103a2366004611986565b610ded565b6101a66103b5366004611986565b610df8565b600054610100900460ff16158080156103da5750600054600160ff909116105b806103f45750303b1580156103f4575060005460ff166001145b6104195760405162461bcd60e51b815260040161041090611ad4565b60405180910390fd5b6000805460ff19166001179055801561043c576000805461ff0019166101001790555b6104468585610f54565b6104508383611061565b8015610496576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105149190611b22565b6001600160a01b0316336001600160a01b0316146105445760405162461bcd60e51b815260040161041090611b3f565b61054d816111b4565b50565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c79190611b22565b6001600160a01b0316336001600160a01b0316146105f75760405162461bcd60e51b815260040161041090611b3f565b6106018282610f54565b5050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610652573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106769190611b89565b6106925760405162461bcd60e51b815260040161041090611bab565b6001548181161461070b5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610410565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b6001805460009182918116141561079e5760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b6044820152606401610410565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108165760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152606401610410565b61082084846112b9565b60335460006108316103e883611c09565b905060006103e861084061139b565b61084a9190611c09565b905060006108588783611c21565b9050806108658489611c38565b61086f9190611c57565b9550856108d55760405162461bcd60e51b815260206004820152602e60248201527f5374726174656779426173652e6465706f7369743a206e65775368617265732060448201526d63616e6e6f74206265207a65726f60901b6064820152608401610410565b6108df8685611c09565b60338190556108fc9083906108f7906103e890611c09565b61140d565b505050505092915050565b600054610100900460ff16158080156109275750600054600160ff909116105b806109415750303b158015610941575060005460ff166001145b61095d5760405162461bcd60e51b815260040161041090611ad4565b6000805460ff191660011790558015610980576000805461ff0019166101001790555b61098a8383611061565b80156109d0576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60006109e36102e783610b39565b92915050565b60005460405163237dfb4760e11b8152336004820152620100009091046001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610a36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5a9190611b89565b610a765760405162461bcd60e51b815260040161041090611bab565b600019600181905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b6000806103e8603354610ac89190611c09565b905060006103e8610ad761139b565b610ae19190611c09565b905081610aee8583611c38565b610af89190611c57565b949350505050565b60006109e382610db4565b60006109e36103a283610b39565b60606040518060800160405280604d8152602001611ccc604d9139905090565b604051633d3f06c960e11b81526001600160a01b0382811660048301523060248301526000917f000000000000000000000000000000000000000000000000000000000000000090911690637a7e0d9290604401602060405180830381865afa158015610baa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e39190611c79565b6001805460029081161415610c215760405162461bcd60e51b815260206004820152601960248201527814185d5cd8589b194e881a5b99195e081a5cc81c185d5cd959603a1b6044820152606401610410565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c995760405162461bcd60e51b815260206004820181905260248201527f5374726174656779426173652e6f6e6c7953747261746567794d616e616765726044820152606401610410565b610ca4848484611459565b60335480831115610d335760405162461bcd60e51b815260206004820152604d60248201527f5374726174656779426173652e77697468647261773a20616d6f756e7453686160448201527f726573206d757374206265206c657373207468616e206f7220657175616c207460648201526c6f20746f74616c53686172657360981b608482015260a401610410565b6000610d416103e883611c09565b905060006103e8610d5061139b565b610d5a9190611c09565b9050600082610d698784611c38565b610d739190611c57565b9050610d7f8685611c21565b603355610d9f610d8f8284611c21565b6103e86033546108f79190611c09565b610daa8888836114dc565b5050505050505050565b6000806103e8603354610dc79190611c09565b905060006103e8610dd661139b565b610de09190611c09565b905080610aee8386611c38565b60006109e382610ab5565b600060029054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e4b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6f9190611b22565b6001600160a01b0316336001600160a01b031614610e9f5760405162461bcd60e51b815260040161041090611b3f565b600154198119600154191614610f1d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610410565b600181905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161073e565b60645460408051918252602082018490527ff97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5910160405180910390a160655460408051918252602082018390527f6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452910160405180910390a1808211156110565760405162461bcd60e51b815260206004820152604b60248201527f53747261746567794261736554564c4c696d6974732e5f73657454564c4c696d60448201527f6974733a206d61785065724465706f7369742065786365656473206d6178546f60648201526a74616c4465706f7369747360a81b608482015260a401610410565b606491909155606555565b600054610100900460ff166110cc5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b6064820152608401610410565b603280546001600160a01b0319166001600160a01b0384161790556110f28160006114f0565b7f1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507603260009054906101000a90046001600160a01b0316836001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015611167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118b9190611c92565b604080516001600160a01b03909316835260ff9091166020830152015b60405180910390a15050565b6001600160a01b0381166112425760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610410565b600054604080516001600160a01b03620100009093048316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b6064548111156113235760405162461bcd60e51b815260206004820152602f60248201527f53747261746567794261736554564c4c696d6974733a206d617820706572206460448201526e195c1bdcda5d08195e18d959591959608a1b6064820152608401610410565b60655461132e61139b565b11156113915760405162461bcd60e51b815260206004820152602c60248201527f53747261746567794261736554564c4c696d6974733a206d6178206465706f7360448201526b1a5d1cc8195e18d95959195960a21b6064820152608401610410565b61060182826115dc565b6032546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa1580156113e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114089190611c79565b905090565b7fd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be88161144184670de0b6b3a7640000611c38565b61144b9190611c57565b6040519081526020016111a8565b6032546001600160a01b038381169116146109d05760405162461bcd60e51b815260206004820152603b60248201527f5374726174656779426173652e77697468647261773a2043616e206f6e6c792060448201527f77697468647261772074686520737472617465677920746f6b656e00000000006064820152608401610410565b6109d06001600160a01b0383168483611658565b6000546201000090046001600160a01b031615801561151757506001600160a01b03821615155b6115995760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610410565b600181905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2610601826111b4565b6032546001600160a01b038381169116146106015760405162461bcd60e51b815260206004820152603660248201527f5374726174656779426173652e6465706f7369743a2043616e206f6e6c79206460448201527532b837b9b4ba103ab73232b9363cb4b733aa37b5b2b760511b6064820152608401610410565b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526109d0928692916000916116e8918516908490611765565b8051909150156109d057808060200190518101906117069190611b89565b6109d05760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610410565b6060611774848460008561177e565b90505b9392505050565b6060824710156117df5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610410565b6001600160a01b0385163b6118365760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610410565b600080866001600160a01b031685876040516118529190611caf565b60006040518083038185875af1925050503d806000811461188f576040519150601f19603f3d011682016040523d82523d6000602084013e611894565b606091505b50915091506118a48282866118af565b979650505050505050565b606083156118be575081611777565b8251156118ce5782518084602001fd5b8160405162461bcd60e51b81526004016104109190611a60565b6001600160a01b038116811461054d57600080fd5b6000806000806080858703121561191357600080fd5b8435935060208501359250604085013561192c816118e8565b9150606085013561193c816118e8565b939692955090935050565b60006020828403121561195957600080fd5b8135611777816118e8565b6000806040838503121561197757600080fd5b50508035926020909101359150565b60006020828403121561199857600080fd5b5035919050565b600080604083850312156119b257600080fd5b82356119bd816118e8565b946020939093013593505050565b600080604083850312156119de57600080fd5b82356119e9816118e8565b915060208301356119f9816118e8565b809150509250929050565b60ff8116811461054d57600080fd5b600060208284031215611a2557600080fd5b813561177781611a04565b60005b83811015611a4b578181015183820152602001611a33565b83811115611a5a576000848401525b50505050565b6020815260008251806020840152611a7f816040850160208701611a30565b601f01601f19169190910160400192915050565b600080600060608486031215611aa857600080fd5b8335611ab3816118e8565b92506020840135611ac3816118e8565b929592945050506040919091013590565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b600060208284031215611b3457600080fd5b8151611777816118e8565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215611b9b57600080fd5b8151801515811461177757600080fd5b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b60008219821115611c1c57611c1c611bf3565b500190565b600082821015611c3357611c33611bf3565b500390565b6000816000190483118215151615611c5257611c52611bf3565b500290565b600082611c7457634e487b7160e01b600052601260045260246000fd5b500490565b600060208284031215611c8b57600080fd5b5051919050565b600060208284031215611ca457600080fd5b815161177781611a04565b60008251611cc1818460208701611a30565b919091019291505056fe4261736520537472617465677920696d706c656d656e746174696f6e20746f20696e68657269742066726f6d20666f72206d6f726520636f6d706c657820696d706c656d656e746174696f6e73a2646970667358221220c4d8a1d1ceb7ae63d68ad92ee64fadf934c6102735251742fa8fae81b99c19c564736f6c634300080c0033",
}

// StrategyBaseTVLLimitsABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyBaseTVLLimitsMetaData.ABI instead.
var StrategyBaseTVLLimitsABI = StrategyBaseTVLLimitsMetaData.ABI

// StrategyBaseTVLLimitsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrategyBaseTVLLimitsMetaData.Bin instead.
var StrategyBaseTVLLimitsBin = StrategyBaseTVLLimitsMetaData.Bin

// DeployStrategyBaseTVLLimits deploys a new Ethereum contract, binding an instance of StrategyBaseTVLLimits to it.
func DeployStrategyBaseTVLLimits(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address) (common.Address, *types.Transaction, *StrategyBaseTVLLimits, error) {
	parsed, err := StrategyBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrategyBaseTVLLimitsBin), backend, _strategyManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrategyBaseTVLLimits{StrategyBaseTVLLimitsCaller: StrategyBaseTVLLimitsCaller{contract: contract}, StrategyBaseTVLLimitsTransactor: StrategyBaseTVLLimitsTransactor{contract: contract}, StrategyBaseTVLLimitsFilterer: StrategyBaseTVLLimitsFilterer{contract: contract}}, nil
}

// StrategyBaseTVLLimits is an auto generated Go binding around an Ethereum contract.
type StrategyBaseTVLLimits struct {
	StrategyBaseTVLLimitsCaller     // Read-only binding to the contract
	StrategyBaseTVLLimitsTransactor // Write-only binding to the contract
	StrategyBaseTVLLimitsFilterer   // Log filterer for contract events
}

// StrategyBaseTVLLimitsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyBaseTVLLimitsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyBaseTVLLimitsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategyBaseTVLLimitsSession struct {
	Contract     *StrategyBaseTVLLimits // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StrategyBaseTVLLimitsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyBaseTVLLimitsCallerSession struct {
	Contract *StrategyBaseTVLLimitsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// StrategyBaseTVLLimitsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyBaseTVLLimitsTransactorSession struct {
	Contract     *StrategyBaseTVLLimitsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StrategyBaseTVLLimitsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsRaw struct {
	Contract *StrategyBaseTVLLimits // Generic contract binding to access the raw methods on
}

// StrategyBaseTVLLimitsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsCallerRaw struct {
	Contract *StrategyBaseTVLLimitsCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyBaseTVLLimitsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyBaseTVLLimitsTransactorRaw struct {
	Contract *StrategyBaseTVLLimitsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategyBaseTVLLimits creates a new instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimits(address common.Address, backend bind.ContractBackend) (*StrategyBaseTVLLimits, error) {
	contract, err := bindStrategyBaseTVLLimits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimits{StrategyBaseTVLLimitsCaller: StrategyBaseTVLLimitsCaller{contract: contract}, StrategyBaseTVLLimitsTransactor: StrategyBaseTVLLimitsTransactor{contract: contract}, StrategyBaseTVLLimitsFilterer: StrategyBaseTVLLimitsFilterer{contract: contract}}, nil
}

// NewStrategyBaseTVLLimitsCaller creates a new read-only instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsCaller(address common.Address, caller bind.ContractCaller) (*StrategyBaseTVLLimitsCaller, error) {
	contract, err := bindStrategyBaseTVLLimits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsCaller{contract: contract}, nil
}

// NewStrategyBaseTVLLimitsTransactor creates a new write-only instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyBaseTVLLimitsTransactor, error) {
	contract, err := bindStrategyBaseTVLLimits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsTransactor{contract: contract}, nil
}

// NewStrategyBaseTVLLimitsFilterer creates a new log filterer instance of StrategyBaseTVLLimits, bound to a specific deployed contract.
func NewStrategyBaseTVLLimitsFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyBaseTVLLimitsFilterer, error) {
	contract, err := bindStrategyBaseTVLLimits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsFilterer{contract: contract}, nil
}

// bindStrategyBaseTVLLimits binds a generic wrapper to an already deployed contract.
func bindStrategyBaseTVLLimits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrategyBaseTVLLimitsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyBaseTVLLimitsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrategyBaseTVLLimits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Explanation() (string, error) {
	return _StrategyBaseTVLLimits.Contract.Explanation(&_StrategyBaseTVLLimits.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Explanation() (string, error) {
	return _StrategyBaseTVLLimits.Contract.Explanation(&_StrategyBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) GetTVLLimits(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "getTVLLimits")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.GetTVLLimits(&_StrategyBaseTVLLimits.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.GetTVLLimits(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) MaxPerDeposit() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxPerDeposit(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxPerDeposit(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) MaxTotalDeposits() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxTotalDeposits(&_StrategyBaseTVLLimits.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.MaxTotalDeposits(&_StrategyBaseTVLLimits.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Paused(index uint8) (bool, error) {
	return _StrategyBaseTVLLimits.Contract.Paused(&_StrategyBaseTVLLimits.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Paused(index uint8) (bool, error) {
	return _StrategyBaseTVLLimits.Contract.Paused(&_StrategyBaseTVLLimits.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Paused0() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Paused0(&_StrategyBaseTVLLimits.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Paused0() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Paused0(&_StrategyBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) PauserRegistry() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.PauserRegistry(&_StrategyBaseTVLLimits.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) PauserRegistry() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.PauserRegistry(&_StrategyBaseTVLLimits.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Shares(&_StrategyBaseTVLLimits.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.Shares(&_StrategyBaseTVLLimits.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlying(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlying(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.SharesToUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) StrategyManager() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyManager(&_StrategyBaseTVLLimits.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) StrategyManager() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.StrategyManager(&_StrategyBaseTVLLimits.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) TotalShares() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.TotalShares(&_StrategyBaseTVLLimits.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) TotalShares() (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.TotalShares(&_StrategyBaseTVLLimits.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToShares(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToShares(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToSharesView(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToSharesView(&_StrategyBaseTVLLimits.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToken(&_StrategyBaseTVLLimits.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UnderlyingToken() (common.Address, error) {
	return _StrategyBaseTVLLimits.Contract.UnderlyingToken(&_StrategyBaseTVLLimits.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StrategyBaseTVLLimits.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlyingView(&_StrategyBaseTVLLimits.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Deposit(&_StrategyBaseTVLLimits.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Deposit(&_StrategyBaseTVLLimits.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Initialize(opts *bind.TransactOpts, _maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "initialize", _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize(&_StrategyBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize(&_StrategyBaseTVLLimits.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Initialize0(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "initialize0", _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Initialize0(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize0(&_StrategyBaseTVLLimits.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Initialize0(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Initialize0(&_StrategyBaseTVLLimits.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Pause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Pause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.PauseAll(&_StrategyBaseTVLLimits.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) PauseAll() (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.PauseAll(&_StrategyBaseTVLLimits.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetPauserRegistry(&_StrategyBaseTVLLimits.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetPauserRegistry(&_StrategyBaseTVLLimits.TransactOpts, newPauserRegistry)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) SetTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "setTVLLimits", newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetTVLLimits(&_StrategyBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.SetTVLLimits(&_StrategyBaseTVLLimits.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Unpause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Unpause(&_StrategyBaseTVLLimits.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlying(&_StrategyBaseTVLLimits.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.UserUnderlying(&_StrategyBaseTVLLimits.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Withdraw(&_StrategyBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _StrategyBaseTVLLimits.Contract.Withdraw(&_StrategyBaseTVLLimits.TransactOpts, recipient, token, amountShares)
}

// StrategyBaseTVLLimitsExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsExchangeRateEmittedIterator struct {
	Event *StrategyBaseTVLLimitsExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsExchangeRateEmitted)
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
		it.Event = new(StrategyBaseTVLLimitsExchangeRateEmitted)
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
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsExchangeRateEmitted represents a ExchangeRateEmitted event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsExchangeRateEmittedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsExchangeRateEmittedIterator{contract: _StrategyBaseTVLLimits.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsExchangeRateEmitted)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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

// ParseExchangeRateEmitted is a log parse operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseExchangeRateEmitted(log types.Log) (*StrategyBaseTVLLimitsExchangeRateEmitted, error) {
	event := new(StrategyBaseTVLLimitsExchangeRateEmitted)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsInitializedIterator struct {
	Event *StrategyBaseTVLLimitsInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsInitialized)
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
		it.Event = new(StrategyBaseTVLLimitsInitialized)
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
func (it *StrategyBaseTVLLimitsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsInitialized represents a Initialized event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsInitializedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsInitializedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsInitialized) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsInitialized)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseInitialized(log types.Log) (*StrategyBaseTVLLimitsInitialized, error) {
	event := new(StrategyBaseTVLLimitsInitialized)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator struct {
	Event *StrategyBaseTVLLimitsMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
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
		it.Event = new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
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
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsMaxPerDepositUpdatedIterator{contract: _StrategyBaseTVLLimits.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseMaxPerDepositUpdated(log types.Log) (*StrategyBaseTVLLimitsMaxPerDepositUpdated, error) {
	event := new(StrategyBaseTVLLimitsMaxPerDepositUpdated)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator struct {
	Event *StrategyBaseTVLLimitsMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
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
		it.Event = new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
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
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsMaxTotalDepositsUpdatedIterator{contract: _StrategyBaseTVLLimits.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*StrategyBaseTVLLimitsMaxTotalDepositsUpdated, error) {
	event := new(StrategyBaseTVLLimitsMaxTotalDepositsUpdated)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPausedIterator struct {
	Event *StrategyBaseTVLLimitsPaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsPaused)
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
		it.Event = new(StrategyBaseTVLLimitsPaused)
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
func (it *StrategyBaseTVLLimitsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsPaused represents a Paused event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseTVLLimitsPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsPausedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsPaused)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParsePaused(log types.Log) (*StrategyBaseTVLLimitsPaused, error) {
	event := new(StrategyBaseTVLLimitsPaused)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPauserRegistrySetIterator struct {
	Event *StrategyBaseTVLLimitsPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsPauserRegistrySet)
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
		it.Event = new(StrategyBaseTVLLimitsPauserRegistrySet)
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
func (it *StrategyBaseTVLLimitsPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsPauserRegistrySet represents a PauserRegistrySet event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsPauserRegistrySetIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsPauserRegistrySetIterator{contract: _StrategyBaseTVLLimits.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsPauserRegistrySet)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParsePauserRegistrySet(log types.Log) (*StrategyBaseTVLLimitsPauserRegistrySet, error) {
	event := new(StrategyBaseTVLLimitsPauserRegistrySet)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsStrategyTokenSetIterator struct {
	Event *StrategyBaseTVLLimitsStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsStrategyTokenSet)
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
		it.Event = new(StrategyBaseTVLLimitsStrategyTokenSet)
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
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsStrategyTokenSet represents a StrategyTokenSet event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*StrategyBaseTVLLimitsStrategyTokenSetIterator, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsStrategyTokenSetIterator{contract: _StrategyBaseTVLLimits.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsStrategyTokenSet)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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

// ParseStrategyTokenSet is a log parse operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseStrategyTokenSet(log types.Log) (*StrategyBaseTVLLimitsStrategyTokenSet, error) {
	event := new(StrategyBaseTVLLimitsStrategyTokenSet)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyBaseTVLLimitsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsUnpausedIterator struct {
	Event *StrategyBaseTVLLimitsUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyBaseTVLLimitsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyBaseTVLLimitsUnpaused)
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
		it.Event = new(StrategyBaseTVLLimitsUnpaused)
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
func (it *StrategyBaseTVLLimitsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyBaseTVLLimitsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyBaseTVLLimitsUnpaused represents a Unpaused event raised by the StrategyBaseTVLLimits contract.
type StrategyBaseTVLLimitsUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyBaseTVLLimitsUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyBaseTVLLimitsUnpausedIterator{contract: _StrategyBaseTVLLimits.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyBaseTVLLimitsUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StrategyBaseTVLLimits.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyBaseTVLLimitsUnpaused)
				if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_StrategyBaseTVLLimits *StrategyBaseTVLLimitsFilterer) ParseUnpaused(log types.Log) (*StrategyBaseTVLLimitsUnpaused, error) {
	event := new(StrategyBaseTVLLimitsUnpaused)
	if err := _StrategyBaseTVLLimits.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
