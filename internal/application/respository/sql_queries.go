package respository

const (
	createApplication = `INSERT INTO applications (user_id, college_name, course_name, intake_date, current_stage, status)
							VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	updateApplication = `UPDATE applications 
						SET current_stage = COALESCE(NULLIF($1, ''), current_stage),
						    status = COALESCE(NULLIF($2, ''), status),
						    college_name = COALESCE(NULLIF($3, ''), college_name),
						    course_name = COALESCE(NULLIF($4, ''), course_name),
						    intake_date = COALESCE(NULLIF($5, ''), intake_date),
						    updated_at = now()
						WHERE user_id = $7
						RETURNING *
				`
	getApplicationStatus = `SELECT 
								a.id,
								a.current_stage,
								a.status,
								a.college_name,
								a.course_name,
								a.intake_date,
								a.created_at,
								a.updated_at,
								COALESCE(
									json_agg(
										json_build_object(
											'stage', sp.stage,
											'status', sp.status,
											'start_date', sp.start_date,
											'completion_date', sp.completion_date
										)
										ORDER BY sp.start_date
									) FILTER (WHERE sp.id IS NOT NULL),
									'[]'
								) as stages
							FROM applications a
							LEFT JOIN stage_progress sp ON sp.application_id = a.id
							WHERE a.user_id = $1
							GROUP BY 
								a.id, 
								a.current_stage, 
								a.status, 
								a.college_name, 
								a.course_name, 
								a.intake_date, 
								a.created_at, 
								a.updated_at
							ORDER BY a.created_at DESC`
	getApplicationByID = `SELECT id, user_id, college_name, course_name, intake_date, current_stage, status, created_at, updated_at
							FROM applications WHERE id = $1`
	getApplicationByUserID = `SELECT id, user_id, college_name, course_name, intake_date, current_stage, status, created_at, updated_at
							FROM applications WHERE user_id = $1`
)
