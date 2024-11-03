// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SankoPredicts is ReentrancyGuard, Ownable {
    using SafeERC20 for IERC20;

    IERC20 public spToken;
    uint256 public feePercentage = 42; // Default fee percentage (4.2%)
    uint256 public creatorFeePercentile = 100; // Creator fee percentage (10% represented as 100 basis points)

    enum GameType {
        Election,
        Coin
    }

    struct Game {
        uint256 id;
        string gameName;
        string banner;
        GameType gameType;
        string candidateA;
        string candidateB;
        uint256 initialPrice;
        uint256 finalPrice;
        uint256 expiryTime;
        uint256 lockTime;
        uint256 totalPoolA;
        uint256 totalPoolB;
        bool resolved;
        uint256 winnerPool;
        uint256 feePercent;
        string market;
        address creator;
        uint status; // 1 pending 2 rejected 3 approved
    }

    struct GameDeposit {
        address user;
        uint256 amount;
        uint256 timestamp;
        uint8 pool;
    }

    struct LeaderboardEntry {
        address player;
        uint256 totalDeposits;
        uint256 totalWinnings;
    }

    mapping(uint256 => Game) public games;
    mapping(uint256 => mapping(address => uint256)) public depositsA;
    mapping(uint256 => mapping(address => uint256)) public depositsB;

    mapping(uint256 => GameDeposit[]) public depositHistory;
    mapping(address => LeaderboardEntry) public leaderboard;

    address[] public players;

    uint256 public gameCount;

    constructor(IERC20 _spToken) Ownable(msg.sender) {
        spToken = _spToken;
    }

    event GameCreated(
        uint256 gameId,
        string gameName,
        GameType gameType,
        string candidateA,
        string candidateB,
        uint256 initialPrice,
        uint256 expiryTime,
        uint256 lockTime,
        uint256 feePercent,
        string banner,
        address creator,
        uint status
    );

    event Deposit(
        address indexed user,
        uint256 gameId,
        uint256 amount,
        uint8 option
    );

    event GameResolved(uint256 gameId, uint8 winner);
    event Payout(address indexed user, uint256 gameId, uint256 amount);
    event FeeUpdated(uint256 gameId, uint256 newFeePercent);
    event FeeUpdated(uint256 newfeePercentage);

    modifier gameNotLocked(uint256 gameId) {
        require(
            block.timestamp < games[gameId].lockTime,
            "Deposits are locked for this game"
        );
        _;
    }

    modifier gameApproved(uint256 gameId) {
        require(games[gameId].status == 3, "Game not approved");
        _;
    }

    modifier gameNotExpired(uint256 gameId) {
        require(block.timestamp < games[gameId].expiryTime, "game has expired");
        _;
    }

    modifier gameResolved(uint256 gameId) {
        require(games[gameId].resolved, "Game has not been resolved yet");
        _;
    }

    function _createGame(
        string memory _gameName,
        GameType _gameType,
        string memory _candidateA,
        string memory _candidateB,
        uint256 _initialPrice,
        uint256 _expiryTime,
        uint256 _lockTime,
        string memory _market,
        string memory _banner,
        address _creator,
        uint _status
    ) internal {
        require(_expiryTime > block.timestamp, "Expiry must be in the future");
        require(_lockTime < _expiryTime, "Lock time must be before expiry");

        gameCount++;
        games[gameCount] = Game({
            id: gameCount,
            gameName: _gameName,
            gameType: _gameType,
            candidateA: _candidateA,
            candidateB: _candidateB,
            initialPrice: _initialPrice,
            finalPrice: 0,
            expiryTime: _expiryTime,
            lockTime: _lockTime,
            totalPoolA: 0,
            totalPoolB: 0,
            resolved: false,
            winnerPool: 0,
            feePercent: feePercentage,
            market: _market,
            banner: _banner,
            creator: _creator,
            status: _status
        });

        emit GameCreated(
            gameCount,
            _gameName,
            _gameType,
            _candidateA,
            _candidateB,
            _initialPrice,
            _expiryTime,
            _lockTime,
            feePercentage,
            _banner,
            _creator,
            _status
        );
    }

    function createElectionGame(
        string memory _gameName,
        uint256 _expiryTime,
        uint256 _lockTime,
        string memory _market,
        string memory _candidateA,
        string memory _candidateB,
        string memory _banner,
        address _creator
    ) public {
        _createGame(
            _gameName,
            GameType.Election,
            _candidateA,
            _candidateB,
            0,
            _expiryTime,
            _lockTime,
            _market,
            _banner,
            owner() == msg.sender ? owner() : _creator,
            owner() == msg.sender ? 3 : 1
        );
    }

    function createCoinGame(
        string memory _gameName,
        uint256 _expiryTime,
        uint256 _lockTime,
        string memory _market,
        uint256 _initialPrice,
        string memory _banner,
        address _creator
    ) public {
        _createGame(
            _gameName,
            GameType.Coin,
            "", // No candidates for Coin game
            "",
            _initialPrice,
            _expiryTime,
            _lockTime,
            _market,
            _banner,
            owner() == msg.sender ? owner() : _creator,
            owner() == msg.sender ? 3 : 1
        );
    }

    function updateGameStatus(uint256 gameId, uint newStatus) public onlyOwner {
        Game storage game = games[gameId];
        game.status = newStatus;
    }

    function resolveGame(
        uint256 gameId,
        uint8 winner,
        uint256 _finalPrice
    ) public gameApproved(gameId) onlyOwner {
        Game storage game = games[gameId];
        require(!game.resolved, "game already resolved");
        require(winner == 1 || winner == 2, "Invalid winner");

        game.resolved = true;
        game.winnerPool = winner;

        if (game.gameType == GameType.Coin) {
            game.finalPrice = _finalPrice;
        }

        uint256 losingPool = winner == 1 ? game.totalPoolB : game.totalPoolA;

        uint256 totalFee = (losingPool * feePercentage) / 1000;
        uint256 creatorFee = (totalFee * creatorFeePercentile) / 1000;
        uint256 ownerFee = totalFee - creatorFee;

        if (game.creator == owner()) {
            spToken.safeTransfer(owner(), totalFee);
        } else {
            spToken.safeTransfer(owner(), ownerFee);
            spToken.safeTransfer(game.creator, creatorFee);
        }

        emit GameResolved(gameId, winner);
    }

    function getGameDetails(
        uint256 gameId
    ) public view returns (Game memory game, GameDeposit[] memory deposits) {
        game = games[gameId];
        deposits = depositHistory[gameId];
    }

    function getGames(
        bool includeResolved,
        uint status,
        address creator
    ) public view returns (Game[] memory) {
        uint256 count = 0;
        for (uint256 i = 1; i <= gameCount; i++) {
            if (
                games[i].resolved == includeResolved &&
                games[i].status == status &&
                (creator == address(0) || games[i].creator == creator)
            ) {
                count++;
            }
        }

        Game[] memory gamesList = new Game[](count);

        uint256 index = 0;
        for (uint256 i = 1; i <= gameCount; i++) {
            if (
                games[i].resolved == includeResolved &&
                games[i].status == status &&
                (creator == address(0) || games[i].creator == creator)
            ) {
                gamesList[index] = games[i];
                index++;
            }
        }

        return gamesList;
    }

    function deposit(
        uint256 gameId,
        uint8 option,
        uint256 amount
    )
        public
        gameApproved(gameId)
        gameNotLocked(gameId)
        gameNotExpired(gameId)
        nonReentrant
    {
        require(option == 1 || option == 2, "Invalid option");
        require(amount > 0, "Deposit must be greater than 0");

        spToken.safeTransferFrom(msg.sender, address(this), amount);

        depositHistory[gameId].push(
            GameDeposit({
                user: msg.sender,
                amount: amount,
                timestamp: block.timestamp,
                pool: option
            })
        );

        if (option == 1) {
            depositsA[gameId][msg.sender] += amount;
            games[gameId].totalPoolA += amount;
        } else {
            depositsB[gameId][msg.sender] += amount;
            games[gameId].totalPoolB += amount;
        }

        if (leaderboard[msg.sender].player == address(0)) {
            players.push(msg.sender);
            leaderboard[msg.sender].player = msg.sender;
        }

        leaderboard[msg.sender].totalDeposits += amount;

        emit Deposit(msg.sender, gameId, amount, option);
    }

    function getDepositsForGame(
        uint256 gameId
    ) public view returns (GameDeposit[] memory) {
        return depositHistory[gameId];
    }

    function getClaimableOrPendingGames(
        address user
    )
        public
        view
        returns (
            uint256[] memory claimableGameIds,
            uint256[] memory pendingGameIds
        )
    {
        uint256 claimableCount = 0;
        uint256 pendingCount = 0;

        for (uint256 i = 1; i <= gameCount; i++) {
            Game storage game = games[i];

            if (!game.resolved) {
                if (depositsA[i][user] > 0 || depositsB[i][user] > 0) {
                    pendingCount++;
                }
            } else {
                if (
                    (game.winnerPool == 1 && depositsA[i][user] > 0) ||
                    (game.winnerPool == 2 && depositsB[i][user] > 0)
                ) {
                    claimableCount++;
                }
            }
        }

        claimableGameIds = new uint256[](claimableCount);
        pendingGameIds = new uint256[](pendingCount);

        uint256 claimableIndex = 0;
        uint256 pendingIndex = 0;

        for (uint256 i = 1; i <= gameCount; i++) {
            Game storage game = games[i];

            if (!game.resolved) {
                if (depositsA[i][user] > 0 || depositsB[i][user] > 0) {
                    pendingGameIds[pendingIndex++] = i;
                }
            } else {
                if (
                    (game.winnerPool == 1 && depositsA[i][user] > 0) ||
                    (game.winnerPool == 2 && depositsB[i][user] > 0)
                ) {
                    claimableGameIds[claimableIndex++] = i;
                }
            }
        }
    }

    function claim(uint256 gameId) public gameResolved(gameId) nonReentrant {
        uint256 userShare = 0;
        Game storage game = games[gameId];

        if (game.winnerPool == 1) {
            uint256 userDeposit = depositsA[gameId][msg.sender];
            require(userDeposit > 0, "No deposit in winning pool");
            uint256 totalWinningPool = game.totalPoolA;
            uint256 losingPool = game.totalPoolB;
            uint256 fee = (losingPool * game.feePercent) / 1000;
            uint256 rewardPool = losingPool - fee;

            userShare =
                userDeposit +
                (userDeposit * rewardPool) /
                totalWinningPool;
            depositsA[gameId][msg.sender] = 0;
        } else if (game.winnerPool == 2) {
            uint256 userDeposit = depositsB[gameId][msg.sender];
            require(userDeposit > 0, "No deposit in winning pool");
            uint256 totalWinningPool = game.totalPoolB;
            uint256 losingPool = game.totalPoolA;
            uint256 fee = (losingPool * game.feePercent) / 1000; // Use 1000 for percentage
            uint256 rewardPool = losingPool - fee;

            userShare =
                userDeposit +
                (userDeposit * rewardPool) /
                totalWinningPool;
            depositsB[gameId][msg.sender] = 0;
        }

        leaderboard[msg.sender].totalWinnings += userShare;
        spToken.safeTransfer(msg.sender, userShare);
        emit Payout(msg.sender, gameId, userShare);
    }

    function updateCreatorFeePercentage(
        uint256 _newPercentile
    ) external onlyOwner {
        require(_newPercentile <= 1000, "Creator fee cannot exceed 100%");
        creatorFeePercentile = _newPercentile;
    }

    function updateFeePercentage(uint256 newfeePercentage) public onlyOwner {
        require(
            newfeePercentage <= 1000,
            "Default fee percent must be <= 1000"
        );
        feePercentage = newfeePercentage;
        emit FeeUpdated(newfeePercentage);
    }

    function withdraw(uint256 amount) public onlyOwner nonReentrant {
        require(
            spToken.balanceOf(address(this)) >= amount,
            "Insufficient balance"
        );
        spToken.safeTransfer(owner(), amount);
    }

    function getLeaderboard() public view returns (LeaderboardEntry[] memory) {
        uint256 entryCount = 0;

        for (uint256 i = 0; i < players.length; i++) {
            if (
                leaderboard[players[i]].totalWinnings > 0 ||
                leaderboard[players[i]].totalDeposits > 0
            ) {
                entryCount++;
            }
        }

        LeaderboardEntry[] memory entries = new LeaderboardEntry[](entryCount);
        uint256 index = 0;

        for (uint256 i = 0; i < players.length; i++) {
            if (
                leaderboard[players[i]].totalWinnings > 0 ||
                leaderboard[players[i]].totalDeposits > 0
            ) {
                entries[index] = LeaderboardEntry({
                    player: players[i],
                    totalDeposits: leaderboard[players[i]].totalDeposits,
                    totalWinnings: leaderboard[players[i]].totalWinnings
                });
                index++;
            }
        }

        return entries;
    }
}
