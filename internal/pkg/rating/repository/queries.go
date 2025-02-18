package ratrepository

const (

	queryPostRating = `
	INSERT INTO
		ratings (user_id, movie_id, rating)
	VALUES
		($1, $2, $3);
	`

	queryChangeRating = `
	UPDATE ratings
	SET rating = $1
	WHERE user_id = $2 AND movie_id = $3;
	`

	queryGetMovieRating = `

	SELECT movies.rating
	FROM movies 
	WHERE movies.id = $1;
	`

	queryGetMovieVotesnum = `
	SELECT movies.votesnum
	FROM movies 
	WHERE movies.id = $1;
	`

	queryGetRatingUserCount = `
	SELECT COUNT(*)
	FROM ratings
	WHERE ratings.user_id = $1 and ratings.movie_id = $2;
	`

	queryGetOldRatingUser = `
	SELECT ratings.rating
	FROM ratings
	WHERE ratings.user_id = $1 AND ratings.movie_id = $2;
	`

	queryIncrementVotesnum = `
	UPDATE movies
	SET votesnum = votesnum + 1
	WHERE id = $1;
	`

	querySetMovieRating = `
	UPDATE movies
	SET rating = $1
	WHERE id = $2;
	`
)

const (
	queryUserExist = `
	SELECT COUNT(*)
	FROM users
	WHERE id = $1;
	`

	queryMovieExist = `
	SELECT COUNT(*)
	FROM movies
	WHERE id = $1;
	`
)
