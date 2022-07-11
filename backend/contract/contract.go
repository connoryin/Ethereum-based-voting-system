package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/model"
	"github.com/6675-voting-system/voting-system/backend/poll"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var addressEvents string = "0x7445f55497b20C05C2695CfaCC8ce6d8F0cAA9eC"
var eventAddress common.Address = common.HexToAddress(addressEvents)

var URL string = "http://localhost:7545"

// var URL string = "https://eth-ropsten.alchemyapi.io/v2/UDpGK4GHxpl4Txw-uKNwqGpeW2nEeAsB"
var privateKey string = "57020558a5472ef3bc9b0c8dffb27f8333469402864eacbd5fbc565d48ec7f82"
var publicKey string = "publicKey"
var publicKeyECDSA *ecdsa.PublicKey
var privateKeyECDSA *ecdsa.PrivateKey
var Client *ethclient.Client
var EventsAddress common.Address

type AdminEvents struct {
	instanceEvents *poll.Events
	events         map[int]Ballot
}

type Ballot struct {
	// url             string
	address      common.Address
	owner        *big.Int
	instancePoll *poll.Poll
	event        data_access.EventPO
	// publicKeyECDSA  *ecdsa.PublicKey
	// privateKeyECDSA *ecdsa.PrivateKey
	// privateKey      string
}

func SetupContract() {
	SetUpECDSA()
	SetUpClient()
}

func SetUpClient() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	var err error
	Client, err = ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}
}

func (ballot *Ballot) SetUpAddress(addressHex string) {
	ballot.address = common.HexToAddress(addressHex)
}

func (ballot *Ballot) DeployPoll(auth *bind.TransactOpts, owner *big.Int, _eventId *big.Int, _totalNumVoter *big.Int, _numBallotPerVoter *big.Int, _voters []poll.PollVoter, _candidates []poll.PollCandidate, _isAnonymous bool) common.Address {
	address, transaction, instancePoll, err := poll.DeployPoll(auth, Client, owner, _eventId, _totalNumVoter, _numBallotPerVoter, _voters, _candidates, _isAnonymous)
	if err != nil {
		log.Fatal(err)
	}
	ballot.owner = owner
	ballot.address = address
	ballot.instancePoll = instancePoll

	fmt.Printf("address: %v\n", address)
	fmt.Printf("transaction: %v\n", transaction)
	return address
}

func (adminEvents *AdminEvents) Init() {
	adminEvents.events = make(map[int]Ballot)
}

func (adminEvents *AdminEvents) DeployEvents(auth *bind.TransactOpts) {
	// addressEvents, _, instanceEvents, err := poll.DeployEvents(auth, )
	instanceEvents, err := poll.NewEvents(eventAddress, Client)
	if err != nil {
		log.Fatal(err)
	}
	adminEvents.instanceEvents = instanceEvents
}

func (adminEvents *AdminEvents) loadEventInstances() {
	instanceEvents, err := poll.NewEvents(EventsAddress, Client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded.")
	adminEvents.instanceEvents = instanceEvents
}

func (ballot *Ballot) LoadPoll() {
	instancePoll, err := poll.NewPoll(ballot.address, Client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded.")
	ballot.instancePoll = instancePoll
}

func SetUpECDSA() {
	var err error
	privateKeyECDSA, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKeyECDSA.Public()
	var ok bool
	publicKeyECDSA, ok = publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// ballot.publicKeyECDSA = publicKeyECDSA
}

//writing to a smart contract
//requires us to sign the sign transaction with our private key

func SetAuthOptions() (*bind.TransactOpts, *ethclient.Client) {
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKeyECDSA)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(40000000) // in units
	auth.GasPrice = gasPrice

	return auth, Client
}

func (ballot *Ballot) EndVote(auth *bind.TransactOpts, adminID int) (*types.Transaction, error) {
	transaction, err := ballot.instancePoll.EndVote(auth, big.NewInt(int64(adminID)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return transaction, err
}

func (ballot *Ballot) Vote(candidatesID []*big.Int, voterID [32]byte, auth *bind.TransactOpts) *types.Transaction {
	transaction, err := ballot.instancePoll.Vote(auth, candidatesID, voterID)
	if err != nil {
		log.Fatal(err)
	}
	return transaction
}

func (adminEvents *AdminEvents) GetInvCodeEvents(invCode string) (*big.Int, *big.Int, bool, common.Address) {
	owner, eventId, isEnd, addr, err := adminEvents.instanceEvents.GetPollInfoByInvCode(nil, str2byte32(invCode))
	if err != nil {
		log.Fatal(err)
	}
	return owner, eventId, isEnd, addr
}

func (ballot *Ballot) GetWinner() (*big.Int, poll.PollCandidate) {
	winnerId, winnerCandidate, err := ballot.instancePoll.GetWinner(nil)
	if err != nil {
		log.Fatal(err)
	}
	return winnerId, winnerCandidate
}

func (ballot *Ballot) GetPollInfo(event *model.Event) {
	finished, numVoterVoted, totalNumVoter, numBallotPerVoter, isAnonymous, err := ballot.instancePoll.GetInfo(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("isAnonymous %v", isAnonymous)

	event.IsEnd = finished
	event.MaxVoteNumPerPerson = int(numBallotPerVoter.Int64())
	event.ReceivedVoteNum = int(numVoterVoted.Int64())
	event.TotalVoteNum = int(totalNumVoter.Int64())
	event.IsAnonymous = isAnonymous

	// event.Candidates = make([]model.Candidate, 0)
	// for _, candidate := range candidates {
	// 	event.Candidates = append(event.Candidates, model.Candidate{
	// 		Id:         int(candidate.Id.Int64()),
	// 		Name:       string(candidate.Name[:]),
	// 		VoteNumGet: int(candidate.VoteNumGet.Int64()),
	// 		IsWinner:   candidate.IsWinner,
	// 	})
	// }
	// return finished, numVoterVoted, totalNumVoter, numBallotPerVoter, isAnonymous
}

func (ballot *Ballot) GetVoterInfo(invCode string) poll.PollVoter {
	voterInfo, _, err := ballot.instancePoll.GetVoterInfo(nil, str2byte32(invCode))
	if err != nil {
		log.Fatal(err)
	}
	return voterInfo
}

func (ballot *Ballot) GetAllVoterDetails() ([]poll.PollVoter, [][][32]byte, error) {
	voters, candidateNames, err := ballot.instancePoll.GetAllVoterDetails(nil)
	return voters, candidateNames, err
}

func (ballot *Ballot) GetCandidateInfo(candidateId int) poll.PollCandidate {
	candidate := big.NewInt(int64(candidateId))
	candidateInfo, err := ballot.instancePoll.GetCandidateInfo(nil, candidate)
	if err != nil {
		log.Fatal(err)
	}
	return candidateInfo
}
