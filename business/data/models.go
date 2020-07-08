package job

import "database/sql"

type Job struct {
	AplctnCD       sql.NullString  `json:"aplctn_cd"`
	S3BktKeyCmbntn sql.NullString  `json:"s3_bkt_key_cmbntn"`
	DomainCD       sql.NullString  `json:"domain_cd"`
	SorCD          sql.NullString  `json:"sor_cd"`
	TrgtTblNm      sql.NullString  `json:"trgt_tbl_nm"`
	TrgtSchma      sql.NullString  `json:"trgt_schma"`
	PrcsngType     sql.NullString  `json:"prcsng_type"`
	JobID          sql.NullInt64   `json:"job_id"`
	JobNm          sql.NullString  `json:"job_nm"`
	LoadType       sql.NullString  `json:"load_type"`
	S3InbndBkt     sql.NullString  `json:"s3_inbnd_bkt"`
	S3InbndKey     sql.NullString  `json:"s3_inbnd_key"`
	S3ArchvBkt     sql.NullString  `json:"s3_archv_bkt"`
	S3ArchvKey     sql.NullString  `json:"s3_archv_key"`
	S3AppBkt       sql.NullString  `json:"s3_app_bkt"`
	S3AppKey       sql.NullString  `json:"s3_app_key"`
	S3CnsmptnBkt   sql.NullString  `json:"s3_cnsmptn_bkt"`
	S3CnsmptnKey   sql.NullString  `json:"s3_cnsmptn_key"`
	TrgtPltfrm     sql.NullString  `json:"trgt_pltfrm"`
	DelTblNm       sql.NullString  `json:"del_tbl_nm"`
	StgTblNm       sql.NullString  `json:"stg_tbl_nm"`
	StgSchma       sql.NullString  `json:"stg_schma"`
	KeyList        sql.NullString  `json:"key_list"`
	DelKeyList     sql.NullString  `json:"del_key_list"`
	SrcFileType    sql.NullString  `json:"src_file_type"`
	UnldFileType   sql.NullString  `json:"unld_file_type"`
	UnldPartnKey   sql.NullString  `json:"unld_partn_key"`
	UnldFrqncy     sql.NullString  `json:"unld_frqncy"`
	UnldType       sql.NullString  `json:"unld_type"`
	Dlmtr          sql.NullString  `json:"dlmtr"`
	VrncAlwdPct    sql.NullFloat64 `json:"vrnc_alwd_pct"`
	PostLoadMthd   sql.NullString  `json:"post_load_mthd"`
	JobType        sql.NullString  `json:"job_type"`
	EtlJobParms    sql.NullString  `json:"etl_job_parms"`
	LoadFrqncy     sql.NullString  `json:"load_frqncy"`
	DscvrSchmaFlag sql.NullString  `json:"dscvr_schma_flag"`
	ActvFlag       sql.NullString  `json:"actv_flag"`
	CreatDtm       sql.NullString  `json:"creat_dtm"`
	RqstrID        sql.NullString  `json:"rqstr_id"`
	OwnrshpTeam    sql.NullString  `json:"ownrshp_team"`
	LastUpdtDtm    sql.NullString  `json:"last_updt_dtm"`
}

type NewJob struct {
	AplctnCD       string `json:"aplctn_cd"`
	S3BktKeyCmbntn string `json:"s3_bkt_key_cmbntn"`
	DomainCD       string `json:"domain_cd"`
	SorCD          string `json:"sor_cd"`
	TrgtTblNm      string `json:"trgt_tbl_nm"`
	TrgtSchma      string `json:"trgt_schma"`
	PrcsngType     string `json:"prcsng_type"`
	JobID          string `json:"job_id"`
	JobNm          string `json:"job_nm"`
	LoadType       string `json:"load_type"`
	S3InbndBkt     string `json:"s3_inbnd_bkt"`
	S3InbndKey     string `json:"s3_inbnd_key"`
	S3ArchvBkt     string `json:"s3_archv_bkt"`
	S3ArchvKey     string `json:"s3_archv_key"`
	S3AppBkt       string `json:"s3_app_bkt"`
	S3AppKey       string `json:"s3_app_key"`
	S3CnsmptnBkt   string `json:"s3_cnsmptn_bkt"`
	S3CnsmptnKey   string `json:"s3_cnsmptn_key"`
	TrgtPltfrm     string `json:"trgt_pltfrm"`
	DelTblNm       string `json:"del_tbl_nm"`
	StgTblNm       string `json:"stg_tbl_nm"`
	StgSchma       string `json:"stg_schma"`
	KeyList        string `json:"key_list"`
	DelKeyList     string `json:"del_key_list"`
	SrcFileType    string `json:"src_file_type"`
	UnldFileType   string `json:"unld_file_type"`
	UnldPartnKey   string `json:"unld_partn_key"`
	UnldFrqncy     string `json:"unld_frqncy"`
	UnldType       string `json:"unld_type"`
	Dlmtr          string `json:"dlmtr"`
	VrncAlwdPct    string `json:"vrnc_alwd_pct"`
	PostLoadMthd   string `json:"post_load_mthd"`
	JobType        string `json:"job_type"`
	EtlJobParms    string `json:"etl_job_parms"`
	LoadFrqncy     string `json:"load_frqncy"`
	DscvrSchmaFlag string `json:"dscvr_schma_flag"`
	ActvFlag       string `json:"actv_flag"`
	CreatDtm       string `json:"creat_dtm"`
	RqstrID        string `json:"rqstr_id"`
	OwnrshpTeam    string `json:"ownrshp_team"`
	LastUpdtDtm    string `json:"last_updt_dtm"`
}
