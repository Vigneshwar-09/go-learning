package model

import (
	"math/rand"
	"strconv"
	"time"
)

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

type OAuth struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	AccessToken string    `bson:"accessToken" json:"accessToken"`
	TokenType   string    `bson:"tokenType" json:"tokenType"`
	ExpiresIn   time.Time `bson:"expiresIn" json:"expiresIn"`
}

func NewOAuth(response OAuthResponse) *OAuth {

	oAuth := new(OAuth)

	oAuth.ID = "TOK" + strconv.Itoa(rand.Intn(10000000))
	oAuth.AccessToken = response.AccessToken
	oAuth.TokenType = response.TokenType
	oAuth.ExpiresIn = time.Now().Add(time.Second * time.Duration(response.ExpiresIn))

	return oAuth
}

type User struct {
	ID                               interface{}       `json:"id" bson:"_id,omitempty"`
	AccountHistory                   []interface{}     `json:"account_history"`
	ActiveTournamentBanner           interface{}       `json:"active_tournament_banner"`
	AvatarURL                        string            `json:"avatar_url"`
	Badges                           []interface{}     `json:"badges"`
	BeatmapPlaycountsCount           int64             `json:"beatmap_playcounts_count"`
	CommentsCount                    int64             `json:"comments_count"`
	Country                          Country           `json:"country"`
	CountryCode                      string            `json:"country_code"`
	Cover                            Cover             `json:"cover"`
	CoverURL                         string            `json:"cover_url"`
	DefaultGroup                     string            `json:"default_group"`
	Discord                          string            `json:"discord"`
	FavouriteBeatmapsetCount         int64             `json:"favourite_beatmapset_count"`
	FollowerCount                    int64             `json:"follower_count"`
	GraveyardBeatmapsetCount         int64             `json:"graveyard_beatmapset_count"`
	Groups                           []interface{}     `json:"groups"`
	GuestBeatmapsetCount             int64             `json:"guest_beatmapset_count"`
	HasSupported                     bool              `json:"has_supported"`
	Interests                        string            `json:"interests"`
	IsActive                         bool              `json:"is_active"`
	IsBot                            bool              `json:"is_bot"`
	IsDeleted                        bool              `json:"is_deleted"`
	IsOnline                         bool              `json:"is_online"`
	IsSupporter                      bool              `json:"is_supporter"`
	JoinDate                         string            `json:"join_date"`
	Kudosu                           Kudosu            `json:"kudosu"`
	LastVisit                        string            `json:"last_visit"`
	Location                         interface{}       `json:"location"`
	LovedBeatmapsetCount             int64             `json:"loved_beatmapset_count"`
	MappingFollowerCount             int64             `json:"mapping_follower_count"`
	MaxBlocks                        int64             `json:"max_blocks"`
	MaxFriends                       int64             `json:"max_friends"`
	MonthlyPlaycounts                []Count           `json:"monthly_playcounts"`
	NominatedBeatmapsetCount         int64             `json:"nominated_beatmapset_count"`
	Occupation                       interface{}       `json:"occupation"`
	Page                             Page              `json:"page"`
	PendingBeatmapsetCount           int64             `json:"pending_beatmapset_count"`
	Playmode                         string            `json:"playmode"`
	Playstyle                        []string          `json:"playstyle"`
	PmFriendsOnly                    bool              `json:"pm_friends_only"`
	PostCount                        int64             `json:"post_count"`
	PreviousUsernames                []interface{}     `json:"previous_usernames"`
	ProfileColour                    interface{}       `json:"profile_colour"`
	ProfileOrder                     []string          `json:"profile_order"`
	RankHistory                      RankHistory       `json:"rankHistory"`
	RankHighest                      RankHighest       `json:"rank_highest"`
	UserRankHistory                  RankHistory       `json:"rank_history"`
	RankedAndApprovedBeatmapsetCount int64             `json:"ranked_and_approved_beatmapset_count"`
	RankedBeatmapsetCount            int64             `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts             []Count           `json:"replays_watched_counts"`
	ScoresBestCount                  int64             `json:"scores_best_count"`
	ScoresFirstCount                 int64             `json:"scores_first_count"`
	ScoresPinnedCount                int64             `json:"scores_pinned_count"`
	ScoresRecentCount                int64             `json:"scores_recent_count"`
	Statistics                       Statistics        `json:"statistics"`
	SupportLevel                     int64             `json:"support_level"`
	Title                            interface{}       `json:"title"`
	TitleURL                         interface{}       `json:"title_url"`
	Twitter                          interface{}       `json:"twitter"`
	UnrankedBeatmapsetCount          int64             `json:"unranked_beatmapset_count"`
	UserAchievements                 []UserAchievement `json:"user_achievements"`
	Username                         string            `json:"username"`
	Website                          interface{}       `json:"website"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Cover struct {
	CustomURL string      `json:"custom_url"`
	ID        interface{} `json:"id"`
	URL       string      `json:"url"`
}

type Kudosu struct {
	Available int64 `json:"available"`
	Total     int64 `json:"total"`
}

type Count struct {
	Count     int64  `json:"count"`
	StartDate string `json:"start_date"`
}

type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

type RankHighest struct {
	Rank      int64  `json:"rank"`
	UpdatedAt string `json:"updated_at"`
}

type RankHistory struct {
	Data []int64 `json:"data"`
	Mode string  `json:"mode"`
}

type Statistics struct {
	Count100               int64       `json:"count_100"`
	Count300               int64       `json:"count_300"`
	Count50                int64       `json:"count_50"`
	CountMiss              int64       `json:"count_miss"`
	CountryRank            int64       `json:"country_rank"`
	GlobalRank             int64       `json:"global_rank"`
	GlobalRankExp          int64       `json:"global_rank_exp"`
	GradeCounts            GradeCounts `json:"grade_counts"`
	HitAccuracy            float64     `json:"hit_accuracy"`
	IsRanked               bool        `json:"is_ranked"`
	Level                  Level       `json:"level"`
	MaximumCombo           int64       `json:"maximum_combo"`
	PlayCount              int64       `json:"play_count"`
	PlayTime               int64       `json:"play_time"`
	Pp                     float64     `json:"pp"`
	PpExp                  float64     `json:"pp_exp"`
	Rank                   Rank        `json:"rank"`
	RankedScore            int64       `json:"ranked_score"`
	ReplaysWatchedByOthers int64       `json:"replays_watched_by_others"`
	TotalHits              int64       `json:"total_hits"`
	TotalScore             int64       `json:"total_score"`
}

type GradeCounts struct {
	A   int64 `json:"a"`
	S   int64 `json:"s"`
	Sh  int64 `json:"sh"`
	Ss  int64 `json:"ss"`
	SSH int64 `json:"ssh"`
}

type Level struct {
	Current  int64 `json:"current"`
	Progress int64 `json:"progress"`
}

type Rank struct {
	Country int64 `json:"country"`
}

type UserAchievement struct {
	AchievedAt    string `json:"achieved_at"`
	AchievementID int64  `json:"achievement_id"`
}
