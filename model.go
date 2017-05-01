package snowball

type CodeList struct {
    Count   map[string]int
    Success bool
    Stocks  [](map[string]string)
}

var detailStringFields = []string{
    "symbol", "name", "code", "exchange",
}

var detailFloatFields = []string{
    // 当前价格
    "current",
    // 涨跌幅度
    "percentage",
    // 涨跌额
    "change",
    // 开、收、最高、最低
    "open", "close", "high", "low",
    // 52周 最高 最低
    "high52week", "low52week",
    // 成交量  平均成交量
    "volume", "volumeAverage",
    // 市值
    "marketCapital",
    // 美股收益
    "eps",
    /*
 市盈率 TTM
 TTM = Trailing Twelve Months，字面翻译为连续12个月内。 一般说的市盈率PE(TTM)指在一个考察期（通常为12个月的时间）内，股票的价格和每股收益的比例。计算方法为：市盈率=普通股每股市场价格÷普通股每年每股盈利。每股盈利的计算方法，是该企业在过去12个月的净收入除以总发行已售出股数。
 */
    "pe_ttm",
    /*
 市盈率 LYR
 LYR=Last Year Ratio，按照去年（最新年报）年度指标进行计算，比如市盈率PE(LTR)则表示用去年年度的每股收益来计算市盈率。
 */
    "pe_lyr",
    // beta值   风险指标？
    "beta",
    // 总股本
    "totalShares",

    // 盘后数据
    "afterHours", "afterHoursPct", "afterHoursChg",

    // 股息/红利  收益
    "dividend", "yield",

    // 换手率
    "turnover_rate",
    // 机构持股
    "instOwn",

    "rise_stop", "fall_stop", "amount",

    //// 美股净资产
    //"net_assets",


    // 市净率MRQ
    "pb",
};

type StockDetail struct {
    Symbol        string
    Exchange      string
    Code          string
    Name          string
    Current       float64
    Percentage    float64
    Change        float64
    Open          float64
    High          float64
    Low           float64
    Close         float64
    LastClose     float64
    High52week    float64
    Low52week     float64
    Volume        float64
    volumeAverage float64
    MarketCapital float64
    Eps           float64
    PeTtm         float64
    PeLyr         float64
    Beta          float64
    TotalShares   float64
    AfterHours    float64
    AfterHoursPct float64
    AfterHoursChg float64
    Dividend      float64
    Yield         float64
    TurnoverRate  float64
    InstOwn       float64
    RiseStop      float64
    FallStop      float64
    Amount        float64
    NetAssets     string
    Pb            float64
}

