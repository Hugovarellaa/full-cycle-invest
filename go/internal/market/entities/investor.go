package entities

type InvestorAssetPossition struct {
	AssetID string
	Shares  int
}

type Investor struct {
	ID            string
	name          string
	assetPosition []*InvestorAssetPossition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		assetPosition: []*InvestorAssetPossition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPossition) {
	i.assetPosition = append(i.assetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.assetPosition = append(i.assetPosition, newInvestorAssetPosition(assetID, qtShares))
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPossition {
	for _, assetPosition := range i.assetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}

	return nil
}

func newInvestorAssetPosition(assetID string, shares int) *InvestorAssetPossition {
	return &InvestorAssetPossition{
		AssetID: assetID,
		Shares:  shares,
	}
}
