# error

```shell
76 passing (48m)
  2 pending
  5 failing

  1) Contract: Staking app, Locking funds flows
       same origin and destiny
         when user has staked
           "before each" hook: stakes for "owner unlocks all and then unstakes":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:235:9
      at async Promise.all (index 1)
      at checkOverSlashing (test/helpers/helpers.js:229:7)
      at checkInvariants (test/helpers/helpers.js:265:5)
      at Context.<anonymous> (test/locking/funds_flows.js:62:9)

  2) Contract: Staking app, Locking funds flows
       different origin and destiny
         when user has staked
           when user hasn’t locked
             to external wallet
               transfers all:
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:193:13
      at async Promise.all (index 0)
      at checkOverUnlocking (test/helpers/helpers.js:187:5)
      at checkInvariants (test/helpers/helpers.js:263:5)
      at Context.<anonymous> (test/locking/funds_flows.js:201:13)

  3) Contract: Staking app, Locking funds flows
       "after each" hook: after test for "transfers all":
     Error: invalid argument 0: hex string without 0x prefix
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)

  4) Contract: Staking app
       fails deploying if token is not a contract:
     AssertionError: Expected error code "revert" but failed with "Error: The contract code couldn't be stored, please check your gas limit. -- Reason given: STAKING_TOKEN_NOT_CONTRACT." instead.
      at assertError (node_modules/@aragon/contract-helpers-test/assertThrow.js:5:3)
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:12:5)
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at Context.<anonymous> (test/staking.js:36:5)

  5) Contract: Staking app
       "after each" hook: after test for "fails deploying if token is not a contract":
     Error: invalid argument 0: hex string without 0x prefix
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)



error Command failed with exit code 5.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
error Command failed with exit code 5.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
Error: Test: staking exited with error code 5
    at ChildProcess.<anonymous> (/Users/mac/Data/Codes/go/ethermint/tests/solidity/test-helper.js:150:16)
    at ChildProcess.emit (node:events:390:28)
    at maybeClose (node:internal/child_process:1064:16)
    at Socket.<anonymous> (node:internal/child_process:450:11)
    at Socket.emit (node:events:390:28)
    at Pipe.<anonymous> (node:net:687:12)
error Command failed with exit code 255.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
make: *** [test-solidity] Error 255
 ▲ ~/D/C/g/ethermint  $  
```


# error 2

```shell

  75 passing (40m)
  2 pending
  4 failing

  1) Contract: Staking app, Locking funds flows
       same origin and destiny
         when user has staked
           "before each" hook: stakes for "unstakes remaining":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:214:9
      at async Promise.all (index 1)
      at checkOverTransferring (test/helpers/helpers.js:206:5)
      at checkInvariants (test/helpers/helpers.js:264:5)
      at Context.<anonymous> (test/locking/funds_flows.js:62:9)

  2) Contract: Staking app, Locking funds flows
       different origin and destiny
         when user has staked
           "before each" hook: stakes for "transfers remaining":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:230:9
      at async Promise.all (index 2)
      at checkOverSlashing (test/helpers/helpers.js:229:7)
      at checkInvariants (test/helpers/helpers.js:265:5)
      at Context.<anonymous> (test/locking/funds_flows.js:177:9)

  3) Contract: Staking app
       fails deploying if token is not a contract:
     AssertionError: Expected error code "revert" but failed with "Error: The contract code couldn't be stored, please check your gas limit. -- Reason given: STAKING_TOKEN_NOT_CONTRACT." instead.
      at assertError (node_modules/@aragon/contract-helpers-test/assertThrow.js:5:3)
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:12:5)
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at Context.<anonymous> (test/staking.js:36:5)

  4) Contract: Staking app
       "after each" hook: after test for "fails deploying if token is not a contract":
     Error: invalid argument 0: hex string without 0x prefix
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)



error Command failed with exit code 4.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
error Command failed with exit code 4.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
Error: Test: staking exited with error code 4
    at ChildProcess.<anonymous> (/Users/mac/Data/Codes/go/ethermint/tests/solidity/test-helper.js:150:16)
    at ChildProcess.emit (node:events:390:28)
    at maybeClose (node:internal/child_process:1064:16)
    at Socket.<anonymous> (node:internal/child_process:450:11)
    at Socket.emit (node:events:390:28)
    at Pipe.<anonymous> (node:net:687:12)
error Command failed with exit code 255.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
make: *** [test-solidity] Error 255
```

# error 3

```shell
5 failing

  1) Contract: Staking app, Locking funds flows
       same origin and destiny
         when user has staked
           when user has locked
             when lock manager is contract
               when lock manager allows to unlock
                 "before each" hook: stakes and locks for "unstakes remaining":
     AssertionError: Expected error code "revert" but failed with "ProviderError: CONNECTION ERROR: Couldn't connect to node http://127.0.0.1:8545." instead.
      at assertError (node_modules/@aragon/contract-helpers-test/assertThrow.js:5:3)
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:12:5)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at checkOverSlashing (test/helpers/helpers.js:249:7)
      at checkInvariants (test/helpers/helpers.js:265:5)
      at Context.<anonymous> (test/locking/funds_flows.js:100:13)

  2) Contract: Staking app, Locking funds flows
       same origin and destiny
         when user has staked
           when user has locked
             when lock manager is contract
               when lock manager doesn’t allow to unlock
                 "before each" hook: stakes and locks for "unstakes remaining":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:193:13
      at async Promise.all (index 1)
      at checkOverUnlocking (test/helpers/helpers.js:187:5)
      at checkInvariants (test/helpers/helpers.js:263:5)
      at Context.<anonymous> (test/locking/funds_flows.js:100:13)

  3) Contract: Staking app, Locking funds flows
       different origin and destiny
         when user has staked
           "before each" hook: stakes for "transfers half":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:214:9
      at async Promise.all (index 1)
      at checkOverTransferring (test/helpers/helpers.js:206:5)
      at checkInvariants (test/helpers/helpers.js:264:5)
      at Context.<anonymous> (test/locking/funds_flows.js:177:9)

  4) Contract: Staking app
       fails deploying if token is not a contract:
     AssertionError: Expected error code "revert" but failed with "Error: The contract code couldn't be stored, please check your gas limit. -- Reason given: STAKING_TOKEN_NOT_CONTRACT." instead.
      at assertError (node_modules/@aragon/contract-helpers-test/assertThrow.js:5:3)
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:12:5)
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at Context.<anonymous> (test/staking.js:36:5)

  5) Contract: Staking app
       "after each" hook: after test for "fails deploying if token is not a contract":
     Error: invalid argument 0: hex string without 0x prefix
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
```

# error 4

```shell
1) Contract: Staking app, Locking funds flows
       same origin and destiny
         when user has staked
           "before each" hook: stakes for "unstakes remaining":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/suites/staking/test/helpers/helpers.js:235:9
      at async Promise.all (index 2)
      at checkOverSlashing (test/helpers/helpers.js:229:7)
      at checkInvariants (test/helpers/helpers.js:265:5)
      at Context.<anonymous> (test/locking/funds_flows.js:64:9)

  2) Contract: Staking app, Locking funds flows
       different origin and destiny
         when user has staked
           "before each" hook: stakes for "transfers remaining":
     AssertionError: Expected "revert" but it did not fail
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:16:3)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at checkOverSlashing (test/helpers/helpers.js:244:7)
      at checkInvariants (test/helpers/helpers.js:265:5)
      at Context.<anonymous> (test/locking/funds_flows.js:177:9)

  3) Contract: Staking app, Time locking
       "before each" hook for "locks using seconds":
     TypeError: Cannot read properties of undefined (reading 'args')
      at getEvent (test/helpers/deploy.js:11:100)
      at deployStaking (test/helpers/deploy.js:26:28)
      at runMicrotasks (<anonymous>)
      at runNextTicks (node:internal/process/task_queues:61:5)
      at listOnTimeout (node:internal/timers:526:9)
      at processTimers (node:internal/timers:500:7)
      at deploy (test/helpers/deploy.js:16:21)
      at Context.<anonymous> (test/locking/locking_time.js:29:24)

  4) Contract: Staking app
       fails deploying if token is not a contract:
     AssertionError: Expected error code "revert" but failed with "Error: The contract code couldn't be stored, please check your gas limit. -- Reason given: STAKING_TOKEN_NOT_CONTRACT." instead.
      at assertError (node_modules/@aragon/contract-helpers-test/assertThrow.js:5:3)
      at assertThrows (node_modules/@aragon/contract-helpers-test/assertThrow.js:12:5)
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
      at assertRevert (node_modules/@aragon/contract-helpers-test/assertThrow.js:33:19)
      at Context.<anonymous> (test/staking.js:36:5)

  5) Contract: Staking app
       "after each" hook: after test for "fails deploying if token is not a contract":
     Error: invalid argument 0: hex string without 0x prefix
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:25:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/packages/provider/wrapper.js:166:1
      at /Users/mac/Data/Codes/go/ethermint/tests/solidity/node_modules/truffle/build/webpack:/node_modules/web3-providers-http/lib/index.js:127:1
      at runMicrotasks (<anonymous>)
      at processTicksAndRejections (node:internal/process/task_queues:96:5)
```