// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.9;

// Import this file to use console.log
import "hardhat/console.sol";

contract Todo {
    address payable public owner;
    enum Status {
        NEW,
        DONE,
        DELETED
    }
    struct TodoItem {
        uint id;
        Status status;
        string title;
        string description;
        address createdBy;
    }
    uint public noOfTodos;
    mapping(uint => TodoItem) todos;
    mapping(uint => address) owners;

    constructor() payable {
        owner = payable(msg.sender);
    }

    modifier onlyCreater(uint id) {
        require(msg.sender == owners[id], "You are not allowed to delete");
        _;
    }

    function createTodo(string memory title, string memory description)
        public
        returns (uint)
    {
        noOfTodos += 1;
        TodoItem memory todo = TodoItem(
            noOfTodos,
            Status.NEW,
            title,
            description,
            msg.sender
        );

        todos[noOfTodos] = todo;
        owners[noOfTodos] = msg.sender;

        return noOfTodos;
    }

    function getTodo(uint id) public view returns (TodoItem memory) {
        return todos[id];
    }

    function deleteTodo(uint id)
        public
        onlyCreater(id)
        returns (TodoItem memory)
    {
        todos[id].status = Status.DELETED;
        return todos[id];
    }

    function updateTodo(
        uint id,
        string memory title,
        string memory description
    ) public onlyCreater(id) returns (TodoItem memory) {
        todos[id].title = title;
        todos[id].description = description;
        return todos[id];
    }

    function updateStatus(uint id, Status status)
        public
        onlyCreater(id)
        returns (TodoItem memory)
    {
        todos[id].status = status;
        return todos[id];
    }

    function getTodos() public view returns (TodoItem[] memory) {
        TodoItem[] memory ltodos = new TodoItem[](noOfTodos);
        for (uint i = 1; i <= noOfTodos; i++) {
            TodoItem storage lBid = todos[i];
            ltodos[i-1] = lBid;
        }
        return ltodos;
    }
}
