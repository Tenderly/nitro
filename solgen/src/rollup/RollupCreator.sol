// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.0;

import "../bridge/Bridge.sol";
import "../bridge/SequencerInbox.sol";
import "../bridge/Inbox.sol";
import "../bridge/Outbox.sol";
import "./RollupEventBridge.sol";
import "./BridgeCreator.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

import "./Rollup.sol";
import "./RollupUserLogic.sol";
import "./RollupAdminLogic.sol";
import "../bridge/IBridge.sol";

import "./RollupLib.sol";
import "../utils/ICloneable.sol";

contract RollupCreator is Ownable {
    event RollupCreated(address indexed rollupAddress, address inboxAddress, address adminProxy);
    event TemplatesUpdated();

    BridgeCreator public bridgeCreator;
    ICloneable public rollupTemplate;
    address public challengeFactory;
    address public rollupAdminLogic;
    address public rollupUserLogic;

    constructor() Ownable() {}

    function setTemplates(
        BridgeCreator _bridgeCreator,
        ICloneable _rollupTemplate,
        address _challengeFactory,
        address _rollupAdminLogic,
        address _rollupUserLogic
    ) external onlyOwner {
        bridgeCreator = _bridgeCreator;
        rollupTemplate = _rollupTemplate;
        challengeFactory = _challengeFactory;
        rollupAdminLogic = _rollupAdminLogic;
        rollupUserLogic = _rollupUserLogic;
        emit TemplatesUpdated();
    }

    // sequencerInboxParams = [ maxDelayBlocks, maxFutureBlocks, maxDelaySeconds, maxFutureSeconds ]
    function createRollup(
        uint256 confirmPeriodBlocks,
        uint256 extraChallengeTimeBlocks,
        address stakeToken,
        uint256 baseStake,
        bytes32 wasmModuleRoot,
        address owner,
        uint256 chainId,
        uint256[4] memory sequencerInboxParams
    ) external returns (address) {
        return
            createRollup(
                RollupLib.Config(
                    confirmPeriodBlocks,
                    extraChallengeTimeBlocks,
                    stakeToken,
                    baseStake,
                    wasmModuleRoot,
                    owner,
                    chainId,
                    sequencerInboxParams
                )
            );
    }

    struct CreateRollupFrame {
        ProxyAdmin admin;
        Bridge delayedBridge;
        SequencerInbox sequencerInbox;
        Inbox inbox;
        RollupEventBridge rollupEventBridge;
        Outbox outbox;
        address rollup;
    }

    // After this setup:
    // Rollup should be the owner of bridge
    // RollupOwner should be the owner of Rollup's ProxyAdmin
    // RollupOwner should be the owner of Rollup
    // Bridge should have a single inbox and outbox
    function createRollup(RollupLib.Config memory config) private returns (address) {
        CreateRollupFrame memory frame;
        frame.admin = new ProxyAdmin();
        frame.rollup = address(
            new TransparentUpgradeableProxy(address(rollupTemplate), address(frame.admin), "")
        );

        (
            frame.delayedBridge,
            frame.sequencerInbox,
            frame.inbox,
            frame.rollupEventBridge,
            frame.outbox
        ) = bridgeCreator.createBridge(address(frame.admin), frame.rollup);

        frame.admin.transferOwnership(config.owner);
        Rollup(payable(frame.rollup)).initialize(
            config.wasmModuleRoot,
            [
                config.confirmPeriodBlocks,
                config.extraChallengeTimeBlocks,
                config.chainId,
                config.baseStake
            ],
            config.stakeToken,
            config.owner,
            [
                address(frame.delayedBridge),
                address(frame.sequencerInbox),
                address(frame.outbox),
                address(frame.rollupEventBridge),
                challengeFactory
            ],
            [rollupAdminLogic, rollupUserLogic],
            config.sequencerInboxParams
        );

        emit RollupCreated(frame.rollup, address(frame.inbox), address(frame.admin));
        return frame.rollup;
    }
}