import zhLocale from 'element-ui/lib/locale/lang/zh-CN'
const messagesCn = {
    menus:{
        home: '首页',
        blockchain: '区块链',
        blocks: '区块',
        transactions:'交易',
        transfers:'转账',
        accounts: '账户',
        statistics: '统计',
        contracts: '合约',
        nodes: '服务节点',
        supernode:'生产节点'
    },
    footer:{
        title:'联系我们'
    },
    tips:{
        error:"数据加载错误，请检查数据接口！",
        nodata:"暂无数据！"
    },
    home:{
        banner_title:'Bottos区块浏览器',
        searchs:'搜索账户、区块高度、区块ID、交易ID',
        searchs_tips:'未找到记录！',
        stats1:'节点',
        stats2:'最新区块',
        stats3:'过去一天交易数',
        stats4:'总账户数',
        card1:'过去14天交易数',
        card1_series_name:'交易数',
        card2:'账户增长',
        card2_series_name:'总账户数',
        card3:'区块',
        card3_block:'区块',
        card3_transactions:'交易',
        card3_delegate:'出块者',
        card3_reward:'区块奖励',
        card4:'转账',
        card4_transactions:'交易',
        card_btn:'查看全部'
    },
    blocks:{
        title:"区块",
        table1:"高度",
        table2:"时间",
        table3:"交易数",
        table4:"出块者",
        table5:"大小",
        table6:"奖励"
    },
    blocksDetail:{
        title:"区块详情",
        state:"状态",
        state_yes:"已确认",
        state_no:"未确认",
        BlockID:"区块ID",
        height:"高度",
        time:"时间",
        TransactionNumber:"交易数",
        ThePreviousBlockID:"上一区块ID",
        Chunk:"出块者",
        Size:"大小",
        tabs1:"交易",
        tabs2:"转账"
    },
    transactions:{
        title:"交易",
        table1:"交易ID",
        table2:"区块",
        table3:"时间",
        table4:"合约名称",
        method_name:"方法名称",
        computing_resource:"计算资源",
        network_resource:"网络资源",
    },
    transactionsDetail:{
        title:"交易详情",
        state:"状态",
        state_yes:"已确认",
        state_no:"未确认",
        transactionsID:"交易ID",
        block:"区块",
        time:"时间",
        computing_resource:"计算资源",
        network_resource:"网络资源",
        title2:"合约",
        name:"合约名称",
        method_name:"方法名称",
        input:"参数"
    },
    transfers:{
        title:"转账",
        table1:"交易ID",
        table2:"区块",
        table3:"时间",
        table4:"发送人->接收人",
        table5:"转账金额"
    },
    transfersDetail:{
        title:"转账详情",
        state:"状态",
        state_yes:"已确认",
        state_no:"未确认",
        transactionsID:"交易ID",
        block:"区块",
        time:"时间",
        computing_resource:"计算资源",
        network_resource:"网络资源",
        title2:"合约 - 转账",
        name:"合约名称",
        method_name:"方法名称",
        input:"参数"
    },
    accounts:{
        title:"账户",
        title_statistics:"实时总账户数",
        table1:"账户",
        table2:"余额",
        table3:"持有量"
    },
    accountsDetail:{
        title:"账户详情",
        accounts:"账户",
        balance:"余额",
        TransactionNumber:"交易数",
        Transfer:"转账",
        tabs1:"交易",
        tabs2:"转账"
    },
    statistics:{
        title:"统计",
        card1:"24小时交易数",
        card2:"账户增长"
    },
    contracts:{
        title:"合约",
        searchs:"搜索合约名称",
        searchsbtn:"搜索",
        table1:"合约名称",
        table2:"合约地址",
        table3:"交易数",
        table4:"调用账户数"
    },
    contractsDetail:{
        title:"合约详情",
        name:"合约名称",
        address:"合约地址",
        TransactionNumber:"交易数",
        AccountSize:"调用账户数",
        Record:"调用记录",
        code:"代码",
        table1:"交易ID",
        table2:"区块",
        table3:"调用账户",
        table4:"调用时间",
        table5:"详情",
        contractsAPI:"合约API",
        contractsByte:"字节码"
    },
    nodes:{
        stats1:"生产节点",
        stats2:"服务节点",
        title:'节点',
        searchs:"节点名称",
        searchsbtn:"搜索",
        table1:'节点',
        table2:'服务类型',
        table3:'地区',
        table4:'当前连接数',
        table5:'剩余连接数',
        table6:'服务API'
    },
    supernode:{
        stats1:"生产节点",
        stats2:"总已投票数",
        stats3:"总已投人数",
        title:"生产节点",
        searchs:"搜索超级节点名称",
        lists_block:"正在出块",
        lists_ticket:"票",
        lists_throw:"人已投",
        lists_more:"更多",
        lists_nodata:"未搜索到相关节点"
    },
    supernodeDetail:{
        ranking:"排名",
        got_vote:"得票",
        ticket:"票",
        cast_vote:"投票",
        people:"人",
        expand:"展开",
        account:"账户",
        quality_deposit:"质押金额",
        number_of_blocks:"出块数",
        block_success_rate:"出块成功率",
        serviceApi:"服务API",
        tabs1:"出块详情",
        tabs2:"投票人",
        table1:"高度",
        table2:"时间",
        table3:"交易数",
        table4:"大小",
        table5:"投票人",
        table6:"投票数",
        table7:"本节点投票占比",
        table8:"整网投票占比"
    },
    ...zhLocale
}

export default messagesCn