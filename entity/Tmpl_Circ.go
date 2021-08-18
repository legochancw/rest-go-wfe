package entity

// `CIRC_ID` int(11) NOT NULL AUTO_INCREMENT,
//   `CIRC_TYPE` varchar(1) NOT NULL DEFAULT "D" COMMENT "D: Dynamic - General Workflow\nH: Hardcode - Mobile hardcoded UI flow",
//   `CIRC_TTL` varchar(100) NOT NULL,
//   `CIRC_CAT` varchar(100) NOT NULL,
//   `CIRC_SUB_CAT` varchar(100) DEFAULT NULL,
//   `STS` varchar(100) NOT NULL DEFAULT "A" COMMENT "A: Active",
//   `PROJ_ID` varchar(100) DEFAULT "ALL" COMMENT "Chun Wo Project ID (at this moment)",
//   `DUR` int(11) NOT NULL DEFAULT "7",

type TMPL_CIRC struct {
	CIRC_ID      NullString `json:"circ_id"`
	CIRC_TYPE    NullString `json:"circ_type"`
	CIRC_TTL     NullString `json:"circ_ttl"`
	CIRC_CAT     NullString `json:"circ_cat"`
	CIRC_SUB_CAT NullString `json:"circ_sub_cat,omitempty"`
	STS          NullString `json:"sts"`
	PROJ_ID      NullString `json:"proj_id"`
	DUR          NullString `json:"dur"`
}
