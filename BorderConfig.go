package texttable

var (
	NoBorders    *BorderConfig = &BorderConfig{}
	OuterBorders *BorderConfig = &BorderConfig{topBorder: true, bottomBorder: true, leftBorder: true, rightBorder: true}
	InnerBorders *BorderConfig = &BorderConfig{columnSeparator: true, rowSeparator: true}
	AllBorders   *BorderConfig = &BorderConfig{topBorder: true, bottomBorder: true, leftBorder: true, rightBorder: true, columnSeparator: true, rowSeparator: true}
)

const (
	BorderRune       = '┼'
	RowSeparatorRune = '─'
	ColSeparatorRune = '│'
)

func IsBorderRune(r rune) bool {
	return r == BorderRune || r == RowSeparatorRune || r == ColSeparatorRune
}

type BorderConfig struct {
	topBorder       bool
	bottomBorder    bool
	leftBorder      bool
	rightBorder     bool
	headerSeparator bool
	footerSeparator bool
	columnSeparator bool
	rowSeparator    bool
}

//	func (conf *BorderConfig) WithXXX(val bool) *BorderConfig {
//		return &BorderConfig{
//			conf.topBorder,
//			conf.bottomBorder,
//			conf.leftBorder,
//			conf.rightBorder,
//			conf.headerSeparator,
//			conf.footerSeparator,
//			conf.columnSeparator,
//			conf.rowSeparator,
//		}
//	}
func (conf *BorderConfig) WithTopBorder(val bool) *BorderConfig {
	return &BorderConfig{
		val,
		conf.bottomBorder,
		conf.leftBorder,
		conf.rightBorder,
		conf.headerSeparator,
		conf.footerSeparator,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithBottomBorder(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		val,
		conf.leftBorder,
		conf.rightBorder,
		conf.headerSeparator,
		conf.footerSeparator,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithLeftBorder(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		val,
		conf.rightBorder,
		conf.headerSeparator,
		conf.footerSeparator,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithRightBorder(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		conf.leftBorder,
		val,
		conf.headerSeparator,
		conf.footerSeparator,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithHeaderSeparator(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		conf.leftBorder,
		conf.rightBorder,
		val,
		conf.footerSeparator,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithFooterSeparator(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		conf.leftBorder,
		conf.rightBorder,
		conf.headerSeparator,
		val,
		conf.columnSeparator,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithColumnSeparator(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		conf.leftBorder,
		conf.rightBorder,
		conf.headerSeparator,
		conf.footerSeparator,
		val,
		conf.rowSeparator,
	}
}
func (conf *BorderConfig) WithRowSeparator(val bool) *BorderConfig {
	return &BorderConfig{
		conf.topBorder,
		conf.bottomBorder,
		conf.leftBorder,
		conf.rightBorder,
		conf.headerSeparator,
		conf.footerSeparator,
		conf.columnSeparator,
		val,
	}
}

func (conf *BorderConfig) String() string {
	t := NewTable(4, 4, conf)
	t.Append("COL1")
	t.Append("COL2")
	t.Append("COL3")
	t.Append("COL4")

	t.Append("val1")
	t.Append("val2")
	t.Append("val3")
	t.Append("val4")

	t.Append("val5")
	t.Append("val6")
	t.Append("val7")
	t.Append("val8")

	t.Append("val9")
	t.Append("val10")
	t.Append("val11")
	t.Append("val12")
	return t.String()

}
func (conf *BorderConfig) GetSeparatorRightOf(col, numberOfColumns int) rune {
	return conf.GetSeparatorLeftOf(col+1, numberOfColumns)
}
func (conf *BorderConfig) GetSeparatorLeftOf(col, numberOfColumns int) rune {
	switch {
	case col == 0 && conf.leftBorder:
		return BorderRune
	case col == numberOfColumns && conf.rightBorder:
		return BorderRune
	case col > 0 && col < numberOfColumns && conf.columnSeparator:
		return ColSeparatorRune
	default:
		return 0
	}
}
func (conf *BorderConfig) GetSeparatorBelow(row, numberOfRows int) rune {
	return conf.GetSeparatorAbove(row+1, numberOfRows)
}
func (conf *BorderConfig) GetSeparatorAbove(row, numberOfRows int) rune {
	switch {
	case row == 0 && conf.topBorder:
		return BorderRune
	case row == 1 && conf.headerSeparator:
		return RowSeparatorRune
	case row == numberOfRows-1 && conf.footerSeparator:
		return RowSeparatorRune
	case row == numberOfRows && conf.bottomBorder:
		return BorderRune
	case row >= 1 && row < numberOfRows && conf.rowSeparator:
		return RowSeparatorRune
	default:
		return 0
	}
}
