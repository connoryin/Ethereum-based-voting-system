pragma solidity ^0.8.10;

contract Poll {
    uint owner;
    uint eventId;
    bool finished;
    Candidate[] candidates;
    // mapping(bytes32 => Candidate) candidates;
    // Voter[] voters;
    // mapping(eventid=>event)
    mapping(bytes32 => Voter) voters;
    // uint[] counts;
    uint numVoterVoted;
    uint totalNumVoter;
    uint numBallotPerVoter;
    bool isAnonymous;
    uint winnerId;
    uint voterNum;

    bytes32[] voterCodes;

    struct Voter {
        bytes32 id;
        bool voted;
        uint[] votedFor;
        bool exists;
    }

    struct Candidate {
        uint id;
        bytes32 name;
        uint voteNumGet;
        bool isWinner;
        bool exists;
    }

    // mapping(uint => Voter) voters;

    constructor (uint _owner, uint _eventId, uint _totalNumVoter, uint _numBallotPerVoter, Voter[] memory _voters, Candidate[] memory _candidates, bool _isAnonymous) {
        owner = _owner;
        voterNum = 0;
        eventId = _eventId;
        totalNumVoter = _totalNumVoter;
        numBallotPerVoter = _numBallotPerVoter;
        for (uint i = 0; i < _candidates.length; i++) {
            candidates.push(_candidates[i]);
        }
        // for (uint i = 0; i < _candidates.length; i++) {
        //     bytes32 candidateID = _candidates[i].id;
        //     candidates[candidateID] = _candidates[i];
        //     candidates[candidateID].exists = true;
        // }
        // candidates = _candidates;

        for (uint i = 0; i < _voters.length; i++) {
            bytes32 voterID = _voters[i].id;
            voters[voterID] = _voters[i];
            voters[voterID].exists = true;
            voterNum += 1;
        }
        isAnonymous = _isAnonymous;
        if (isAnonymous == false) {
            for (uint i = 0; i < _voters.length; i++) {
                voterCodes.push(_voters[i].id);
            }
        }
    }

    function vote(uint[] memory candidatesID, bytes32 voterID) public {
        require(voters[voterID].exists, "Cannot find voter!");
        require(!voters[voterID].voted, "Voter has already voted!");
        require(!finished, "Voting has already finished!");
        voters[voterID].votedFor = candidatesID;
        voters[voterID].voted = true;
        numVoterVoted += 1;
        for (uint i = 0; i < candidatesID.length; i++) {
            candidates[candidatesID[i]].voteNumGet += 1;
        }
    }

    function endVote(uint adminID) public {
        require(!finished, "Voting has already finished!");
        require(adminID == owner, "Only creator can stop the vote!");
        finished = true;
        uint maxVotes = 0;
        winnerId = 0;
        for (uint i = 0; i < candidates.length; i++) {
            Candidate memory candidate = candidates[i];
            if (candidate.voteNumGet > maxVotes) {
                maxVotes = candidate.voteNumGet;
                winnerId = i;
            }
        }
    }

    function getWinner() public view returns (uint, Candidate memory){
        require(finished, "Voting is not completed yet!");
        return (winnerId, candidates[winnerId]);
    }

    function getInfo() public view returns (bool, uint, uint, uint, bool) {
        return (finished, numVoterVoted, totalNumVoter, numBallotPerVoter, isAnonymous);
    }

    function getVoterInfo(bytes32 voterId) public view returns (Voter memory, Candidate[] memory){
        Voter memory voter = voters[voterId];
        require(voter.exists, "Voter does not exist");
        // if (isAnonymous){
        //     require(owner != msg.sender, "Cannot access by owner");
        // }
        Candidate[] memory candidateVoted = new Candidate[](voter.votedFor.length);
        for (uint i = 0; i < voter.votedFor.length; i++) {
            candidateVoted[i] = candidates[voter.votedFor[i]];
        }
        return (voter, candidateVoted);
    }

    function getCandidateInfo(uint candidateId) public view returns (Candidate memory){
        Candidate memory candidate = candidates[candidateId];
        require(candidate.exists, "Candidate does not exist");
        return candidate;
    }

    function getAllVoterDetails() public view returns (Voter[] memory, bytes32[][] memory candidateNames){
        require(!isAnonymous, "anonymous details are not visible.");
        Voter[] memory results = new Voter[](voterNum);
        bytes32[][] memory candidateResults = new bytes32[][](voterNum);
        for (uint i = 0; i < voterCodes.length; i++) {
            results[i] = voters[voterCodes[i]];
            bytes32[] memory candidateBytes = new bytes32[](results[i].votedFor.length);
            for (uint j = 0; j < results[i].votedFor.length; j++) {
                candidateBytes[j] = candidates[results[i].votedFor[j]].name;
            }
            candidateResults[i] = candidateBytes;
        }
        return (results, candidateResults);
    }
}
