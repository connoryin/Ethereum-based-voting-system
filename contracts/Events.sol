pragma solidity ^0.8.10;

// import './Poll.sol';


contract Events {
    struct pollInfo {
        uint owner;
        uint eventId;
        address addr;
        bool exists;
        bool isEnd;
    }

    mapping(bytes32 => pollInfo) events;
    //    address owner;

    mapping(uint => pollInfo[]) adminEvents;

    //    constructor (uint adminID) {
    //        owner = adminID;
    //    }

    function uploadPollInfo(bytes32[]memory invCodes, pollInfo memory info, uint adminID) public {
        //        require(owner == msg.sender, "sender is not owner.");
        for (uint i = 0; i < invCodes.length; i++) {
            events[invCodes[i]] = info;
            events[invCodes[i]].exists = true;
        }
        adminEvents[adminID].push(info);
    }

    function getPollInfoByInvCode(bytes32 invCode) public view returns (uint, uint, bool, address){
        require(events[invCode].exists, "Event does not exist");
        return (events[invCode].owner, events[invCode].eventId, events[invCode].isEnd, events[invCode].addr);
    }

    function getPollInfoByAdminID(uint adminID) public view returns (pollInfo[] memory){
        return adminEvents[adminID];
    }
}