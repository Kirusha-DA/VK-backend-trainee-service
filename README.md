# REST API

POST /actor -- create actor
PUT /actors/{id} -- update the whole actor info
PATCH /actors/{id} -- update some actor info
DELETE /actors/{id} -- delete actor

GET /movies?sort_by=rating -- (default) get movies sorted by rating desc
GET /movies?sort_by=name -- get movies sorted by name desc
GET /movies?sort_by=release_date -- get movies sorted by release date desc
GET /movies/{name} -- get movie by id
POST /movie -- create movie
PUT /movies/{id} -- update the whole movie info
PATCH /movies/{id} -- update some movie info
DELETE /movies/{id} -- delete movie

GET /movies/{actor_name} -- get movie by actor name
GET /actors/movies -- get actors with list of movies
