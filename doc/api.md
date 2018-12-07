#####`/explorer/burn/summary`

Return the count of burn BCH. Total and Avail of the burn BCH will be returned.

| Protocol | Method | API                    |
| -------- | ------ | ---------------------- |
| HTTP     | GET    | /explorer/burn/summary |

**Argument**

None

##### Result

`Total` Total amount of burning BCH

`Avail` Available number of BCH

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/burn/summary

//Result
{"result":{"Total":"47063.45013","Avail":"47063.45013"}}
```

***

##### `/explorer/burn/list`

Return the list of transactions burning BCH.

| Protocol | Method | API                 |
| -------- | ------ | ------------------- |
| HTTP     | GET    | /explorer/burn/list |

**Argument**

- pageNo `string` required:the number of page(can be null)
- pageSize `string` required:the size of page(can be null)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list`  Burning BCH transaction list

- `TxHash` transaction hash
- `BlockTime` block time
- `Address` Account address
- `Whc` Number of whc
- `Bch` Number of bch
- `TxBlockHeight` Block height of the transaction
- `Process` Progress of the transaction
- `MatureTime` Mature Time

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/burn/list

//Result
{"result":{"pageNo":1,"pageSize":4,"total":357,"list":[{"TxHash":"cf340df7ac330ee731f2293dfeee8c3f9931bd055b5ba0d233ee1f720edb7c45","BlockTime":1543635258,"Address":"bchtest:qrz6dw5zau3kpahk97kmzcu8smg58xc4ny6l7z94s9","Whc":"100","Bch":"1","TxBlockHeight":1271064,"Process":"3/3","MatureTime":1543637543},{"TxHash":"5c2c8d7e2b5240c8da9070da05a3625bfde274628dce211fbd2047ed3d2d3fad","BlockTime":1543635258,"Address":"bchtest:qpwy20dcqxlr48d5tlvtdd5wfdltdlz8w5evuce5ee","Whc":"200","Bch":"2","TxBlockHeight":1271064,"Process":"3/3","MatureTime":1543637543},{"TxHash":"1feb82be24c8cf4889da26314e10fe8b69633aae4a5799993852f2e9f8dc6c53","BlockTime":1543496311,"Address":"bchtest:qqxugl4zj53tftj0wnwkz4970s6dxhkvf5xjctfsae","Whc":"100","Bch":"1","TxBlockHeight":1270840,"Process":"3/3","MatureTime":1543499989},{"TxHash":"a148db68c98f144b13005f69e7ba4a319a22eb88732636a8e70395b0db1d8b66","BlockTime":1543482317,"Address":"bchtest:qq4ea22nt0tjzne8zx4psh035cemk7w25yv5zeq4vs","Whc":"100","Bch":"1","TxBlockHeight":1270830,"Process":"3/3","MatureTime":1543486447}]}}
```

------

##### `/info`

Return blockchain and wormhole information

| Protocol | Method | API   |
| -------- | ------ | ----- |
| HTTP     | GET    | /info |

**Argument**

None

##### Result

`block_height` the block height of the blockchain tip

`block_time` the block time of the blockchain tip

`current_whc_count` the wormhole transactions count in the current block

`one_day_whc_count` the wormhole transactions count in 24 hours

`total_whc_count` the wormhole transactions count util now in wormhole system

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/info

//Result
{"result":{"block_height":1271081,"block_time":1543644521,"current_whc_count":0,"one_day_whc_count":28,"total_whc_count":7464}}
```

------

##### `/properties`

Return the list of wormhole properties and page info. 

| Protocol | Method | API         |
| -------- | ------ | ----------- |
| HTTP     | GET    | /properties |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)
- propertyType `string` required:the type of property(can be null, default:-1, 50:fixed property, 51:Crowdsale, 54:managed property, -1:all)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list`  Property list

- `propertyId` Property Id
- `propertyName` Property Name
- `propertyCategory` Property class classification name
- `propertySubcategory` Property secondary classification name
- `propertyServiceUrl` Property URL
- `issuer` Issuer address
- `txType` Property type

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/properties

//Result
{"result":{"pageNo":1,"pageSize":2,"total":90,"list":[{"propertyId":7,"propertyName":"crowsale_5","propertyCategory":"test crowdsale token 5","propertySubcategory":"test","propertyServiceUrl":"www.testcrowdsaletoken.com","issuer":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","txType":"51"},{"propertyId":11,"propertyName":"3","propertyCategory":"test3","propertySubcategory":"pricision_3","propertyServiceUrl":"www.pricision.3","issuer":"bchtest:qpalmy832fp9ytdlx444sehajljnm554dulckcvjl5","txType":"51"}]}}

```

------

##### `/property/:propertyId`

Return property information with propertyId.

| Protocol | Method | API                   |
| -------- | ------ | --------------------- |
| HTTP     | GET    | /property/:propertyId |

**Argument**

None

##### Result

`active` Crowdfunding extra fields - active

`addedissuertokens` Added issuer tokens

`amountraised` Amount raised

`category ` Property class classification name

`closetx` Close property transaction hash

`creationtxid` Create property transaction hash

`data`

`deadline` Crowdfunding extra field - end time

`earlybonus` Crowdfunding Extra Fields - Early Bird Reward Percentage

`fixedissuance` Crowdfunding extra field - Fixed issuance

`freezingenabled` Freezing enabled

`issuer` Issuer address

`managedissuance` Managed issuance

`name` Property Name

`participanttransactions`  participant transactions

`precision` Number of digits after the decimal point

`propertyid` Property id

`propertyiddesired` property id desired

`purchasedtokens` Crowdfunding Extra Fields - Crowdfunded Quantity

`starttime` Start timestamp

`subcategory` Property secondary classification name

`tokensissued` Crowdfunding Extra Fields - Total Crowdfunding

`tokensperunit` Crowdfunding Extra Fields - Ratio to WHC

`totaltokens` Total of tokens

`txType` Property Type

`url` URL

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/property/27

//Result
{"active":true,"addedissuertokens":"0.0000000","amountraised":"61.77041916","category":"","closetx":null,"creationtxid":"5cba679e754f038062da9becf7f2f41d5cd093cde8bd306179611d27a02f08cd","data":"hi","deadline":1596127632,"earlybonus":10,"fixedissuance":false,"freezingenabled":null,"issuer":"bchtest:qz3fgledq0tgl0ry6pn0c5nmufspfrr8aqsyuc39yl","managedissuance":false,"name":"zc_crow2","participanttransactions":null,"precision":"7","propertyid":27,"propertyiddesired":1,"purchasedtokens":637.0081442,"starttime":1532970004,"subcategory":"","tokensissued":"100000","tokensperunit":"1","totaltokens":"100000","txType":"51","url":"www.zc"}

```

------

##### `/property/:propertyId/history`

Return history info list of with propertyId.

| Protocol | Method | API                           |
| -------- | ------ | ----------------------------- |
| HTTP     | GET    | /property/:propertyId/history |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list` Changed property transaction list

- `blockTime` Block timestamp
- `txType` Event type
- `txHash` Transaction hash

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/property/27/history

//Result
{"result":{"pageNo":1,"pageSize":50,"total":1,"list":[{"blockTime":1532970004,"txType":"51","txHash":"5cba679e754f038062da9becf7f2f41d5cd093cde8bd306179611d27a02f08cd"}]}}

```

------

##### `/property/:propertyId/txs`

Return page info and transactions info list of with propertyId. 

| Protocol | Method | API                       |
| -------- | ------ | ------------------------- |
| HTTP     | GET    | /property/:propertyId/txs |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)
- eventType `string` required: the type of transaction event(can be null, default:-1,SimpleSend（0）、SendAll（4）、
AdditionalIssuance（55）、Participate in crowdfunding（1）、airdrop（3）)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list` Property-related transaction list

- `actualinvested` actual invested
- `amount` property amount
- `block` block height
- `blockhash` block hash
- `blocktime` block timestamp
- `confirmations` Confirmed number
- `fee` The cost of this transaction
- `ismine` 
- `issuertokens` Number of issuer tokens
- `precision` Number of digits after the decimal point
- `propertyid` property id
- `purchasedpropertyid` Purchased Property ID
- `purchasedpropertyname`  Purchased Property Name
- `purchasedpropertyprecision` Purchased property precision
- `purchasedtokens` Number of tokens purchased
- `referenceaddress` Reference address
- `sendingaddress` Sending address
- `subsends`
- `txid` transaction hash
- `type` transaction type
- `type_int` Enumeration value of transaction type
- `valid` 
- `version`

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/property/27/txs

//Result
{"pageNo":1,"pageSize":3,"total":1,"list":[{"block":1262739,"blockhash":"0000000000bf2d8a357f1a7753e4b82d764dc327805c84e0b6a5ca955a892a4d","blocktime":1539608725,"confirmations":5317,"ecosystem":"main","fee":"498","fee_rate":"0.00002024","ismine":false,"referenceaddress":"bchtest:qqu9lh4jpc05p59pfhu9amyv9uvder8j3sa2up95vs","sendingaddress":"bchtest:qpz8py2yqyp7x2aaqfvrwdk4jf2c23ypzczr6weclj","subsends":[{"amount":"190.01025814","precision":"8","propertyid":1},{"amount":"10.3775469","precision":"7","propertyid":27},{"amount":"1000","precision":"0","propertyid":228},{"amount":"200.0","precision":"1","propertyid":305},{"amount":"100.00000000","precision":"8","propertyid":369},{"amount":"100.00000000","precision":"8","propertyid":378},{"amount":"1100.00000000","precision":"8","propertyid":381},{"amount":"1000.00000000","precision":"8","propertyid":382}],"txid":"4d325a1de59d5c0eb4d6ee1e77c8fa8ad0a7f0468ac240b0c6033fd5db47fd8d","type":"Send All","type_int":4,"valid":true,"version":0}]}

```

------

##### /property/:propertyId/listowners

Return the owner list of the property

| Protocol | Method | API                              |
| -------- | ------ | -------------------------------- |
| HTTP     | GET    | /property/:propertyId/listowners |

**Argument**

None

##### Result

`total` Total number of results

`list` List of owner

- `Address` Owner address
- `balance_available` Owner available balance
- `Status` Owner status

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET http://127.0.0.1:8080/explorer/property/4/listowners

//Result
{"result":{"list":[{"Address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","balance_available":"110.1","Status":0},{"Address":"bchtest:qpalmy832fp9ytdlx444sehajljnm554dulckcvjl5","balance_available":"12.12","Status":0}],"total":2}}
```

------

#####/properties/query

Return a list of property, fuzzy search according to the ID or name of the property

| Protocol | Method | API             |
| -------- | ------ | --------------- |
| HTTP     | GET    | /property/query |

**Argument**

- keyword `string` required: Property ID or fuzzy name

##### Result

[properties](/properties)  response info.

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET "https://127.0.0.1:8080/explorer/properties/query?keyword=aa"

//Result
{"result":{"pageNo":1,"pageSize":50,"total":2,"list":[{"propertyId":65,"propertyName":"aaaa test token","propertyCategory":"aaaatest","propertySubcategory":"aaaa test","propertyServiceUrl":"www.aaaatest.com","issuer":"bchtest:qpv62pqneh59wvm59c8p3sxu9u00n86psysnr9a909","txType":"54"},{"propertyId":463,"propertyName":"aacoimn","propertyCategory":"Accommodation and food","propertySubcategory":"Accommodation","propertyServiceUrl":"http://aa.com","issuer":"bchtest:qrtyx3a40esw5zvd00wd30cv64654vkgn5mf0hk8pw","txType":"51"}]}}
```

-----

##### /address/:address`

Return address details info.

| Protocol | Method | API               |
| -------- | ------ | ----------------- |
| HTTP     | GET    | /address/:address |

**Argument**

None

##### Result

`balanceAvail` Whc available balance

`balanceTotal` Whc contains the total balance of the immature part

`consumedBal` Total whc consumed

`receivedBal` Total whc received

`sendedBal` Total whc transferred out

`whcTxCount` Whc transaction number

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/address/bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5

//Result
{"result":{"balanceAvail":"127.87841411","balanceTotal":"127.87841411","consumedBal":"0","receivedBal":"143.58530041","sendedBal":"-10.7068863","whcTxCount":36}}


```

------

##### /address/:address/txs

Return transactions information about address

| Protocol | Method | API                   |
| -------- | ------ | --------------------- |
| HTTP     | GET    | /address/:address/txs |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)
- txType `string` required: transaction type(default:-1)
- beginTime `string` required: begin time(default:-1)
- endTime `string` required: end time(default:-1)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list` List of transactions under this address

- `tx_hash`  transaction hash

- `type_int` Enumeration value of transaction type

- `type_str` The name of the transaction type

- `sending_address` sending address

- `reference_address` reference address list

- - `address` receiver address
  - `amount ` Received quantity

- `amount` The transaction amount

- `block_height` block height

- `block_time` block timestamp

- `state` transaction state

- `property_id` Property ID of the transaction

- `property_name` The name of the received property

- `sending_property_name` The name of the sent property

- `fee_rate` Fee rate

- `miner_fee` Miner fee

- `whc_fee`  Whc fee

- `confirmations` Confirmation times

- `direction` The direction of property trading


##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/address/bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5/txs

//Result
{"pageNo":1,"pageSize":2,"total":14,"list":[{"tx_hash":"cf340df7ac330ee731f2293dfeee8c3f9931bd055b5ba0d233ee1f720edb7c45","type_int":68,"type_str":"Burn BCH Get WHC","sending_address":"bchtest:qrz6dw5zau3kpahk97kmzcu8smg58xc4ny6l7z94s9","reference_address":[{"address":"bchtest:qrz6dw5zau3kpahk97kmzcu8smg58xc4ny6l7z94s9","amount":"100"}],"amount":"1","block_height":1271064,"block_time":1543635258,"state":"unmature","property_id":1,"property_name":"WHC","sending_property_name":"BCH","fee_rate":"0.00002642","miner_fee":"1036","whc_fee":"","confirmations":36,"direction":1},{"tx_hash":"ba3c8f4dc64fbbbf885b3b0d420f570372c2204c8dd8fb6d9ea92d5e9c7e675a","type_int":3,"type_str":"Send To Owners","sending_address":"bchtest:qrz6dw5zau3kpahk97kmzcu8smg58xc4ny6l7z94s9","reference_address":[{"address":"bchtest:qz3aj8ctfmeefk4epsg03vqakhy87aetqqfegm0hfl","amount":"0.00001000"}],"amount":"0.00001","block_height":1270162,"block_time":1543406252,"state":"true","property_id":532,"property_name":"HaiTaoCoin","sending_property_name":"HaiTaoCoin","fee_rate":"0.00002035","miner_fee":"460","whc_fee":"0.00000001","confirmations":938,"direction":-1}]}


```

------

##### `/address/:address/properties`

Return the properties details of address

| Protocol | Method | API                          |
| -------- | ------ | ---------------------------- |
| HTTP     | GET    | /address/:address/properties |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)
- propertyType `string` required: property type(default:-1)

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list` List of property under current address

- `propertyId` property Id
- `propertyName` property name
- `txType` property type
- `balanceAvailable` Available balance of property

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/address/bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5/properties

//Result
{"result":{"pageNo":1,"pageSize":50,"total":7,"list":[{"propertyId":3,"propertyName":"test_token1","txType":54,"balanceAvailable":80},{"propertyId":4,"propertyName":"test_token1","txType":54,"balanceAvailable":110.1},{"propertyId":5,"propertyName":"test_token1","txType":54,"balanceAvailable":1000.23456},{"propertyId":7,"propertyName":"crowsale_5","txType":51,"balanceAvailable":0},{"propertyId":13,"propertyName":"test_managed","txType":54,"balanceAvailable":0},{"propertyId":18,"propertyName":"test_bitcoin","txType":51,"balanceAvailable":6534.9},{"propertyId":115,"propertyName":"luzhiyao","txType":54,"balanceAvailable":8.477}]}}

```

------

##### `/address/:address/property/:propertyId/txs`

Returns the transactions for the propertyId under this address

| Protocol | Method | API                                        |
| -------- | ------ | ------------------------------------------ |
| HTTP     | GET    | /address/:address/property/:propertyId/txs |

**Argument**

- pageNo `string` required: the number of page(can be null)
- pageSize `string` required: the size of page(can be null)

##### Result

[transactions](#/address/:address/txs)  response info.

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/address/bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5/property/3/txs

//Result
{"result":{"pageNo":1,"pageSize":20,"total":4,"list":[{"tx_hash":"a1f2d0126a04296aad6f492a0ef8c1c1afb781efc6c5f37de105790a7debcf87","type_int":56,"type_str":"Revoke Property Tokens","sending_address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","reference_address":[{"address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5"}],"amount":"10","block_height":1249134,"block_time":1532837426,"state":"true","property_id":3,"property_name":"test_token1","sending_property_name":"test_token1","fee_rate":"0.0000434","miner_fee":"968","whc_fee":"","confirmations":21985,"direction":-1},{"tx_hash":"27a2adc609f0d9c5276be4a5b0727574a1bc7b32d17d48f7506889bcadf643de","type_int":0,"type_str":"Simple Send","sending_address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","reference_address":[{"address":"bchtest:qpalmy832fp9ytdlx444sehajljnm554dulckcvjl5"}],"amount":"10","block_height":1249134,"block_time":1532837426,"state":"true","property_id":3,"property_name":"test_token1","sending_property_name":"test_token1","fee_rate":"0.00004335","miner_fee":"1110","whc_fee":"","confirmations":21985,"direction":-1},{"tx_hash":"6afffd7d14060b6e79c504c5f17596616eb99356b71d301b5b37df1df065b9a0","type_int":55,"type_str":"Grant Property Tokens","sending_address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","reference_address":[{"address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5"}],"amount":"100","block_height":1249133,"block_time":1532837111,"state":"true","property_id":3,"property_name":"test_token1","sending_property_name":"test_token1","fee_rate":"0.00004334","miner_fee":"1114","whc_fee":"","confirmations":21986,"direction":1},{"tx_hash":"1c3f95acbd6eb38e2a7c26b12dc9138b4523c355a20944874bdc3c82f4c5e4e1","type_int":54,"type_str":"Create Property - Manual","sending_address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","reference_address":[{"address":"bchtest:qz04wg2jj75x34tge2v8w0l6r0repfcvcygv3t7sg5","amount":"0"}],"amount":"1","block_height":1249129,"block_time":1532836350,"state":"true","property_id":3,"property_name":"test_token1","sending_property_name":"WHC","fee_rate":"0.00004332","miner_fee":"1252","whc_fee":"1","confirmations":21990,"direction":0}]}}
```

------

##### `/search/:keyword`

Return information such as block hash/tx hash/ block height

| Protocol | Method | API              |
| -------- | ------ | ---------------- |
| HTTP     | GET    | /search/:keyword |

**Argument**

None

##### Result

[block](#/block/:block)  response info.

[transactions](#/address/:address/txs)  response info.

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/search/0000000000000109af442b5eb6344b919854c1498b4176ccd9ffb2b03bc76bbd

//Result
{"result":{"version":536870912,"block_height":1249126,"block_hash":"0000000000000109af442b5eb6344b919854c1498b4176ccd9ffb2b03bc76bbd","nonce":3369635936,"bits":436340075,"prev_block":"00000000000000a4a44f0c161745d3f0420e158cf44b34588d883e37f8523b19","merkleroot":"ad7702a1821e79e6e0d70bf1972003fa43d9b1ba4e0d3cf479ae9070d71b85b5","block_time":1532835195,"txcount":3,"whccount":0,"size":988,"nonce_str":"0xc8d89060","bits_str":"0x1a02056b","difficulty":8300642.84465382,"confirmations":22298,"fee_rates":null}}
```

------

##### `/tx/list`

Pagination returns transaction information for the height block

| Protocol | Method | API      |
| -------- | ------ | -------- |
| HTTP     | GET    | /tx/list |

**Argument**

- block_height `string` required: block height
- pageSize `string` required: the size of page(can be null)
- pageNo `string` required: the number of page(can be null)

##### Result

[transactions](#/address/:address/txs)  response info.

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/tx/list?block_height=1251104&txType4&pageSize=10&pageNo=1

//Result
{"result":{"pageNo":1,"pageSize":10,"total":3,"list":[{"tx_hash":"91c56172524fe11fabb7f954cf893a7c42be31e79f549031d11b89a5ea7d4581","type_int":56,"type_str":"Revoke Property Tokens","sending_address":"bchtest:qz08vwmzp6zy6h5jvgrt556d9f9e08a32y5eqaqztq","reference_address":[{"address":"bchtest:qz08vwmzp6zy6h5jvgrt556d9f9e08a32y5eqaqztq"}],"amount":"100","block_height":1251104,"block_time":1533891668,"state":"true","property_id":115,"property_name":"luzhiyao","sending_property_name":"luzhiyao","fee_rate":"0.0002","miner_fee":"4480","whc_fee":"","confirmations":20320,"direction":0},{"tx_hash":"d87ae34ed64e23087228eba458af1ebaf94f0db04912c59f6531f2b8c5c72f91","type_int":0,"type_str":"Simple Send","sending_address":"bchtest:qz08vwmzp6zy6h5jvgrt556d9f9e08a32y5eqaqztq","reference_address":[{"address":"bchtest:qpalmy832fp9ytdlx444sehajljnm554dulckcvjl5"}],"amount":"100","block_height":1251104,"block_time":1533891668,"state":"true","property_id":115,"property_name":"luzhiyao","sending_property_name":"luzhiyao","fee_rate":"0.0002","miner_fee":"5140","whc_fee":"","confirmations":20320,"direction":0},{"tx_hash":"ee9b0eacc5c8e9a564fa7ae5487bd706ee20dffe655cf6b1a91b69f232b195fc","type_int":1,"type_str":"Crowdsale Purchase","sending_address":"bchtest:qztr275nr423fd79ap4ewlewktawz6t7xux7gdlm0t","reference_address":[{"address":"bchtest:qr4g79cjapp02s3zs59gtu3dxu7sgwvp8gmnh9rw97"}],"amount":"0","block_height":1251104,"block_time":1533891668,"state":"false","property_id":0,"property_name":"","sending_property_name":"","fee_rate":"0.00002035","miner_fee":"521","whc_fee":"1","confirmations":20320,"direction":0}]}}
```

------

##### `/tx/hash/:txhash`

Return details of a wormhole transaction

| Protocol | Method | API              |
| -------- | ------ | ---------------- |
| HTTP     | GET    | /tx/hash/:txhash |

**Argument**

None

##### Result

`raw_data` the raw hexadecimal string for the transaction

[transactions](#/address/:address/txs)  response info.

##### Example

```json
//Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/tx/hash/fb9330f94e5460053c8f46088df8767ba3a82f7830e45287d647dd8ecc9596b6

//Result
{"result":{"tx_hash":"fb9330f94e5460053c8f46088df8767ba3a82f7830e45287d647dd8ecc9596b6","type_int":3,"type_str":"Send To Owners","sending_address":"bchtest:qqrj9ekzdrh5vex3w76yp33sxqk4gctxpu90g3wcue","reference_address":[{"address":"bchtest:qpncl685d5eys3tj4vykrhw4507fcuas9scm4qd70e","amount":"5002.04440452"},{"address":"bchtest:qq5uyyd0gzpv289vmkxcn86wvlc845yq3gf3a2dukm","amount":"4997.95559548"}],"amount":"10000","block_height":1250295,"block_time":1533459341,"state":"true","property_id":67,"property_name":"fixed issuance","sending_property_name":"fixed issuance","fee_rate":"0.00020088","miner_fee":"4540","whc_fee":"0.00000002","confirmations":21129,"raw_data":"02000000014b3cd0482575d0b22e2b1656049679076d7c8e0b9a41da79f6c84bd25a67d446010000006a47304402203642f98b1e934603b3f0e98d91117874fb2aded6a38889efe011efef5d3b14ef022049804dd6f28c75248d0b9ebbdb6dcbb773327f953db3328d383e381b0e76fead412103f6bcead2d91eacdbaa2e99c5a8b40fc399561cc7cde0b27a6c966f122c57041cfeffffff0200000000000000001a6a18087768630000000300000043000000e8d4a5100000000042de10f505000000001976a9140722e6c268ef4664d177b440c630302d5461660f88acf6131300","direction":0}}
```

------

##### `/tx/latest`

Return the trend of the number of wormhole transactions in the last seven days

| Protocol | Method | API        |
| -------- | ------ | ---------- |
| HTTP     | GET    | /tx/latest |

**Argument**

- timeOffset `string` required: Time offset in minutes(default: 0)

##### Result

`result` Number of seven days of wormhole trading

##### Example

```json
#Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/tx/latest?timeOffset=-480


# Result
{"result":{"11-27":15,"11-28":16,"11-29":31,"11-30":25,"12-01":44,"12-02":1,"12-03":5}}
```

------

##### `/block/:block`

Returns the details of the block, supports block height and block hash search

| Protocol | Method | API           |
| -------- | ------ | ------------- |
| HTTP     | GET    | /block/:block |

**Argument**

None

##### Result

`version` block version

`block_height` block height

`block_hash` block hash

`nonce` nonce integer

`bits` difficulty target for the current block

`prev_block` the previous block hash

`merkleroot` the summary for all transactions in this block. Please refer to the relative information.

`block_time` the confirmation time for this transaction

`txcount` bitcoin cash transactions count

`whccount` wormhole transactions count

`size` size in byte

`nonce_str` nonce string in hexadecimal format

`bits_str` bits string in hexadecimal format

`difficulty` difficulty string format

`confirmations` confirmed times for this transaction

`fee_rates` fee rates

##### Example

```json
#Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/block/block/1249142

curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/block/block/0000000094913296f661df435f2850490963fb3f638d51c32a40ef719a7ae5af

#Result
{"result":{"version":536870912,"block_height":1249142,"block_hash":"0000000094913296f661df435f2850490963fb3f638d51c32a40ef719a7ae5af","nonce":3057240586,"bits":486604799,"prev_block":"00000000005a15e4b49a073a3836a64bbdc9e0175f67228b5e744538a3352733","merkleroot":"450a32bb70adc79863cd711bfff5ea109b7ca3aac70c8c5cdd521dd59ab7b7cb","block_time":1532842563,"txcount":27,"whccount":3,"size":46031,"nonce_str":"0xb639ca0a","bits_str":"0x1d00ffff","difficulty":1,"confirmations":22283,"fee_rates":{"min_fee_rate":"0.00065758","avg_fee_rate":"0.0007028","max_fee_rate":"0.00073197"}}}
```

------

##### `/block/list`

Returns the list of blocks in a period of time, from, to set the time period, pageNo, pageSize to set the page number and page size

| Protocol | Method | API         |
| -------- | ------ | ----------- |
| HTTP     | GET    | /block/list |

**Argument**

- from `string` required: start time
- to `string` required: end time
- pageSize `string` required: the size of page
- pageNo `string` required: the number of page

##### Result

`pageNo` Current result page number

`pageSize` Page size

`total` Total number of eligible results

`list` List of block

[block](#/block/:block)  response info.

##### Example

```json
#Request
curl -i -H 'Content-Type: application/json' -X GET https://127.0.0.1:8080/explorer/block/list?from=1540278395&to=1540279295&pageSize=10&pageNo=1

#Result
{"result":{"pageNo":1,"pageSize":10,"total":2,"list":[{"version":536870912,"block_height":1264052,"block_hash":"0000000000000253f43eef2c668930febcdbc5d73fd50879bcd15e169314e0c2","nonce":3818756700,"bits":436483450,"prev_block":"00000000000001306b93070672ca9a31e6401f808594e9f77515a08cc599b4e8","merkleroot":"8c5c9adcafd530479d92dd5cddf0ff0ea92494f2571be2770faf476585da6946","block_time":1540279295,"txcount":6,"whccount":0,"size":1646,"nonce_str":"0xe39d9a5c","bits_str":"0x1a04357a","difficulty":3986074.41635186,"confirmations":7374,"fee_rates":null},{"version":536870912,"block_height":1264051,"block_hash":"00000000000001306b93070672ca9a31e6401f808594e9f77515a08cc599b4e8","nonce":3074014702,"bits":436488044,"prev_block":"00000000000003e826b395acbb495a1941ae7e2eae99fd2651fcf4d27c912448","merkleroot":"6903642a4c8a86338c6772382c1c369c65ee842b6ff881051ca888e5a5d4ac48","block_time":1540279230,"txcount":72,"whccount":0,"size":43137,"nonce_str":"0xb739bdee","bits_str":"0x1a04476c","difficulty":3920774.14010013,"confirmations":7375,"fee_rates":null}]}}

```





