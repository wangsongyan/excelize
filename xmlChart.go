// Copyright 2016 - 2024 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.

package excelize

import "encoding/xml"

// xlsxChartSpace directly maps the chartSpace element. The chart namespace in
// DrawingML is for representing visualizations of numeric data with column
// charts, pie charts, scatter charts, or other types of charts.
type xlsxChartSpace struct {
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/drawingml/2006/chart chartSpace"`
	XMLNSa         string          `xml:"xmlns:a,attr"`
	Date1904       *attrValBool    `xml:"date1904"`
	Lang           *attrValString  `xml:"lang"`
	RoundedCorners *attrValBool    `xml:"roundedCorners"`
	Chart          cChart          `xml:"chart"`
	SpPr           *cSpPr          `xml:"spPr"`
	TxPr           *cTxPr          `xml:"txPr"`
	PrintSettings  *cPrintSettings `xml:"printSettings"`
}

// cThicknessSpPr directly maps the element that specifies the thickness of
// the walls or floor as a percentage of the largest dimension of the plot
// volume and SpPr element.
type cThicknessSpPr struct {
	Thickness *attrValInt `xml:"thickness"`
	SpPr      *cSpPr      `xml:"spPr"`
}

// cChart (Chart) directly maps the chart element. This element specifies a
// title.
type cChart struct {
	Title            *cTitle            `xml:"title"`
	AutoTitleDeleted *cAutoTitleDeleted `xml:"autoTitleDeleted"`
	View3D           *cView3D           `xml:"view3D"`
	Floor            *cThicknessSpPr    `xml:"floor"`
	SideWall         *cThicknessSpPr    `xml:"sideWall"`
	BackWall         *cThicknessSpPr    `xml:"backWall"`
	PlotArea         *cPlotArea         `xml:"plotArea"`
	Legend           *cLegend           `xml:"legend"`
	PlotVisOnly      *attrValBool       `xml:"plotVisOnly"`
	DispBlanksAs     *attrValString     `xml:"dispBlanksAs"`
	ShowDLblsOverMax *attrValBool       `xml:"showDLblsOverMax"`
}

// cTitle (Title) directly maps the title element. This element specifies a
// title.
type cTitle struct {
	Tx      cTx          `xml:"tx,omitempty"`
	Layout  string       `xml:"layout,omitempty"`
	Overlay *attrValBool `xml:"overlay"`
	SpPr    cSpPr        `xml:"spPr,omitempty"`
	TxPr    cTxPr        `xml:"txPr,omitempty"`
}

// cTx (Chart Text) directly maps the tx element. This element specifies text
// to use on a chart, including rich text formatting.
type cTx struct {
	StrRef *cStrRef `xml:"strRef"`
	Rich   *cRich   `xml:"rich,omitempty"`
}

// cRich (Rich Text) directly maps the rich element. This element contains a
// string with rich text formatting.
type cRich struct {
	BodyPr   aBodyPr `xml:"a:bodyPr,omitempty"`
	LstStyle string  `xml:"a:lstStyle,omitempty"`
	P        []aP    `xml:"a:p"`
}

// aBodyPr (Body Properties) directly maps the a:bodyPr element. This element
// defines the body properties for the text body within a shape.
type aBodyPr struct {
	Anchor           string  `xml:"anchor,attr,omitempty"`
	AnchorCtr        bool    `xml:"anchorCtr,attr"`
	Rot              int     `xml:"rot,attr"`
	BIns             float64 `xml:"bIns,attr,omitempty"`
	CompatLnSpc      bool    `xml:"compatLnSpc,attr,omitempty"`
	ForceAA          bool    `xml:"forceAA,attr,omitempty"`
	FromWordArt      bool    `xml:"fromWordArt,attr,omitempty"`
	HorzOverflow     string  `xml:"horzOverflow,attr,omitempty"`
	LIns             float64 `xml:"lIns,attr,omitempty"`
	NumCol           int     `xml:"numCol,attr,omitempty"`
	RIns             float64 `xml:"rIns,attr,omitempty"`
	RtlCol           bool    `xml:"rtlCol,attr,omitempty"`
	SpcCol           int     `xml:"spcCol,attr,omitempty"`
	SpcFirstLastPara bool    `xml:"spcFirstLastPara,attr"`
	TIns             float64 `xml:"tIns,attr,omitempty"`
	Upright          bool    `xml:"upright,attr,omitempty"`
	Vert             string  `xml:"vert,attr,omitempty"`
	VertOverflow     string  `xml:"vertOverflow,attr,omitempty"`
	Wrap             string  `xml:"wrap,attr,omitempty"`
}

// aP (Paragraph) directly maps the a:p element. This element specifies a
// paragraph of content in the document.
type aP struct {
	PPr        *aPPr        `xml:"a:pPr"`
	R          *aR          `xml:"a:r"`
	EndParaRPr *aEndParaRPr `xml:"a:endParaRPr"`
}

// aPPr (Paragraph Properties) directly maps the a:pPr element. This element
// specifies a set of paragraph properties which shall be applied to the
// contents of the parent paragraph after all style/numbering/table properties
// have been applied to the text. These properties are defined as direct
// formatting, since they are directly applied to the paragraph and supersede
// any formatting from styles.
type aPPr struct {
	DefRPr aRPr `xml:"a:defRPr"`
}

// aSolidFill (Solid Fill) directly maps the solidFill element. This element
// specifies a solid color fill. The shape is filled entirely with the specified
// color.
type aSolidFill struct {
	SchemeClr *aSchemeClr    `xml:"a:schemeClr"`
	SrgbClr   *attrValString `xml:"a:srgbClr"`
}

// aSchemeClr (Scheme Color) directly maps the a:schemeClr element. This
// element specifies a color bound to a user's theme. As with all elements which
// define a color, it is possible to apply a list of color transforms to the
// base color defined.
type aSchemeClr struct {
	Val    string      `xml:"val,attr,omitempty"`
	LumMod *attrValInt `xml:"a:lumMod"`
	LumOff *attrValInt `xml:"a:lumOff"`
}

// attrValInt directly maps the val element with integer data type as an
// attribute.
type attrValInt struct {
	Val *int `xml:"val,attr"`
}

// attrValFloat directly maps the val element with float64 data type as an
// attribute.
type attrValFloat struct {
	Val *float64 `xml:"val,attr"`
}

// attrValBool directly maps the val element with boolean data type as an
// attribute.
type attrValBool struct {
	Val *bool `xml:"val,attr"`
}

// attrValString directly maps the val element with string data type as an
// attribute.
type attrValString struct {
	Val *string `xml:"val,attr"`
}

// aCs directly maps the a:cs element.
type aCs struct {
	Typeface string `xml:"typeface,attr"`
}

// aEa directly maps the a:ea element.
type aEa struct {
	Typeface string `xml:"typeface,attr"`
}

type xlsxCTTextFont struct {
	Typeface    string `xml:"typeface,attr"`
	Panose      string `xml:"panose,attr,omitempty"`
	PitchFamily string `xml:"pitchFamily,attr,omitempty"`
	Charset     string `xml:"Charset,attr,omitempty"`
}

// aR directly maps the a:r element.
type aR struct {
	RPr aRPr   `xml:"a:rPr,omitempty"`
	T   string `xml:"a:t,omitempty"`
}

// aRPr (Run Properties) directly maps the rPr element. This element
// specifies a set of run properties which shall be applied to the contents of
// the parent run after all style formatting has been applied to the text. These
// properties are defined as direct formatting, since they are directly applied
// to the run and supersede any formatting from styles.
type aRPr struct {
	AltLang    string          `xml:"altLang,attr,omitempty"`
	B          bool            `xml:"b,attr"`
	Baseline   int             `xml:"baseline,attr"`
	Bmk        string          `xml:"bmk,attr,omitempty"`
	Cap        string          `xml:"cap,attr,omitempty"`
	Dirty      bool            `xml:"dirty,attr,omitempty"`
	Err        bool            `xml:"err,attr,omitempty"`
	I          bool            `xml:"i,attr"`
	Kern       int             `xml:"kern,attr"`
	Kumimoji   bool            `xml:"kumimoji,attr,omitempty"`
	Lang       string          `xml:"lang,attr,omitempty"`
	NoProof    bool            `xml:"noProof,attr,omitempty"`
	NormalizeH bool            `xml:"normalizeH,attr,omitempty"`
	SmtClean   bool            `xml:"smtClean,attr,omitempty"`
	SmtID      uint64          `xml:"smtId,attr,omitempty"`
	Spc        int             `xml:"spc,attr"`
	Strike     string          `xml:"strike,attr,omitempty"`
	Sz         float64         `xml:"sz,attr,omitempty"`
	U          string          `xml:"u,attr,omitempty"`
	SolidFill  *aSolidFill     `xml:"a:solidFill"`
	Latin      *xlsxCTTextFont `xml:"a:latin"`
	Ea         *aEa            `xml:"a:ea"`
	Cs         *aCs            `xml:"a:cs"`
}

// cSpPr (Shape Properties) directly maps the spPr element. This element
// specifies the visual shape properties that can be applied to a shape. These
// properties include the shape fill, outline, geometry, effects, and 3D
// orientation.
type cSpPr struct {
	NoFill    *string     `xml:"a:noFill"`
	SolidFill *aSolidFill `xml:"a:solidFill"`
	Ln        *aLn        `xml:"a:ln"`
	Sp3D      *aSp3D      `xml:"a:sp3d"`
	EffectLst *string     `xml:"a:effectLst"`
}

// aSp3D (3-D Shape Properties) directly maps the a:sp3d element. This element
// defines the 3D properties associated with a particular shape in DrawingML.
// The 3D properties which can be applied to a shape are top and bottom bevels,
// a contour and an extrusion.
type aSp3D struct {
	ContourW   int          `xml:"contourW,attr"`
	ContourClr *aContourClr `xml:"a:contourClr"`
}

// aContourClr (Contour Color) directly maps the a:contourClr element. This
// element defines the color for the contour on a shape. The contour of a shape
// is a solid filled line which surrounds the outer edges of the shape.
type aContourClr struct {
	SchemeClr *aSchemeClr `xml:"a:schemeClr"`
}

// aLn (Outline) directly maps the a:ln element. This element specifies an
// outline style that can be applied to a number of different objects such as
// shapes and text. The line allows for the specifying of many different types
// of outlines including even line dashes and bevels.
type aLn struct {
	Algn      string         `xml:"algn,attr,omitempty"`
	Cap       string         `xml:"cap,attr,omitempty"`
	Cmpd      string         `xml:"cmpd,attr,omitempty"`
	W         int            `xml:"w,attr,omitempty"`
	NoFill    *attrValString `xml:"a:noFill"`
	Round     string         `xml:"a:round,omitempty"`
	SolidFill *aSolidFill    `xml:"a:solidFill"`
}

// cTxPr (Text Properties) directly maps the txPr element. This element
// specifies text formatting. The lstStyle element is not supported.
type cTxPr struct {
	BodyPr   aBodyPr `xml:"a:bodyPr,omitempty"`
	LstStyle string  `xml:"a:lstStyle,omitempty"`
	P        aP      `xml:"a:p,omitempty"`
}

// aEndParaRPr (End Paragraph Run Properties) directly maps the a:endParaRPr
// element. This element specifies the text run properties that are to be used
// if another run is inserted after the last run specified. This effectively
// saves the run property state so that it can be applied when the user enters
// additional text. If this element is omitted, then the application can
// determine which default properties to apply. It is recommended that this
// element be specified at the end of the list of text runs within the paragraph
// so that an orderly list is maintained.
type aEndParaRPr struct {
	Lang    string `xml:"lang,attr"`
	AltLang string `xml:"altLang,attr,omitempty"`
	Sz      int    `xml:"sz,attr,omitempty"`
}

// cAutoTitleDeleted (Auto Title Is Deleted) directly maps the
// autoTitleDeleted element. This element specifies the title shall not be
// shown for this chart.
type cAutoTitleDeleted struct {
	Val bool `xml:"val,attr"`
}

// cView3D (View In 3D) directly maps the view3D element. This element
// specifies the 3-D view of the chart.
type cView3D struct {
	RotX         *attrValInt `xml:"rotX"`
	RotY         *attrValInt `xml:"rotY"`
	RAngAx       *attrValInt `xml:"rAngAx"`
	DepthPercent *attrValInt `xml:"depthPercent"`
	Perspective  *attrValInt `xml:"perspective"`
	ExtLst       *xlsxExtLst `xml:"extLst"`
}

// cPlotArea directly maps the plotArea element. This element specifies the
// plot area of the chart.
type cPlotArea struct {
	Layout         *string  `xml:"layout"`
	AreaChart      *cCharts `xml:"areaChart"`
	Area3DChart    *cCharts `xml:"area3DChart"`
	BarChart       *cCharts `xml:"barChart"`
	Bar3DChart     *cCharts `xml:"bar3DChart"`
	BubbleChart    *cCharts `xml:"bubbleChart"`
	DoughnutChart  *cCharts `xml:"doughnutChart"`
	LineChart      *cCharts `xml:"lineChart"`
	Line3DChart    *cCharts `xml:"line3DChart"`
	PieChart       *cCharts `xml:"pieChart"`
	Pie3DChart     *cCharts `xml:"pie3DChart"`
	OfPieChart     *cCharts `xml:"ofPieChart"`
	RadarChart     *cCharts `xml:"radarChart"`
	ScatterChart   *cCharts `xml:"scatterChart"`
	Surface3DChart *cCharts `xml:"surface3DChart"`
	SurfaceChart   *cCharts `xml:"surfaceChart"`
	CatAx          []*cAxs  `xml:"catAx"`
	ValAx          []*cAxs  `xml:"valAx"`
	SerAx          []*cAxs  `xml:"serAx"`
	SpPr           *cSpPr   `xml:"spPr"`
}

// cCharts specifies the common element of the chart.
type cCharts struct {
	BarDir       *attrValString `xml:"barDir"`
	BubbleScale  *attrValFloat  `xml:"bubbleScale"`
	Grouping     *attrValString `xml:"grouping"`
	RadarStyle   *attrValString `xml:"radarStyle"`
	ScatterStyle *attrValString `xml:"scatterStyle"`
	OfPieType    *attrValString `xml:"ofPieType"`
	VaryColors   *attrValBool   `xml:"varyColors"`
	Wireframe    *attrValBool   `xml:"wireframe"`
	Ser          *[]cSer        `xml:"ser"`
	SplitPos     *attrValInt    `xml:"splitPos"`
	SerLines     *attrValString `xml:"serLines"`
	DLbls        *cDLbls        `xml:"dLbls"`
	Shape        *attrValString `xml:"shape"`
	HoleSize     *attrValInt    `xml:"holeSize"`
	Smooth       *attrValBool   `xml:"smooth"`
	Overlap      *attrValInt    `xml:"overlap"`
	AxID         []*attrValInt  `xml:"axId"`
}

// cAxs directly maps the catAx and valAx element.
type cAxs struct {
	AxID           *attrValInt    `xml:"axId"`
	Scaling        *cScaling      `xml:"scaling"`
	Delete         *attrValBool   `xml:"delete"`
	AxPos          *attrValString `xml:"axPos"`
	MajorGridlines *cChartLines   `xml:"majorGridlines"`
	MinorGridlines *cChartLines   `xml:"minorGridlines"`
	Title          *cTitle        `xml:"title"`
	NumFmt         *cNumFmt       `xml:"numFmt"`
	MajorTickMark  *attrValString `xml:"majorTickMark"`
	MinorTickMark  *attrValString `xml:"minorTickMark"`
	TickLblPos     *attrValString `xml:"tickLblPos"`
	SpPr           *cSpPr         `xml:"spPr"`
	TxPr           *cTxPr         `xml:"txPr"`
	CrossAx        *attrValInt    `xml:"crossAx"`
	Crosses        *attrValString `xml:"crosses"`
	CrossBetween   *attrValString `xml:"crossBetween"`
	MajorUnit      *attrValFloat  `xml:"majorUnit"`
	MinorUnit      *attrValFloat  `xml:"minorUnit"`
	Auto           *attrValBool   `xml:"auto"`
	LblAlgn        *attrValString `xml:"lblAlgn"`
	LblOffset      *attrValInt    `xml:"lblOffset"`
	TickLblSkip    *attrValInt    `xml:"tickLblSkip"`
	TickMarkSkip   *attrValInt    `xml:"tickMarkSkip"`
	NoMultiLvlLbl  *attrValBool   `xml:"noMultiLvlLbl"`
}

// cChartLines directly maps the chart lines content model.
type cChartLines struct {
	SpPr *cSpPr `xml:"spPr"`
}

// cScaling directly maps the scaling element. This element contains
// additional axis settings.
type cScaling struct {
	LogBase     *attrValFloat  `xml:"logBase"`
	Orientation *attrValString `xml:"orientation"`
	Max         *attrValFloat  `xml:"max"`
	Min         *attrValFloat  `xml:"min"`
}

// cNumFmt (Numbering Format) directly maps the numFmt element. This element
// specifies number formatting for the parent element.
type cNumFmt struct {
	FormatCode   string `xml:"formatCode,attr"`
	SourceLinked bool   `xml:"sourceLinked,attr"`
}

// cSer directly maps the ser element. This element specifies a series on a
// chart.
type cSer struct {
	IDx              *attrValInt  `xml:"idx"`
	Order            *attrValInt  `xml:"order"`
	Tx               *cTx         `xml:"tx"`
	SpPr             *cSpPr       `xml:"spPr"`
	DPt              []*cDPt      `xml:"dPt"`
	DLbls            *cDLbls      `xml:"dLbls"`
	Marker           *cMarker     `xml:"marker"`
	InvertIfNegative *attrValBool `xml:"invertIfNegative"`
	Cat              *cCat        `xml:"cat"`
	Val              *cVal        `xml:"val"`
	XVal             *cCat        `xml:"xVal"`
	YVal             *cVal        `xml:"yVal"`
	Smooth           *attrValBool `xml:"smooth"`
	BubbleSize       *cVal        `xml:"bubbleSize"`
	Bubble3D         *attrValBool `xml:"bubble3D"`
}

// cMarker (Marker) directly maps the marker element. This element specifies a
// data marker.
type cMarker struct {
	Symbol *attrValString `xml:"symbol"`
	Size   *attrValInt    `xml:"size"`
	SpPr   *cSpPr         `xml:"spPr"`
}

// cDPt (Data Point) directly maps the dPt element. This element specifies a
// single data point.
type cDPt struct {
	IDx      *attrValInt  `xml:"idx"`
	Bubble3D *attrValBool `xml:"bubble3D"`
	SpPr     *cSpPr       `xml:"spPr"`
}

// cCat (Category Axis Data) directly maps the cat element. This element
// specifies the data used for the category axis.
type cCat struct {
	StrRef *cStrRef `xml:"strRef"`
}

// cStrRef (String Reference) directly maps the strRef element. This element
// specifies a reference to data for a single data label or title with a cache
// of the last values used.
type cStrRef struct {
	F        string     `xml:"f"`
	StrCache *cStrCache `xml:"strCache"`
}

// cStrCache (String Cache) directly maps the strCache element. This element
// specifies the last string data used for a chart.
type cStrCache struct {
	Pt      []*cPt      `xml:"pt"`
	PtCount *attrValInt `xml:"ptCount"`
}

// cPt directly maps the pt element. This element specifies data for a
// particular data point.
type cPt struct {
	IDx int     `xml:"idx,attr"`
	V   *string `xml:"v"`
}

// cVal directly maps the val element. This element specifies the data values
// which shall be used to define the location of data markers on a chart.
type cVal struct {
	NumRef *cNumRef `xml:"numRef"`
}

// cNumRef directly maps the numRef element. This element specifies a
// reference to numeric data with a cache of the last values used.
type cNumRef struct {
	F        string     `xml:"f"`
	NumCache *cNumCache `xml:"numCache"`
}

// cNumCache directly maps the numCache element. This element specifies the
// last data shown on the chart for a series.
type cNumCache struct {
	FormatCode string      `xml:"formatCode"`
	Pt         []*cPt      `xml:"pt"`
	PtCount    *attrValInt `xml:"ptCount"`
}

// cDLbls (Data Labels) directly maps the dLbls element. This element serves
// as a root element that specifies the settings for the data labels for an
// entire series or the entire chart. It contains child elements that specify
// the specific formatting and positioning settings.
type cDLbls struct {
	NumFmt          *cNumFmt       `xml:"numFmt"`
	DLblPos         *attrValString `xml:"dLblPos"`
	ShowLegendKey   *attrValBool   `xml:"showLegendKey"`
	ShowVal         *attrValBool   `xml:"showVal"`
	ShowCatName     *attrValBool   `xml:"showCatName"`
	ShowSerName     *attrValBool   `xml:"showSerName"`
	ShowPercent     *attrValBool   `xml:"showPercent"`
	ShowBubbleSize  *attrValBool   `xml:"showBubbleSize"`
	ShowLeaderLines *attrValBool   `xml:"showLeaderLines"`
}

// cLegend (Legend) directly maps the legend element. This element specifies
// the legend.
type cLegend struct {
	Layout    *string        `xml:"layout"`
	LegendPos *attrValString `xml:"legendPos"`
	Overlay   *attrValBool   `xml:"overlay"`
	SpPr      *cSpPr         `xml:"spPr"`
	TxPr      *cTxPr         `xml:"txPr"`
}

// cPrintSettings directly maps the printSettings element. This element
// specifies the print settings for the chart.
type cPrintSettings struct {
	HeaderFooter *string       `xml:"headerFooter"`
	PageMargins  *cPageMargins `xml:"pageMargins"`
	PageSetup    *string       `xml:"pageSetup"`
}

// cPageMargins directly maps the pageMargins element. This element specifies
// the page margins for a chart.
type cPageMargins struct {
	B      float64 `xml:"b,attr"`
	Footer float64 `xml:"footer,attr"`
	Header float64 `xml:"header,attr"`
	L      float64 `xml:"l,attr"`
	R      float64 `xml:"r,attr"`
	T      float64 `xml:"t,attr"`
}

// ChartNumFmt directly maps the number format settings of the chart.
type ChartNumFmt struct {
	CustomNumFmt string
	SourceLinked bool
}

// ChartAxis directly maps the format settings of the chart axis.
type ChartAxis struct {
	None           bool
	MajorGridLines bool
	MinorGridLines bool
	MajorUnit      float64
	TickLabelSkip  int
	ReverseOrder   bool
	Secondary      bool
	Maximum        *float64
	Minimum        *float64
	Font           Font
	LogBase        float64
	NumFmt         ChartNumFmt
	Title          []RichTextRun
	axID           int
}

// ChartDimension directly maps the dimension of the chart.
type ChartDimension struct {
	Width  uint
	Height uint
}

// ChartPlotArea directly maps the format settings of the plot area.
type ChartPlotArea struct {
	SecondPlotValues int
	ShowBubbleSize   bool
	ShowCatName      bool
	ShowLeaderLines  bool
	ShowPercent      bool
	ShowSerName      bool
	ShowVal          bool
	NumFmt           ChartNumFmt
}

// Chart directly maps the format settings of the chart.
type Chart struct {
	Type         ChartType
	Series       []ChartSeries
	Format       GraphicOptions
	Dimension    ChartDimension
	Legend       ChartLegend
	Title        []RichTextRun
	VaryColors   *bool
	XAxis        ChartAxis
	YAxis        ChartAxis
	PlotArea     ChartPlotArea
	Border       ChartLine
	ShowBlanksAs string
	BubbleSize   int
	HoleSize     int
	order        int
}

// ChartLegend directly maps the format settings of the chart legend.
type ChartLegend struct {
	Position      string
	ShowLegendKey bool
}

// ChartMarker directly maps the format settings of the chart marker.
type ChartMarker struct {
	Symbol string
	Size   int
}

// ChartLine directly maps the format settings of the chart line.
type ChartLine struct {
	Type   ChartLineType
	Smooth bool
	Width  float64
}

// ChartSeries directly maps the format settings of the chart series.
type ChartSeries struct {
	Name              string
	Categories        string
	Values            string
	Sizes             string
	Fill              Fill
	Line              ChartLine
	Marker            ChartMarker
	DataLabelPosition ChartDataLabelPositionType
}
