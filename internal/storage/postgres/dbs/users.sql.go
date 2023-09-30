// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: users.sql

package dbs

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const socialUserProfiles = `-- name: SocialUserProfiles :many
SELECT u.id,
       s.social_provider,
       s.social_provider_user_id,
       s.username,
       s.name,
       u.created_at
FROM users u
         INNER JOIN user_social_profiles s ON u.id = s.user_id
WHERE s.deleted_at IS NULL
ORDER BY s.id DESC
LIMIT $1
`

type SocialUserProfilesRow struct {
	ID                   int64
	SocialProvider       SocialProvider
	SocialProviderUserID string
	Username             string
	Name                 string
	CreatedAt            pgtype.Timestamp
}

func (q *Queries) SocialUserProfiles(ctx context.Context, limit int32) ([]SocialUserProfilesRow, error) {
	rows, err := q.db.Query(ctx, socialUserProfiles, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SocialUserProfilesRow
	for rows.Next() {
		var i SocialUserProfilesRow
		if err := rows.Scan(
			&i.ID,
			&i.SocialProvider,
			&i.SocialProviderUserID,
			&i.Username,
			&i.Name,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userRegistrationDailyCountStats = `-- name: UserRegistrationDailyCountStats :many
SELECT days.day::DATE                    AS day,
       COALESCE(s.user_count, 0)::BIGINT AS user_count
FROM GENERATE_SERIES(
    $1::DATE,
    $2::DATE,
    '1 DAY'::INTERVAL
) AS days (day)
    LEFT JOIN (
        SELECT DATE_TRUNC('DAY', created_at) AS day,
               COUNT(*)                      AS user_count
        FROM users
        WHERE created_at >= $1::DATE
        GROUP BY day
    ) AS s ON (days.day = s.day)
ORDER BY days.day
`

type UserRegistrationDailyCountStatsParams struct {
	From pgtype.Date
	To   pgtype.Date
}

type UserRegistrationDailyCountStatsRow struct {
	Day       pgtype.Date
	UserCount int64
}

func (q *Queries) UserRegistrationDailyCountStats(ctx context.Context, arg UserRegistrationDailyCountStatsParams) ([]UserRegistrationDailyCountStatsRow, error) {
	rows, err := q.db.Query(ctx, userRegistrationDailyCountStats, arg.From, arg.To)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserRegistrationDailyCountStatsRow
	for rows.Next() {
		var i UserRegistrationDailyCountStatsRow
		if err := rows.Scan(&i.Day, &i.UserCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userSocialProfileChangeHistoryNew = `-- name: UserSocialProfileChangeHistoryNew :exec
INSERT INTO user_social_profile_change_history (user_id, user_social_profile_id, email, username, name, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
`

type UserSocialProfileChangeHistoryNewParams struct {
	UserID              int64
	UserSocialProfileID int64
	Email               string
	Username            string
	Name                string
	CreatedAt           pgtype.Timestamp
}

func (q *Queries) UserSocialProfileChangeHistoryNew(ctx context.Context, arg UserSocialProfileChangeHistoryNewParams) error {
	_, err := q.db.Exec(ctx, userSocialProfileChangeHistoryNew,
		arg.UserID,
		arg.UserSocialProfileID,
		arg.Email,
		arg.Username,
		arg.Name,
		arg.CreatedAt,
	)
	return err
}

const userSocialProfileGet = `-- name: UserSocialProfileGet :one
SELECT id, user_id, email, username, name
FROM user_social_profiles usp
WHERE usp.social_provider = $1
  AND usp.social_provider_user_id = $2
  AND usp.deleted_at IS NULL
    FOR UPDATE
`

type UserSocialProfileGetParams struct {
	SocialProvider       SocialProvider
	SocialProviderUserID string
}

type UserSocialProfileGetRow struct {
	ID       int64
	UserID   int64
	Email    string
	Username string
	Name     string
}

func (q *Queries) UserSocialProfileGet(ctx context.Context, arg UserSocialProfileGetParams) (UserSocialProfileGetRow, error) {
	row := q.db.QueryRow(ctx, userSocialProfileGet, arg.SocialProvider, arg.SocialProviderUserID)
	var i UserSocialProfileGetRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Email,
		&i.Username,
		&i.Name,
	)
	return i, err
}

const userSocialProfileUpdate = `-- name: UserSocialProfileUpdate :exec
UPDATE user_social_profiles
SET email      = $1,
    username   = $2,
    name       = $3,
    updated_at = $4
WHERE id = $5
`

type UserSocialProfileUpdateParams struct {
	Email     string
	Username  string
	Name      string
	UpdatedAt pgtype.Timestamp
	ID        int64
}

func (q *Queries) UserSocialProfileUpdate(ctx context.Context, arg UserSocialProfileUpdateParams) error {
	_, err := q.db.Exec(ctx, userSocialProfileUpdate,
		arg.Email,
		arg.Username,
		arg.Name,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
