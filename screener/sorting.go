package screener

type Sorting string

func (s Sorting) Desc() Sorting {
	if s != "" {
		return "-" + s
	}
	return ""

}

const (
	EPSgrowththisyear         Sorting = "epsyoy"
	EPSgrowthnextyear         Sorting = "epsyoy1"
	EPSgrowthpast5years       Sorting = "eps5years"
	EPSgrowthnext5years       Sorting = "estltgrowth"
	Salesgrowthpast5years     Sorting = "sales5years"
	EPSgrowthqtroverqtr       Sorting = "epsqoq"
	Salesgrowthqtroverqtr     Sorting = "salesqoq"
	SharesOutstanding         Sorting = "sharesoutstanding2"
	SharesFloat               Sorting = "sharesfloat"
	InsiderOwnership          Sorting = "insiderown"
	InsiderTransactions       Sorting = "insidertrans"
	InstitutionalOwnership    Sorting = "instown"
	InstitutionalTransactions Sorting = "insttrans"
	ShortInterestShare        Sorting = "shortinterestshare"
	ShortInterestRatio        Sorting = "shortinterestratio"
	EarningsDate              Sorting = "earningsdate"
	ReturnonAssets            Sorting = "roa"
	ReturnonEquity            Sorting = "roe"
	ReturnonInvestment        Sorting = "roi"
	CurrentRatio              Sorting = "curratio"
	QuickRatio                Sorting = "quickratio"
	LTDebtEquity              Sorting = "ltdebteq"
	TotalDebtEquity           Sorting = "debteq"
	GrossMargin               Sorting = "grossmargin"
	OperatingMargin           Sorting = "opermargin"
	NetProfitMargin           Sorting = "netmargin"
	AnalystRecommendation     Sorting = "recom"
	PerformanceWeek           Sorting = "perf1w"
	PerformanceMonth          Sorting = "perf4w"
	PerformanceQuarter        Sorting = "perf13w"
	PerformanceHalfYear       Sorting = "perf26w"
	PerformanceYear           Sorting = "perf52w"
	PerformanceYearToDate     Sorting = "perfytd"
	Beta                      Sorting = "beta"
	AverageTrueRange          Sorting = "averagetruerange"
	VolatilityWeek            Sorting = "volatility1w"
	VolatilityMonth           Sorting = "volatility4w"
	SMARelative20Day          Sorting = "sma20"
	SMARelative50Day          Sorting = "sma50"
	SMARelative200Day         Sorting = "sma200"
	HighRelative50Day         Sorting = "high50d"
	LowRelative50Day          Sorting = "low50d"
	HighRelative52Week        Sorting = "high52w"
	LowRelative52Week         Sorting = "low52w"
	RelativeStrengthIndex14   Sorting = "rsi"
	AverageVolume3Month       Sorting = "averagevolume"
	RelativeVolume            Sorting = "relativevolume"
	Change                    Sorting = "change"
	ChangefromOpen            Sorting = "changeopen"
	Gap                       Sorting = "gap"
	Volume                    Sorting = "volume"
	Price                     Sorting = "price"
	TargetPrice               Sorting = "targetprice"
	IPODate                   Sorting = "ipodate"
)
