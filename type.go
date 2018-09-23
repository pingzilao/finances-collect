package main

type EnumRiskLevel int

const(
	EnumRiskLevel_None EnumRiskLevel = 0
	EnumRiskLevel_BaoBenBaoXi EnumRiskLevel = 1
	EnumRiskLevel_BaoBen EnumRiskLevel = 2
	EnumRiskLevel_R1 EnumRiskLevel = 3
	EnumRiskLevel_R2 EnumRiskLevel = 4
	EnumRiskLevel_R3 EnumRiskLevel = 5
	EnumRiskLevel_R4 EnumRiskLevel = 6
	EnumRiskLevel_R5 EnumRiskLevel = 7
)




type Rong360LiCaiProductInfo struct{
	Name string `json"name"`
	Bank string `json"bank"`
	Coin string `json"coin"`
	StartTime uint64
	EndTime uint64
	Term int
	MinAmount uint64
	EarnByYear uint
	RiskLevel  uint
}
