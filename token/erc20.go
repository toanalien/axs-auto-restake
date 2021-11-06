// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// IERC20StakingManagerBlockReward is an auto generated low-level Go binding around an user-defined struct.
type IERC20StakingManagerBlockReward struct {
	FromBlock      *big.Int
	RewardPerBlock *big.Int
}

// IERC20StakingManagerPool is an auto generated low-level Go binding around an user-defined struct.
type IERC20StakingManagerPool struct {
	StakingToken               common.Address
	RewardToken                common.Address
	AccumulatedRewardsPerShare *big.Int
	LastSyncedBlock            *big.Int
	StartedAtBlock             *big.Int
}

// IERC20StakingManagerUserReward is an auto generated low-level Go binding around an user-defined struct.
type IERC20StakingManagerUserReward struct {
	DebitedRewards   *big.Int
	CreditedRewards  *big.Int
	LastClaimedBlock *big.Int
}

// TokenMetaData contains all meta data concerning the Token contract.
var TokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_new\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_old\",\"type\":\"uint256\"}],\"name\":\"MinClaimedBlocksUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accumulatedPerShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lastSyncedBlock\",\"type\":\"uint256\"}],\"name\":\"PoolSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"PoolWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIERC20StakingManager.BlockReward[]\",\"name\":\"_blockRewards\",\"type\":\"tuple[]\"}],\"name\":\"RewardPerBlockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"debitedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIERC20StakingManager.UserReward\",\"name\":\"_rewardInfo\",\"type\":\"tuple\"}],\"name\":\"UserRewardUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"allocateRewards\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_earnedRewards\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"canObtainRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_toBlock\",\"type\":\"uint256\"}],\"name\":\"getIntervalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_result\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"getPoolInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewardsPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSyncedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAtBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIERC20StakingManager.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"isPoolStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"isPoolWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minClaimedBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedRewardsPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSyncedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAtBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"resetRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fromBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"setFutureBlockReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minClaimedBlocks\",\"type\":\"uint256\"}],\"name\":\"setMinClaimedBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool\",\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"syncPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newStakingAmount\",\"type\":\"uint256\"}],\"name\":\"syncRewardInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debitedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimedBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20StakingPool[]\",\"name\":\"_pools\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_startedAtBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardPerBlocks\",\"type\":\"uint256[]\"}],\"name\":\"whitelistPools\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimPendingRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"emergencyUnstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getStakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakingTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"methodPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_method\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_method\",\"type\":\"bytes4\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restakeRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_method\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstakeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMetaData.ABI instead.
var TokenABI = TokenMetaData.ABI

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCallerSession) Admin() (common.Address, error) {
	return _Token.Contract.Admin(&_Token.CallOpts)
}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCaller) Admin0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "admin0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenSession) Admin0() (common.Address, error) {
	return _Token.Contract.Admin0(&_Token.CallOpts)
}

// Admin0 is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Token *TokenCallerSession) Admin0() (common.Address, error) {
	return _Token.Contract.Admin0(&_Token.CallOpts)
}

// BlockRewards is a free data retrieval call binding the contract method 0x97c58892.
//
// Solidity: function blockRewards(address , uint256 ) view returns(uint256 fromBlock, uint256 rewardPerBlock)
func (_Token *TokenCaller) BlockRewards(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock      *big.Int
	RewardPerBlock *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "blockRewards", arg0, arg1)

	outstruct := new(struct {
		FromBlock      *big.Int
		RewardPerBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardPerBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BlockRewards is a free data retrieval call binding the contract method 0x97c58892.
//
// Solidity: function blockRewards(address , uint256 ) view returns(uint256 fromBlock, uint256 rewardPerBlock)
func (_Token *TokenSession) BlockRewards(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock      *big.Int
	RewardPerBlock *big.Int
}, error) {
	return _Token.Contract.BlockRewards(&_Token.CallOpts, arg0, arg1)
}

// BlockRewards is a free data retrieval call binding the contract method 0x97c58892.
//
// Solidity: function blockRewards(address , uint256 ) view returns(uint256 fromBlock, uint256 rewardPerBlock)
func (_Token *TokenCallerSession) BlockRewards(arg0 common.Address, arg1 *big.Int) (struct {
	FromBlock      *big.Int
	RewardPerBlock *big.Int
}, error) {
	return _Token.Contract.BlockRewards(&_Token.CallOpts, arg0, arg1)
}

// CanObtainRewards is a free data retrieval call binding the contract method 0xd96f0723.
//
// Solidity: function canObtainRewards(address _pool, address _user) view returns(bool)
func (_Token *TokenCaller) CanObtainRewards(opts *bind.CallOpts, _pool common.Address, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "canObtainRewards", _pool, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanObtainRewards is a free data retrieval call binding the contract method 0xd96f0723.
//
// Solidity: function canObtainRewards(address _pool, address _user) view returns(bool)
func (_Token *TokenSession) CanObtainRewards(_pool common.Address, _user common.Address) (bool, error) {
	return _Token.Contract.CanObtainRewards(&_Token.CallOpts, _pool, _user)
}

// CanObtainRewards is a free data retrieval call binding the contract method 0xd96f0723.
//
// Solidity: function canObtainRewards(address _pool, address _user) view returns(bool)
func (_Token *TokenCallerSession) CanObtainRewards(_pool common.Address, _user common.Address) (bool, error) {
	return _Token.Contract.CanObtainRewards(&_Token.CallOpts, _pool, _user)
}

// GetBlockReward is a free data retrieval call binding the contract method 0x3b268950.
//
// Solidity: function getBlockReward(address _pool, uint256 _blockNumber) view returns(uint256 _reward)
func (_Token *TokenCaller) GetBlockReward(opts *bind.CallOpts, _pool common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getBlockReward", _pool, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockReward is a free data retrieval call binding the contract method 0x3b268950.
//
// Solidity: function getBlockReward(address _pool, uint256 _blockNumber) view returns(uint256 _reward)
func (_Token *TokenSession) GetBlockReward(_pool common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Token.Contract.GetBlockReward(&_Token.CallOpts, _pool, _blockNumber)
}

// GetBlockReward is a free data retrieval call binding the contract method 0x3b268950.
//
// Solidity: function getBlockReward(address _pool, uint256 _blockNumber) view returns(uint256 _reward)
func (_Token *TokenCallerSession) GetBlockReward(_pool common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _Token.Contract.GetBlockReward(&_Token.CallOpts, _pool, _blockNumber)
}

// GetIntervalRewards is a free data retrieval call binding the contract method 0xae019e83.
//
// Solidity: function getIntervalRewards(address _pool, uint256 _fromBlock, uint256 _toBlock) view returns(uint256 _result)
func (_Token *TokenCaller) GetIntervalRewards(opts *bind.CallOpts, _pool common.Address, _fromBlock *big.Int, _toBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getIntervalRewards", _pool, _fromBlock, _toBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIntervalRewards is a free data retrieval call binding the contract method 0xae019e83.
//
// Solidity: function getIntervalRewards(address _pool, uint256 _fromBlock, uint256 _toBlock) view returns(uint256 _result)
func (_Token *TokenSession) GetIntervalRewards(_pool common.Address, _fromBlock *big.Int, _toBlock *big.Int) (*big.Int, error) {
	return _Token.Contract.GetIntervalRewards(&_Token.CallOpts, _pool, _fromBlock, _toBlock)
}

// GetIntervalRewards is a free data retrieval call binding the contract method 0xae019e83.
//
// Solidity: function getIntervalRewards(address _pool, uint256 _fromBlock, uint256 _toBlock) view returns(uint256 _result)
func (_Token *TokenCallerSession) GetIntervalRewards(_pool common.Address, _fromBlock *big.Int, _toBlock *big.Int) (*big.Int, error) {
	return _Token.Contract.GetIntervalRewards(&_Token.CallOpts, _pool, _fromBlock, _toBlock)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0x7a27db57.
//
// Solidity: function getPendingRewards(address _pool, address _user) view returns(uint256)
func (_Token *TokenCaller) GetPendingRewards(opts *bind.CallOpts, _pool common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getPendingRewards", _pool, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards is a free data retrieval call binding the contract method 0x7a27db57.
//
// Solidity: function getPendingRewards(address _pool, address _user) view returns(uint256)
func (_Token *TokenSession) GetPendingRewards(_pool common.Address, _user common.Address) (*big.Int, error) {
	return _Token.Contract.GetPendingRewards(&_Token.CallOpts, _pool, _user)
}

// GetPendingRewards is a free data retrieval call binding the contract method 0x7a27db57.
//
// Solidity: function getPendingRewards(address _pool, address _user) view returns(uint256)
func (_Token *TokenCallerSession) GetPendingRewards(_pool common.Address, _user common.Address) (*big.Int, error) {
	return _Token.Contract.GetPendingRewards(&_Token.CallOpts, _pool, _user)
}

// GetPendingRewards0 is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Token *TokenCaller) GetPendingRewards0(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getPendingRewards0", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingRewards0 is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Token *TokenSession) GetPendingRewards0(_user common.Address) (*big.Int, error) {
	return _Token.Contract.GetPendingRewards0(&_Token.CallOpts, _user)
}

// GetPendingRewards0 is a free data retrieval call binding the contract method 0xf6ed2017.
//
// Solidity: function getPendingRewards(address _user) view returns(uint256)
func (_Token *TokenCallerSession) GetPendingRewards0(_user common.Address) (*big.Int, error) {
	return _Token.Contract.GetPendingRewards0(&_Token.CallOpts, _user)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _pool) view returns((address,address,uint256,uint256,uint256))
func (_Token *TokenCaller) GetPoolInfo(opts *bind.CallOpts, _pool common.Address) (IERC20StakingManagerPool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getPoolInfo", _pool)

	if err != nil {
		return *new(IERC20StakingManagerPool), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC20StakingManagerPool)).(*IERC20StakingManagerPool)

	return out0, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _pool) view returns((address,address,uint256,uint256,uint256))
func (_Token *TokenSession) GetPoolInfo(_pool common.Address) (IERC20StakingManagerPool, error) {
	return _Token.Contract.GetPoolInfo(&_Token.CallOpts, _pool)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x06bfa938.
//
// Solidity: function getPoolInfo(address _pool) view returns((address,address,uint256,uint256,uint256))
func (_Token *TokenCallerSession) GetPoolInfo(_pool common.Address) (IERC20StakingManagerPool, error) {
	return _Token.Contract.GetPoolInfo(&_Token.CallOpts, _pool)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Token *TokenCaller) GetRewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getRewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Token *TokenSession) GetRewardToken() (common.Address, error) {
	return _Token.Contract.GetRewardToken(&_Token.CallOpts)
}

// GetRewardToken is a free data retrieval call binding the contract method 0x69940d79.
//
// Solidity: function getRewardToken() view returns(address)
func (_Token *TokenCallerSession) GetRewardToken() (common.Address, error) {
	return _Token.Contract.GetRewardToken(&_Token.CallOpts)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Token *TokenCaller) GetStakingAmount(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getStakingAmount", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Token *TokenSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Token.Contract.GetStakingAmount(&_Token.CallOpts, _user)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x74363daa.
//
// Solidity: function getStakingAmount(address _user) view returns(uint256)
func (_Token *TokenCallerSession) GetStakingAmount(_user common.Address) (*big.Int, error) {
	return _Token.Contract.GetStakingAmount(&_Token.CallOpts, _user)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Token *TokenCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Token *TokenSession) GetStakingToken() (common.Address, error) {
	return _Token.Contract.GetStakingToken(&_Token.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Token *TokenCallerSession) GetStakingToken() (common.Address, error) {
	return _Token.Contract.GetStakingToken(&_Token.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Token *TokenCaller) GetStakingTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "getStakingTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Token *TokenSession) GetStakingTotal() (*big.Int, error) {
	return _Token.Contract.GetStakingTotal(&_Token.CallOpts)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x61ee98dc.
//
// Solidity: function getStakingTotal() view returns(uint256)
func (_Token *TokenCallerSession) GetStakingTotal() (*big.Int, error) {
	return _Token.Contract.GetStakingTotal(&_Token.CallOpts)
}

// IsPoolStarted is a free data retrieval call binding the contract method 0x7bc6d2bb.
//
// Solidity: function isPoolStarted(address _pool) view returns(bool)
func (_Token *TokenCaller) IsPoolStarted(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPoolStarted", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolStarted is a free data retrieval call binding the contract method 0x7bc6d2bb.
//
// Solidity: function isPoolStarted(address _pool) view returns(bool)
func (_Token *TokenSession) IsPoolStarted(_pool common.Address) (bool, error) {
	return _Token.Contract.IsPoolStarted(&_Token.CallOpts, _pool)
}

// IsPoolStarted is a free data retrieval call binding the contract method 0x7bc6d2bb.
//
// Solidity: function isPoolStarted(address _pool) view returns(bool)
func (_Token *TokenCallerSession) IsPoolStarted(_pool common.Address) (bool, error) {
	return _Token.Contract.IsPoolStarted(&_Token.CallOpts, _pool)
}

// IsPoolWhitelisted is a free data retrieval call binding the contract method 0x2b26a982.
//
// Solidity: function isPoolWhitelisted(address _pool) view returns(bool)
func (_Token *TokenCaller) IsPoolWhitelisted(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "isPoolWhitelisted", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolWhitelisted is a free data retrieval call binding the contract method 0x2b26a982.
//
// Solidity: function isPoolWhitelisted(address _pool) view returns(bool)
func (_Token *TokenSession) IsPoolWhitelisted(_pool common.Address) (bool, error) {
	return _Token.Contract.IsPoolWhitelisted(&_Token.CallOpts, _pool)
}

// IsPoolWhitelisted is a free data retrieval call binding the contract method 0x2b26a982.
//
// Solidity: function isPoolWhitelisted(address _pool) view returns(bool)
func (_Token *TokenCallerSession) IsPoolWhitelisted(_pool common.Address) (bool, error) {
	return _Token.Contract.IsPoolWhitelisted(&_Token.CallOpts, _pool)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Token *TokenCaller) MethodPaused(opts *bind.CallOpts, arg0 [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "methodPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Token *TokenSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Token.Contract.MethodPaused(&_Token.CallOpts, arg0)
}

// MethodPaused is a free data retrieval call binding the contract method 0xfc9dd5c4.
//
// Solidity: function methodPaused(bytes4 ) view returns(bool)
func (_Token *TokenCallerSession) MethodPaused(arg0 [4]byte) (bool, error) {
	return _Token.Contract.MethodPaused(&_Token.CallOpts, arg0)
}

// MinClaimedBlocks is a free data retrieval call binding the contract method 0x9bfc4e6e.
//
// Solidity: function minClaimedBlocks() view returns(uint256)
func (_Token *TokenCaller) MinClaimedBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "minClaimedBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinClaimedBlocks is a free data retrieval call binding the contract method 0x9bfc4e6e.
//
// Solidity: function minClaimedBlocks() view returns(uint256)
func (_Token *TokenSession) MinClaimedBlocks() (*big.Int, error) {
	return _Token.Contract.MinClaimedBlocks(&_Token.CallOpts)
}

// MinClaimedBlocks is a free data retrieval call binding the contract method 0x9bfc4e6e.
//
// Solidity: function minClaimedBlocks() view returns(uint256)
func (_Token *TokenCallerSession) MinClaimedBlocks() (*big.Int, error) {
	return _Token.Contract.MinClaimedBlocks(&_Token.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Token *TokenCaller) Paused(opts *bind.CallOpts, _method [4]byte) (bool, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "paused", _method)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Token *TokenSession) Paused(_method [4]byte) (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts, _method)
}

// Paused is a free data retrieval call binding the contract method 0x59237eba.
//
// Solidity: function paused(bytes4 _method) view returns(bool)
func (_Token *TokenCallerSession) Paused(_method [4]byte) (bool, error) {
	return _Token.Contract.Paused(&_Token.CallOpts, _method)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address ) view returns(address stakingToken, address rewardToken, uint256 accumulatedRewardsPerShare, uint256 lastSyncedBlock, uint256 startedAtBlock)
func (_Token *TokenCaller) PoolInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	StakingToken               common.Address
	RewardToken                common.Address
	AccumulatedRewardsPerShare *big.Int
	LastSyncedBlock            *big.Int
	StartedAtBlock             *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		StakingToken               common.Address
		RewardToken                common.Address
		AccumulatedRewardsPerShare *big.Int
		LastSyncedBlock            *big.Int
		StartedAtBlock             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakingToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AccumulatedRewardsPerShare = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastSyncedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartedAtBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address ) view returns(address stakingToken, address rewardToken, uint256 accumulatedRewardsPerShare, uint256 lastSyncedBlock, uint256 startedAtBlock)
func (_Token *TokenSession) PoolInfo(arg0 common.Address) (struct {
	StakingToken               common.Address
	RewardToken                common.Address
	AccumulatedRewardsPerShare *big.Int
	LastSyncedBlock            *big.Int
	StartedAtBlock             *big.Int
}, error) {
	return _Token.Contract.PoolInfo(&_Token.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address ) view returns(address stakingToken, address rewardToken, uint256 accumulatedRewardsPerShare, uint256 lastSyncedBlock, uint256 startedAtBlock)
func (_Token *TokenCallerSession) PoolInfo(arg0 common.Address) (struct {
	StakingToken               common.Address
	RewardToken                common.Address
	AccumulatedRewardsPerShare *big.Int
	LastSyncedBlock            *big.Int
	StartedAtBlock             *big.Int
}, error) {
	return _Token.Contract.PoolInfo(&_Token.CallOpts, arg0)
}

// UserRewardInfo is a free data retrieval call binding the contract method 0x66993646.
//
// Solidity: function userRewardInfo(address , address ) view returns(uint256 debitedRewards, uint256 creditedRewards, uint256 lastClaimedBlock)
func (_Token *TokenCaller) UserRewardInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	DebitedRewards   *big.Int
	CreditedRewards  *big.Int
	LastClaimedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Token.contract.Call(opts, &out, "userRewardInfo", arg0, arg1)

	outstruct := new(struct {
		DebitedRewards   *big.Int
		CreditedRewards  *big.Int
		LastClaimedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DebitedRewards = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CreditedRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastClaimedBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserRewardInfo is a free data retrieval call binding the contract method 0x66993646.
//
// Solidity: function userRewardInfo(address , address ) view returns(uint256 debitedRewards, uint256 creditedRewards, uint256 lastClaimedBlock)
func (_Token *TokenSession) UserRewardInfo(arg0 common.Address, arg1 common.Address) (struct {
	DebitedRewards   *big.Int
	CreditedRewards  *big.Int
	LastClaimedBlock *big.Int
}, error) {
	return _Token.Contract.UserRewardInfo(&_Token.CallOpts, arg0, arg1)
}

// UserRewardInfo is a free data retrieval call binding the contract method 0x66993646.
//
// Solidity: function userRewardInfo(address , address ) view returns(uint256 debitedRewards, uint256 creditedRewards, uint256 lastClaimedBlock)
func (_Token *TokenCallerSession) UserRewardInfo(arg0 common.Address, arg1 common.Address) (struct {
	DebitedRewards   *big.Int
	CreditedRewards  *big.Int
	LastClaimedBlock *big.Int
}, error) {
	return _Token.Contract.UserRewardInfo(&_Token.CallOpts, arg0, arg1)
}

// AllocateRewards is a paid mutator transaction binding the contract method 0x80cde84d.
//
// Solidity: function allocateRewards(address _user) returns(address _rewardToken, uint256 _earnedRewards)
func (_Token *TokenTransactor) AllocateRewards(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "allocateRewards", _user)
}

// AllocateRewards is a paid mutator transaction binding the contract method 0x80cde84d.
//
// Solidity: function allocateRewards(address _user) returns(address _rewardToken, uint256 _earnedRewards)
func (_Token *TokenSession) AllocateRewards(_user common.Address) (*types.Transaction, error) {
	return _Token.Contract.AllocateRewards(&_Token.TransactOpts, _user)
}

// AllocateRewards is a paid mutator transaction binding the contract method 0x80cde84d.
//
// Solidity: function allocateRewards(address _user) returns(address _rewardToken, uint256 _earnedRewards)
func (_Token *TokenTransactorSession) AllocateRewards(_user common.Address) (*types.Transaction, error) {
	return _Token.Contract.AllocateRewards(&_Token.TransactOpts, _user)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeAdmin(&_Token.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeAdmin(&_Token.TransactOpts, _newAdmin)
}

// ChangeAdmin0 is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenTransactor) ChangeAdmin0(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "changeAdmin0", _newAdmin)
}

// ChangeAdmin0 is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenSession) ChangeAdmin0(_newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeAdmin0(&_Token.TransactOpts, _newAdmin)
}

// ChangeAdmin0 is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Token *TokenTransactorSession) ChangeAdmin0(_newAdmin common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeAdmin0(&_Token.TransactOpts, _newAdmin)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Token *TokenTransactor) ClaimPendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "claimPendingRewards")
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Token *TokenSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Token.Contract.ClaimPendingRewards(&_Token.TransactOpts)
}

// ClaimPendingRewards is a paid mutator transaction binding the contract method 0x92bd7b2c.
//
// Solidity: function claimPendingRewards() returns()
func (_Token *TokenTransactorSession) ClaimPendingRewards() (*types.Transaction, error) {
	return _Token.Contract.ClaimPendingRewards(&_Token.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Token *TokenTransactor) EmergencyUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "emergencyUnstake")
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Token *TokenSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Token.Contract.EmergencyUnstake(&_Token.TransactOpts)
}

// EmergencyUnstake is a paid mutator transaction binding the contract method 0x7589cf2f.
//
// Solidity: function emergencyUnstake() returns()
func (_Token *TokenTransactorSession) EmergencyUnstake() (*types.Transaction, error) {
	return _Token.Contract.EmergencyUnstake(&_Token.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Token *TokenTransactor) Pause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pause", _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Token *TokenSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts, _method)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _method) returns()
func (_Token *TokenTransactorSession) Pause(_method [4]byte) (*types.Transaction, error) {
	return _Token.Contract.Pause(&_Token.TransactOpts, _method)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Token *TokenTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Token *TokenSession) PauseAll() (*types.Transaction, error) {
	return _Token.Contract.PauseAll(&_Token.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Token *TokenTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Token.Contract.PauseAll(&_Token.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenSession) RemoveAdmin() (*types.Transaction, error) {
	return _Token.Contract.RemoveAdmin(&_Token.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Token.Contract.RemoveAdmin(&_Token.TransactOpts)
}

// RemoveAdmin0 is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenTransactor) RemoveAdmin0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeAdmin0")
}

// RemoveAdmin0 is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenSession) RemoveAdmin0() (*types.Transaction, error) {
	return _Token.Contract.RemoveAdmin0(&_Token.TransactOpts)
}

// RemoveAdmin0 is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Token *TokenTransactorSession) RemoveAdmin0() (*types.Transaction, error) {
	return _Token.Contract.RemoveAdmin0(&_Token.TransactOpts)
}

// ResetRewards is a paid mutator transaction binding the contract method 0x5aa0c7f8.
//
// Solidity: function resetRewards(address _user) returns()
func (_Token *TokenTransactor) ResetRewards(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "resetRewards", _user)
}

// ResetRewards is a paid mutator transaction binding the contract method 0x5aa0c7f8.
//
// Solidity: function resetRewards(address _user) returns()
func (_Token *TokenSession) ResetRewards(_user common.Address) (*types.Transaction, error) {
	return _Token.Contract.ResetRewards(&_Token.TransactOpts, _user)
}

// ResetRewards is a paid mutator transaction binding the contract method 0x5aa0c7f8.
//
// Solidity: function resetRewards(address _user) returns()
func (_Token *TokenTransactorSession) ResetRewards(_user common.Address) (*types.Transaction, error) {
	return _Token.Contract.ResetRewards(&_Token.TransactOpts, _user)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Token *TokenTransactor) RestakeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "restakeRewards")
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Token *TokenSession) RestakeRewards() (*types.Transaction, error) {
	return _Token.Contract.RestakeRewards(&_Token.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x3d8527ba.
//
// Solidity: function restakeRewards() returns()
func (_Token *TokenTransactorSession) RestakeRewards() (*types.Transaction, error) {
	return _Token.Contract.RestakeRewards(&_Token.TransactOpts)
}

// SetFutureBlockReward is a paid mutator transaction binding the contract method 0x39768e6a.
//
// Solidity: function setFutureBlockReward(address _pool, uint256 _fromBlock, uint256 _rewardPerBlock) returns()
func (_Token *TokenTransactor) SetFutureBlockReward(opts *bind.TransactOpts, _pool common.Address, _fromBlock *big.Int, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setFutureBlockReward", _pool, _fromBlock, _rewardPerBlock)
}

// SetFutureBlockReward is a paid mutator transaction binding the contract method 0x39768e6a.
//
// Solidity: function setFutureBlockReward(address _pool, uint256 _fromBlock, uint256 _rewardPerBlock) returns()
func (_Token *TokenSession) SetFutureBlockReward(_pool common.Address, _fromBlock *big.Int, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetFutureBlockReward(&_Token.TransactOpts, _pool, _fromBlock, _rewardPerBlock)
}

// SetFutureBlockReward is a paid mutator transaction binding the contract method 0x39768e6a.
//
// Solidity: function setFutureBlockReward(address _pool, uint256 _fromBlock, uint256 _rewardPerBlock) returns()
func (_Token *TokenTransactorSession) SetFutureBlockReward(_pool common.Address, _fromBlock *big.Int, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetFutureBlockReward(&_Token.TransactOpts, _pool, _fromBlock, _rewardPerBlock)
}

// SetMinClaimedBlocks is a paid mutator transaction binding the contract method 0x405f5b07.
//
// Solidity: function setMinClaimedBlocks(uint256 _minClaimedBlocks) returns()
func (_Token *TokenTransactor) SetMinClaimedBlocks(opts *bind.TransactOpts, _minClaimedBlocks *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setMinClaimedBlocks", _minClaimedBlocks)
}

// SetMinClaimedBlocks is a paid mutator transaction binding the contract method 0x405f5b07.
//
// Solidity: function setMinClaimedBlocks(uint256 _minClaimedBlocks) returns()
func (_Token *TokenSession) SetMinClaimedBlocks(_minClaimedBlocks *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMinClaimedBlocks(&_Token.TransactOpts, _minClaimedBlocks)
}

// SetMinClaimedBlocks is a paid mutator transaction binding the contract method 0x405f5b07.
//
// Solidity: function setMinClaimedBlocks(uint256 _minClaimedBlocks) returns()
func (_Token *TokenTransactorSession) SetMinClaimedBlocks(_minClaimedBlocks *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SetMinClaimedBlocks(&_Token.TransactOpts, _minClaimedBlocks)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Token *TokenTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Token *TokenSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Stake(&_Token.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_Token *TokenTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Stake(&_Token.TransactOpts, _amount)
}

// SyncPool is a paid mutator transaction binding the contract method 0x8fb2e2c2.
//
// Solidity: function syncPool(address _pool) returns()
func (_Token *TokenTransactor) SyncPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "syncPool", _pool)
}

// SyncPool is a paid mutator transaction binding the contract method 0x8fb2e2c2.
//
// Solidity: function syncPool(address _pool) returns()
func (_Token *TokenSession) SyncPool(_pool common.Address) (*types.Transaction, error) {
	return _Token.Contract.SyncPool(&_Token.TransactOpts, _pool)
}

// SyncPool is a paid mutator transaction binding the contract method 0x8fb2e2c2.
//
// Solidity: function syncPool(address _pool) returns()
func (_Token *TokenTransactorSession) SyncPool(_pool common.Address) (*types.Transaction, error) {
	return _Token.Contract.SyncPool(&_Token.TransactOpts, _pool)
}

// SyncRewardInfo is a paid mutator transaction binding the contract method 0x73939bb5.
//
// Solidity: function syncRewardInfo(address _user, uint256 _newStakingAmount) returns()
func (_Token *TokenTransactor) SyncRewardInfo(opts *bind.TransactOpts, _user common.Address, _newStakingAmount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "syncRewardInfo", _user, _newStakingAmount)
}

// SyncRewardInfo is a paid mutator transaction binding the contract method 0x73939bb5.
//
// Solidity: function syncRewardInfo(address _user, uint256 _newStakingAmount) returns()
func (_Token *TokenSession) SyncRewardInfo(_user common.Address, _newStakingAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SyncRewardInfo(&_Token.TransactOpts, _user, _newStakingAmount)
}

// SyncRewardInfo is a paid mutator transaction binding the contract method 0x73939bb5.
//
// Solidity: function syncRewardInfo(address _user, uint256 _newStakingAmount) returns()
func (_Token *TokenTransactorSession) SyncRewardInfo(_user common.Address, _newStakingAmount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.SyncRewardInfo(&_Token.TransactOpts, _user, _newStakingAmount)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Token *TokenTransactor) Unpause(opts *bind.TransactOpts, _method [4]byte) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpause", _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Token *TokenSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts, _method)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _method) returns()
func (_Token *TokenTransactorSession) Unpause(_method [4]byte) (*types.Transaction, error) {
	return _Token.Contract.Unpause(&_Token.TransactOpts, _method)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Token *TokenTransactor) UnpauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unpauseAll")
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Token *TokenSession) UnpauseAll() (*types.Transaction, error) {
	return _Token.Contract.UnpauseAll(&_Token.TransactOpts)
}

// UnpauseAll is a paid mutator transaction binding the contract method 0x8a2ddd03.
//
// Solidity: function unpauseAll() returns()
func (_Token *TokenTransactorSession) UnpauseAll() (*types.Transaction, error) {
	return _Token.Contract.UnpauseAll(&_Token.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Token *TokenTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Token *TokenSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Unstake(&_Token.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_Token *TokenTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Unstake(&_Token.TransactOpts, _amount)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Token *TokenTransactor) UnstakeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "unstakeAll")
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Token *TokenSession) UnstakeAll() (*types.Transaction, error) {
	return _Token.Contract.UnstakeAll(&_Token.TransactOpts)
}

// UnstakeAll is a paid mutator transaction binding the contract method 0x35322f37.
//
// Solidity: function unstakeAll() returns()
func (_Token *TokenTransactorSession) UnstakeAll() (*types.Transaction, error) {
	return _Token.Contract.UnstakeAll(&_Token.TransactOpts)
}

// WhitelistPools is a paid mutator transaction binding the contract method 0xaef85c90.
//
// Solidity: function whitelistPools(address[] _pools, uint256[] _startedAtBlocks, uint256[] _rewardPerBlocks) returns()
func (_Token *TokenTransactor) WhitelistPools(opts *bind.TransactOpts, _pools []common.Address, _startedAtBlocks []*big.Int, _rewardPerBlocks []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "whitelistPools", _pools, _startedAtBlocks, _rewardPerBlocks)
}

// WhitelistPools is a paid mutator transaction binding the contract method 0xaef85c90.
//
// Solidity: function whitelistPools(address[] _pools, uint256[] _startedAtBlocks, uint256[] _rewardPerBlocks) returns()
func (_Token *TokenSession) WhitelistPools(_pools []common.Address, _startedAtBlocks []*big.Int, _rewardPerBlocks []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.WhitelistPools(&_Token.TransactOpts, _pools, _startedAtBlocks, _rewardPerBlocks)
}

// WhitelistPools is a paid mutator transaction binding the contract method 0xaef85c90.
//
// Solidity: function whitelistPools(address[] _pools, uint256[] _startedAtBlocks, uint256[] _rewardPerBlocks) returns()
func (_Token *TokenTransactorSession) WhitelistPools(_pools []common.Address, _startedAtBlocks []*big.Int, _rewardPerBlocks []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.WhitelistPools(&_Token.TransactOpts, _pools, _startedAtBlocks, _rewardPerBlocks)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Token *TokenTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Token *TokenSession) WithdrawEther() (*types.Transaction, error) {
	return _Token.Contract.WithdrawEther(&_Token.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Token *TokenTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _Token.Contract.WithdrawEther(&_Token.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Token *TokenTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "withdrawToken", _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Token *TokenSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawToken(&_Token.TransactOpts, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Token *TokenTransactorSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Token.Contract.WithdrawToken(&_Token.TransactOpts, _token)
}

// TokenAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Token contract.
type TokenAdminChangedIterator struct {
	Event *TokenAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminChanged)
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
		it.Event = new(TokenAdminChanged)
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
func (it *TokenAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminChanged represents a AdminChanged event raised by the Token contract.
type TokenAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*TokenAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminChangedIterator{contract: _Token.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminChanged)
				if err := _Token.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) ParseAdminChanged(log types.Log) (*TokenAdminChanged, error) {
	event := new(TokenAdminChanged)
	if err := _Token.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAdminChanged0Iterator is returned from FilterAdminChanged0 and is used to iterate over the raw logs and unpacked data for AdminChanged0 events raised by the Token contract.
type TokenAdminChanged0Iterator struct {
	Event *TokenAdminChanged0 // Event containing the contract specifics and raw log

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
func (it *TokenAdminChanged0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminChanged0)
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
		it.Event = new(TokenAdminChanged0)
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
func (it *TokenAdminChanged0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminChanged0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminChanged0 represents a AdminChanged0 event raised by the Token contract.
type TokenAdminChanged0 struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged0 is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) FilterAdminChanged0(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*TokenAdminChanged0Iterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdminChanged0", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminChanged0Iterator{contract: _Token.contract, event: "AdminChanged0", logs: logs, sub: sub}, nil
}

// WatchAdminChanged0 is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) WatchAdminChanged0(opts *bind.WatchOpts, sink chan<- *TokenAdminChanged0, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdminChanged0", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminChanged0)
				if err := _Token.contract.UnpackLog(event, "AdminChanged0", log); err != nil {
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

// ParseAdminChanged0 is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Token *TokenFilterer) ParseAdminChanged0(log types.Log) (*TokenAdminChanged0, error) {
	event := new(TokenAdminChanged0)
	if err := _Token.contract.UnpackLog(event, "AdminChanged0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Token contract.
type TokenAdminRemovedIterator struct {
	Event *TokenAdminRemoved // Event containing the contract specifics and raw log

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
func (it *TokenAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRemoved)
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
		it.Event = new(TokenAdminRemoved)
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
func (it *TokenAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminRemoved represents a AdminRemoved event raised by the Token contract.
type TokenAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*TokenAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRemovedIterator{contract: _Token.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *TokenAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminRemoved)
				if err := _Token.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) ParseAdminRemoved(log types.Log) (*TokenAdminRemoved, error) {
	event := new(TokenAdminRemoved)
	if err := _Token.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenAdminRemoved0Iterator is returned from FilterAdminRemoved0 and is used to iterate over the raw logs and unpacked data for AdminRemoved0 events raised by the Token contract.
type TokenAdminRemoved0Iterator struct {
	Event *TokenAdminRemoved0 // Event containing the contract specifics and raw log

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
func (it *TokenAdminRemoved0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAdminRemoved0)
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
		it.Event = new(TokenAdminRemoved0)
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
func (it *TokenAdminRemoved0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAdminRemoved0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAdminRemoved0 represents a AdminRemoved0 event raised by the Token contract.
type TokenAdminRemoved0 struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved0 is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) FilterAdminRemoved0(opts *bind.FilterOpts, _oldAdmin []common.Address) (*TokenAdminRemoved0Iterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "AdminRemoved0", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenAdminRemoved0Iterator{contract: _Token.contract, event: "AdminRemoved0", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved0 is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) WatchAdminRemoved0(opts *bind.WatchOpts, sink chan<- *TokenAdminRemoved0, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "AdminRemoved0", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAdminRemoved0)
				if err := _Token.contract.UnpackLog(event, "AdminRemoved0", log); err != nil {
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

// ParseAdminRemoved0 is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Token *TokenFilterer) ParseAdminRemoved0(log types.Log) (*TokenAdminRemoved0, error) {
	event := new(TokenAdminRemoved0)
	if err := _Token.contract.UnpackLog(event, "AdminRemoved0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenEmergencyUnstakedIterator is returned from FilterEmergencyUnstaked and is used to iterate over the raw logs and unpacked data for EmergencyUnstaked events raised by the Token contract.
type TokenEmergencyUnstakedIterator struct {
	Event *TokenEmergencyUnstaked // Event containing the contract specifics and raw log

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
func (it *TokenEmergencyUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenEmergencyUnstaked)
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
		it.Event = new(TokenEmergencyUnstaked)
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
func (it *TokenEmergencyUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenEmergencyUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenEmergencyUnstaked represents a EmergencyUnstaked event raised by the Token contract.
type TokenEmergencyUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnstaked is a free log retrieval operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) FilterEmergencyUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*TokenEmergencyUnstakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenEmergencyUnstakedIterator{contract: _Token.contract, event: "EmergencyUnstaked", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnstaked is a free log subscription operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) WatchEmergencyUnstaked(opts *bind.WatchOpts, sink chan<- *TokenEmergencyUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "EmergencyUnstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenEmergencyUnstaked)
				if err := _Token.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
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

// ParseEmergencyUnstaked is a log parse operation binding the contract event 0x589526ce978dd18660ed3d203132d4f86762231a31e8fe21896e2ee637069551.
//
// Solidity: event EmergencyUnstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) ParseEmergencyUnstaked(log types.Log) (*TokenEmergencyUnstaked, error) {
	event := new(TokenEmergencyUnstaked)
	if err := _Token.contract.UnpackLog(event, "EmergencyUnstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMinClaimedBlocksUpdatedIterator is returned from FilterMinClaimedBlocksUpdated and is used to iterate over the raw logs and unpacked data for MinClaimedBlocksUpdated events raised by the Token contract.
type TokenMinClaimedBlocksUpdatedIterator struct {
	Event *TokenMinClaimedBlocksUpdated // Event containing the contract specifics and raw log

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
func (it *TokenMinClaimedBlocksUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinClaimedBlocksUpdated)
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
		it.Event = new(TokenMinClaimedBlocksUpdated)
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
func (it *TokenMinClaimedBlocksUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinClaimedBlocksUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinClaimedBlocksUpdated represents a MinClaimedBlocksUpdated event raised by the Token contract.
type TokenMinClaimedBlocksUpdated struct {
	New *big.Int
	Old *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinClaimedBlocksUpdated is a free log retrieval operation binding the contract event 0xe5da2dc1975dc4accc01d49cad544016c79215ed2e6ed8f843583deea3071e95.
//
// Solidity: event MinClaimedBlocksUpdated(uint256 indexed _new, uint256 indexed _old)
func (_Token *TokenFilterer) FilterMinClaimedBlocksUpdated(opts *bind.FilterOpts, _new []*big.Int, _old []*big.Int) (*TokenMinClaimedBlocksUpdatedIterator, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinClaimedBlocksUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return &TokenMinClaimedBlocksUpdatedIterator{contract: _Token.contract, event: "MinClaimedBlocksUpdated", logs: logs, sub: sub}, nil
}

// WatchMinClaimedBlocksUpdated is a free log subscription operation binding the contract event 0xe5da2dc1975dc4accc01d49cad544016c79215ed2e6ed8f843583deea3071e95.
//
// Solidity: event MinClaimedBlocksUpdated(uint256 indexed _new, uint256 indexed _old)
func (_Token *TokenFilterer) WatchMinClaimedBlocksUpdated(opts *bind.WatchOpts, sink chan<- *TokenMinClaimedBlocksUpdated, _new []*big.Int, _old []*big.Int) (event.Subscription, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinClaimedBlocksUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinClaimedBlocksUpdated)
				if err := _Token.contract.UnpackLog(event, "MinClaimedBlocksUpdated", log); err != nil {
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

// ParseMinClaimedBlocksUpdated is a log parse operation binding the contract event 0xe5da2dc1975dc4accc01d49cad544016c79215ed2e6ed8f843583deea3071e95.
//
// Solidity: event MinClaimedBlocksUpdated(uint256 indexed _new, uint256 indexed _old)
func (_Token *TokenFilterer) ParseMinClaimedBlocksUpdated(log types.Log) (*TokenMinClaimedBlocksUpdated, error) {
	event := new(TokenMinClaimedBlocksUpdated)
	if err := _Token.contract.UnpackLog(event, "MinClaimedBlocksUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPoolSyncedIterator is returned from FilterPoolSynced and is used to iterate over the raw logs and unpacked data for PoolSynced events raised by the Token contract.
type TokenPoolSyncedIterator struct {
	Event *TokenPoolSynced // Event containing the contract specifics and raw log

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
func (it *TokenPoolSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolSynced)
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
		it.Event = new(TokenPoolSynced)
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
func (it *TokenPoolSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPoolSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPoolSynced represents a PoolSynced event raised by the Token contract.
type TokenPoolSynced struct {
	Pool                common.Address
	AccumulatedPerShare *big.Int
	LastSyncedBlock     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPoolSynced is a free log retrieval operation binding the contract event 0xd247cac827464dafee8ef292281a2959c39bb57f4cfe99c6e4bde901bfff5781.
//
// Solidity: event PoolSynced(address _pool, uint256 _accumulatedPerShare, uint256 _lastSyncedBlock)
func (_Token *TokenFilterer) FilterPoolSynced(opts *bind.FilterOpts) (*TokenPoolSyncedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "PoolSynced")
	if err != nil {
		return nil, err
	}
	return &TokenPoolSyncedIterator{contract: _Token.contract, event: "PoolSynced", logs: logs, sub: sub}, nil
}

// WatchPoolSynced is a free log subscription operation binding the contract event 0xd247cac827464dafee8ef292281a2959c39bb57f4cfe99c6e4bde901bfff5781.
//
// Solidity: event PoolSynced(address _pool, uint256 _accumulatedPerShare, uint256 _lastSyncedBlock)
func (_Token *TokenFilterer) WatchPoolSynced(opts *bind.WatchOpts, sink chan<- *TokenPoolSynced) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "PoolSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPoolSynced)
				if err := _Token.contract.UnpackLog(event, "PoolSynced", log); err != nil {
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

// ParsePoolSynced is a log parse operation binding the contract event 0xd247cac827464dafee8ef292281a2959c39bb57f4cfe99c6e4bde901bfff5781.
//
// Solidity: event PoolSynced(address _pool, uint256 _accumulatedPerShare, uint256 _lastSyncedBlock)
func (_Token *TokenFilterer) ParsePoolSynced(log types.Log) (*TokenPoolSynced, error) {
	event := new(TokenPoolSynced)
	if err := _Token.contract.UnpackLog(event, "PoolSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenPoolWhitelistedIterator is returned from FilterPoolWhitelisted and is used to iterate over the raw logs and unpacked data for PoolWhitelisted events raised by the Token contract.
type TokenPoolWhitelistedIterator struct {
	Event *TokenPoolWhitelisted // Event containing the contract specifics and raw log

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
func (it *TokenPoolWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenPoolWhitelisted)
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
		it.Event = new(TokenPoolWhitelisted)
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
func (it *TokenPoolWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenPoolWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenPoolWhitelisted represents a PoolWhitelisted event raised by the Token contract.
type TokenPoolWhitelisted struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolWhitelisted is a free log retrieval operation binding the contract event 0xe6432174fa1253e6c61176cce257657dfbe2d386e217d070bd9360f1a76d2103.
//
// Solidity: event PoolWhitelisted(address _pool)
func (_Token *TokenFilterer) FilterPoolWhitelisted(opts *bind.FilterOpts) (*TokenPoolWhitelistedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "PoolWhitelisted")
	if err != nil {
		return nil, err
	}
	return &TokenPoolWhitelistedIterator{contract: _Token.contract, event: "PoolWhitelisted", logs: logs, sub: sub}, nil
}

// WatchPoolWhitelisted is a free log subscription operation binding the contract event 0xe6432174fa1253e6c61176cce257657dfbe2d386e217d070bd9360f1a76d2103.
//
// Solidity: event PoolWhitelisted(address _pool)
func (_Token *TokenFilterer) WatchPoolWhitelisted(opts *bind.WatchOpts, sink chan<- *TokenPoolWhitelisted) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "PoolWhitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenPoolWhitelisted)
				if err := _Token.contract.UnpackLog(event, "PoolWhitelisted", log); err != nil {
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

// ParsePoolWhitelisted is a log parse operation binding the contract event 0xe6432174fa1253e6c61176cce257657dfbe2d386e217d070bd9360f1a76d2103.
//
// Solidity: event PoolWhitelisted(address _pool)
func (_Token *TokenFilterer) ParsePoolWhitelisted(log types.Log) (*TokenPoolWhitelisted, error) {
	event := new(TokenPoolWhitelisted)
	if err := _Token.contract.UnpackLog(event, "PoolWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Token contract.
type TokenRewardClaimedIterator struct {
	Event *TokenRewardClaimed // Event containing the contract specifics and raw log

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
func (it *TokenRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRewardClaimed)
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
		it.Event = new(TokenRewardClaimed)
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
func (it *TokenRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRewardClaimed represents a RewardClaimed event raised by the Token contract.
type TokenRewardClaimed struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) FilterRewardClaimed(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*TokenRewardClaimedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenRewardClaimedIterator{contract: _Token.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *TokenRewardClaimed, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "RewardClaimed", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRewardClaimed)
				if err := _Token.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
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

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) ParseRewardClaimed(log types.Log) (*TokenRewardClaimed, error) {
	event := new(TokenRewardClaimed)
	if err := _Token.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenRewardPerBlockUpdatedIterator is returned from FilterRewardPerBlockUpdated and is used to iterate over the raw logs and unpacked data for RewardPerBlockUpdated events raised by the Token contract.
type TokenRewardPerBlockUpdatedIterator struct {
	Event *TokenRewardPerBlockUpdated // Event containing the contract specifics and raw log

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
func (it *TokenRewardPerBlockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRewardPerBlockUpdated)
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
		it.Event = new(TokenRewardPerBlockUpdated)
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
func (it *TokenRewardPerBlockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRewardPerBlockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRewardPerBlockUpdated represents a RewardPerBlockUpdated event raised by the Token contract.
type TokenRewardPerBlockUpdated struct {
	Pool         common.Address
	BlockRewards []IERC20StakingManagerBlockReward
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardPerBlockUpdated is a free log retrieval operation binding the contract event 0x1c1ff54b17c7e688c5c98e727cc95e474dd35e66e81fece270d178ea4c32324d.
//
// Solidity: event RewardPerBlockUpdated(address _pool, (uint256,uint256)[] _blockRewards)
func (_Token *TokenFilterer) FilterRewardPerBlockUpdated(opts *bind.FilterOpts) (*TokenRewardPerBlockUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "RewardPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenRewardPerBlockUpdatedIterator{contract: _Token.contract, event: "RewardPerBlockUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardPerBlockUpdated is a free log subscription operation binding the contract event 0x1c1ff54b17c7e688c5c98e727cc95e474dd35e66e81fece270d178ea4c32324d.
//
// Solidity: event RewardPerBlockUpdated(address _pool, (uint256,uint256)[] _blockRewards)
func (_Token *TokenFilterer) WatchRewardPerBlockUpdated(opts *bind.WatchOpts, sink chan<- *TokenRewardPerBlockUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "RewardPerBlockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRewardPerBlockUpdated)
				if err := _Token.contract.UnpackLog(event, "RewardPerBlockUpdated", log); err != nil {
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

// ParseRewardPerBlockUpdated is a log parse operation binding the contract event 0x1c1ff54b17c7e688c5c98e727cc95e474dd35e66e81fece270d178ea4c32324d.
//
// Solidity: event RewardPerBlockUpdated(address _pool, (uint256,uint256)[] _blockRewards)
func (_Token *TokenFilterer) ParseRewardPerBlockUpdated(log types.Log) (*TokenRewardPerBlockUpdated, error) {
	event := new(TokenRewardPerBlockUpdated)
	if err := _Token.contract.UnpackLog(event, "RewardPerBlockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Token contract.
type TokenStakedIterator struct {
	Event *TokenStaked // Event containing the contract specifics and raw log

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
func (it *TokenStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenStaked)
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
		it.Event = new(TokenStaked)
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
func (it *TokenStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenStaked represents a Staked event raised by the Token contract.
type TokenStaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) FilterStaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*TokenStakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenStakedIterator{contract: _Token.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *TokenStaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Staked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenStaked)
				if err := _Token.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) ParseStaked(log types.Log) (*TokenStaked, error) {
	event := new(TokenStaked)
	if err := _Token.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Token contract.
type TokenUnstakedIterator struct {
	Event *TokenUnstaked // Event containing the contract specifics and raw log

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
func (it *TokenUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUnstaked)
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
		it.Event = new(TokenUnstaked)
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
func (it *TokenUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUnstaked represents a Unstaked event raised by the Token contract.
type TokenUnstaked struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) FilterUnstaked(opts *bind.FilterOpts, _user []common.Address, _token []common.Address, _amount []*big.Int) (*TokenUnstakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenUnstakedIterator{contract: _Token.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *TokenUnstaked, _user []common.Address, _token []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Unstaked", _userRule, _tokenRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUnstaked)
				if err := _Token.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed _user, address indexed _token, uint256 indexed _amount)
func (_Token *TokenFilterer) ParseUnstaked(log types.Log) (*TokenUnstaked, error) {
	event := new(TokenUnstaked)
	if err := _Token.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenUserRewardUpdatedIterator is returned from FilterUserRewardUpdated and is used to iterate over the raw logs and unpacked data for UserRewardUpdated events raised by the Token contract.
type TokenUserRewardUpdatedIterator struct {
	Event *TokenUserRewardUpdated // Event containing the contract specifics and raw log

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
func (it *TokenUserRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenUserRewardUpdated)
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
		it.Event = new(TokenUserRewardUpdated)
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
func (it *TokenUserRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenUserRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenUserRewardUpdated represents a UserRewardUpdated event raised by the Token contract.
type TokenUserRewardUpdated struct {
	Pool       common.Address
	User       common.Address
	RewardInfo IERC20StakingManagerUserReward
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserRewardUpdated is a free log retrieval operation binding the contract event 0x027f73145bb86dfcdffa5fae931b3cab5ab93c376099cc84b6d2e4985f10e14b.
//
// Solidity: event UserRewardUpdated(address _pool, address _user, (uint256,uint256,uint256) _rewardInfo)
func (_Token *TokenFilterer) FilterUserRewardUpdated(opts *bind.FilterOpts) (*TokenUserRewardUpdatedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "UserRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenUserRewardUpdatedIterator{contract: _Token.contract, event: "UserRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRewardUpdated is a free log subscription operation binding the contract event 0x027f73145bb86dfcdffa5fae931b3cab5ab93c376099cc84b6d2e4985f10e14b.
//
// Solidity: event UserRewardUpdated(address _pool, address _user, (uint256,uint256,uint256) _rewardInfo)
func (_Token *TokenFilterer) WatchUserRewardUpdated(opts *bind.WatchOpts, sink chan<- *TokenUserRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "UserRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenUserRewardUpdated)
				if err := _Token.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
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

// ParseUserRewardUpdated is a log parse operation binding the contract event 0x027f73145bb86dfcdffa5fae931b3cab5ab93c376099cc84b6d2e4985f10e14b.
//
// Solidity: event UserRewardUpdated(address _pool, address _user, (uint256,uint256,uint256) _rewardInfo)
func (_Token *TokenFilterer) ParseUserRewardUpdated(log types.Log) (*TokenUserRewardUpdated, error) {
	event := new(TokenUserRewardUpdated)
	if err := _Token.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
