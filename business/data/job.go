package job

import (
	"context"
	"database/sql"
	"errors"
)

var (
	// ErrNotFound is used when a specific Product is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

// List returns all records
func List(ctx context.Context, db *sql.DB) ([]Job, error) {
	stmt, err := db.Prepare(`SELECT aplctn_cd, s3_bkt_key_cmbntn, domain_cd, sor_cd, trgt_tbl_nm, trgt_schma, prcsng_type, job_id, job_nm, load_type, s3_inbnd_bkt, s3_inbnd_key, s3_archv_bkt, s3_archv_key, s3_app_bkt, s3_app_key, s3_cnsmptn_bkt, s3_cnsmptn_key, trgt_pltfrm, del_tbl_nm, stg_tbl_nm, stg_schma, load_type, key_list, del_key_list, src_file_type, unld_file_type, unld_partn_key, unld_frqncy, unld_type, dlmtr, vrnc_alwd_pct, post_load_mthd, job_type, etl_job_parms, load_frqncy, dscvr_schma_flag, actv_flag, creat_dtm, rqstr_id, ownrshp_team, last_updt_dtm FROM edl_job_cnfgrn ORDER BY domain_cd, sor_cd, trgt_tbl_nm`)
	if err != nil {
		return []Job{}, err
	}
	defer stmt.Close()

	// Returns the results of the query with passed in id
	rows, err := stmt.Query()

	if err != nil {
		return []Job{}, err
	}

	defer rows.Close()

	returnObject := []Job{}
	// .Next() and .Scan() are used in the database/sql library, functioning as loop-like
	for rows.Next() {
		// create a temp struct to scan rows into and add to our tableObject slice
		var newRow Job

		err := rows.Scan(&newRow.AplctnCD,
			&newRow.S3BktKeyCmbntn,
			&newRow.DomainCD,
			&newRow.SorCD,
			&newRow.TrgtTblNm,
			&newRow.TrgtSchma,
			&newRow.PrcsngType,
			&newRow.JobID,
			&newRow.JobNm,
			&newRow.LoadType,
			&newRow.S3InbndBkt,
			&newRow.S3InbndKey,
			&newRow.S3ArchvBkt,
			&newRow.S3ArchvKey,
			&newRow.S3AppBkt,
			&newRow.S3AppKey,
			&newRow.S3CnsmptnBkt,
			&newRow.S3CnsmptnKey,
			&newRow.TrgtPltfrm,
			&newRow.DelTblNm,
			&newRow.StgTblNm,
			&newRow.StgSchma,
			&newRow.LoadType,
			&newRow.KeyList,
			&newRow.DelKeyList,
			&newRow.SrcFileType,
			&newRow.UnldFileType,
			&newRow.UnldPartnKey,
			&newRow.UnldFrqncy,
			&newRow.UnldType,
			&newRow.Dlmtr,
			&newRow.VrncAlwdPct,
			&newRow.PostLoadMthd,
			&newRow.JobType,
			&newRow.EtlJobParms,
			&newRow.LoadFrqncy,
			&newRow.DscvrSchmaFlag,
			&newRow.ActvFlag,
			&newRow.CreatDtm,
			&newRow.RqstrID,
			&newRow.OwnrshpTeam,
			&newRow.LastUpdtDtm,
		)
		if err != nil {
			return []Job{}, err
		}

		returnObject = append(returnObject, newRow)
	}

	if err = rows.Err(); err != nil {
		return []Job{}, err
	}

	return returnObject, nil
}
