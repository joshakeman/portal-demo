package job

import "database/sql"

type Job struct {
	AplctnCD       sql.NullString  `db:"aplctn_cd" json:"aplctn_cd"`
	S3BktKeyCmbntn sql.NullString  `db:"s3_bkt_key_cmbntn" json:"s3_bkt_key_cmbntn"`
	DomainCD       sql.NullString  `db:"domain_cd" json:"domain_cd"`
	SorCD          sql.NullString  `db:"sor_cd" json:"sor_cd"`
	TrgtTblNm      sql.NullString  `db:"trgt_tbl_nm" json:"trgt_tbl_nm"`
	TrgtSchma      sql.NullString  `db:"trgt_schma" json:"trgt_schma"`
	PrcsngType     sql.NullString  `db:"prcsng_type" json:"prcsng_type"`
	JobID          sql.NullInt64   `db:"job_id" json:"job_id"`
	JobNm          sql.NullString  `db:"job_nm" json:"job_nm"`
	LoadType       sql.NullString  `db:"load_type" json:"load_type"`
	S3InbndBkt     sql.NullString  `db:"s3_inbnd_bkt" json:"s3_inbnd_bkt"`
	S3InbndKey     sql.NullString  `db:"s3_inbnd_key" json:"s3_inbnd_key"`
	S3ArchvBkt     sql.NullString  `db:"s3_archv_bkt" json:"s3_archv_bkt"`
	S3ArchvKey     sql.NullString  `db:"s3_archv_key" json:"s3_archv_key"`
	S3AppBkt       sql.NullString  `db:"s3_app_bkt" json:"s3_app_bkt"`
	S3AppKey       sql.NullString  `db:"s3_app_key" json:"s3_app_key"`
	S3CnsmptnBkt   sql.NullString  `db:"s3_cnsmptn_bkt" json:"s3_cnsmptn_bkt"`
	S3CnsmptnKey   sql.NullString  `db:"s3_cnsmptn_key" json:"s3_cnsmptn_key"`
	TrgtPltfrm     sql.NullString  `db:"trgt_pltfrm" json:"trgt_pltfrm"`
	DelTblNm       sql.NullString  `db:"del_tbl_nm" json:"del_tbl_nm"`
	StgTblNm       sql.NullString  `db:"stg_tbl_nm" json:"stg_tbl_nm"`
	StgSchma       sql.NullString  `db:"stg_schma" json:"stg_schma"`
	KeyList        sql.NullString  `db:"key_list" json:"key_list"`
	DelKeyList     sql.NullString  `db:"del_key_list" json:"del_key_list"`
	SrcFileType    sql.NullString  `db:"src_file_type" json:"src_file_type"`
	UnldFileType   sql.NullString  `db:"unld_file_type" json:"unld_file_type"`
	UnldPartnKey   sql.NullString  `db:"unld_partn_key" json:"unld_partn_key"`
	UnldFrqncy     sql.NullString  `db:"unld_frqncy" json:"unld_frqncy"`
	UnldType       sql.NullString  `db:"unld_type" json:"unld_type"`
	Dlmtr          sql.NullString  `db:"dlmtr" json:"dlmtr"`
	VrncAlwdPct    sql.NullFloat64 `db:"vrnc_alwd_pct" json:"vrnc_alwd_pct"`
	PostLoadMthd   sql.NullString  `db:"post_load_mthd" json:"post_load_mthd"`
	JobType        sql.NullString  `db:"job_type" json:"job_type"`
	EtlJobParms    sql.NullString  `db:"etl_job_parms" json:"etl_job_parms"`
	LoadFrqncy     sql.NullString  `db:"load_frqncy" json:"load_frqncy"`
	DscvrSchmaFlag sql.NullString  `db:"dscvr_schma_flag" json:"dscvr_schma_flag"`
	ActvFlag       sql.NullString  `db:"actv_flag" json:"actv_flag"`
	CreatDtm       sql.NullString  `db:"creat_dtm" json:"creat_dtm"`
	RqstrID        sql.NullString  `db:"rqstr_id" json:"rqstr_id"`
	OwnrshpTeam    sql.NullString  `db:"ownrshp_team" json:"ownrshp_team"`
	LastUpdtDtm    sql.NullString  `db:"last_updt_dtm" json:"last_updt_dtm"`
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
