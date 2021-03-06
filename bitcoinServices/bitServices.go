//此处写RPC通信方法

package bitcoinServices

import (
	"bitcoin-Core/models/blockchain"
	"bitcoin-Core/utils"
	"bitcoin-Core/utils/Rpc"
	"strings"
)

type btcSer struct {
	BlockChahin
	Control
	Generating
	Mining
	Network
	Util
	Wallet
	Zmq
}

func GetBC() btcSer {
	return btcSer{}
}

//获取最高区块Hash,成功返回最高区块Hash,否则返回空字符串
func (bc btcSer) GetBestBlockHahs() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBESTBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//根据区块Hash获取区块信息
func (bc btcSer) GetBlockInfoByHash(blockHash string) blockchain.BlockInfo {
	paramsSlice := []interface{}{blockHash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCK, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockInfo := blockchain.BlockInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockInfo.Time = resBytes["time"].(float64)
		blockInfo.Bits = resBytes["bits"].(string)
		blockInfo.NTx = resBytes["nTx"].(float64)
		blockInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockInfo.Strippedsize = resBytes["strippedsize"].(float64)
		blockInfo.Size = resBytes["size"].(float64)
		blockInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockInfo.Weight = resBytes["weight"].(float64)
		blockInfo.Version = resBytes["version"].(float64)
		blockInfo.VersionHex = resBytes["versionHex"].(string)
		blockInfo.Nonce = resBytes["nonce"].(float64)
		blockInfo.Difficulty = resBytes["difficulty"].(float64)
		blockInfo.Hash = resBytes["hash"].(string)
		blockInfo.Height = resBytes["height"].(float64)
		blockInfo.Mediantime = resBytes["mediantime"].(float64)
		blockInfo.Chainwork = resBytes["chainwork"].(string)
		blockInfo.Confirmations = resBytes["confirmations"].(float64)
		blockInfo.Tx = resBytes["tx"].([]interface{})
	}

	return blockInfo
}

//根据区块hash获取区块Info
func (bc btcSer) GetBlockByHeight(height int) blockchain.BlockInfo {
	blockInfo := blockchain.BlockInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockInfo
	}

	paramsSlice := []interface{}{height}

	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockInfoByHash(hash)
	}

	return blockInfo
}

//获取区块链信息
func (bc btcSer) GetBlockChainInfo() blockchain.BlockChainInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKCHAININFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockChainInfo := blockchain.BlockChainInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockChainInfo.Size_on_disk = resBytes["size_on_disk"].(float64)
		blockChainInfo.Blocks = resBytes["blocks"].(float64)
		blockChainInfo.Mediantime = resBytes["mediantime"].(float64)
		blockChainInfo.Chainwork = resBytes["chainwork"].(string)
		blockChainInfo.Automatic_pruning = resBytes["automatic_pruning"].(bool)
		blockChainInfo.Pruneheight = resBytes["pruneheight"].(float64)
		blockChainInfo.Bestblockhash = resBytes["bestblockhash"].(string)
		blockChainInfo.Difficulty = resBytes["difficulty"].(float64)
		blockChainInfo.Initialblockdownload = resBytes["initialblockdownload"].(bool)
		blockChainInfo.Pruned = resBytes["pruned"].(bool)
		blockChainInfo.Prune_target_size = resBytes["prune_target_size"].(float64)
		blockChainInfo.Chain = resBytes["chain"].(string)
		blockChainInfo.Warnings = resBytes["warnings"].(string)
		blockChainInfo.Headers = resBytes["headers"].(float64)

		blockChainInfo.Softforks_, ok = resBytes["softforks"].(map[string]interface{})
		if ok {
			blockChainInfo.Softforks.Bip34_, ok = blockChainInfo.Softforks_["bip34"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip34.Height = blockChainInfo.Softforks.Bip34_["height"].(float64)
				blockChainInfo.Softforks.Bip34.Type = blockChainInfo.Softforks.Bip34_["type"].(string)
				blockChainInfo.Softforks.Bip34.Active = blockChainInfo.Softforks.Bip34_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip65_, ok = blockChainInfo.Softforks_["bip65"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip65.Height = blockChainInfo.Softforks.Bip65_["height"].(float64)
				blockChainInfo.Softforks.Bip65.Type = blockChainInfo.Softforks.Bip65_["type"].(string)
				blockChainInfo.Softforks.Bip65.Active = blockChainInfo.Softforks.Bip65_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip66_, ok = blockChainInfo.Softforks_["Bip66"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip66.Height = blockChainInfo.Softforks.Bip66_["height"].(float64)
				blockChainInfo.Softforks.Bip66.Type = blockChainInfo.Softforks.Bip66_["type"].(string)
				blockChainInfo.Softforks.Bip66.Active = blockChainInfo.Softforks.Bip66_["active"].(bool)
			}

			blockChainInfo.Softforks.Segwit_, ok = blockChainInfo.Softforks_["Segwit"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Segwit.Height = blockChainInfo.Softforks.Segwit_["height"].(float64)
				blockChainInfo.Softforks.Segwit.Type = blockChainInfo.Softforks.Segwit_["type"].(string)
				blockChainInfo.Softforks.Segwit.Active = blockChainInfo.Softforks.Segwit_["active"].(bool)
			}

			blockChainInfo.Softforks.Csv_, ok = blockChainInfo.Softforks_["Csv"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Csv.Height = blockChainInfo.Softforks.Csv_["height"].(float64)
				blockChainInfo.Softforks.Csv.Type = blockChainInfo.Softforks.Csv_["type"].(string)
				blockChainInfo.Softforks.Csv.Active = blockChainInfo.Softforks.Csv_["active"].(bool)
			}

		}
	}

	return blockChainInfo
}

//获取网络信息
func (bc btcSer) GetNetWorkInfo() blockchain.NetWorkInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETNETWORKINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	netWorkInfo := blockchain.NetWorkInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		netWorkInfo.Relayfee = resBytes["relayfee"].(float64)
		netWorkInfo.Warnings = resBytes["warnings"].(string)
		netWorkInfo.Localrelay = resBytes["localrelay"].(bool)
		netWorkInfo.Networks_, ok = resBytes["networks"].([]interface{})
		if ok {
			for i := 0; i < len(netWorkInfo.Networks_); i++ {
				mapValue, ok := netWorkInfo.Networks_[i].(map[string]interface{})
				if ok {
					var network blockchain.NetWork
					network.Name = mapValue["name"].(string)
					network.Limited = mapValue["limited"].(bool)
					network.Reachable = mapValue["reachable"].(bool)
					network.Proxy = mapValue["proxy"].(string)
					network.Proxy_randomize_credentials = mapValue["proxy_randomize_credentials"].(bool)
					netWorkInfo.Networks = append(netWorkInfo.Networks, network)
				}
			}
		}

		netWorkInfo.Version = resBytes["version"].(float64)
		netWorkInfo.Subversion = resBytes["subversion"].(string)
		netWorkInfo.Protocolversion = resBytes["protocolversion"].(float64)
		netWorkInfo.Timeoffset = resBytes["timeoffset"].(float64)
		netWorkInfo.Networkactive = resBytes["networkactive"].(bool)
		netWorkInfo.Localaddresses = resBytes["localaddresses"].([]interface{})
		netWorkInfo.Localservices = resBytes["localservices"].(string)
		netWorkInfo.Localservicesnames = resBytes["localservicesnames"].([]interface{})
		netWorkInfo.Connections = resBytes["connections"].(float64)
		netWorkInfo.Incrementalfee = resBytes["incrementalfee"].(float64)
	}

	return netWorkInfo
}

//获取区块总数
func (bc btcSer) GetBlockCount() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKCOUNT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}

//根据区块高度获取区块的hash
func (bc btcSer) GetBlockHash(height int) string {
	if float64(height) > bc.GetBlockCount() {
		return ""
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return hash
	}

	return ""
}

//根据区块Hash获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHash(hash string) blockchain.BlockHeaderInfo {
	paramsSlice := []interface{}{hash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHEADER, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列华操作
	blockHeadInfo := blockchain.BlockHeaderInfo{}

	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockHeadInfo.Height = resBytes["height"].(float64)
		blockHeadInfo.Mediantime = resBytes["mediantime"].(float64)
		blockHeadInfo.Chainwork = resBytes["chainwork"].(string)
		blockHeadInfo.NTx = resBytes["nTx"].(float64)
		blockHeadInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockHeadInfo.Hash = resBytes["hash"].(string)
		blockHeadInfo.Version = resBytes["version"].(float64)
		blockHeadInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockHeadInfo.Time = resBytes["time"].(float64)
		blockHeadInfo.Difficulty = resBytes["difficulty"].(float64)
		blockHeadInfo.Confirmations = resBytes["confirmations"].(float64)
		blockHeadInfo.VersionHex = resBytes["versionHex"].(string)
		blockHeadInfo.Nonce = resBytes["nonce"].(float64)
		blockHeadInfo.Bits = resBytes["bits"].(string)
	}

	return blockHeadInfo
}

//根据区块Height获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHeight(height float64) blockchain.BlockHeaderInfo {
	blockHeaderInfo := blockchain.BlockHeaderInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockHeaderInfo
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockHeaderInfoByHash(hash)
	}

	return blockHeaderInfo
}

//根据区块高度获取区块状态
func (bc btcSer) GetBlockStatsInfoByHeight(height float64) blockchain.BlockStats {
	blockStats := blockchain.BlockStats{}
	if height > bc.GetBlockCount() {
		return blockStats
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON("getblockstats", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockStats.Subsidy = res["subsidy"].(float64)
		blockStats.Total_weight = res["total_weight"].(float64)
		blockStats.Height = res["height"].(float64)
		blockStats.Medianfee = res["medianfee"].(float64)
		blockStats.Mediantxsize = res["mediantxsize"].(float64)
		blockStats.Totalfee = res["totalfee"].(float64)
		blockStats.Avgfee = res["avgfee"].(float64)
		blockStats.Avgfeerate = res["avgfeerate"].(float64)
		blockStats.Minfee = res["minfee"].(float64)
		blockStats.Swtxs = res["swtxs"].(float64)
		blockStats.Time = res["time"].(float64)
		blockStats.Total_size = res["total_size"].(float64)
		blockStats.Blockhash = res["blockhash"].(string)
		blockStats.Feerate_percentiles_, ok = res["feerate_percentiles"].([]interface{})
		if ok {
			for i := 0; i < len(blockStats.Feerate_percentiles_); i++ {
				value, ok := blockStats.Feerate_percentiles_[i].(float64)
				if ok {
					blockStats.Feerate_percentiles = append(blockStats.Feerate_percentiles, value)
				}
			}
		}

		blockStats.Ins = res["ins"].(float64)
		blockStats.Minfeerate = res["minfeerate"].(float64)
		blockStats.Avgtxsize = res["avgtxsize"].(float64)
		blockStats.Maxfeerate = res["maxfeerate"].(float64)
		blockStats.Mintxsize = res["mintxsize"].(float64)
		blockStats.Outs = res["outs"].(float64)
		blockStats.Utxo_increase = res["utxo_increase"].(float64)
		blockStats.Mediantime = res["mediantime"].(float64)
		blockStats.Swtotal_weight = res["swtotal_weight"].(float64)
		blockStats.Txs = res["txs"].(float64)
		blockStats.Utxo_size_inc = res["utxo_size_inc"].(float64)
		blockStats.Maxfee = res["maxfee"].(float64)
		blockStats.Swtotal_size = res["swtotal_size"].(float64)
		blockStats.Total_out = res["total_out"].(float64)
	}

	return blockStats
}

//根据区块Hash获取区块状态
func (bc btcSer) GetBlockStatsInfoByHash(hash string) blockchain.BlockStats {
	blockStats := blockchain.BlockStats{}
	if len(hash) != 64 || hash[0] != 48{
		return blockStats
	}

	var height float64 = -1
	if hash == bc.GetBlockHash(0) {
		return bc.GetBlockStatsInfoByHeight(0)
	}

	blockInfo := bc.GetBlockInfoByHash(hash)
	height = blockInfo.Height

	if height > 0 {
		return bc.GetBlockStatsInfoByHeight(height)
	}

	return blockStats
}

//获取区块链Tip信息
func (bc btcSer) GetChainTips() blockchain.ChainTips {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETCHAINTIPS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTips := blockchain.ChainTips{}

	res, ok := rpcResult.Data.Result.([]interface{})
	if ok {
		for i := 0; i < len(res); i++ {
			var tip blockchain.Tip
			m, ok := res[i].(map[string]interface{})
			if ok {
				tip.Height = m["height"].(float64)
				tip.Hash = m["hash"].(string)
				tip.Branchlen = m["branchlen"].(float64)
				tip.Status = m["status"].(string)
			}
			chainTips.Tips = append(chainTips.Tips, tip)
		}
	}

	return chainTips
}

//获取区块链的交易状态
func (bc btcSer) GetChainTxStats() blockchain.ChainTxStats {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETCHAINTXSTATS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTxStats := blockchain.ChainTxStats{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		chainTxStats.Window_final_block_hash = res["window_final_block_hash"].(string)
		chainTxStats.Window_final_block_height = res["window_final_block_height"].(float64)
		chainTxStats.Window_block_count = res["window_block_count"].(float64)
		chainTxStats.Window_tx_count = res["window_tx_count"].(float64)
		chainTxStats.Window_interval = res["window_interval"].(float64)
		chainTxStats.Txrate = res["txrate"].(float64)
		chainTxStats.Time = res["time"].(float64)
		chainTxStats.Txcount = res["txcount"].(float64)
	}

	return chainTxStats
}

//获取当前挖矿难度,成功返回挖矿难度,否则返回-1
func (bc btcSer) GetDifficulty() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETDIFFICULTY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -0
}

//获取回收内存信息
func (bc btcSer) GetMempoolInfo() blockchain.MempoolInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON("getmempoolinfo", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	mempoolInfo := blockchain.MempoolInfo{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		mempoolInfo.Loaded = res["loaded"].(bool)
		mempoolInfo.Size = res["size"].(float64)
		mempoolInfo.Bytes = res["bytes"].(float64)
		mempoolInfo.Usage = res["usage"].(float64)
		mempoolInfo.Maxmempool = res["maxmempool"].(float64)
		mempoolInfo.Mempoolminfee = res["mempoolminfee"].(float64)
		mempoolInfo.Minrelaytxfee = res["minrelaytxfee"].(float64)
	}

	return mempoolInfo
}