// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Context.sol";
import "@openzeppelin/contracts/utils/Counters.sol";

contract ERC721Minter is Context, ERC721URIStorage {
    address public current_minter = address(0);
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    modifier onlyMinter() {
        require(
            current_minter == _msgSender(),
            "onlyMinter: caller is not the minter"
        );
        _;
    }

    constructor(
        string memory name,
        string memory symbol,
        address minter
    ) ERC721(name, symbol) {
        require(minter != address(0), "ERROR: Zero address");
        current_minter = minter;
    }

    function mint(address player, string memory tokenURI)
        external
        onlyMinter
        returns (uint256)
    {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _mint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    function burn(uint256 tokenId) external onlyMinter {
        //solhint-disable-next-line max-line-length
        require(
            _isApprovedOrOwner(_msgSender(), tokenId),
            "ERC721Burnable: caller is not owner nor approved"
        );
        _burn(tokenId);
    }

    function replaceMinter(address newMinter) external onlyMinter {
        current_minter = newMinter;
    }

    function _transfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual override(ERC721) {
        super._transfer(from, to, tokenId);
        if (_msgSender() != current_minter && to == current_minter) {
            _burn(tokenId);
        }
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual override(ERC721) {
        super._beforeTokenTransfer(from, to, tokenId);
    }
}
