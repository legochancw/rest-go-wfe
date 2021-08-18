package entity

//CIRC_TASK_ID
//CIRC_ID
//CIRC_TASK_TYPE
//CIRC_TASK_TTL
//CIRC_TASK_REM

type TMPL_CIRC_TASK struct {
	CIRC_TASK_ID   NullString `json:"circ_task_id"`
	CIRC_ID        NullString `json:"circ_id"`
	CIRC_TASK_TYPE NullString `json:"circ_task_type"`
	CIRC_TASK_TTL  NullString `json:"circ_task_ttl"`
	CIRC_TASK_REM  NullString `json:"circ_task_rem,omitempty"`
}
