// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract ERC1155Minter is ERC1155 {
    address public current_minter = address(0);

    modifier onlyMinter() {
        require(
            current_minter == _msgSender(),
            "onlyMinter: caller is not the minter"
        );
        _;
    }

    constructor(string memory uri_, address minter) ERC1155(uri_) {
        require(minter != address(0), "ERROR: Zero address");
        current_minter = minter;
    }

    function mint(
        address account,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) external onlyMinter {
       _mint(account, id, amount, data);
    }

     function mintBatch(
        address to,
        uint256[] memory ids,
        uint256[] memory amounts,
        bytes memory data
    ) external onlyMinter {
        _mintBatch(to, ids, amounts, data);
    }

    function burn(
        address account,
        uint256 id,
        uint256 value
    ) external onlyMinter {
        _burn(account, id, value);
    }

    function burnBatch(
        address account,
        uint256[] memory ids,
        uint256[] memory values
    ) external onlyMinter {
        _burnBatch(account, ids, values);
    }

    function replaceMinter(address newMinter) external onlyMinter {
        current_minter = newMinter;
    }

}
