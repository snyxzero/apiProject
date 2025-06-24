CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
COMMENT ON TABLE users IS 'Таблица для хранения информации о пользователях';
COMMENT ON COLUMN users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN users.name IS 'Имя пользователя';

CREATE TABLE breweries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);
COMMENT ON TABLE breweries IS 'Таблица для хранения информации о пивоварнях';
COMMENT ON COLUMN breweries.id IS 'Уникальный идентификатор пивоварни';
COMMENT ON COLUMN breweries.name IS 'Название пивоварни';


CREATE TABLE beers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    breweries_id INTEGER NOT NULL,
    CONSTRAINT fk_brewery
        FOREIGN KEY(breweries_id) REFERENCES breweries(id) ON DELETE CASCADE
);
COMMENT ON TABLE beers IS 'Таблица для хранения информации о сортах пива';
COMMENT ON COLUMN beers.id IS 'Уникальный идентификатор сорта пива';
COMMENT ON COLUMN beers.name IS 'Название сорта пива';
COMMENT ON COLUMN beers.breweries_id IS 'Идентификатор пивоварни, выпускающей этот сорт';

CREATE TABLE user_beer_ratings (
    id SERIAL PRIMARY KEY,
    users_id INTEGER NOT NULL,
    beers_id INTEGER NOT NULL,
    rating SMALLINT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    CONSTRAINT fk_users
        FOREIGN KEY(users_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_beers
        FOREIGN KEY(beers_id) REFERENCES beers(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_beer_rating UNIQUE (users_id, beers_id)
);
COMMENT ON TABLE user_beer_ratings IS 'Таблица для хранения оценок сортов пива пользователями';
COMMENT ON COLUMN user_beer_ratings.id IS 'Уникальный идентификатор оценки';
COMMENT ON COLUMN user_beer_ratings.users_id IS 'Идентификатор пользователя, поставившего оценку';
COMMENT ON COLUMN user_beer_ratings.beers_id IS 'Идентификатор сорта пива, который оценили';
COMMENT ON COLUMN user_beer_ratings.rating IS 'Оценка от 1 до 5';
COMMENT ON CONSTRAINT unique_user_beer_rating ON user_beer_ratings IS 'Обеспечивает уникальность оценки пользователя для конкретного сорта пива';
--